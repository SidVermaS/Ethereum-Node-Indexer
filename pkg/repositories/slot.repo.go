package repositories

import (
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SlotRepo struct {
	Db *gorm.DB
}

func (slotRepo *SlotRepo) CreateMany(slots []*models.Slot) error {
	for index, slotItem := range slots {
		slots[index] = &models.Slot{Eid: slotItem.Eid,StateId: slotItem.StateId, Index: slotItem.Index, }
	}

	result := slotRepo.Db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).CreateInBatches(slots, 100)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (slotRepo *SlotRepo) Create(slot *models.Slot) (uint, error) {
	slot = &models.Slot{Eid: slot.Eid, Index: slot.Index}
	result := slotRepo.Db.Create(slot)
	if result.Error != nil {
		return 0, result.Error
	}
	return slot.ID, nil
}
