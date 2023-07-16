package config

import (
	"fmt"

	"github.com/SidVermaS/Ethereum-Consensus-Layer/pkg/types/structs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBInstance struct {
	DB *gorm.DB
}

var Repository DBInstance = DBInstance{}

// Connecting to the Database
func CreateConnection(config *structs.DbConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", config.Host, config.User, config.Password, config.DBName, config.Port, config.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	Repository.DB = db
	return db, nil
}
