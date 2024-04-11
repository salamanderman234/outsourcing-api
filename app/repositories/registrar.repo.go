package repositories

import "github.com/salamanderman234/outsourcing-api/app/domains"

func RegisterAllRepos() {
	domains.RepoRegistry.BaseRepo = NewBaseRepo()
}
