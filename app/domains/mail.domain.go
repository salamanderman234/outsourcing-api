package domains

type MailInterface interface {
	GetTemplate() (string, error)
	GetData() map[string]any
	GetSubject() string
}
