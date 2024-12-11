package handlers

import (
	"fmt"
	"strings"

	"github.com/fairmath/shuttle/internal/server/handlers/dto"
)

const txDebugRPCName = "debug"

type Debug struct {
	name   string
	target EthProxy
}

func NewDebug(target EthProxy) *Debug {
	return &Debug{
		name:   txDebugRPCName,
		target: target,
	}
}

func (t *Debug) Name() string {
	return t.name
}

func (t *Debug) TraceTransaction(hash string, _ any) (dto.TraceTx, error) {
	hash, _ = strings.CutPrefix(hash, "0x")

	txInfo, err := t.target.GetTx(hash)
	if err != nil {
		return dto.TraceTx{}, fmt.Errorf("get tx by hash: %w", err)
	}

	block, err := t.target.GetBlockTxsByHeight(txInfo.TxResponse.Height)
	if err != nil {
		return dto.TraceTx{}, fmt.Errorf("get tx by height: %s: %w", txInfo.TxResponse.Height, err)
	}

	traceTx, err := dto.EthTrace(block, txInfo)
	if err != nil {
		return dto.TraceTx{}, fmt.Errorf("eth trace: %w", err)
	}

	return traceTx, nil
}
