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
var WaitGroup *sync.WaitGroup = &sync.WaitGroup{}

func UseServices() {
	ConsensysVendor = helpers.GetVendor(consts.Consensys)

	go StreamConsensysNode(ConsensysVendor, consensysconsts.AllConsensysTopics)
}
func InitializeAll() {
	// Load the .env file
	godotenv.Load(".env")

	WaitGroup.Add(2)
	// Initializing the Database
	go configs.InitializeDB(WaitGroup)
	// Connecting to the Cache
	go configs.CreateCacheConnection(WaitGroup)
	WaitGroup.Wait()

	// Initialize Vendor Configuration
	consts.InitializeVendorConfig()

	//	Accesses various services
	UseServices()

}
