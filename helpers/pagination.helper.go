package helpers

import "github.com/salamanderman234/outsourcing-api/domains"

func MakePagination(maxPage uint, currentPage uint, queries map[string]any) domains.Pagination {
	nextPage := min(maxPage, currentPage)
	previousPage := max(1, currentPage)

	return domains.Pagination{
		Next:     nextPage,
		Previous: previousPage,
		Current:  currentPage,
		MaxPage:  maxPage,
		Queries:  queries,
	}
}

func MakeDefaultGetPaginationQueries(q string, id uint, page uint, orderBy string, desc bool) map[string]any {
	paginationQueries := map[string]any{}
	if q != "" {
		paginationQueries["q"] = q
	}
	if id != 0 {
		paginationQueries["id"] = id
	}
	if page > 1 {
		paginationQueries["page"] = page
	}
	if orderBy != "" {
		paginationQueries["order_by"] = orderBy
	}
	if !desc {
		paginationQueries["desc"] = 0
	}
	return paginationQueries
}