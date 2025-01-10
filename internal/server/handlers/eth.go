package handlers

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/go-openapi/strfmt"

	"github.com/fairmath/shuttle/internal/client/tendermint/api"
	"github.com/fairmath/shuttle/internal/server/handlers/dto"
)

const ethRPCName = "eth"

type EthProxy interface {
	GetBlockByHeight(height string) (*api.Block, error)
	GetBlockTxsByHeight(height string) (*api.BlockWithTxs, error)
	GetBalance(prefix, denom string, addr strfmt.Base64) (string, error)
	GetTx(hash string) (*api.TxInfo, error)

	BlockByHash(ctx context.Context, hash []byte) (api.Block, error)
}

type EthServer struct {
	name               string
	target             EthProxy
	cosmosDenom        string
	cosmosBech32Prefix string
}

func NewEthServer(proxyTo EthProxy, cfg Config) *EthServer {
	return &EthServer{
		name:               ethRPCName,
		target:             proxyTo,
		cosmosDenom:        cfg.CosmosDenom,
		cosmosBech32Prefix: cfg.Bech32AddrPrefix,
	}
}

func (e *EthServer) Name() string {
	return e.name
}

func (e *EthServer) BlockNumber() (any, error) {
	const latest = "latest"

	b, err := e.target.GetBlockByHeight(latest)
	if err != nil {
		return nil, fmt.Errorf("get block '%s': %w", latest, err)
	}

	num, err := strconv.ParseInt(b.Block.Header.Height, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("parse block number: %w", err)
	}

	return "0x" + strconv.FormatInt(num, 16), nil
}

func (e *EthServer) GetBlockByNumber(height string, fullTransactions bool) (any, error) {
	var (
		block *api.BlockWithTxs
		err   error
	)

	if height == "latest" {
		b, err := e.target.GetBlockByHeight(height)
		if err != nil {
			return nil, fmt.Errorf("get block '%s': %w", height, err)
		}

		height = b.Block.Header.Height
	}

	if block, err = e.target.GetBlockTxsByHeight(height); err != nil {
		return nil, fmt.Errorf("get %s txs: %w", height, err)
	}

	ethBlock, err := dto.FromCosmosBlockWithTxs(block, fullTransactions)
	if err != nil {
		return nil, fmt.Errorf("block conversion: %w", err)
	}

	return ethBlock, nil
}

func (e *EthServer) GetBalance(addr string, _ string) (string, error) {
	addr, _ = strings.CutPrefix(addr, "0x")

	addrBytes, err := hex.DecodeString(addr)
	if err != nil {
		return "", fmt.Errorf("decode addr '%s': %w", addr, err)
	}

	amount, err := e.target.GetBalance(e.cosmosBech32Prefix, e.cosmosDenom, addrBytes)
	if err != nil {
		return "", fmt.Errorf("get balance: %w", err)
	}

	var ok bool

	value := big.NewInt(0)

	if value, ok = value.SetString(amount, 10); !ok { //nolint:mnd // base
		return "", fmt.Errorf("parse balance: %w", err)
	}

	return "0x" + value.Text(16), nil //nolint:mnd // base
}

func (e *EthServer) GetTransactionReceipt(hash string) (any, error) {
	hash, _ = strings.CutPrefix(hash, "0x")

	txInfo, err := e.target.GetTx(hash)
	if err != nil {
		return nil, fmt.Errorf("get tx by hash: %w", err)
	}

	block, err := e.target.GetBlockTxsByHeight(txInfo.TxResponse.Height)
	if err != nil {
		return nil, fmt.Errorf("get tx by height: %s: %w", txInfo.TxResponse.Height, err)
	}

	res, err := dto.ToEthTxReceipt(block, txInfo)
	if err != nil {
		return nil, fmt.Errorf("eth receipt converter: %w", err)
	}

	return res, nil
}

func (e *EthServer) GetTransactionByHash(hash string) (any, error) {
	trimHash, _ := strings.CutPrefix(hash, "0x")

	txInfo, err := e.target.GetTx(trimHash)
	if err != nil {
		return nil, fmt.Errorf("get tx by hash: %w", err)
	}

	block, err := e.target.GetBlockTxsByHeight(txInfo.TxResponse.Height)
	if err != nil {
		return nil, fmt.Errorf("get tx by height: %s: %w", txInfo.TxResponse.Height, err)
	}

	txs, err := dto.ToEthTxs(block)
	if err != nil {
		return nil, fmt.Errorf("eth tx converter: %w", err)
	}

	for _, t := range txs {
		if t.Hash.String() == hash {
			return t, nil
		}
	}

	return nil, errors.New("not found")
}

func (e *EthServer) GetBlockByHash(hash string, fullTransactions bool) (any, error) {
	trimHash, _ := strings.CutPrefix(hash, "0x")

	h, err := hex.DecodeString(trimHash)
	if err != nil {
		return nil, fmt.Errorf("parse hash: %s: %w", trimHash, err)
	}

	cosmosBlock, err := e.target.BlockByHash(context.Background(), h)
	if err != nil {
		return nil, fmt.Errorf("get block by hash: %w", err)
	}

	num, err := strconv.ParseUint(cosmosBlock.Block.Header.Height, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("parse block height: %w", err)
	}

	res, err := e.GetBlockByNumber("0x"+strconv.FormatUint(num, 16), fullTransactions)
	if err != nil {
		return nil, fmt.Errorf("block by num: %w", err)
	}

	return res, nil
}
