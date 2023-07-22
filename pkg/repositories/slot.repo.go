package repositories

import (
	"github.com/SidVermaS/Ethereum-Consensus/pkg/models"
	"gorm.io/gorm"
)

type SlotRepo struct {
	Db *gorm.DB
}

func (slotRepo *SlotRepo) CreateMany(slots []*models.Slot) error {
	for index, slotItem := range slots {
		slots[index] = &models.Slot{Eid: slotItem.Eid, Index: slotItem.Index}
	}
	result := slotRepo.Db.Create(slots)
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
