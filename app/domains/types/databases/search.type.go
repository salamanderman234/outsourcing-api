package database_types

import "github.com/salamanderman234/outsourcing-api/app/models"

type DBSearchParam struct {
	Field    string
	Operator string
}

type DBSearchConfig struct {
	SortBy         string
	IsDesc         bool
	WithPagination bool
	Params         []DBSearchParam
	Query          string
	Page           uint
	Preloads       []string
	Model          models.ModelInterface
	Limit          int
}
