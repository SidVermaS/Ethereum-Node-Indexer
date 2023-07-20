package migrations

import (
	"github.com/SidVermaS/Ethereum-Consensus/pkg/models"
	"gorm.io/gorm"
)

// Migrates all the tables (It needs to be executed only for the first time)
func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&models.Block{})
}
