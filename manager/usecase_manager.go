package manager

import "go_livecode_persiapan/usecase"

type UseCaseManager interface {
	CrudMenuUseCase() usecase.CrudMenuUseCase

	CrudMenuPriceUseCase() usecase.CrudMenuPriceUseCase
}

type useCaseManager struct {
	repoManager RepositoryManager
}

func (u *useCaseManager) CrudMenuPriceUseCase() usecase.CrudMenuPriceUseCase {
	return usecase.NewCrudMenuPriceUseCase(u.repoManager.MenuPriceRepo())
}

func (u *useCaseManager) CrudMenuUseCase() usecase.CrudMenuUseCase {
	return usecase.NewCrudMenuUsecase(u.repoManager.MenuRepo())
}

func NewUseCaseManager(repoManager RepositoryManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repoManager,
	}
}