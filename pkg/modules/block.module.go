package modules

import (
	configs "github.com/SidVermaS/Ethereum-Node-Indexer/pkg/config"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/models"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/repositories"
)

// Fetches the paginated blocks
func GetBlocks(offset int, limit int) ([]*models.Block, error) {
	blockRepo := &repositories.BlockRepo{
		Db: configs.GetDBInstance(),
	}
	blocks, err := blockRepo.FetchPaginatedData(offset, limit)

	if err != nil {
		return nil, err
	}
	return blocks, nil
}
