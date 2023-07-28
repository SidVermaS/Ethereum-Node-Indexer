package modules

import (
	configs "github.com/SidVermaS/Ethereum-Node-Indexer/pkg/config"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/helpers"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/models"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/repositories"
)

// Fetches the paginated validator statuses
func GetValidatorStatuses(filter map[string]string, offset int, limit int) ([]*models.ValidatorStatus, error) {
	validatorStatusesRepo := &repositories.ValidatorStatusRepo{
		Db: configs.GetDBInstance(),
	}
	var validatorFilter *models.ValidatorStatus = &models.ValidatorStatus{}

	epochIdQuery, exists := filter["epoch_id"]
	if exists {
		epochId, _ := helpers.ConvertStringToUInt(epochIdQuery)
		validatorFilter.EpochId =  &epochId
	}
	stateIdQuery, exists := filter["state_id"]
	if exists {
		stateId, _ := helpers.ConvertStringToUInt(stateIdQuery)
		validatorFilter.StateId =  &stateId
	}
	validatorIdQuery, exists := filter["validator_id"]
	if exists {
		validatorId, _ := helpers.ConvertStringToUInt(validatorIdQuery)
		validatorFilter.ValidatorId =  &validatorId
	}

	validatorStatuses, err := validatorStatusesRepo.FetchFilteredPaginatedData(validatorFilter, offset, limit)

	if err != nil {
		return nil, err
	}
	return validatorStatuses, nil
}
