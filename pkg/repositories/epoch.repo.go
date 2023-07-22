package repositories

import (
	"github.com/SidVermaS/Ethereum-Consensus/pkg/models"
	"gorm.io/gorm"
)

type EpochRepo struct {
	Db *gorm.DB
}

func (epochRepo *EpochRepo) CreateMany(epochs []*models.Epoch) error {
	for index, epochItem := range epochs {
		epochs[index] = &models.Epoch{Bid: epochItem.Bid,Epoch: epochItem.Epoch}
	}
	result := epochRepo.Db.Create(epochs)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (epochRepo *EpochRepo) Create(epoch *models.Epoch) (uint, error) {
	epoch =  &models.Epoch{Bid: epoch.Bid,Epoch: epoch.Epoch}
	result := epochRepo.Db.Create(epoch)
	if result.Error != nil {
		return 0, result.Error
	}
	return epoch.ID, nil
}

// func (epochRepo *EpochRepo) FetchAll()	{
// 	result := epochRepo.Db.Find()
// }