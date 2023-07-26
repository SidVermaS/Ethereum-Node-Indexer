package repositories

import (
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SlotRepo struct {
	Db *gorm.DB
}

type SlotsErrorChannelStruct struct {
	Slots []*models.Slot
	Err   error
}
// Inserts multiple slots in batches
func (slotRepo *SlotRepo) CreateMany(slots []*models.Slot) error {
	for index, slotItem := range slots {
		slots[index] = &models.Slot{Eid: slotItem.Eid, StateId: slotItem.StateId, Index: slotItem.Index}
	}

	result := slotRepo.Db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).CreateInBatches(slots, 100)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
// Inserts an individual slot
func (slotRepo *SlotRepo) Create(slot *models.Slot) (uint, error) {
	slot = &models.Slot{Eid: slot.Eid, Index: slot.Index}
	result := slotRepo.Db.Create(slot)
	if result.Error != nil {
		return 0, result.Error
	}
	return slot.ID, nil
}

// Fetches slots based on an array of Epoch IDs
func (slotRepo *SlotRepo) FetchByEids(eids []uint) ([]*models.Slot, error) {
	var slots []*models.Slot
	result := slotRepo.Db.Where("eid IN ?", eids).Find(&slots)
	if result.Error != nil {
		return nil, result.Error
	}
	return slots, nil
}
// Fetches slots and communicates via. a channel for a goRoutine
func (slotRepo *SlotRepo) FetchByEidsFromChannel(eids []uint, slotsErrorChannel chan *SlotsErrorChannelStruct) {
	slots, err := slotRepo.FetchByEids(eids)
	slotsErrorChannel <- &SlotsErrorChannelStruct{
		Slots: slots,
		Err:   err,
	}
}

