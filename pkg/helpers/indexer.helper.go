package helpers

import (
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/structs"
)

func CalculateNetworksParticipationRate(data *structs.CalculateParticipatiRateStruct) float64 {
	participationRate := 1 - float64(data.MissedAttestations/((data.Epochs)*data.SlotsPerEpoch*data.ValidatorSetSize))
	return participationRate*100
}
