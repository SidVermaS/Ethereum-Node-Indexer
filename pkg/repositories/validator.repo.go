package repositories

import (
	"github.com/SidVermaS/Ethereum-Consensus/pkg/models"
	"gorm.io/gorm"
)

type ValidatorRepo struct {
	Db *gorm.DB
}

func (validatorRepo *ValidatorRepo) CreateMany(validators []*models.Validator) error {
	for index, validatorItem := range validators {
		validators[index] = &models.Validator{PublicKey: validatorItem.PublicKey, Status: validatorItem.Status, IsSlashed: validatorItem.IsSlashed}
	}
	result := validatorRepo.Db.Create(validators)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (validatorRepo *ValidatorRepo) Create(validator *models.Validator) (uint, error) {
	validator = &models.Validator{PublicKey: validator.PublicKey, Status: validator.Status, IsSlashed: validator.IsSlashed}
	result := validatorRepo.Db.Create(validator)
	if result.Error != nil {
		return 0, result.Error
	}
	return validator.ID, nil
}
