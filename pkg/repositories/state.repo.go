package repositories

import (
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type StateRepo struct {
	Db *gorm.DB
}

func (stateRepo *StateRepo) CreateMany(states []*models.State) error {
	for index, stateItem := range states {
		states[index] = &models.State{Eid: stateItem.Eid, Bid: stateItem.Bid, StateStored: stateItem.StateStored}
	}
	result := stateRepo.Db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).CreateInBatches(states, 100)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (stateRepo *StateRepo) Create(state *models.State) (uint, error) {
	state = &models.State{Eid: state.Eid, Bid: state.Bid, StateStored: state.StateStored}
	result := stateRepo.Db.Create(state)
	if result.Error != nil {
		return 0, result.Error
	}
	return state.ID, nil
}
