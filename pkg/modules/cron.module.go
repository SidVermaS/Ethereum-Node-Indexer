package modules

import (
	"sync"
	"time"

	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/consts"
	"github.com/go-co-op/gocron"
)

type Cron struct {
	// Schedulers for different functions
	CheckNewBlockAddedScheduler *gocron.Scheduler
}

// Define StartScheduler() method on the Cron struct
func (cronInstance *Cron) StartScheduler(cronType interface{}) {

	if cronType == nil || cronType == consts.CheckNewBlockAddedScheduler {
		var expression string = consts.CronSchedulerExpressions[consts.CheckNewBlockAddedScheduler]
		cronInstance.CheckNewBlockAddedScheduler.Cron(expression).Do(ResetEpochsMetaData)

	}
}

// Define StartScheduler() method on the Cron struct
func (cronInstance *Cron) StopScheduler(cronType interface{}) {

	if cronType == nil || cronType == consts.CheckNewBlockAddedScheduler {
		cronInstance.CheckNewBlockAddedScheduler.Stop()
	}
}

// Common cron instance which has all the schedulers needed on the app
var cronInstance = &Cron{}

// Starts the Cron Schedulers
func StartCronSchedulers() {
	//	Initializing a new Scheduler
	cronInstance.CheckNewBlockAddedScheduler = gocron.NewScheduler(time.UTC)
	cronInstance.StartScheduler(consts.CheckNewBlockAddedScheduler)
}

// Stops the Cron Schedulers
func StopCronSchedulers(waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	cronInstance.StopScheduler(nil)
}
