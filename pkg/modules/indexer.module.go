package modules

import (
	configs "github.com/SidVermaS/Ethereum-Node-Indexer/pkg/config"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/consts"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/helpers"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/repositories"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/structs"
	consensysconsts "github.com/SidVermaS/Ethereum-Node-Indexer/pkg/vendors/consensys/consts"
)

func GetNetworksParticipationRate() (float64, error) {
	epochRepo := &repositories.EpochRepo{
		Db: configs.GetDBInstance(),
	}
	epochs, err := epochRepo.FetchWithLimit(int(consts.EpochLimit))
	if err != nil {
		return -1, err
	}
	slotRepo := &repositories.SlotRepo{
		Db: configs.GetDBInstance(),
	}
	var eids []uint
	for _, epochItem := range epochs {
		eids = append(eids, epochItem.ID)
	}
	slots, err := slotRepo.FetchByEids(eids)
	if err != nil {
		return -1, err
	}

	var slotsIds []uint
	for _, slotItem := range slots {
		slotsIds = append(slotsIds, slotItem.ID)
	}
	validatorStatusRepo := &repositories.ValidatorStatusRepo{
		Db: configs.GetDBInstance(),
	}
	validatorStatuses, err := validatorStatusRepo.FetchByEidsAndSlotsIds(eids, slotsIds)
	if err != nil {
		return -1, err
	}
	var missedAttestations uint = 0
	for _, validatorStatusItem := range validatorStatuses {
		if validatorStatusItem.IsSlashed || validatorStatusItem.Status != string(consensysconsts.ActiveOngoing) {
			missedAttestations++
		}
	}
	networksParticipationRate := helpers.CalculateNetworksParticipationRate(&structs.CalculateParticipatiRateStruct{
		MissedAttestations: int(missedAttestations),
		SlotsPerEpoch:      len(slots),
		ValidatorSetSize:   len(validatorStatuses),
		Epochs:             len(epochs),
	})
	return networksParticipationRate, nil
}
