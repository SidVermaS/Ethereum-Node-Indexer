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
		slots[index] = &models.Slot{EpochId: slotItem.EpochId, StateId: slotItem.StateId, BlockId: slotItem.BlockId, Index: slotItem.Index}
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
	slot = &models.Slot{EpochId: slot.EpochId, BlockId: slot.BlockId, Index: slot.Index}
	result := slotRepo.Db.Create(slot)
	if result.Error != nil {
		return 0, result.Error
	}
	return slot.ID, nil
}

// Fetches slots based on an array of Epoch IDs
func (slotRepo *SlotRepo) FetchByEids(eids []uint) ([]*models.Slot, error) {
	var slots []*models.Slot
	result := slotRepo.Db.Where("epoch_id IN ?", eids).Find(&slots)
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

// Fetches filtered paginated slots
func (slotRepo *SlotRepo) FetchFilteredPaginatedData(filter *models.Slot, offset int, limit int) ([]*models.Slot, error) {
	var slots []*models.Slot

	result := slotRepo.Db.Table("slots sl").InnerJoins("Epoch").InnerJoins("Block").Select("sl.id","index").Where(filter).Offset(offset).Limit(limit).Find(&slots)
	if result.Error != nil {
		return nil, result.Error
	}
	return slots, nil
}
