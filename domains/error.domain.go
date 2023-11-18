package domains

import "errors"

var (
	RepositoryErr                    = errors.New("gorm error")
	RepositoryInterfaceConversionErr = errors.New("repository interface error")
)
