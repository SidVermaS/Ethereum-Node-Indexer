package modules

import (
	"fmt"
	"log"
	"strconv"
	"sync"

	"golang.org/x/exp/maps"

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

// Creates a channel to communicate between goRoutines
// These functions ran with "go" keyword will run them in GoRoutines (SaveBlocks(), SaveEpochs())
var blocksEpochsChannel chan []*models.Block = make(chan []*models.Block)
var epochsStatesChannel chan []*models.Epoch = make(chan []*models.Epoch)
var statesAndValidatorsChannel chan []*models.State = make(chan []*models.State)
var validatorToValidatorsStatusChannel chan *structs.ValidatorToValidatorsStatusChannelStruct = make(chan *structs.ValidatorToValidatorsStatusChannelStruct)
var validatorsStatusAndSaveSlotsAndCommitteesChannel chan bool = make(chan bool)

// Fetches all the validators based on the stateId
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
// Fetches Committies in a specific epoch 
func GetCommittieesFromStateAndEpoch(stateId uint, stateIdentifierOrHex string, epochId uint, epochPeriod uint, consensysInstance *consensys.Consensys, queueForCommittieesFromStateAndEpochDataChannel chan *structs.CommittieesFromStateAndEpochData) {

	var getCommitteesAtStateResponse *consensysstructs.GetCommitteesAtStateResponse = consensysInstance.GetCommitteesAtState(stateIdentifierOrHex, epochPeriod)

	queueForCommittieesFromStateAndEpochDataChannel <- &structs.CommittieesFromStateAndEpochData{Eid: epochId, StateId: stateId, SlotData: getCommitteesAtStateResponse.Data}
	return
}
// Entry point for saving the data needed for indexing
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
// Saves the blocks that came during their respective epochs
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
// The epochs are mapped with blocks and saved
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

// The states are mapped with the blocks and epochs
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
// The unique validators are saved
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
// The status of the validator in a slot is saved. It also saves whether the validator was saved or not.
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
// Mapping the committies with the validators and saving them
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

	var consensysInstance *consensys.Consensys = &consensys.Consensys{
		Vendor: consts.VendorConfigMap[consts.Consensys],
	}
	getCommittieesWaitGroup := &sync.WaitGroup{}
	mux := &sync.Mutex{}
	var CommittieesFromStateAndEpochDataArray []*structs.CommittieesFromStateAndEpochData
	queueForCommittieesFromStateAndEpochDataChannel := make(chan *structs.CommittieesFromStateAndEpochData)
	for _, stateItem := range states {
		getCommittieesWaitGroup.Add(1)
		// StateStored can be passed for the results in that specifice state
		go GetCommittieesFromStateAndEpoch(stateItem.ID, string(consensysconsts.Finalized), stateItem.Eid, stateItem.Epoch.Period, consensysInstance, queueForCommittieesFromStateAndEpochDataChannel)
	}

	go func() {
		for queueForValidatorToValidatorsStatusChannelReceived := range queueForCommittieesFromStateAndEpochDataChannel {
			mux.Lock()
			CommittieesFromStateAndEpochDataArray = append(CommittieesFromStateAndEpochDataArray, queueForValidatorToValidatorsStatusChannelReceived)
			mux.Unlock()
			getCommittieesWaitGroup.Done()
		}
	}()

	getCommittieesWaitGroup.Wait()
	var slotsToBeCreated []*models.Slot
	var validatorsIndexesMap = map[uint64]struct{}{}
	for _, committieesFromStateAndEpochDataArrayItem := range CommittieesFromStateAndEpochDataArray {
		for _, slotDataItem := range committieesFromStateAndEpochDataArrayItem.SlotData {
			index,_ := helpers.ConvertStringToUInt64(slotDataItem.Index)
			slotsToBeCreated = append(slotsToBeCreated, &models.Slot{
				Index:   index,
				Eid:     committieesFromStateAndEpochDataArrayItem.Eid,
				StateId: committieesFromStateAndEpochDataArrayItem.StateId,
			})
			for _, validatorItem := range slotDataItem.Validators {
				validatorsIndexKey,_:=helpers.ConvertStringToUInt64(validatorItem)
				validatorsIndexesMap[validatorsIndexKey] = struct{}{}
			}
		}
	}
	createManySlotsChannel := make(chan []*models.Slot)
	go CreateManySlots(slotsToBeCreated, createManySlotsChannel)
	slots := <-createManySlotsChannel

	defer close(createManySlotsChannel)
	fetchManyValidatorsChannel := make(chan []*models.Validator)
	go FetchManyValidators(maps.Keys(validatorsIndexesMap), fetchManyValidatorsChannel)
	validators := <-fetchManyValidatorsChannel
	defer close(fetchManyValidatorsChannel)
	var validatorsMap = map[string]*models.Validator{}
	for _, validatorItem := range validators {
		validatorsMap[fmt.Sprint(validatorItem.Index)] = validatorItem
	}
	var committees []*models.Committee
	for committieesFromStateAndEpochDataArrayIndex, committieesFromStateAndEpochDataArrayItem := range CommittieesFromStateAndEpochDataArray {
		for slotIndex, slotDataItem := range committieesFromStateAndEpochDataArrayItem.SlotData {
			for _, validatorItem := range slotDataItem.Validators {
				committees = append(committees, &models.Committee{
					Eid:     states[committieesFromStateAndEpochDataArrayIndex].Eid,
					StateId: states[committieesFromStateAndEpochDataArrayIndex].ID,
					SlotId:  slots[slotIndex].ID,
					Vid:     validatorsMap[validatorItem].ID,
				})
			}
		}
	}

	committeeRepo := &repositories.CommitteeRepo{
		Db: configs.GetDBInstance(),
	}
	committeeRepo.CreateMany(committees)
}
// Slots are insereted in to the DB
func CreateManySlots(slotsToBeCreated []*models.Slot, createManySlotsChannel chan []*models.Slot) {
	slotRepo := &repositories.SlotRepo{
		Db: configs.GetDBInstance(),
	}
	slotRepo.CreateMany(slotsToBeCreated)
	createManySlotsChannel <- slotsToBeCreated
}

// Fetch the validators from their indexes
func FetchManyValidators(indexes []uint64, fetchManyValidatorsChannel chan []*models.Validator) {
	validatorRepo := &repositories.ValidatorRepo{
		Db: configs.GetDBInstance(),
	}
	validators, _ := validatorRepo.FetchFromIndexes(indexes)
	fetchManyValidatorsChannel <- validators
}
