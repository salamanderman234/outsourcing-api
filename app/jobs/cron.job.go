package jobs

import "github.com/salamanderman234/outsourcing-api/app/helpers"

func addCronJob() {
	helpers.Cron.AddDurationJob(300, helpers.JobManager.ExecuteQueue)
	helpers.Cron.AddScheduleJob(20, 58, 0, func() {
		performanceJob := NewGeneratePerformanceFormJob()
		err := make(chan error)
		performanceJob.Handle(err)
	})
}

func RunCron() {
	helpers.Cron.SetupCron()
	addCronJob()
	helpers.Cron.StartCron()
}
