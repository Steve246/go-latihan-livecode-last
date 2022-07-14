package usecase

import (
	"go_livecode_persiapan/model"
	"go_livecode_persiapan/repo"
)

type CrudDiscountUseCase interface {
	CreateDiscount(disc *model.Discount) error
	// DeleteMenu(menuFood *model.Menu) error
	DeleteDiscount(id string) error
	UpdateDiscount(disc *model.Discount, id string) error
}

type crudDiscountUseCase struct {
	repo repo.DiscountRepository
}

func (c *crudDiscountUseCase) UpdateDiscount(disc *model.Discount, id string) error {
	return c.repo.Update(disc, id)

}

func (c *crudDiscountUseCase) DeleteDiscount(id string) error {
	return c.repo.Delete(id)

}

func (c *crudDiscountUseCase) CreateDiscount(disc *model.Discount) error {
	return c.repo.Create(disc)

}

func NewCrudDiscountUseCase(repo repo.DiscountRepository) CrudDiscountUseCase {
	return &crudDiscountUseCase{
		repo:repo,
	}
}