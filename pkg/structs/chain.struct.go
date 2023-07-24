package structs

import (
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/models"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/vendors/consensys/consensysstructs"
)

type ValidatorsStatusAndSaveSlotsAndCommitteesChannelStruct struct {
	Eid        uint
	StateId uint
}


type ValidatorToValidatorsStatusChannelStruct struct {
	Validators        []*models.Validator
	ValidatorStatuses []*models.ValidatorStatus
}

type CommittieesFromStateAndEpochData struct	{
	Eid        uint
	StateId uint
	SlotData []consensysstructs.SlotData
}

