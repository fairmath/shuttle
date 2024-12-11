package api

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"strings"

	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/go-openapi/strfmt"

	bankClient "github.com/fairmath/shuttle/internal/client/tendermint/api/internal/bank/client"
	bankService "github.com/fairmath/shuttle/internal/client/tendermint/api/internal/bank/client/query"
	bclient "github.com/fairmath/shuttle/internal/client/tendermint/api/internal/blocks/client"
	bservice "github.com/fairmath/shuttle/internal/client/tendermint/api/internal/blocks/client/service"
	tclient "github.com/fairmath/shuttle/internal/client/tendermint/api/internal/txs/client"
	tservice "github.com/fairmath/shuttle/internal/client/tendermint/api/internal/txs/client/service"
)

type CosmosApi struct {
	blocksApi bservice.ClientService
	txsApi    tservice.ClientService
	bankApi   bankService.ClientService
}

func NewCosmosApi(serverURL string) (*CosmosApi, error) {
	parsedUrl, err := url.Parse(serverURL)
	if err != nil {
		return nil, fmt.Errorf("parse url: '%s': %w", serverURL, err)
	}

	return &CosmosApi{
		blocksApi: bclient.NewHTTPClientWithConfig(strfmt.Default, &bclient.TransportConfig{
			Host:     parsedUrl.Host,
			BasePath: parsedUrl.Path,
			Schemes:  []string{parsedUrl.Scheme},
		}).Service,
		txsApi: tclient.NewHTTPClientWithConfig(strfmt.Default, &tclient.TransportConfig{
			Host:     parsedUrl.Host,
			BasePath: parsedUrl.Path,
			Schemes:  []string{parsedUrl.Scheme},
		}).Service,
		bankApi: bankClient.NewHTTPClientWithConfig(strfmt.Default, &bankClient.TransportConfig{
			Host:     parsedUrl.Host,
			BasePath: parsedUrl.Path,
			Schemes:  []string{parsedUrl.Scheme},
		}).Query,
	}, nil
}

func (t *CosmosApi) GetBlockByHeight(height string) (*Block, error) {
	if height == "latest" {
		response, err := t.blocksApi.ServiceGetLatestBlock(bservice.NewServiceGetLatestBlockParams())
		if err != nil {
			return nil, fmt.Errorf("get %s block: %w", height, err)
		}

		return (*Block)(response.Payload), nil
	}

	if strings.HasPrefix(height, "0x") {
		var n uint64
		_, _ = fmt.Sscanf(height, "0x%x", &n)
		height = fmt.Sprintf("%d", n)
	}

	response, err := t.blocksApi.ServiceGetBlockByHeight(bservice.NewServiceGetBlockByHeightParams().WithHeight(height))
	if err != nil {
		return nil, fmt.Errorf("get %s block: %w", height, err)
	}

	return (*Block)(response.Payload), nil
}

func (t *CosmosApi) GetBlockTxsByHeight(height string) (*BlockWithTxs, error) {
	res, err := t.txsApi.ServiceGetBlockWithTxs(tservice.NewServiceGetBlockWithTxsParams().WithHeight(height))
	if err != nil {
		return nil, fmt.Errorf("get block %s txs: %w", height, err)
	}

	blockTxs := (*BlockWithTxs)(res.Payload)

	txHashes := make([]strfmt.Base64, 0, len(res.Payload.Block.Data.Txs))

	for _, tx := range blockTxs.Block.Data.Txs {
		hash := sha256.Sum256(tx)
		txHashes = append(txHashes, strfmt.Base64(hash[:]))
	}

	blockTxs.Block.Data.Txs = txHashes

	return blockTxs, nil
}

func (t *CosmosApi) GetTx(hash string) (*TxInfo, error) {
	res, err := t.txsApi.ServiceGetTx(tservice.NewServiceGetTxParams().WithHash(hash))
	if err != nil {
		return nil, fmt.Errorf("get tx: %w", err)
	}

	return (*TxInfo)(res.Payload), nil
}

func (t *CosmosApi) GetBalance(prefix, denom string, addr strfmt.Base64) (string, error) {
	strAddr, err := bech32.ConvertAndEncode(prefix, addr)
	if err != nil {
		return "", fmt.Errorf("decode addr: '0x%s'", hex.EncodeToString(addr))
	}

	res, err := t.bankApi.QueryBalance(
		bankService.NewQueryBalanceParams().
			WithAddress(strAddr).
			WithDenom(&denom),
	)
	if err != nil {
		return "", fmt.Errorf("query balance: %w", err)
	}

	return res.Payload.Balance.Amount, nil
}
