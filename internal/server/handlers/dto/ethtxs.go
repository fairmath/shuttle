package dto

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"slices"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/fairmath/shuttle/internal/client/tendermint/api"
)

var ErrKeyNotFound = errors.New("public key not found")

type TxLog struct { //nolint:tagliatelle // ethereum serialized json
	Address          common.Address `json:"address"`
	Topics           []common.Hash  `json:"topics"`
	Data             string         `json:"data"`
	BlockNumber      HexUint64      `json:"blockNumber"`
	TransactionHash  common.Hash    `json:"transactionHash"`
	TransactionIndex HexUint64      `json:"transactionIndex"`
	BlockHash        common.Hash    `json:"blockHash"`
	LogIndex         HexUint64      `json:"logIndex"`
	Removed          bool           `json:"removed"`
}

type TxReceipt struct { //nolint:tagliatelle // ethereum serialized json
	BlockHash         common.Hash    `json:"blockHash"`
	BlockNumber       HexUint64      `json:"blockNumber"`
	ContractAddress   common.Address `json:"contractAddress"`
	CumulativeGasUsed HexUint64      `json:"cumulativeGasUsed"`
	EffectiveGasPrice HexUint64      `json:"effectiveGasPrice"`
	From              common.Address `json:"from"`
	GasUsed           HexUint64      `json:"gasUsed"`
	Logs              []TxLog        `json:"logs"`
	Bloom             ethtypes.Bloom `json:"logsBloom"`
	Status            HexUint64      `json:"status"`
	To                common.Address `json:"to"`
	TransactionHash   common.Hash    `json:"transactionHash"`
	TransactionIndex  HexUint64      `json:"transactionIndex"`
	Type              HexUint64      `json:"type"`
}

type Tx struct { //nolint:tagliatelle // ethereum serialized json
	BlockHash        common.Hash    `json:"blockHash"`
	BlockNumber      HexUint64      `json:"blockNumber"`
	Creates          common.Address `json:"creates"`
	From             common.Address `json:"from"`
	Gas              HexUint64      `json:"gas"`
	GasPrice         HexUint64      `json:"gasPrice"`
	Hash             common.Hash    `json:"hash"`
	Input            string         `json:"input"`
	Nonce            HexUint64      `json:"nonce"`
	To               common.Address `json:"to"`
	TransactionIndex HexUint64      `json:"transactionIndex"`
	Value            HexUint64      `json:"value"`
	Type             HexUint64      `json:"type"`
	ChainID          HexUint64      `json:"chainId"`
	V                HexUint64      `json:"v"`
	R                string         `json:"r"`
	S                string         `json:"s"`
	StandardV        HexUint64      `json:"standardV"`
	PublickKey       string         `json:"publicKey"`
	Raw              string         `json:"raw"`
}

func ToEthTxs(blockTx *api.BlockWithTxs) ([]Tx, error) {
	num, err := strconv.ParseUint(blockTx.Block.Header.Height, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("height %s convert: %w", blockTx.Block.Header.Height, err)
	}

	rtx := make([]Tx, 0, len(blockTx.Txs))

	for i, tx := range blockTx.Txs {
		gas, err := strconv.ParseInt(tx.AuthInfo.Fee.GasLimit, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("gas limit convert '%s': %w", tx.AuthInfo.Fee.GasLimit, err)
		}

		msgs := make([]byte, 0, len(tx.Body.Messages))

		for _, data := range tx.Body.Messages {
			bindata, _ := data.MarshalBinary()
			msgs = append(msgs, bindata...)
		}

		var amount int64

		if len(tx.AuthInfo.Fee.Amount) > 0 {
			amount, err = strconv.ParseInt(tx.AuthInfo.Fee.Amount[0].Amount, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("parse amount '%s': %w", tx.AuthInfo.Fee.Amount[0].Amount, err)
			}
		}

		signer, err := signerAddress(api.Tx(*tx))
		if err != nil {
			return nil, fmt.Errorf("signer addr: %w", err)
		}

		pk, _ := tx.AuthInfo.SignerInfos[0].PublicKey.MarshalBinary()
		rtx = append(rtx, Tx{
			BlockHash:        common.Hash(blockTx.BlockID.Hash),
			BlockNumber:      HexUint64(num),
			Creates:          signer,
			From:             signer,
			To:               convertAddress(tx.AuthInfo.Fee.Granter),
			Gas:              HexUint64(gas),          //nolint:gosec // always positive or 0
			GasPrice:         HexUint64(amount / gas), //nolint:gosec // always positive or 0
			Hash:             common.Hash(blockTx.Block.Data.Txs[i]),
			Input:            "0x" + hex.EncodeToString(msgs),
			TransactionIndex: HexUint64(i),      //nolint:gosec // always positive or 0
			Value:            HexUint64(amount), //nolint:gosec // always positive or 0
			Type:             2,
			ChainID:          1,
			V:                0xbd, //nolint:mnd // seems like there is no V, R and S in cosmos public keys, need investigate in details
			R:                "0xad3733df250c87556335ffe46c23e34dbaffde93097ef92f52c88632a40f0c75",
			S:                "0x72caddc0371451a58de2ca6ab64e0f586ccdb9465ff54e1c82564940e89291e3",
			StandardV:        0,
			PublickKey:       "0x" + hex.EncodeToString(pk),
			Raw:              "0x" + hex.EncodeToString(blockTx.Block.Data.Txs[i]),
		})
	}

	return rtx, nil
}

