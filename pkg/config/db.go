package configs

import (
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/consts"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/migrations"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/structs"
)

var Repository *structs.DBRepository = &structs.DBRepository{}

func GetDBInstance() *gorm.DB {
	return Repository.DB
}

// The gorm DB Connection is closed
func CloseDBConnection(waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	db, _ := Repository.DB.DB()
	db.Close()
}

// Connecting to the Database
func CreateDBConnection(config *structs.DbConfig, repository *structs.DBRepository) (*gorm.DB, error) {
	// Data source name to connect to the database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", config.Host, config.User, config.Password, config.DBName, config.Port, config.SSLMode)

	// Connecting to the DB via gorm
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return db, err
	}
	// Initialize the DB instance of the repository
	repository.DB = db
	return db, nil
}

func InitializeDB(wg *sync.WaitGroup) {
	// WaitGroup is deferred and Done() is executed, once all the statements are executed
	defer wg.Done()
	// Configuration for the database
	dbConfig := &structs.DbConfig{
		Host:     os.Getenv(string(consts.POSTGRES_HOST)),
		User:     os.Getenv(string(consts.POSTGRES_USER)),
		Password: os.Getenv(string(consts.POSTGRES_PASSWORD)),
		DBName:   os.Getenv(string(consts.POSTGRES_DB)),
		Port:     os.Getenv(string(consts.POSTGRES_PORT)),
		SSLMode:  os.Getenv(string(consts.POSTGRES_SSL_MODE)),
	}
	// Passed the configuration and the DBRepository to initialize the gorm.DB instance
	CreateDBConnection(dbConfig, Repository)

	// It needs to be executed only for the first time
	migrations.InitialMigration(Repository.DB)

}


