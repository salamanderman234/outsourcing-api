package domains

type ResourceInterface interface {
	GetFullPath() string
	GetName() (string, error)
	GetFieldName() string
	GetFieldValue() string
	GetFileConfig() any
	GetPolicy() Policy
	GetData() ModelInterface
	SetData(data ModelInterface)
}
