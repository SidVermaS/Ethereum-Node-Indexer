package modules

import (
	configs "github.com/SidVermaS/Ethereum-Node-Indexer/pkg/config"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/helpers"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/models"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/repositories"
)

// Fetches the lastest epoch entry
func GetLastEpoch() (*models.Epoch, error) {
	epochRepo := &repositories.EpochRepo{
		Db: configs.GetDBInstance(),
	}
	epochs, err := epochRepo.FetchWithLimit(1)
	if err != nil {
		return nil, err
	}
	if len(epochs) <= 0 {
		return nil, nil
	}
	return epochs[0], nil
}

// Fetches the filterd paginated epochs
func GetEpochs(filter map[string]string, offset int, limit int) ([]*models.Epoch, error) {
	epochRepo := &repositories.EpochRepo{
		Db: configs.GetDBInstance(),
	}
	epochFilter := &models.Epoch{}

	blockIdQuery, exists := filter["block_id"]
	if exists {
		blockId, _ := helpers.ConvertStringToUInt(blockIdQuery)
		epochFilter.BlockId =  &blockId
	}
	epochs, err := epochRepo.FetchPaginatedData(epochFilter, offset, limit)

	if err != nil {
		return nil, err
	}
	return epochs, nil
}
