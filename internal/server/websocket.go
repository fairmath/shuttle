package server

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math/big"
	"math/rand"
	"net/http"
	"net/url"
	"sync"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	bfttypes "github.com/cometbft/cometbft/types"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

const ETHSubscribeMethod = "eth_subscribe"

// todo: use tendermint ws connection instead of WebsocketPool(see client/tendermint/api/internal/tendermint.go)
type WebsocketPool struct {
	log          *zap.Logger
	tendermintWS string
	cosmosClient *rpchttp.HTTP
	events       <-chan coretypes.ResultEvent
}

type JSONRpcMsg struct { //nolint:tagliatelle // ethereum serialized json
	ID       uint64 `json:"id"`
	Params   any    `json:"params"`
	Method   string `json:"method"`
	Protocol string `json:"jsonrpc"`
}

type JSONRpcSubscriptionMsg struct { //nolint:tagliatelle // ethereum serialized json
	Result       any    `json:"result"`
	Subscription string `json:"subscription"`
}

func NewWebsocketPool(tendermintWS string, log *zap.Logger) (*WebsocketPool, error) {
	parsed, err := url.Parse(tendermintWS)
	if err != nil {
		return nil, fmt.Errorf("parse tendermint ws url: '%s': %w", tendermintWS, err)
	}

	log.Info("creating ws connection", zap.String("host", parsed.Host), zap.String("path", parsed.Path))

	cosmosClient, err := rpchttp.New(parsed.Scheme+"://"+parsed.Host, parsed.Path)
	if err != nil {
		return nil, fmt.Errorf("create tendermint websocket: %w", err)
	}

	if err := cosmosClient.Start(); err != nil {
		return nil, fmt.Errorf("start listening tendermint websocket: %w", err)
	}

	return &WebsocketPool{
		log:          log,
		cosmosClient: cosmosClient,
		tendermintWS: tendermintWS,
	}, nil
}

func (wp *WebsocketPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(_ *http.Request) bool {
			return true
		},
	}

	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		wp.log.Error("upgrade ws connection", zap.Error(err))

		return
	}

	wp.serveConnection(r.Context(), connection)
}

//nolint:funlen,gocognit // websocket should be refactored later
func (wp *WebsocketPool) serveConnection(ctx context.Context, connection *websocket.Conn) {
	wp.log.Info("serve connection")

	defer connection.Close()

	done := make(chan struct{})
	once := sync.Once{}
	subscribers := make(chan uint32)
	proxyDone := make(chan struct{})

	//nolint:contextcheck // fix after refactoring
	go func() {
		defer close(done)

		for {
			mt, message, err := connection.ReadMessage()
			if err != nil || mt == websocket.CloseMessage {
				wp.log.Warn("read message", zap.Error(err))

				break
			}

			if mt == websocket.PingMessage {
				if err := connection.WriteMessage(websocket.PongMessage, []byte("")); err != nil {
					wp.log.Warn("pong message", zap.Error(err))
				}

				continue
			}

			msg := &JSONRpcMsg{}
			if mt == websocket.TextMessage {
				if err := json.Unmarshal(message, msg); err != nil {
					wp.log.Error("unmarshal ws message", zap.Error(err))

					continue
				}
			}

			if msg.Method == ETHSubscribeMethod {
				once.Do(func() {
					if err := wp.prepareSubscription(); err != nil {
						wp.log.Error("subscribe on cosmos events", zap.Error(err))

						return
					}

					go func() {
						defer close(proxyDone)

						wp.startProxy(connection, msg.ID, subscribers, done)
					}()
				})

				subscriptionID := rand.Uint32() //nolint:gosec //no reason for crypto rand

				response := fmt.Sprintf(`{"jsonrpc":"2.0","result":"0x%x","id":%d}`, subscriptionID, msg.ID)
				if err := connection.WriteMessage(websocket.TextMessage, []byte(response)); err != nil {
					wp.log.Error("approve subscription", zap.Error(err))

					continue
				}

				subscribers <- subscriptionID
			}
		}
	}()

	select {
	case <-ctx.Done():
	case <-done:
	}

	<-proxyDone
}

func (wp *WebsocketPool) prepareSubscription() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var err error

	query := "tm.event = 'NewBlock'"

	wp.events, err = wp.cosmosClient.Subscribe(ctx, "shuttle", query)
	if err != nil {
		return fmt.Errorf("subscribe: %w", err)
	}

	return nil
}

func (wp *WebsocketPool) startProxy(conn *websocket.Conn, id uint64, subscribers chan uint32, done chan struct{}) {
	wp.log.Info("start proxying")

	subscriberIDs := []uint32{}

	for {
		select {
		case <-done:
			return
		case id := <-subscribers:
			subscriberIDs = append(subscriberIDs, id)
		case e := <-wp.events:
			data, ok := e.Data.(bfttypes.EventDataNewBlock)
			if !ok {
				continue
			}

			blockHeader := ethtypes.Header{
				Number:      big.NewInt(data.Block.Height),
				Coinbase:    common.Address(data.Block.Header.ProposerAddress),
				ParentHash:  common.Hash(data.Block.Header.LastCommitHash),
				Nonce:       ethtypes.BlockNonce(binary.LittleEndian.AppendUint64([]byte{}, rand.Uint64())), //nolint:gosec,lll //no reason for crypto rand
				MixDigest:   common.Hash(data.Block.Header.ConsensusHash),
				ReceiptHash: common.Hash{},
				UncleHash:   common.Hash{},
				Root:        common.Hash(data.Block.Header.EvidenceHash),
				Extra:       []byte{},
				Difficulty:  big.NewInt(int64(0x1046bb7e3f8)),      //nolint:mnd // need to translate from cosmos chain
				Time:        uint64(data.Block.Header.Time.Unix()), //nolint:gosec // always positive or 0
			}

			msg := JSONRpcMsg{
				ID:       id,
				Method:   "eth_subscription",
				Protocol: "2.0",
				Params: JSONRpcSubscriptionMsg{
					Subscription: "0x%x",
					Result:       blockHeader,
				},
			}

			respTemplate, err := json.Marshal(msg)
			if err != nil {
				wp.log.Error("marshal eth ws event", zap.Error(err))

				continue
			}

			for _, sID := range subscriberIDs {
				resp := fmt.Sprintf(string(respTemplate), sID)

				if err := conn.WriteMessage(websocket.TextMessage, []byte(resp)); err != nil {
					wp.log.Error("write err", zap.Error(err))

					return
				}

				wp.log.Debug("cosmos block forwarded", zap.Int64("height", data.Block.Header.Height), zap.Uint32s("to", subscriberIDs))
			}
		}
	}
}
