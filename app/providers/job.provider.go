package providers

import (
	"reflect"

	"github.com/salamanderman234/outsourcing-api/app/domains"
)

type generateNewJobInstance func() domains.JobInterface
type jobProvider struct {
	list map[string]generateNewJobInstance
}

func (j *jobProvider) RegisterJob(job domains.JobInterface) {
	jobName := reflect.TypeOf(job).Name()
	j.list[jobName] = func() domains.JobInterface {
		e := reflect.New(reflect.TypeOf(job)).Elem().Interface()
		return e.(domains.JobInterface)
	}
}

func (j *jobProvider) NewJobInstance(name string) domains.JobInterface {
	fun, ok := j.list[name]
	if !ok {
		return nil
	}
	return fun()
}

var JobProvider = jobProvider{
	list: make(map[string]generateNewJobInstance),
}
