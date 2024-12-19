package api

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"strconv"
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

type CosmosAPI struct {
	blocksAPI bservice.ClientService
	txsAPI    tservice.ClientService
	bankAPI   bankService.ClientService
}

func NewCosmosAPI(serverURL string) (*CosmosAPI, error) {
	parsedURL, err := url.Parse(serverURL)
	if err != nil {
		return nil, fmt.Errorf("parse url: '%s': %w", serverURL, err)
	}

	return &CosmosAPI{
		blocksAPI: bclient.NewHTTPClientWithConfig(strfmt.Default, &bclient.TransportConfig{
			Host:     parsedURL.Host,
			BasePath: parsedURL.Path,
			Schemes:  []string{parsedURL.Scheme},
		}).Service,
		txsAPI: tclient.NewHTTPClientWithConfig(strfmt.Default, &tclient.TransportConfig{
			Host:     parsedURL.Host,
			BasePath: parsedURL.Path,
			Schemes:  []string{parsedURL.Scheme},
		}).Service,
		bankAPI: bankClient.NewHTTPClientWithConfig(strfmt.Default, &bankClient.TransportConfig{
			Host:     parsedURL.Host,
			BasePath: parsedURL.Path,
			Schemes:  []string{parsedURL.Scheme},
		}).Query,
	}, nil
}

func (t *CosmosAPI) GetBlockByHeight(height string) (*Block, error) {
	if height == "latest" {
		response, err := t.blocksAPI.ServiceGetLatestBlock(bservice.NewServiceGetLatestBlockParams())
		if err != nil {
			return nil, fmt.Errorf("get %s block: %w", height, err)
		}

		return (*Block)(response.Payload), nil
	}

	if strings.HasPrefix(height, "0x") {
		var n uint64
		_, _ = fmt.Sscanf(height, "0x%x", &n)
		height = strconv.FormatUint(n, 10)
	}

	response, err := t.blocksAPI.ServiceGetBlockByHeight(bservice.NewServiceGetBlockByHeightParams().WithHeight(height))
	if err != nil {
		return nil, fmt.Errorf("get %s block: %w", height, err)
	}

	return (*Block)(response.Payload), nil
}

func (t *CosmosAPI) GetBlockTxsByHeight(height string) (*BlockWithTxs, error) {
	res, err := t.txsAPI.ServiceGetBlockWithTxs(tservice.NewServiceGetBlockWithTxsParams().WithHeight(height))
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

func (t *CosmosAPI) GetTx(hash string) (*TxInfo, error) {
	res, err := t.txsAPI.ServiceGetTx(tservice.NewServiceGetTxParams().WithHash(hash))
	if err != nil {
		return nil, fmt.Errorf("get tx: %w", err)
	}

	return (*TxInfo)(res.Payload), nil
}

func (t *CosmosAPI) GetBalance(prefix, denom string, addr strfmt.Base64) (string, error) {
	strAddr, err := bech32.ConvertAndEncode(prefix, addr)
	if err != nil {
		return "", fmt.Errorf("decode addr: '0x%s'", hex.EncodeToString(addr))
	}

	res, err := t.bankAPI.QueryBalance(
		bankService.NewQueryBalanceParams().
			WithAddress(strAddr).
			WithDenom(&denom),
	)
	if err != nil {
		return "", fmt.Errorf("query balance: %w", err)
	}

	return res.Payload.Balance.Amount, nil
}
