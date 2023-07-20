package config

import (
	"os"

	"github.com/SidVermaS/Ethereum-Consensus/pkg/consts"
	"github.com/SidVermaS/Ethereum-Consensus/pkg/migrations"
	"github.com/SidVermaS/Ethereum-Consensus/pkg/types/structs"
	"github.com/SidVermaS/Ethereum-Consensus/pkg/utils"
	"github.com/SidVermaS/Ethereum-Consensus/pkg/vendors"
	"github.com/joho/godotenv"
)

var Repository *structs.DBRepository = &structs.DBRepository{}
var CronInstance *structs.Cron = &structs.Cron{}
var ConsensysVendor *vendors.Consensys = &vendors.Consensys{}

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
}
func UseVendorsService() {
	ConsensysVendor.AccessConsensysNode()
}
func InitializeAll() {
	// Load the .env file
	godotenv.Load(".env")

	// Initializing the Database
	InitializeDB()
	// It needs to be executed only for the first time
	migrations.InitialMigration(Repository.DB)

	// Initialize & Start Crons Scheduler
	utils.InitializeCron(CronInstance)

	//	Accesses vendors' services
	// UseVendorsService()
}
