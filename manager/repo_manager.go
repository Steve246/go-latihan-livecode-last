package manager

import "go_livecode_persiapan/repo"

type RepositoryManager interface {
	MenuRepo() repo.MenuRepository
	MenuPriceRepo() repo.MenuPriceRepository
}

type repositoryManager struct {
	infra Infra
}

func (r *repositoryManager) MenuPriceRepo() repo.MenuPriceRepository {
	return repo.NewMenuPriceRepository(r.infra.SqlDb())
}

func (r *repositoryManager) MenuRepo() repo.MenuRepository {
	return repo.NewMenuRepository(r.infra.SqlDb())
}

func NewRepositoryManager(infra Infra) RepositoryManager {
	return &repositoryManager{
		infra: infra,
	}
}