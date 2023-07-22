package consts

type CronTypesE string

const (
	CheckNewBlockAddedScheduler CronTypesE = "CheckNewBlockAddedScheduler"
)

var CronSchedulerExpressions = map[CronTypesE]string{
	// 3 minutes interval
	CheckNewBlockAddedScheduler: "*/3 * * * *",
}
