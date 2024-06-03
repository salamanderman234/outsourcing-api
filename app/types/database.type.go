package types

import "github.com/salamanderman234/outsourcing-api/app/domains"

type WhereQuery struct {
	Field    string
	Operator string
	Str      string
	IsOr     bool
}

type DBSearchParams struct {
	SortBy         string
	IsDesc         bool
	WithPagination bool
	Params         []WhereQuery
	Query          string
	Page           uint
	Preloads       []string
	Model          domains.ModelInterface
	Limit          int
	Joins          []string
}
