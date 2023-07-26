package helpers

import (
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/structs"
)

func CalculateNetworksParticipationRate(data *structs.CalculateParticipatiRateStruct) float64 {
	var denominator = ((data.Epochs) * data.SlotsPerEpoch * data.ValidatorSetSize)
	if denominator == 0 {
		return 0
	} else {
		participationRate := 1 - float64(data.MissedAttestations/denominator)

		return participationRate * 100
	}
}

func CalculateValidatorParticipationRate(data *structs.CalculateParticipatiRateStruct) float64 {
	participationRate := 1 - float64(data.MissedAttestations/(data.Epochs*data.SlotsPerEpoch))
	return participationRate * 100
}
