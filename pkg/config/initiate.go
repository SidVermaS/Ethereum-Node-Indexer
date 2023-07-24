package configs

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/consts"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/migrations"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/structs"
)

var Repository *structs.DBRepository = &structs.DBRepository{}

func GetDBInstance() *gorm.DB {
	return Repository.DB
}

// The gorm DB Connection is closed
func CloseDBConnection() {
	db, _ := Repository.DB.DB()
	db.Close()
}

// Connecting to the Database
func CreateConnection(config *structs.DbConfig, repository *structs.DBRepository) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", config.Host, config.User, config.Password, config.DBName, config.Port, config.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Silent)
	})
	if err != nil {
		return db, err
	}
	repository.DB = db
	return db, nil
}

func InitializeDB() {
	dbConfig := &structs.DbConfig{
		Host:     os.Getenv(string(consts.POSTGRES_HOST)),
		User:     os.Getenv(string(consts.POSTGRES_USER)),
		Password: os.Getenv(string(consts.POSTGRES_PASSWORD)),
		DBName:   os.Getenv(string(consts.POSTGRES_DB)),
		Port:     os.Getenv(string(consts.POSTGRES_PORT)),
		SSLMode:  os.Getenv(string(consts.POSTGRES_SSL_MODE)),
	}
	// Passed the configuration and the DBRepository to initialize the gorm.DB instance
	CreateConnection(dbConfig, Repository)

	// It needs to be executed only for the first time
	migrations.InitialMigration(Repository.DB)
}
