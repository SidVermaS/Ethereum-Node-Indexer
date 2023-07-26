package modules

import (
	configs "github.com/SidVermaS/Ethereum-Node-Indexer/pkg/config"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/models"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/repositories"
)

func GetLastEpoch() (*models.Epoch, error) {
	epochRepo := &repositories.EpochRepo{
		Db: configs.GetDBInstance(),
	}
	epochs, err := epochRepo.FetchWithLimit(1)
	if err != nil {
		return nil, err
	}
	if len(epochs) <= 0 {
		return nil, nil
	}
	return epochs[0], nil
}
