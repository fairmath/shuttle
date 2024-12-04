package server

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math/big"
	"math/rand/v2"
	"net/http"
	"sync"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	bfttypes "github.com/cometbft/cometbft/types"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type WebsocketPool struct {
	log          *zap.Logger
	cosmosClient *rpchttp.HTTP
	events       <-chan coretypes.ResultEvent
}

type JSONRpcMsg struct {
	ID       uint64 `json:"id"`
	Params   []any  `json:"params"`
	Method   string `json:"method"`
	Protocol string `json:"jsonrpc"`
}

func NewWebsocketPool(log *zap.Logger) *WebsocketPool {
	return &WebsocketPool{
		log: log,
	}
}

func (wp *WebsocketPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		wp.log.Error("upgrade ws connection", zap.Error(err))

		return
	}

	if wp.cosmosClient, err = rpchttp.New("https://testnet.computer.fairmath.xyz", "/websocket"); err != nil {
		wp.log.Error("open cosmos ws connection", zap.Error(err))

		return
	}

	if err := wp.cosmosClient.Start(); err != nil {
		wp.log.Error("start cosmos ws client", zap.Error(err))

		return
	}

	wp.serveConnection(r.Context(), connection)
}

func (wp *WebsocketPool) serveConnection(ctx context.Context, connection *websocket.Conn) {
	wp.log.Info("serve connection")

	defer connection.Close()

	done := make(chan struct{})

	once := sync.Once{}
	subscribers := make(chan uint32)
	proxyDone := make(chan struct{})

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

			if msg.Method == "eth_subscribe" {
				once.Do(func() {
					ctx, cancel := context.WithCancel(context.Background())
					defer cancel()

					query := "tm.event = 'NewBlock'"

					wp.events, err = wp.cosmosClient.Subscribe(ctx, "github.com/fairmath/shuttle", query)
					if err != nil {
						wp.log.Error("subscribe on cosmos events", zap.Error(err))

						return
					}

					go func() {
						defer close(proxyDone)

						wp.startProxy(connection, subscribers, done)
					}()
				})

				wp.log.Info("request subscription", zap.String("msg", string(message)))
				subscriptionID := rand.Uint32()

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

func (wp *WebsocketPool) startProxy(conn *websocket.Conn, subscribers chan uint32, done chan struct{}) {
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
				Nonce:       ethtypes.BlockNonce(binary.LittleEndian.AppendUint64([]byte{}, rand.Uint64())),
				MixDigest:   common.Hash(data.Block.Header.ConsensusHash),
				ReceiptHash: common.Hash(data.Block.Header.AppHash),
				UncleHash:   common.Hash(data.Block.Header.NextValidatorsHash),
				Root:        common.Hash(data.Block.Header.EvidenceHash),
				TxHash:      common.Hash(data.Block.Header.DataHash),
				Extra:       []byte("empty"),
				Difficulty:  big.NewInt(int64(0x1046bb7e3f8)),

				Time: uint64(data.Block.Header.Time.Unix()),
			}

			marshaled, err := json.Marshal(blockHeader)
			if err != nil {
				wp.log.Error("marshal block", zap.Error(err))

				continue
			}

			for _, sID := range subscriberIDs {
				resp := fmt.Sprintf(`{"id":"%d","jsonrpc":"2.0","method":"eth_subscription","params":{"result":%s,"subscription":"0x%x"}}`,
					1,
					string(marshaled),
					sID,
				)
				if err := conn.WriteMessage(websocket.TextMessage, []byte(resp)); err != nil {
					wp.log.Error("write err", zap.Error(err))

					return
				}

				wp.log.Debug("cosmos block forwarded", zap.Int64("height", data.Block.Header.Height), zap.Uint32s("to", subscriberIDs))
			}
		}
	}
}