func ToEthTxReceipt(block *api.BlockWithTxs, tx *api.TxInfo) (*TxReceipt, error) {
	num, err := strconv.ParseUint(block.Block.Header.Height, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("height %s convert: %w", block.Block.Header.Height, err)
	}

	gas, err := strconv.ParseInt(tx.TxResponse.GasUsed, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("gas limit convert '%s': %w", tx.TxResponse.GasUsed, err)
	}

	status := 0
	if tx.TxResponse.Code == 0 {
		status = 1
	}

	txHash, err := hex.DecodeString(tx.TxResponse.Txhash)
	if err != nil {
		return nil, fmt.Errorf("decode tx hash: %w", err)
	}

	var txIndex uint64

	for i, t := range block.Block.Data.Txs {
		if slices.Equal(t, txHash) {
			txIndex = uint64(i) //nolint:gosec // counter

			break
		}
	}

	var amount int64

	if len(tx.Tx.AuthInfo.Fee.Amount) > 0 {
		amount, err = strconv.ParseInt(tx.Tx.AuthInfo.Fee.Amount[0].Amount, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("parse amount '%s': %w", tx.Tx.AuthInfo.Fee.Amount[0].Amount, err)
		}
	}

	signer, err := signerAddress(api.Tx(*tx.Tx))
	if err != nil {
		return nil, fmt.Errorf("signer addr: %w", err)
	}

	return &TxReceipt{
		BlockHash:         common.Hash(block.BlockID.Hash),
		BlockNumber:       HexUint64(num),
		ContractAddress:   signer,
		CumulativeGasUsed: HexUint64(gas),          //nolint:gosec // always positive or 0
		EffectiveGasPrice: HexUint64(amount / gas), //nolint:gosec // always positive or 0
		From:              signer,
		GasUsed:           HexUint64(gas), //nolint:gosec // always positive or 0
		Logs:              []TxLog{},
		Status:            HexUint64(status), //nolint:gosec // always positive or 0
		To:                convertAddress(tx.Tx.AuthInfo.Fee.Granter),
		TransactionHash:   common.Hash(txHash),
		TransactionIndex:  HexUint64(txIndex),
		Type:              2,
	}, nil
}

func convertAddress(addr string) common.Address {
	if addr != "" {
		_, result, err := bech32.DecodeAndConvert(addr)
		if err != nil {
			return common.Address{}
		}

		return common.Address(result)
	}

	return common.Address{}
}

func signerAddress(tx api.Tx) (common.Address, error) {
	publicKey, ok := tx.AuthInfo.SignerInfos[0].PublicKey.SignerInfoPublicKey["key"].(string)
	if !ok {
		return common.Address{}, ErrKeyNotFound
	}

	key, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return common.Address{}, fmt.Errorf("parse public key: %w", err)
	}

	return common.Address(sdk.AccAddress(key).Bytes()), nil
}
