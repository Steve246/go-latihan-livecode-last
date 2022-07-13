package manager

import "go_livecode_persiapan/usecase"

type UseCaseManager interface {
	CrudMenuUseCase() usecase.CrudMenuUseCase
}

type useCaseManager struct {
	repoManager RepositoryManager
}

func (u *useCaseManager) CrudMenuUseCase() usecase.CrudMenuUseCase {
	return usecase.NewCrudMenuUsecase(u.repoManager.MenuRepo())
}

func NewUseCaseManager(repoManager RepositoryManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repoManager,
	}
}