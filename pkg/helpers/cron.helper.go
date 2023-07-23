package helpers

import (
	"time"

	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/consts"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/structs"
	"github.com/go-co-op/gocron"
)

func InitializeCron(cronInstance *structs.Cron) {
	cronInstance = &structs.Cron{}

	cronInstance.CheckNewBlockAddedScheduler = gocron.NewScheduler(time.UTC)
	cronInstance.StartScheduler(consts.CheckNewBlockAddedScheduler)
}
