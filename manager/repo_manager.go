package manager

import "go_livecode_persiapan/repo"

type RepositoryManager interface {
	MenuRepo() repo.MenuRepository
}

type repositoryManager struct {
	infra Infra
}

func (r *repositoryManager) MenuRepo() repo.MenuRepository {
	return repo.NewMenuRepository(r.infra.SqlDb())
}

func NewRepositoryManager(infra Infra) RepositoryManager {
	return &repositoryManager{
		infra: infra,
	}
}