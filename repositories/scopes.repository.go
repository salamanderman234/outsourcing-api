package repositories

import (
	"context"
	"sync"

	"github.com/salamanderman234/outsourcing-api/configs"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

func paginateScope(page uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := configs.PAGINATION_PER_PAGE * (page - 1)
		take := configs.PAGINATION_PER_PAGE
		return db.Offset(int(offset)).Take(take)
	}
}

func orderScope(model any, orderBy string, desc bool) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		s, err := schema.Parse(model, &sync.Map{}, schema.NamingStrategy{})
		if err != nil {
			return db
		}
		column := clause.Column{Name: s.Table + ".updated_at"}
		sortStatement := clause.OrderByColumn{Column: column, Desc: desc}
		if orderBy != "" {
			valid := false
			for _, field := range s.Fields {
				if orderBy == field.Name {
					valid = true
					break
				}
			}
			if valid {
				column.Name = s.Table + "." + orderBy
			}
		}
		return db.Order(sortStatement)
	}
}

func userSearchScope(usernameField string, username string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(usernameField+" = ?", username)
	}
}

func whereIdEqualScope(id any) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	}
}

func usingContextScope(c context.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.WithContext(c)
	}
}

func usingModelScope(model any) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Model(model)
	}
}
