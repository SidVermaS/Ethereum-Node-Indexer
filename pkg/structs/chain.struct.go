package structs

import (
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/models"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/vendors/consensys/consensysstructs"
)

// Stucts used by the channels for communication between 2 goRoutines
type ValidatorsStatusAndSaveSlotsAndCommitteesChannelStruct struct {
	EpochId *uint
	StateId *uint
}

type ValidatorToValidatorsStatusChannelStruct struct {
	Validators        []*models.Validator
	ValidatorStatuses []*models.ValidatorStatus
}

type CommittieesFromStateAndEpochData struct {
	EpochId  *uint
	StateId  *uint
	BlockId  *uint
	SlotData []consensysstructs.SlotData
}
