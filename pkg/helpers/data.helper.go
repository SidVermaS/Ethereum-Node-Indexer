package helpers

import (
	"strconv"
	"sync"

	"github.com/SidVermaS/Ethereum-Consensus/pkg/consts"
	"github.com/SidVermaS/Ethereum-Consensus/pkg/models"
	"github.com/SidVermaS/Ethereum-Consensus/pkg/repositories"
	"github.com/SidVermaS/Ethereum-Consensus/pkg/vendors/consensys"
	"github.com/SidVermaS/Ethereum-Consensus/pkg/vendors/consensys/consensysstructs"
	consensysconsts "github.com/SidVermaS/Ethereum-Consensus/pkg/vendors/consensys/consts"
)

var waitGroup = &sync.WaitGroup{}

// create a channel to communicate between goroutines (SaveBlocks(), SaveEpochs())
var blocksEpochsChannel chan []*models.Block = make(chan []*models.Block)

func ProcessToSaveDataForIndexing(finalizedCheckpoints []*consensysstructs.FinalizedCheckpoint) {
	waitGroup.Add(3)
	go SaveBlocks(finalizedCheckpoints, waitGroup)
	go SaveEpochs(finalizedCheckpoints, waitGroup)
	go SaveValidators(consensysconsts.Finalized, waitGroup)
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
		epoch, _ := strconv.Atoi(finalizedCheckpoints[index].Epoch)
		epochs = append(epochs, &models.Epoch{
			Epoch: uint(epoch),
			Bid:   uint(blockItem.ID),
		})
	}
	epochRepo := &repositories.EpochRepo{
		Db: GetDBInstance(),
	}
	err := epochRepo.CreateMany(epochs)
	if err != nil {
		panic(err)
	}
}
func SaveValidators(stateId consensysconsts.StateIdsE, processWaitGroup *sync.WaitGroup) {
	defer processWaitGroup.Done()
	var consensys *consensys.Consensys = &consensys.Consensys{
		Vendor: consts.VendorConfigMap[consts.Consensys],
	}
	var getValidatorsFromStateResponse *consensysstructs.GetValidatorsFromStateResponse = consensys.GetValidatorsFromState(stateId)
	validatorRepo := &repositories.ValidatorRepo{
		Db: GetDBInstance(),
	}
	var validators []*models.Validator

	for _, validatorItem := range getValidatorsFromStateResponse.Data {
		validators = append(validators, &models.Validator{PublicKey: validatorItem.Validator.Pubkey, IsSlashed: validatorItem.Validator.Slashed, Status: validatorItem.Status})
	}
	err := validatorRepo.CreateMany(validators)
	if err != nil {
		panic(err)
	}

}
