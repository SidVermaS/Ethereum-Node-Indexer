package repositories

import (
	"github.com/SidVermaS/Ethereum-Consensus/pkg/models"
	"gorm.io/gorm"
)

type CommitteeRepo struct {
	Db *gorm.DB
}

func (committeeRepo *CommitteeRepo) CreateMany(committees []*models.Committee) error {
	for index, committeeItem := range committees {
		committees[index] = &models.Committee{Eid: committeeItem.Eid,Sid: committeeItem.Sid,Vid: committeeItem.Vid}
	}
	result := committeeRepo.Db.Create(committees)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (committeeRepo *CommitteeRepo) Create(committee *models.Committee) (uint, error) {
	committee = &models.Committee{Eid: committee.Eid,Sid: committee.Sid,Vid: committee.Vid}
	result := committeeRepo.Db.Create(committee)
	if result.Error != nil {
		return 0, result.Error
	}
	return committee.ID, nil
}
