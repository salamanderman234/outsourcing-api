package jobs

import "reflect"

type generateNewJobInstance func() JobInterface

var jobRegistry map[string]generateNewJobInstance

type JobInterface interface {
	GetData() map[string]any
	SetData(data map[string]any)
	Handle() error
}

func registerJob(job JobInterface) {
	jobName := reflect.TypeOf(job).Name()
	jobRegistry[jobName] = func() JobInterface {
		e := reflect.New(reflect.TypeOf(job)).Elem().Interface()
		return e.(JobInterface)
	}
}

func NewJobInstance(name string) JobInterface {
	fun, ok := jobRegistry[name]
	if !ok {
		return nil
	}
	return fun()
}
