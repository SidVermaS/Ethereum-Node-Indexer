package helpers

import (
	"os"
	"sync"

	"github.com/SidVermaS/Ethereum-Consensus/pkg/consts"
	"github.com/SidVermaS/Ethereum-Consensus/pkg/migrations"
	"github.com/SidVermaS/Ethereum-Consensus/pkg/structs"
	"github.com/SidVermaS/Ethereum-Consensus/pkg/vendors/consensys"

	"github.com/joho/godotenv"
)

var Repository *structs.DBRepository = &structs.DBRepository{}
var ConsensysVendor *consensys.Consensys
var Wg *sync.WaitGroup=&sync.WaitGroup{}

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

func UseServices() {
	ConsensysVendor = GetVendor(consts.Consensys)

	// go StreamConsensysNode(ConsensysVendor, consensysconsts.AllConsensysTopics)
}
func InitializeAll() {
	// Load the .env file
	godotenv.Load(".env")

	// Initializing the Database
	InitializeDB()
	// It needs to be executed only for the first time
	migrations.InitialMigration(Repository.DB)

	// Initialize Vendor Configuration
	consts.InitializeVendorConfig()

	//	Accesses various services
	UseServices()
	
}
