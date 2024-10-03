package repositories

import (
	"github.com/salamanderman234/outsourcing-api/app/providers"
)

func RegisterAllRepos() {
	providers.RepoProvider.BaseRepo = NewBaseRepo()
	providers.RepoProvider.UserRepo = NewUserRepo()
	providers.RepoProvider.PerformanceRepo = NewPerformanceRepo()
}
