package repositories

import (
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type EpochRepo struct {
	Db *gorm.DB
}

func (epochRepo *EpochRepo) CreateMany(epochs []*models.Epoch) error {
	for index, epochItem := range epochs {
		epochs[index] = &models.Epoch{Bid: epochItem.Bid, Period: epochItem.Period}
	}
	result := epochRepo.Db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).CreateInBatches(epochs, 100)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (epochRepo *EpochRepo) Create(epoch *models.Epoch) (uint, error) {
	epoch = &models.Epoch{Bid: epoch.Bid, Period: epoch.Period}
	result := epochRepo.Db.Create(epoch)
	if result.Error != nil {
		return 0, result.Error
	}
	return epoch.ID, nil
}

func (epochRepo *EpochRepo) FetchFromIDs(ids []uint) ([]*models.Epoch, error) {
	var epochs []*models.Epoch
	result := epochRepo.Db.Where("id IN ?", ids).Find(&epochs)
	if result.Error != nil {
		return nil, result.Error
	}
	return epochs, nil
}
