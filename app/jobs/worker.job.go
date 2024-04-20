package jobs

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/salamanderman234/outsourcing-api/app/domains"
	database_types "github.com/salamanderman234/outsourcing-api/app/domains/types/databases"
	"github.com/salamanderman234/outsourcing-api/app/models"
)

type jobManager struct{}

func (jobManager) DispatchNow(job JobInterface) {
	go job.Handle()
}
func (jobManager) DispacthLater(job JobInterface, date time.Time) error {
	ctx := context.Background()
	payload, _ := json.Marshal(job.GetData())
	jobName := reflect.TypeOf(job).Name()
	datas := []models.Job{
		{Payload: string(payload), JobName: jobName, ReservedAt: date},
	}
	if domains.RepoRegistry.BaseRepo != nil {
		domains.RepoRegistry.BaseRepo.Create(ctx, datas)
	}

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
	cont := []models.Job{}
	if domains.RepoRegistry.BaseRepo != nil {
		domains.RepoRegistry.BaseRepo.ReadAll(ctx, &cont, database_types.DBSearchConfig{
			Limit: n,
			Params: []database_types.DBSearchParam{
				{Field: "reserved_at", Operator: "<="},
			},
			Query: time.Now().GoString(),
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
	jobInstance := NewJobInstance(jobName)
	jobInstance.SetData(data)

	err := jobInstance.Handle()

	ctx := context.Background()
	if err != nil {
		job.Attempts += 1
		go domains.RepoRegistry.BaseRepo.Update(ctx, []uint{job.ID}, job)
		return err
	}

	go domains.RepoRegistry.BaseRepo.Delete(ctx, []uint{job.ID}, &job)
	return nil
}

func (j jobManager) worker(id int, jobs <-chan models.Job) {
	for job := range jobs {
		fmt.Printf("Worker %d is running job-%d\n", id, int(job.ID))
		err := j.ExecuteJob(job)
		if err != nil {
			fmt.Printf("Worker %d is failing to execute job-%d, msg : %s\n", id, int(job.ID), err.Error())
		} else {
			fmt.Printf("Worker %d is fsuccessfully to execute job-%d\n", id, int(job.ID))
		}
	}
}

var JobManager = jobManager{}
