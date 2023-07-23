package modules

import (
	
	"strconv"

	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/models"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/structs"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/vendors/consensys"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/vendors/consensys/consensysstructs"
)

func GetValidatorsFromState(stateDbId uint, epochDbId uint, stateIdentifierOrHex string, consensysInstance *consensys.Consensys,queueForValidatorToValidatorsStatusChannel chan *structs.ValidatorToValidatorsStatusChannelStruct) {
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
