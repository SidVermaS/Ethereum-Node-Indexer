package modules

import (
	configs "github.com/SidVermaS/Ethereum-Node-Indexer/pkg/config"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/helpers"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/models"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/repositories"
)

// Fetches the filtered paginated states
func GetFilteredStates(filter map[string]string, offset int, limit int) ([]*models.State, error) {
	stateRepo := &repositories.StateRepo{
		Db: configs.GetDBInstance(),
	}

	stateFilter := &models.State{}

	epochIdQuery, exists := filter["epoch_id"]
	if exists {
		epochId, _ := helpers.ConvertStringToUInt(epochIdQuery)
		stateFilter.EpochId =  &epochId
	}
	blockIdQuery, exists := filter["block_id"]
	if exists {
		blockId, _ := helpers.ConvertStringToUInt(blockIdQuery)
		stateFilter.BlockId =  &blockId
	}
	states, err := stateRepo.FetchFilteredPaginatedData(stateFilter, offset, limit)

	if err != nil {
		return nil, err
	}
	return states, nil
}
