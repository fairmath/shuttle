package api

import (
	blockModel "github.com/fairmath/shuttle/internal/client/tendermint/api/internal/blocks/models"
	txModel "github.com/fairmath/shuttle/internal/client/tendermint/api/internal/txs/models"
)

type (
	Block        blockModel.CosmosBlock
	BlockWithTxs txModel.BlockWithTxs
)
