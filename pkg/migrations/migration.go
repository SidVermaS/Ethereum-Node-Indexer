package migrations

import (
	"log"

	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/models"
	"gorm.io/gorm"
)

// Migrates all the tables (It needs to be executed only for the first time)
func InitialMigration(db *gorm.DB) {
	err:=db.AutoMigrate(&models.Block{}, &models.Epoch{},&models.State{}, &models.Validator{},&models.ValidatorStatus{}, &models.Slot{}, &models.Committee{},)
	if err != nil {
		log.Printf("~~~ InitialMigration() err %s",err)
	}
}
