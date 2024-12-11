package dto

import (
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fairmath/shuttle/internal/client/tendermint/api"
)

type Call struct { //nolint:tagliatelle // ethereum serialized json
	Gas     HexUint64      `json:"gas"`
	GasUsed HexUint64      `json:"gasUsed"`
	Type    string         `json:"type"`
	From    common.Address `json:"from"`
	To      common.Address `json:"to"`
	Input   string         `json:"input"`
	Output  string         `json:"output"`
}

type TraceTx struct { //nolint:tagliatelle // ethereum serialized json
	Gas          HexUint64      `json:"gas"`
	From         common.Address `json:"from"`
	GasUsed      HexUint64      `json:"gasUsed"`
	Status       HexUint64      `json:"status"`
	To           common.Address `json:"to"`
	Type         string         `json:"type"`
	Value        HexUint64      `json:"value"`
	Input        string         `json:"input"`
	Output       string         `json:"output"`
	Error        *string        `json:"error,omitempty"`
	RevertReason *string        `json:"revertReason,omitempty"`
	Calls        []Call         `json:"calls"`
}

func EthTrace(block *api.BlockWithTxs, txInfo *api.TxInfo) (TraceTx, error) {
	res, err := ToEthTxReceipt(block, txInfo)
	if err != nil {
		return TraceTx{}, fmt.Errorf("eth receipt converter: %w", err)
	}

	gasLimit, err := strconv.ParseUint(txInfo.Tx.AuthInfo.Fee.GasLimit, 10, 64)
	if err != nil {
		return TraceTx{}, fmt.Errorf("gas limit '%s' conversion: %w", txInfo.Tx.AuthInfo.Fee.GasLimit, err)
	}

	var amount uint64

	if len(txInfo.Tx.AuthInfo.Fee.Amount) > 0 {
		amount, err = strconv.ParseUint(txInfo.Tx.AuthInfo.Fee.Amount[0].Amount, 10, 64)
		if err != nil {
			return TraceTx{}, fmt.Errorf("parse amount '%s': %w", txInfo.Tx.AuthInfo.Fee.Amount[0].Amount, err)
		}
	}

	msgs := make([]byte, 0, len(txInfo.Tx.Body.Messages))

	for _, data := range txInfo.Tx.Body.Messages {
		bindata, _ := data.MarshalJSON()
		msgs = append(msgs, bindata...)
	}

	return TraceTx{
		Gas:     HexUint64(gasLimit),
		GasUsed: res.GasUsed,
		From:    res.From,
		To:      res.To,
		Type:    "CALL",
		Value:   HexUint64(amount),
		Input:   "0x" + hex.EncodeToString(msgs),
		Output:  "0x" + hex.EncodeToString([]byte(txInfo.TxResponse.Data)),
		Status:  res.Status,
		Calls: []Call{
			{
				Type:    "STATICCALL",
				Gas:     HexUint64(gasLimit),
				GasUsed: res.GasUsed,
				To:      res.To,
				From:    res.From,
				Input:   "0x" + hex.EncodeToString(msgs),
				Output:  "0x" + hex.EncodeToString([]byte(txInfo.TxResponse.Data)),
			},
		},
	}, nil
}
