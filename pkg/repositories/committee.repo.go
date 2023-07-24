package repositories

import (
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CommitteeRepo struct {
	Db *gorm.DB
}

func (committeeRepo *CommitteeRepo) CreateMany(committees []*models.Committee) error {
	for index, committeeItem := range committees {
		committees[index] = &models.Committee{Eid: committeeItem.Eid,StateId: committeeItem.StateId, SlotId:  committeeItem.SlotId, Vid: committeeItem.Vid}
	}
	result := committeeRepo.Db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).CreateInBatches(committees, 100)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (committeeRepo *CommitteeRepo) Create(committee *models.Committee) (uint, error) {
	committee = &models.Committee{Eid: committee.Eid, SlotId: committee.SlotId, Vid: committee.Vid}
	result := committeeRepo.Db.Create(committee)
	if result.Error != nil {
		return 0, result.Error
	}
	return committee.ID, nil
}
