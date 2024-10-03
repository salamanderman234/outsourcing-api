package domains

type JobInterface interface {
	GetData() map[string]any
	SetData(data map[string]any)
	Handle(err chan<- error)
}
