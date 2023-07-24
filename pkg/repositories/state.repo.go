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

func (stateRepo *StateRepo) FetchByID(id uint, limit int) (*models.State, error) {
	var state *models.State
	result := stateRepo.Db.Where("id = ?", id).Last(state).Limit(limit)
	if result.Error != nil {
		return nil, result.Error
	}
	return state, nil
}
func (stateRepo *StateRepo) FetchByIDs(ids []uint) ([]*models.State, error) {
	var states []*models.State
	result := stateRepo.Db.Where("id IN ?", ids).Find(&states)
	if result.Error != nil {
		return nil, result.Error
	}
	return states, nil
}
func (stateRepo *StateRepo) FetchWithLimit(limit int) ([]*models.State, error) {
	var states []*models.State
	result := stateRepo.Db.Last(states).Limit(limit)
	if result.Error != nil {
		return nil, result.Error
	}
	return states, nil
}

func (stateRepo *StateRepo) FetchStatesAndEpochs(epochsIDs []uint, limit int) ([]*models.State, error) {
	var states []*models.State
	result := stateRepo.Db.InnerJoins("Epoch",).Where("eid in ?", epochsIDs).Limit(limit).Find(&states)
	if result.Error != nil {
		return nil, result.Error
	}
	return states, nil
}
