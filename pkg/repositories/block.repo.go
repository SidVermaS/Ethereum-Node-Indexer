package repositories

import (
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BlockRepo struct {
	Db *gorm.DB
}

// Inserts multiple blocks in batches
func (blockRepo *BlockRepo) CreateMany(blocks []*models.Block) error {
	for index, blockItem := range blocks {
		blocks[index] = &models.Block{Root: blockItem.Root}
	}
	result := blockRepo.Db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).CreateInBatches(blocks, 100)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
// Inserts an individual block
func (blockRepo *BlockRepo) Create(block *models.Block) (uint, error) {
	block = &models.Block{Root: block.Root}
	result := blockRepo.Db.Create(block)
	if result.Error != nil {
		return 0, result.Error
	}
	return block.ID, nil
}


// Fetches paginated blocks
func (blockRepo *BlockRepo) FetchPaginatedData(offset int, limit int) ([]*models.Block, error) {
	var blocks []*models.Block

	result := blockRepo.Db.Offset(offset).Limit(limit).Find(&blocks)
	if result.Error != nil {
		return nil, result.Error
	}
	return blocks, nil
}