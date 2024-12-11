package handlers

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/fairmath/shuttle/internal/client/tendermint/api"
	"github.com/fairmath/shuttle/internal/server/handlers/dto"

	"github.com/go-openapi/strfmt"
)

const ethRPCName = "eth"

type EthProxy interface {
	GetBlockByHeight(height string) (*api.Block, error)
	GetBlockTxsByHeight(height string) (*api.BlockWithTxs, error)
	GetBalance(prefix, denom string, addr strfmt.Base64) (string, error)
	GetTx(hash string) (*api.TxInfo, error)
}

type EthServer struct {
	name   string
	target EthProxy
}

func NewEthServer(proxyTo EthProxy) *EthServer {
	return &EthServer{
		name:   ethRPCName,
		target: proxyTo,
	}
}

func (e *EthServer) Name() string {
	return e.name
}

func (e *EthServer) GetBlockByNumber(height string, _ bool) (any, error) {
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

	ethBlock, err := dto.FromCosmosBlockWithTxs(block)
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

	amount, err := e.target.GetBalance("fairmath", "fmth", addrBytes)
	if err != nil {
		return "", fmt.Errorf("get balance: %w", err)
	}

	var ok bool

	value := big.NewInt(0)

	if value, ok = value.SetString(amount, 10); !ok { //nolint:gomnd // base
		return "", fmt.Errorf("parse balance: %w", err)
	}

	return "0x" + value.Text(16), nil //nolint:gomnd // base
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
