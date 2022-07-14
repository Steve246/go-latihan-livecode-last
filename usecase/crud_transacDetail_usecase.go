package usecase

import (
	"go_livecode_persiapan/model"
	"go_livecode_persiapan/repo"
)

type CrudTransacDetailUseCase interface {
	CreateBillDetail(billDetail *model.Bill_Detail) error

	DeleteBillDetail(id string) error
	UpdateBillDetail(billDetail *model.Bill_Detail, id string) error
}

type crudTransacDetailUseCase struct {
	repo repo.TransacDetailRepository
}

func (c *crudTransacDetailUseCase) UpdateBillDetail(billDetail *model.Bill_Detail, id string) error {
	return c.repo.Update(billDetail, id)

}

func (c *crudTransacDetailUseCase) DeleteBillDetail(id string) error {
	return c.repo.Delete(id)

}

func (c *crudTransacDetailUseCase) CreateBillDetail(billDetail *model.Bill_Detail) error {
	return c.repo.Create(billDetail)
}

func NewCrudTransacDetailUseCase(repo repo.TransacDetailRepository) CrudTransacDetailUseCase {
	return &crudTransacDetailUseCase{
		repo:repo,
	}
}
