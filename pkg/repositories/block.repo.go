package repositories

import (
	"github.com/SidVermaS/Ethereum-Consensus/pkg/models"
	"gorm.io/gorm"
)

type BlockRepo struct {
	Db *gorm.DB
}

func (blockRepo *BlockRepo) CreateMany(blocks []*models.Block) error {
	for index, blockItem := range blocks {
		blocks[index] = &models.Block{Root: blockItem.Root}
	}
	result := blockRepo.Db.Create(blocks)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (blockRepo *BlockRepo) Create(block *models.Block) (uint, error) {
	block = &models.Block{Root: block.Root}
	result := blockRepo.Db.Create(block)
	if result.Error != nil {
		return 0, result.Error
	}
	return block.ID, nil
}

 