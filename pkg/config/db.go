package config

import (
	"fmt"

	"github.com/SidVermaS/Ethereum-Consensus/pkg/types/structs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connecting to the Database
func CreateConnection(config *structs.DbConfig, repository *structs.DBRepository) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", config.Host, config.User, config.Password, config.DBName, config.Port, config.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	Repository.DB = db
	return db, nil
}
