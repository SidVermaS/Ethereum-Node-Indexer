package utils

import (
	"time"

	"github.com/SidVermaS/Ethereum-Consensus/pkg/consts"
	"github.com/SidVermaS/Ethereum-Consensus/pkg/types/structs"
	"github.com/go-co-op/gocron"
)

func InitializeCron(cronInstance *structs.Cron) {
	cronInstance = &structs.Cron{}

	cronInstance.CheckNewBlockAddedScheduler = gocron.NewScheduler(time.UTC)
	cronInstance.StartScheduler(consts.CheckNewBlockAddedScheduler)
}
