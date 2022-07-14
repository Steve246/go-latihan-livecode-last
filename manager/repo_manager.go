package manager

import "go_livecode_persiapan/repo"

type RepositoryManager interface {
	MenuRepo() repo.MenuRepository
	MenuPriceRepo() repo.MenuPriceRepository
	CustomerRepo() repo.CustomerRepository
	DiscountRepo() repo.DiscountRepository

	TableRepo() repo.TableRepository
	TransTypeRepo() repo.TransTypeRepository

}

type repositoryManager struct {
	infra Infra
}

func (r *repositoryManager) TransTypeRepo() repo.TransTypeRepository {
	return repo.NewTransTypeRepository(r.infra.SqlDb())
}

func (r *repositoryManager) TableRepo() repo.TableRepository {
	return repo.NewTableRepository(r.infra.SqlDb())
}

func (r *repositoryManager) DiscountRepo() repo.DiscountRepository {
	return repo.NewDiscountRepository(r.infra.SqlDb())
}


func (r *repositoryManager) CustomerRepo() repo.CustomerRepository {
	return repo.NewCustomerRepository(r.infra.SqlDb())
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