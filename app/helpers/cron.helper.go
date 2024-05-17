package helpers

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron/v2"
)

type cron struct {
	scheduler gocron.Scheduler
}

func (c *cron) SetupCron() error {
	Logger.Info("(Queue) Setup a new scheduler...")
	cron, err := gocron.NewScheduler()
	if err != nil {
		Logger.Fatal(fmt.Sprintf("(Queue) Failed to setup scheduler : %s", err.Error()))
		return err
	}
	c.scheduler = cron
	Logger.Info("(Queue) Successfully setup a new scheduler !")
	return nil
}

func (c cron) StartCron() {
	Logger.Info("(Queue) Starting the scheduler...")
	c.scheduler.Start()
}

func (c *cron) AddDurationJob(every time.Duration, executed func()) {
	interval := gocron.DurationJob(every * time.Second)
	job := gocron.NewTask(executed)
	c.scheduler.NewJob(interval, job)
}

var Cron = cron{}
