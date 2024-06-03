package helpers

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/salamanderman234/outsourcing-api/app/domains"
	"github.com/salamanderman234/outsourcing-api/app/models"
	"github.com/salamanderman234/outsourcing-api/app/providers"
	"github.com/salamanderman234/outsourcing-api/app/types"
)

type jobManager struct{}

func (jobManager) DispatchNow(job domains.JobInterface) error {
	errs := make(chan error)
	go job.Handle(errs)
	// err := <-errs
	return nil
}
func (jobManager) DispacthLater(job domains.JobInterface, date time.Time) error {
	ctx := context.Background()
	payload, _ := json.Marshal(job.GetData())
	jobName := reflect.TypeOf(job).Name()
	datas := []models.Job{
		{Payload: string(payload), JobName: jobName, ReservedAt: date},
	}
	if providers.RepoProvider.BaseRepo != nil {
		providers.RepoProvider.BaseRepo.Create(ctx, datas)
	}
	Logger.Info(
		fmt.Sprintf("(Queue) Dispatching a new job (%s) into database that will execute on %s", jobName, date.String()),
	)
	return nil
}

func (j jobManager) ExecuteQueue() {
	nWorker := 10
	buff := 100
	jobs := make(chan models.Job, buff)

	for i := 1; i <= nWorker; i++ {
		go j.worker(i, jobs)
	}
	j.fetchJob(buff, jobs)
}

func (j jobManager) fetchJob(n int, jobs chan<- models.Job) {
	ctx := context.Background()
	Logger.Info("(Queue) Fetching a new job list...")
	cont := []models.Job{}
	if providers.RepoProvider.BaseRepo != nil {
		providers.RepoProvider.BaseRepo.ReadAll(ctx, &cont, types.DBSearchParams{
			Limit: n,
			Params: []types.WhereQuery{
				{Field: "reserved_at", Operator: "<=", Str: time.Now().GoString()},
			},
		})
	}

	for _, job := range cont {
		jobs <- job
	}
}

func (jobManager) ExecuteJob(job models.Job) error {
	var data map[string]any
	payload, _ := json.Marshal(job.Payload)
	json.Unmarshal(payload, &data)
	jobName := job.JobName
	jobInstance := providers.JobProvider.NewJobInstance(jobName)
	jobInstance.SetData(data)

	Logger.Info(
		fmt.Sprintf("(Queue) Attempting to execute job (%s)#%d...", jobName, job.ID),
	)

	errs := make(chan error)
	go jobInstance.Handle(errs)

	err := <-errs

	ctx := context.Background()
	if err != nil {
		Logger.Warning(
			fmt.Sprintf("(Queue) Failing to execute job (%s)#%d, err: %s", jobName, job.ID, err.Error()),
		)
		job.Attempts += 1
		go providers.RepoProvider.BaseRepo.Update(ctx, []uint{job.ID}, &job)
		return err
	}
	Logger.Info(
		fmt.Sprintf("(Queue) Successfully to execute job(%s)#%d !", jobName, job.ID),
	)
	go providers.RepoProvider.BaseRepo.Delete(ctx, []uint{job.ID}, &job)
	return nil
}

func (j jobManager) worker(id int, jobs <-chan models.Job) {
	for job := range jobs {
		j.ExecuteJob(job)
	}
}

var JobManager = jobManager{}
