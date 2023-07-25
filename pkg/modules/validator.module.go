package modules

import (
	configs "github.com/SidVermaS/Ethereum-Node-Indexer/pkg/config"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/models"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/repositories"
)

func GetValidators(offset int, limit int) ([]*models.Validator, error) {
	validatorRepo := &repositories.ValidatorRepo{
		Db: configs.GetDBInstance(),
	}
	validators, err := validatorRepo.FetchPaginatedData(offset, limit)
	if err != nil {
		return nil, err
	}
	return validators, nil
}
