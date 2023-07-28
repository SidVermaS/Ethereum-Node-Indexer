package modules

import (
	configs "github.com/SidVermaS/Ethereum-Node-Indexer/pkg/config"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/helpers"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/models"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/repositories"
)

// Fetches the filtered paginated slots
func GetFilteredSlots(filter map[string]string, offset int, limit int) ([]*models.Slot, error) {
	slotRepo := &repositories.SlotRepo{
		Db: configs.GetDBInstance(),
	}

	slotFilter := &models.Slot{}

	epochIdQuery, exists := filter["epoch_id"]
	if exists {
		epochId, _ := helpers.ConvertStringToUInt(epochIdQuery)
		slotFilter.EpochId =  &epochId
	}
	stateIdQuery, exists := filter["state_id"]
	if exists {
		stateId, _ := helpers.ConvertStringToUInt(stateIdQuery)
		slotFilter.StateId =  &stateId
	}
	slots, err := slotRepo.FetchFilteredPaginatedData(slotFilter, offset, limit)

	if err != nil {
		return nil, err
	}
	return slots, nil
}
