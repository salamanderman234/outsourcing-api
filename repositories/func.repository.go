package repositories

import (
	"context"
	"errors"
	"math"

	"github.com/salamanderman234/outsourcing-api/configs"
	"github.com/salamanderman234/outsourcing-api/domains"
	"gorm.io/gorm"
)

func basicCreateRepoFunc(c context.Context, db *gorm.DB, model any, data any) error {
	result := db.Scopes(usingContextScope(c), usingModelScope(model)).Create(data)
	return convertRepoError(result)
}

func basicFindRepoFunc(c context.Context, db *gorm.DB, model any, id uint, result any, preloads ...string) error {
	queryResult := db.Scopes(usingContextScope(c), usingModelScope(model), whereIdEqualScope(id))
	for _, preload := range preloads {
		queryResult.Preload(preload)
	}
	queryResult.First(result)
	return convertRepoError(queryResult)
}

func basicUpdateRepoFunc(c context.Context, db *gorm.DB, model any, id uint, data any) (int64, error) {
	result := db.Scopes(usingContextScope(c), usingModelScope(model), whereIdEqualScope(id)).Updates(data)
	if result.RowsAffected == 0 {
		return 0, domains.ErrRecordNotFound
	}
	return result.RowsAffected, convertRepoError(result)
}

func basicDeleteRepoFunc(c context.Context, db *gorm.DB, model any, id uint) (int64, error) {
	result := db.Scopes(usingContextScope(c), whereIdEqualScope(id)).Delete(model)
	if result.RowsAffected == 0 {
		return 0, domains.ErrRecordNotFound
	}
	return result.RowsAffected, convertRepoError(result)
}

func convertRepoError(q *gorm.DB) error {
	err := q.Error
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return domains.ErrDuplicateEntries
	}
	if errors.Is(err, gorm.ErrForeignKeyViolated) {
		return domains.ErrForeignKeyViolated
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return domains.ErrRecordNotFound
	}
	if err != nil {
		return domains.ErrRepository
	}
	return nil
}

func getMaxPage(countRes uint) uint {
	return uint(math.Ceil(float64(countRes) / configs.PAGINATION_PER_PAGE))
}
