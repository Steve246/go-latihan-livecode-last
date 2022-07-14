package manager

import "go_livecode_persiapan/usecase"

type UseCaseManager interface {
	CrudMenuUseCase() usecase.CrudMenuUseCase

	CrudMenuPriceUseCase() usecase.CrudMenuPriceUseCase

	CrudCustomerUseCase() usecase.CrudCustomerUseCase

	CrudDiscountUseCase() usecase.CrudDiscountUseCase

	CrudTableUseCase() usecase.CrudTableUseCase

	CrudTransTypeUseCase() usecase.CrudTransTypeUseCase

	CrudTransacDetailUseCase() usecase.CrudTransacDetailUseCase

	CrudTransacUseCase() usecase.CrudTransacUseCase
	
}

type useCaseManager struct {
	repoManager RepositoryManager
}

func (u *useCaseManager) CrudTransacUseCase() usecase.CrudTransacUseCase {
	return usecase.NewCrudTransacUseCase(u.repoManager.TransacRepo())
}

func (u *useCaseManager) CrudTransacDetailUseCase() usecase.CrudTransacDetailUseCase {
	return usecase.NewCrudTransacDetailUseCase(u.repoManager.TransacDetailRepo())
}

func (u *useCaseManager) CrudTransTypeUseCase() usecase.CrudTransTypeUseCase {
	return usecase.NewCrudTransTypeUseCase(u.repoManager.TransTypeRepo())
}

func (u *useCaseManager) CrudTableUseCase() usecase.CrudTableUseCase {
	return usecase.NewCrudTableUseCase(u.repoManager.TableRepo())
}

func (u *useCaseManager) CrudDiscountUseCase() usecase.CrudDiscountUseCase {
	return usecase.NewCrudDiscountUseCase(u.repoManager.DiscountRepo())

}

func (u *useCaseManager) CrudCustomerUseCase() usecase.CrudCustomerUseCase {
	return usecase.NewCrudCustomerUseCase(u.repoManager.CustomerRepo())
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