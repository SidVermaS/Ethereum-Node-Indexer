package consts

type CronTypesE string

const (
	CheckNewBlockAddedScheduler CronTypesE = "CheckNewBlockAddedScheduler"
)

var CronSchedulerExpressions = map[CronTypesE]string{
	// 13 minutes interval
	CheckNewBlockAddedScheduler: "*/13 * * * *",
}
