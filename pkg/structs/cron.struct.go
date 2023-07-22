package structs

import (
	"fmt"

	"github.com/SidVermaS/Ethereum-Consensus/pkg/consts"
	"github.com/go-co-op/gocron"
)

type Cron struct {
	// Schedulers for different functions
	CheckNewBlockAddedScheduler *gocron.Scheduler
}

func Greet() {
	fmt.Println("Hello Ether Consensus!")
}

// Define StartScheduler() method on the Cron struct
func (cronInstance *Cron) StartScheduler(cronType interface{}) {

	if cronType == nil || cronType == consts.CheckNewBlockAddedScheduler {
		var expression string = consts.CronSchedulerExpressions[consts.CheckNewBlockAddedScheduler]
		cronInstance.CheckNewBlockAddedScheduler.Cron(expression).Do(Greet)

	}
}