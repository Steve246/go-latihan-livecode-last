package usecase

import (
	"go_livecode_persiapan/model"
	"go_livecode_persiapan/repo"
)

type CrudTransacUseCase interface {
	CreateBill(bill *model.Bill) error
	// DeleteMenu(menuFood *model.Menu) error
	DeleteBill(id string) error
	UpdateBill(bill *model.Bill, id string) error
}

type crudTransacUseCase struct {
	repo repo.TransacRepository
}

func(c *crudTransacUseCase) UpdateBill(bill *model.Bill, id string) error {
	return c.repo.Update(bill, id)

}

func(c *crudTransacUseCase) DeleteBill(id string) error {
	return c.repo.Delete(id)

}

func(c *crudTransacUseCase) CreateBill(bill *model.Bill) error {
	return c.repo.Create(bill)
}

func NewCrudTransacUseCase(	repo repo.TransacRepository) CrudTransacUseCase {
	return &crudTransacUseCase{
		repo:repo,
	}
}