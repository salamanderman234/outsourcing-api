package domains

type MailInterface interface {
	GetTemplate() (string, error)
	GetSubject() string
	GetData() map[string]any
}
