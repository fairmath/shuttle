package dto

import (
	"encoding/binary"
	"fmt"
	"math/big"
	"math/rand/v2"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/go-openapi/strfmt"

	"github.com/fairmath/shuttle/internal/client/tendermint/api"
)

type HexUint64 uint64

func (h HexUint64) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"0x%x"`, h)), nil
}

type Header struct { //nolint:tagliatelle // ethereum serialized json
	Hash            common.Hash         `json:"hash"`
	ParentHash      common.Hash         `json:"parentHash"`
	UncleHash       common.Hash         `json:"sha3Uncles"`
	Coinbase        common.Address      `json:"miner"`
	Root            common.Hash         `json:"stateRoot"`
	TxHash          common.Hash         `json:"transactionsRoot"`
	ReceiptHash     common.Hash         `json:"receiptsRoot"`
	Bloom           ethtypes.Bloom      `json:"logsBloom"`
	Difficulty      *big.Int            `json:"difficulty"`
	Number          *big.Int            `json:"number"`
	GasLimit        HexUint64           `json:"gasLimit"`
	GasUsed         HexUint64           `json:"gasUsed"`
	Time            HexUint64           `json:"timestamp"`
	Extra           []byte              `json:"extraData"`
	MixDigest       common.Hash         `json:"mixHash"`
	Nonce           ethtypes.BlockNonce `json:"nonce"`
	Size            HexUint64           `json:"size"`
	TotalDifficulty *big.Int            `json:"total_difficulty"`
	Transactions    []Tx                `json:"transactions,omitempty"`
	Uncles          []string            `json:"uncles"`
}

func FromCosmosBlock(block *api.Block) (Header, error) {
	num, err := strconv.ParseInt(block.Block.Header.Height, 10, 64)
	if err != nil {
		return Header{}, fmt.Errorf("height conversion '%s' -> int64: %w", block.Block.Header.Height, err)
	}

	blockBin, _ := block.Block.MarshalBinary()
	blockSz := len(blockBin)

	return Header{
		Hash:            common.Hash(block.BlockID.Hash),
		Number:          big.NewInt(num),
		Coinbase:        common.Address(block.Block.Header.ProposerAddress),
		ParentHash:      common.Hash(block.Block.Header.LastBlockID.Hash),
		Nonce:           ethtypes.BlockNonce(binary.LittleEndian.AppendUint64([]byte{}, rand.Uint64())),
		MixDigest:       common.Hash(block.Block.Header.ConsensusHash),
		ReceiptHash:     common.Hash(block.Block.Header.DataHash),
		UncleHash:       common.Hash(block.Block.Header.NextValidatorsHash),
		Root:            common.Hash(block.Block.Header.EvidenceHash),
		TxHash:          common.Hash{},
		Extra:           []byte{},
		Difficulty:      big.NewInt(int64(0x1046bb7e3f8)),
		TotalDifficulty: big.NewInt(int64(0x1046bb7e3f8)),
		Size:            HexUint64(blockSz),
		Transactions:    nil,
		Uncles:          []string{},

		Time: HexUint64(time.Time(block.Block.Header.Time).Unix()),
	}, nil
}

func FromCosmosBlockWithTxs(block *api.BlockWithTxs) (Header, error) {
	num, err := strconv.ParseInt(block.Block.Header.Height, 10, 64)
	if err != nil {
		return Header{}, fmt.Errorf("height conversion '%s' -> int64: %w", block.Block.Header.Height, err)
	}

	blockBin, _ := block.Block.MarshalBinary()
	blockSz := len(blockBin)

	txs, err := ToEthTxs(block)
	if err != nil {
		return Header{}, fmt.Errorf("tx convert: %w", err)
	}

	return Header{
		Hash:            common.Hash(block.BlockID.Hash),
		Number:          big.NewInt(num),
		Coinbase:        common.Address(block.Block.Header.ProposerAddress),
		ParentHash:      common.Hash(block.Block.Header.LastBlockID.Hash),
		Nonce:           ethtypes.BlockNonce(binary.LittleEndian.AppendUint64([]byte{}, rand.Uint64())),
		MixDigest:       common.Hash(block.Block.Header.ConsensusHash),
		ReceiptHash:     common.Hash(block.Block.Header.DataHash),
		UncleHash:       common.Hash(block.Block.Header.NextValidatorsHash),
		Root:            common.Hash(block.Block.Header.EvidenceHash),
		TxHash:          firstTxHash(block.Block.Data.Txs),
		Extra:           []byte{},
		Difficulty:      big.NewInt(int64(0x1046bb7e3f8)),
		TotalDifficulty: big.NewInt(int64(0x1046bb7e3f8)),
		Size:            HexUint64(blockSz),
		Transactions:    txs,
		Uncles:          []string{},

		Time: HexUint64(time.Time(block.Block.Header.Time).Unix()),
	}, nil
}

func firstTxHash(data []strfmt.Base64) common.Hash {
	if len(data) > 0 {
		return common.Hash(data[0])
	}

	return common.Hash{}
}
