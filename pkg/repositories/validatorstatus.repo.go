package repositories

import (
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ValidatorStatusRepo struct {
	Db *gorm.DB
}
// Inserts multiple validators' statuses in a particular state in batches
func (validatorStatusRepo *ValidatorStatusRepo) CreateMany(validatorStatuss []*models.ValidatorStatus) error {
	for index, validatorStatusItem := range validatorStatuss {
		validatorStatuss[index] = &models.ValidatorStatus{Vid: validatorStatusItem.Vid, StateId: validatorStatusItem.StateId, Eid: validatorStatusItem.Eid, Status: validatorStatusItem.Status, IsSlashed: validatorStatusItem.IsSlashed}
	}

	result := validatorStatusRepo.Db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).CreateInBatches(validatorStatuss, 100)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
// Inserts an individual validator's status
func (validatorStatusRepo *ValidatorStatusRepo) Create(validatorStatus *models.ValidatorStatus) (uint, error) {
	validatorStatus = &models.ValidatorStatus{Vid: validatorStatus.Vid, StateId: validatorStatus.StateId, Eid: validatorStatus.Eid, Status: validatorStatus.Status, IsSlashed: validatorStatus.IsSlashed}
	result := validatorStatusRepo.Db.Create(validatorStatus)
	if result.Error != nil {
		return 0, result.Error
	}
	return validatorStatus.ID, nil
}
// Fetches multiple validators' statuses from the epoch's ID and the state's ID
func (validatorStatusRepo *ValidatorStatusRepo) FetchAllValidatorsStatusByEidsAndSlotsIds(eids []uint, slotsIds []uint) ([]*models.ValidatorStatus, error) {
	var validatorStatuses []*models.ValidatorStatus
	result := validatorStatusRepo.Db.Select("slots.eid, slots.state_id, validator_statuses.eid, validator_statuses.state_id, validator_statuses.is_slashed, validator_statuses.status").Joins("INNER JOIN slots ON validator_statuses.eid=slots.eid AND validator_statuses.state_id=slots.state_id").Where("validator_statuses.eid IN (?) AND slots.id IN (?)", eids, slotsIds).Find(&validatorStatuses)

	if result.Error != nil {
		return nil, result.Error
	}

	return validatorStatuses, nil
}

// Fetches specific validators' statuses from the epoch's ID and the state's ID
func (validatorStatusRepo *ValidatorStatusRepo) FetchSingleValidatorsStatusByEidsAndSlotsIds(validatorId uint,eids []uint, slotsIds []uint) ([]*models.ValidatorStatus, error) {
	var validatorStatuses []*models.ValidatorStatus
	result := validatorStatusRepo.Db.Select("slots.eid, slots.state_id, validator_statuses.eid, validator_statuses.state_id, validator_statuses.is_slashed, validator_statuses.status").Joins("INNER JOIN slots ON validator_statuses.eid=slots.eid AND validator_statuses.state_id=slots.state_id").Where("validator_statuses.eid IN (?) AND slots.id IN (?) AND validator_statuses.vid IN (?)", eids, slotsIds, validatorId).Find(&validatorStatuses)

	if result.Error != nil {
		return nil, result.Error
	}

	return validatorStatuses, nil
}