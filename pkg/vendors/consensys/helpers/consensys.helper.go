package consensyshelpers

import (
	consensysconsts "github.com/SidVermaS/Ethereum-Node-Indexer/pkg/vendors/consensys/consts"
)

func ConvertTopicsSliceToStringSlice(topics []consensysconsts.ConsensysTopicsE) []string {
	var stringSlice []string
	for _, item := range topics {
		stringSlice = append(stringSlice, string(item))
	}
	return stringSlice
}
