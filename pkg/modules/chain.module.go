package modules

import (
	"fmt"
	"log"
	"strconv"
	"sync"

	configs "github.com/SidVermaS/Ethereum-Node-Indexer/pkg/config"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/consts"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/helpers"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/models"
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
var validatorsStatusAndSaveSlotsAndCommitteesChannel chan bool = make(chan bool)

func GetValidatorsFromState(stateDbId uint, epochDbId uint, stateIdentifierOrHex string, consensysInstance *consensys.Consensys, queueForValidatorToValidatorsStatusChannel chan *structs.ValidatorToValidatorsStatusChannelStruct) {
	var validators []*models.Validator
	var validatorStatuses []*models.ValidatorStatus
	var getValidatorsFromStateResponse *consensysstructs.GetValidatorsFromStateResponse = consensysInstance.GetValidatorsFromState(stateIdentifierOrHex)

	for _, validatorItem := range getValidatorsFromStateResponse.Data {
		validatorItemIndex, _ := strconv.ParseUint(validatorItem.Index, 10, 64)
		validators = append(validators, &models.Validator{PublicKey: validatorItem.Validator.Pubkey, Index: validatorItemIndex})
		validatorStatuses = append(validatorStatuses, &models.ValidatorStatus{IsSlashed: validatorItem.Validator.Slashed, Status: validatorItem.Status, Eid: epochDbId, StateId: stateDbId})
	}
	queueForValidatorToValidatorsStatusChannel <- &structs.ValidatorToValidatorsStatusChannelStruct{
		Validators:        validators,
		ValidatorStatuses: validatorStatuses,
	}
	return
}

func GetCommittiesFromStateAndEpoch(stateId uint,  stateIdentifierOrHex string, epochId uint,epochPeriod uint, consensysInstance *consensys.Consensys, queueForCommittiesFromStateAndEpochDataChannel chan *structs.CommittiesFromStateAndEpochData) {

	var getCommitteesAtStateResponse *consensysstructs.GetCommitteesAtStateResponse = consensysInstance.GetCommitteesAtState(stateIdentifierOrHex, epochPeriod)

	queueForCommittiesFromStateAndEpochDataChannel <- &structs.CommittiesFromStateAndEpochData{Eid: epochId, StateId: stateId, SlotData: getCommitteesAtStateResponse.Data}
	return
}
func ProcessToSaveDataForIndexing(finalizedCheckpoints []*consensysstructs.FinalizedCheckpoint) {
	waitGroup.Add(6)
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
		Db: configs.GetDBInstance(),
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

	defer close(blocksEpochsChannel)
	var epochs []*models.Epoch
	for index, blockItem := range blocks {
		period, _ := strconv.Atoi(finalizedCheckpoints[index].Epoch)
		epochs = append(epochs, &models.Epoch{
			Period: uint(period),
			Bid:    uint(blockItem.ID),
		})
	}
	epochRepo := &repositories.EpochRepo{
		Db: configs.GetDBInstance(),
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
	defer close(epochsStatesChannel)
	var states []*models.State
	for index, epochItem := range epochs {
		states = append(states, &models.State{
			StateStored: finalizedCheckpoints[index].State,
			Bid:         epochItem.Bid,
			Eid:         epochItem.ID,
		})
	}
	stateRepo := &repositories.StateRepo{
		Db: configs.GetDBInstance(),
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
	defer close(statesAndValidatorsChannel)
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
		go GetValidatorsFromState(stateItem.ID, stateItem.Eid, string(consensysconsts.Finalized), consensysInstance, queueForValidatorToValidatorsStatusChannel)
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
		Db: configs.GetDBInstance(),
	}

	err := validatorRepo.CreateMany(validatorsToBeCreated)
	if err != nil {
		log.Printf("~~~ SaveValidators err: %s \n", err.Error())
		// panic(err)
	}
	defer close(queueForValidatorToValidatorsStatusChannel)
	validatorToValidatorsStatusChannel <- &structs.ValidatorToValidatorsStatusChannelStruct{ValidatorStatuses: validatorStatusToBeCreated, Validators: validatorsToBeCreated}
}

func SaveValidatorsStatus(processWaitGroup *sync.WaitGroup) {
	defer processWaitGroup.Done()
	validatorToValidatorsStatusChannelData := <-validatorToValidatorsStatusChannel
	defer close(validatorToValidatorsStatusChannel)

	var validatorStatuses []*models.ValidatorStatus
	var indexes []uint64
	for _, validatorsItem := range validatorToValidatorsStatusChannelData.Validators {
		indexes = append(indexes, validatorsItem.Index)
	}
	validatorRepo := &repositories.ValidatorRepo{
		Db: configs.GetDBInstance(),
	}
	validators, _ := validatorRepo.FetchFromIndexes(indexes)

	indexValidatorsMap := map[uint64]uint{}

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
		Db: configs.GetDBInstance(),
	}
	err := validatorStatusRepo.CreateMany(validatorStatuses)
	if err != nil {
		log.Fatalf("~~~ SaveValidatorsStatus err: %s \n", err.Error())
		// panic(err)
	}
	validatorsStatusAndSaveSlotsAndCommitteesChannel <- true
}

func SaveSlotsAndCommittees(processWaitGroup *sync.WaitGroup) {
	<-validatorsStatusAndSaveSlotsAndCommitteesChannel
	defer close(validatorsStatusAndSaveSlotsAndCommitteesChannel)
	epochRepo := &repositories.EpochRepo{
		Db: configs.GetDBInstance(),
	}
	epochs, err := epochRepo.FetchWithLimit(int(consts.EpochLimit))
	if err != nil {
		log.Fatalf("~~~ SaveSlotsAndCommittees() err: %s\n", err.Error())
	}
	var epochsIDs []uint
	for _, epochItem := range epochs {
		epochsIDs = append(epochsIDs, epochItem.ID)
	}

	stateRepo := &repositories.StateRepo{
		Db: configs.GetDBInstance(),
	}
	states, _ := stateRepo.FetchStatesAndEpochs(epochsIDs, int(consts.EpochLimit))
	// var statesStored []string
	// for _, stateItem := range states {
	// 	statesStored = append(statesStored, stateItem.StateStored)
	// }
	var consensysInstance *consensys.Consensys = &consensys.Consensys{
		Vendor: consts.VendorConfigMap[consts.Consensys],
	}
	getCommittieesWaitGroup := &sync.WaitGroup{}
	mux := &sync.Mutex{}
	var committiesFromStateAndEpochDataArray []*structs.CommittiesFromStateAndEpochData
	queueForCommittiesFromStateAndEpochDataChannel := make(chan *structs.CommittiesFromStateAndEpochData)
	for _, stateItem := range states {
		getCommittieesWaitGroup.Add(1)
		// StateStored can be passed for the results in that specifice state
		go GetCommittiesFromStateAndEpoch( stateItem.ID, string(consensysconsts.Finalized), stateItem.Eid,stateItem.Epoch.Period, consensysInstance, queueForCommittiesFromStateAndEpochDataChannel)
	}

	go func() {
		for queueForValidatorToValidatorsStatusChannelReceived := range queueForCommittiesFromStateAndEpochDataChannel {
			mux.Lock()
			committiesFromStateAndEpochDataArray = append(committiesFromStateAndEpochDataArray, queueForValidatorToValidatorsStatusChannelReceived)
			mux.Unlock()
			getCommittieesWaitGroup.Done()
		}
	}()

	getCommittieesWaitGroup.Wait()
	var slotsToBeCreated []*models.Slot
	for _, committiesFromStateAndEpochDataArrayItem := range committiesFromStateAndEpochDataArray {
		for _, slotDataItem := range committiesFromStateAndEpochDataArrayItem.SlotData {
			index := helpers.ConvertStringToUInt64(slotDataItem.Index)
			slotsToBeCreated = append(slotsToBeCreated, &models.Slot{
				Index:   index,
				Eid:     committiesFromStateAndEpochDataArrayItem.Eid,
				StateId: committiesFromStateAndEpochDataArrayItem.StateId,
			})
		}
	}
	slotsChannel := make(chan []*models.Slot)
	go CreateManySlots(slotsToBeCreated, slotsChannel)
	slots := <-slotsChannel
	if false {
		fmt.Printf("slots: %+v", slots)
	}
}

func CreateManySlots(slotsToBeCreated []*models.Slot, slotsChannel chan []*models.Slot) {
	slotRepo := &repositories.SlotRepo{
		Db: configs.GetDBInstance(),
	}
	slotRepo.CreateMany(slotsToBeCreated)
	slotsChannel <- slotsToBeCreated
}
