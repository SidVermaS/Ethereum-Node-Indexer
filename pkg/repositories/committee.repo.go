package repositories

import (
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CommitteeRepo struct {
	Db *gorm.DB
}

// Inserts multiple committees in batches
func (committeeRepo *CommitteeRepo) CreateMany(committees []*models.Committee) error {
	for index, committeeItem := range committees {
		committees[index] = &models.Committee{EpochId: committeeItem.EpochId, StateId: committeeItem.StateId, BlockId: committeeItem.BlockId, SlotId: committeeItem.SlotId, ValidatorId: committeeItem.ValidatorId}
	}
	result := committeeRepo.Db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).CreateInBatches(committees, 100)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Inserts an individual committee
func (committeeRepo *CommitteeRepo) Create(committee *models.Committee) (uint, error) {
	committee = &models.Committee{EpochId: committee.EpochId, SlotId: committee.SlotId, ValidatorId: committee.ValidatorId, BlockId: committee.BlockId}
	result := committeeRepo.Db.Create(committee)
	if result.Error != nil {
		return 0, result.Error
	}
	return committee.ID, nil
}

// Fetches filtered paginated committee
func (committeeRepo *CommitteeRepo) FetchFilteredPaginatedData(filter *models.Committee, offset int, limit int) ([]*models.Committee, error) {
	var committees []*models.Committee
	result := committeeRepo.Db.Table("committees c").InnerJoins("Slot").InnerJoins("Block").InnerJoins("Validator").Select("c.id").Where(filter).Offset(offset).Limit(limit).Find(&committees)
	if result.Error != nil {
		return nil, result.Error
	}
	return committees, nil
}
