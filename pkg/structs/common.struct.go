package structs

import "github.com/SidVermaS/Ethereum-Node-Indexer/pkg/models"

type ValidatorToValidatorsStatusChannelStruct struct {
	Validators        []*models.Validator
	ValidatorStatuses []*models.ValidatorStatus
}
