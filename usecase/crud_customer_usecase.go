package usecase

import (
	"go_livecode_persiapan/model"
	"go_livecode_persiapan/repo"
)

type CrudCustomerUseCase interface {
	CreateCustomer(customer *model.Customer) error
	// DeleteMenu(menuFood *model.Menu) error
	DeleteCustomer(id string) error
	UpdateCustomer(customer *model.Customer, id string) error
}

type crudCustomerUseCase struct {
	repo repo.CustomerRepository
}

func (c *crudCustomerUseCase) UpdateCustomer(customer *model.Customer, id string) error{
	return c.repo.Update(customer, id)
}

func (c *crudCustomerUseCase) DeleteCustomer(id string) error {
	return c.repo.Delete(id)
}

func (c *crudCustomerUseCase) CreateCustomer(customer *model.Customer) error {
	return c.repo.Create(customer)
}

func NewCrudCustomerUseCase(repo repo.CustomerRepository) CrudCustomerUseCase {
	return &crudCustomerUseCase{
		repo:repo,
	}
}