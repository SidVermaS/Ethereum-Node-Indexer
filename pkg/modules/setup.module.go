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
// Vendor instance needed for accessing the third party APIs
var ConsensysVendor *consensys.Consensys

func UseServices() {
	ConsensysVendor = helpers.GetVendor(consts.Consensys)
	// Start Crons job
	go StartCronSchedulers()
	// Start an event listener to listen to the incoming epochs, slots, states and block.
	go StreamConsensysNode(ConsensysVendor, consensysconsts.AllConsensysTopics)
}
func ActivateAll() {
	// Load the .env file
	godotenv.Load(".env")

	// WaitGroup is needed to wait for the GoRoutines to be executed completely.
	var waitGroup *sync.WaitGroup = &sync.WaitGroup{}
	// We passed 2 as an argument because there were 2 goroutines which were executed
	waitGroup.Add(2)
	// Initializing the Database
	go configs.InitializeDB(waitGroup)
	// Connecting to the Cache
	go configs.CreateCacheConnection(waitGroup)
	// It will wait on this line until all the GoRoutines are executed completely
	waitGroup.Wait()

	// Initialize Vendor Configuration
	consts.InitializeVendorConfig()

	//	Accesses various services
	UseServices()
}

// Closes the connections to the database, cache and stops the cron schedulers
func DeactivateAll() {
	var waitGroup *sync.WaitGroup = &sync.WaitGroup{}
	waitGroup.Add(3)
	go configs.CloseDBConnection(waitGroup)
	go configs.CloseCacheConnection(waitGroup)
	go StopCronSchedulers(waitGroup)
	waitGroup.Wait()
}
