package jobs

import "github.com/salamanderman234/outsourcing-api/app/helpers"

func addCronJob() {
	// run every 5 minute
	helpers.Cron.AddDurationJob(300, helpers.JobManager.ExecuteQueue)
	helpers.Cron.AddDurationJob(300, func() {
		cancelTransactionJob := NewSetTransactionCancelJob()
		err := make(chan error)
		cancelTransactionJob.Handle(err)
	})
	helpers.Cron.AddScheduleJob(03, 0, 0, func() {
		performanceJob := NewGeneratePerformanceFormJob()
		err := make(chan error)
		performanceJob.Handle(err)
	})
	// run at 3 am everyday
	helpers.Cron.AddScheduleJob(03, 0, 0, func() {
		setTransactionJob := NewSetTransactionStatusJob()
		err := make(chan error)
		setTransactionJob.Handle(err)
	})
}

func RunCron() {
	helpers.Cron.SetupCron()
	addCronJob()
	helpers.Cron.StartCron()
}
