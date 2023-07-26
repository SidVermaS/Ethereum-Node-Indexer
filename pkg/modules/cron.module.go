package modules

import (
	"time"

	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/consts"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/structs"
	"github.com/go-co-op/gocron"
)

var CronInstance = &structs.Cron{}

func InitializeCron(cronInstance *structs.Cron) {
	cronInstance.CheckNewBlockAddedScheduler = gocron.NewScheduler(time.UTC)
	cronInstance.StartScheduler(consts.CheckNewBlockAddedScheduler)
}
