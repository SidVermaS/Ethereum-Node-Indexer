package helpers

import (
	"log"
	"strconv"
	"sync"

	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/consts"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/models"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/modules"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/repositories"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/structs"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/vendors/consensys"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/vendors/consensys/consensysstructs"
	consensysconsts "github.com/SidVermaS/Ethereum-Node-Indexer/pkg/vendors/consensys/consts"
)

var waitGroup = &sync.WaitGroup{}

// create a channel to communicate between goroutines (SaveBlocks(), SaveEpochs())
var blocksEpochsChannel chan []*models.Block = make(chan []*models.Block)
var epochsStatesChannel chan []*models.Epoch = make(chan []*models.Epoch)
var statesAndValidatorsChannel chan []*models.State = make(chan []*models.State)

var validatorToValidatorsStatusChannel chan *structs.ValidatorToValidatorsStatusChannelStruct = make(chan *structs.ValidatorToValidatorsStatusChannelStruct)

func ProcessToSaveDataForIndexing(finalizedCheckpoints []*consensysstructs.FinalizedCheckpoint) {
	waitGroup.Add(5)
	go SaveBlocks(finalizedCheckpoints, waitGroup)
	go SaveEpochs(finalizedCheckpoints, waitGroup)
	go SaveStates(finalizedCheckpoints, waitGroup)
	go SaveValidators(waitGroup)
	go SaveValidatorsStatus(waitGroup)
	go SaveSlotsAndCommittees(waitGroup)
	waitGroup.Wait()
}

func SaveBlocks(finalizedCheckpoints []*consensysstructs.FinalizedCheckpoint, processWaitGroup *sync.WaitGroup) {
	defer processWaitGroup.Done()
	var blocks []*models.Block
	for _, finalizedCheckpointItem := range finalizedCheckpoints {
		blocks = append(blocks, &models.Block{Root: finalizedCheckpointItem.Block})
	}
	blockRepo := &repositories.BlockRepo{
		Db: GetDBInstance(),
	}
	err := blockRepo.CreateMany(blocks)
	blocksEpochsChannel <- blocks
	if err != nil {
		panic(err)
	}
}
func SaveEpochs(finalizedCheckpoints []*consensysstructs.FinalizedCheckpoint, processWaitGroup *sync.WaitGroup) {

	defer processWaitGroup.Done()
	blocks := <-blocksEpochsChannel

	var epochs []*models.Epoch
	for index, blockItem := range blocks {
		period, _ := strconv.Atoi(finalizedCheckpoints[index].Epoch)
		epochs = append(epochs, &models.Epoch{
			Period: uint(period),
			Bid:    uint(blockItem.ID),
		})
	}
	epochRepo := &repositories.EpochRepo{
		Db: GetDBInstance(),
	}

	err := epochRepo.CreateMany(epochs)
	if err != nil {
		log.Printf("~~~ SaveEpochs() err: %s", err.Error())
	}
	epochsStatesChannel <- epochs
}

func SaveStates(finalizedCheckpoints []*consensysstructs.FinalizedCheckpoint, processWaitGroup *sync.WaitGroup) {
	defer processWaitGroup.Done()
	epochs := <-epochsStatesChannel

	var states []*models.State
	for index, epochItem := range epochs {
		states = append(states, &models.State{
			StateStored: finalizedCheckpoints[index].State,
			Bid:         epochItem.Bid,
			Eid:         epochItem.ID,
		})
	}
	stateRepo := &repositories.StateRepo{
		Db: GetDBInstance(),
	}
	err := stateRepo.CreateMany(states)
	if err != nil {
		log.Printf("~~~ SaveStates err: %s \n", err.Error())
	}
	statesAndValidatorsChannel <- states
}
func SaveValidators(processWaitGroup *sync.WaitGroup) {
	defer processWaitGroup.Done()
	states := <-statesAndValidatorsChannel
	var consensysInstance *consensys.Consensys = &consensys.Consensys{
		Vendor: consts.VendorConfigMap[consts.Consensys],
	}
	getValidatorsFromStateWaitGroup := &sync.WaitGroup{}
	mux := &sync.Mutex{}
	var validatorsToBeCreated []*models.Validator
	var validatorStatusToBeCreated []*models.ValidatorStatus
	queueForValidatorToValidatorsStatusChannel := make(chan *structs.ValidatorToValidatorsStatusChannelStruct)
	for _, stateItem := range states {
		getValidatorsFromStateWaitGroup.Add(1)
		// StateStored can be passed for the results in that specifice state
		go modules.GetValidatorsFromState(stateItem.ID, stateItem.Eid, string(consensysconsts.Finalized), consensysInstance, queueForValidatorToValidatorsStatusChannel)
	}
	go func() {
		for queueForValidatorToValidatorsStatusChannelReceived := range queueForValidatorToValidatorsStatusChannel {
			mux.Lock()
			validatorsToBeCreated = append(validatorsToBeCreated, queueForValidatorToValidatorsStatusChannelReceived.Validators...)
			validatorStatusToBeCreated = append(validatorStatusToBeCreated, queueForValidatorToValidatorsStatusChannelReceived.ValidatorStatuses...)
			mux.Unlock()
			getValidatorsFromStateWaitGroup.Done()
		}
	}()
	getValidatorsFromStateWaitGroup.Wait()
	validatorRepo := &repositories.ValidatorRepo{
		Db: GetDBInstance(),
	}

	err := validatorRepo.CreateMany(validatorsToBeCreated)
	if err != nil {
		log.Printf("~~~ SaveValidators err: %s \n", err.Error())
		// panic(err)
	}
	validatorToValidatorsStatusChannel <- &structs.ValidatorToValidatorsStatusChannelStruct{ValidatorStatuses: validatorStatusToBeCreated, Validators: validatorsToBeCreated}
}

func SaveValidatorsStatus(processWaitGroup *sync.WaitGroup) {
	defer processWaitGroup.Done()
	validatorToValidatorsStatusChannelData := <-validatorToValidatorsStatusChannel
	var validatorStatuses []*models.ValidatorStatus
	var indexes []uint64
	for _, validatorsItem := range validatorToValidatorsStatusChannelData.Validators {
		indexes = append(indexes, validatorsItem.Index)
	}
	validatorRepo := &repositories.ValidatorRepo{
		Db: GetDBInstance(),
	}
	validators, _ := validatorRepo.FetchFromIndexes(indexes)

	 indexValidatorsMap:= map[uint64]uint{}

	for _, validatorsItem := range validators {
		indexValidatorsMap[validatorsItem.Index] = validatorsItem.ID
	}
	for iteratedIndex, validatorStatusItem := range validatorToValidatorsStatusChannelData.ValidatorStatuses {
		indexValidatorsMapValue, exists := indexValidatorsMap[validatorToValidatorsStatusChannelData.Validators[iteratedIndex].Index]
		if exists {
			validatorStatuses = append(validatorStatuses, &models.ValidatorStatus{
				StateId:   validatorStatusItem.StateId,
				Eid:       validatorStatusItem.Eid,
				IsSlashed: validatorStatusItem.IsSlashed,
				Status:    validatorStatusItem.Status,
				Vid:       indexValidatorsMapValue,
			})
		}
	}
	validatorStatusRepo := &repositories.ValidatorStatusRepo{
		Db: GetDBInstance(),
	}
	err := validatorStatusRepo.CreateMany(validatorStatuses)
	if err != nil {
		log.Printf("~~~ SaveValidatorsStatus err: %s \n", err.Error())
		// panic(err)
	}
}

func SaveSlotsAndCommittees(processWaitGroup *sync.WaitGroup) {}
