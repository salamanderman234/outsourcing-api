package jobs

import "github.com/salamanderman234/outsourcing-api/app/helpers"

func addCronJob() {
	helpers.Cron.AddDurationJob(300, helpers.JobManager.ExecuteQueue)
}

func RunCron() {
	helpers.Cron.SetupCron()
	addCronJob()
	helpers.Cron.StartCron()
}
