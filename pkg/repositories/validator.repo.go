package repositories

import (
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ValidatorRepo struct {
	Db *gorm.DB
}
// Inserts multiple validators in batches
func (validatorRepo *ValidatorRepo) CreateMany(validators []*models.Validator) error {
	for index, validatorItem := range validators {
		validators[index] = &models.Validator{Index: validatorItem.Index, PublicKey: validatorItem.PublicKey}
	}

	result := validatorRepo.Db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).CreateInBatches(validators, 100)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
// Inserts an individual validator
func (validatorRepo *ValidatorRepo) Create(validator *models.Validator) (uint, error) {
	validator = &models.Validator{Index: validator.Index, PublicKey: validator.PublicKey}
	result := validatorRepo.Db.Create(validator)
	if result.Error != nil {
		return 0, result.Error
	}
	return validator.ID, nil
}

// Fetches validators based on their indexes
func (validatorRepo *ValidatorRepo) FetchFromIndexes(indexes []uint64) ([]*models.Validator, error) {
	var validators []*models.Validator
	result := validatorRepo.Db.Where("index IN ?", indexes).Find(&validators)
	if result.Error != nil {
		return nil, result.Error
	}
	return validators, nil
}

// Fetches paginated validators
func (validatorRepo *ValidatorRepo) FetchPaginatedData(offset int, limit int) ([]*models.Validator, error) {
	var validators []*models.Validator

	result := validatorRepo.Db.Offset(offset).Limit(limit).Find(&validators)
	if result.Error != nil {
		return nil, result.Error
	}
	return validators, nil
}
