package repositories

import (
	"log"

	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type EpochRepo struct {
	Db *gorm.DB
}
// Inserts multiple epochs in batches
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
// Inserts an individual epoch
func (epochRepo *EpochRepo) Create(epoch *models.Epoch) (uint, error) {
	epoch = &models.Epoch{Bid: epoch.Bid, Period: epoch.Period}
	result := epochRepo.Db.Create(epoch)
	if result.Error != nil {
		return 0, result.Error
	}
	return epoch.ID, nil
}
// Fetches epochs based on an array of IDs
func (epochRepo *EpochRepo) FetchByIDs(ids []uint) ([]*models.Epoch, error) {
	var epochs []*models.Epoch
	result := epochRepo.Db.Where("id IN ?", ids).Find(&epochs)
	if result.Error != nil {
		return nil, result.Error
	}
	return epochs, nil
}
// Fetches the latest epochs according to the limit in the parameter
func (epochRepo *EpochRepo) FetchWithLimit(limit int) ([]*models.Epoch, error) {
	var epochs []*models.Epoch
	result := epochRepo.Db.Order("id desc").Limit(5).Find(&epochs)
	if result.Error != nil {
		log.Printf("~~~ FetchWithLimit() %s",result.Error)
		return nil, result.Error
	}
	return epochs, nil
}
