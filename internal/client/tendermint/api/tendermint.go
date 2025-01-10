package api

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cometbft/cometbft/rpc/client/http"
	"github.com/go-openapi/strfmt"

	"github.com/fairmath/shuttle/internal/client/tendermint/api/internal/blocks/models"
)

type TendermintRPC struct {
	client *http.HTTP
}

func NewTendermintRPC(rpcServerURL, wsServerURL string) (*TendermintRPC, error) {
	client, err := http.New(rpcServerURL, wsServerURL)
	if err != nil {
		return nil, fmt.Errorf("connect to tendermint rpc: %w", err)
	}

	return &TendermintRPC{
		client: client, // todo: use tendermint ws connection instead of WebsocketPool
	}, nil
}

func (t *TendermintRPC) BlockByHash(ctx context.Context, hash []byte) (Block, error) {
	block, err := t.client.BlockByHash(ctx, hash)
	if err != nil {
		return Block{}, fmt.Errorf("block by hash: %w", err)
	}

	return Block{
		BlockID: &models.CosmosBlockBlockID{
			Hash: block.BlockID.Hash.Bytes(),
		},
		Block: &models.CosmosBlockBlock{
			Header: &models.CosmosBlockBlockHeader{
				AppHash:       strfmt.Base64(block.Block.AppHash),
				ChainID:       block.Block.ChainID,
				ConsensusHash: block.Block.ConsensusHash.Bytes(),
				DataHash:      block.Block.DataHash.Bytes(),
				EvidenceHash:  block.Block.Evidence.Evidence.Hash(),
				Height:        strconv.FormatInt(block.Block.Height, 10),
				LastBlockID: &models.CosmosBlockBlockHeaderLastBlockID{
					Hash: strfmt.Base64(block.Block.LastCommit.BlockID.Hash),
				},
				LastCommitHash:     block.Block.LastCommitHash.Bytes(),
				LastResultsHash:    strfmt.Base64(block.Block.LastResultsHash),
				NextValidatorsHash: block.Block.NextValidatorsHash.Bytes(),
				ProposerAddress:    block.Block.ProposerAddress.Bytes(),
				Time:               strfmt.DateTime(block.Block.Time),
				ValidatorsHash:     block.Block.ValidatorsHash.Bytes(),
			},
		},
	}, nil
}
