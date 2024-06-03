package domains

type ResourceInterface interface {
	GetFullPath() string
	GetName() (string, error)
	GetFieldName() string
	GetFieldValue() string
	GetFileConfig() any
	GetPolicy() any
	GetData() ModelInterface
	SetData(data ModelInterface)
}
