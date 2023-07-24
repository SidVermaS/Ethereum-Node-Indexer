package modules

import (
	"sync"

	configs "github.com/SidVermaS/Ethereum-Node-Indexer/pkg/config"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/consts"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/helpers"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/vendors/consensys"
	consensysconsts "github.com/SidVermaS/Ethereum-Node-Indexer/pkg/vendors/consensys/consts"
	"github.com/joho/godotenv"
)

var ConsensysVendor *consensys.Consensys
var Wg *sync.WaitGroup = &sync.WaitGroup{}

func UseServices() {
	ConsensysVendor = helpers.GetVendor(consts.Consensys)

	StreamConsensysNode(ConsensysVendor, consensysconsts.AllConsensysTopics)
}
func InitializeAll() {
	// Load the .env file
	godotenv.Load(".env")

	// Initializing the Database
	configs.InitializeDB()

	// Initialize Vendor Configuration
	consts.InitializeVendorConfig()

	//	Accesses various services
	go UseServices()

}
