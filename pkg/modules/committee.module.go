package modules

import (
	configs "github.com/SidVermaS/Ethereum-Node-Indexer/pkg/config"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/helpers"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/models"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/repositories"
)

// Fetches the paginated validator statuses
func GetCommitteees(filter map[string]string, offset int, limit int) ([]*models.Committee, error) {
	committeesRepo := &repositories.CommitteeRepo{
		Db: configs.GetDBInstance(),
	}
	var validatorFilter *models.Committee = &models.Committee{}

	epochIdQuery, exists := filter["epoch_id"]
	if exists {
		epochId, _ := helpers.ConvertStringToUInt(epochIdQuery)
		validatorFilter.EpochId = &epochId
	}

	slotIdQuery, exists := filter["slot_id"]
	if exists {
		slotId, _ := helpers.ConvertStringToUInt(slotIdQuery)
		validatorFilter.SlotId = &slotId
	}
	stateIdQuery, exists := filter["state_id"]
	if exists {
		stateId, _ := helpers.ConvertStringToUInt(stateIdQuery)
		validatorFilter.StateId = &stateId
	}
	validatorIdQuery, exists := filter["validator_id"]
	if exists {
		validatorId, _ := helpers.ConvertStringToUInt(validatorIdQuery)
		validatorFilter.ValidatorId = &validatorId
	}

	committees, err := committeesRepo.FetchFilteredPaginatedData(validatorFilter, offset, limit)

	if err != nil {
		return nil, err
	}
	return committees, nil
}
