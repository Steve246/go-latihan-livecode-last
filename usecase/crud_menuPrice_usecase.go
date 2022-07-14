package usecase

import (
	"go_livecode_persiapan/model"
	"go_livecode_persiapan/repo"
)

type CrudMenuPriceUseCase interface {
	CreateMenu(menuFood *model.Menu_Price) error
	// DeleteMenu(menuFood *model.Menu) error
	DeleteMenu(id string) error
	UpdateMenu(menuFood *model.Menu_Price, id string) error
}

type crudMenuPriceUseCase struct {
	repo repo.MenuPriceRepository
}

func (c *crudMenuPriceUseCase) UpdateMenu(menuFood *model.Menu_Price, id string) error{
	return c.repo.Update(menuFood, id)
}

func (c *crudMenuPriceUseCase) DeleteMenu(id string) error{
	return c.repo.Delete(id)
}

func (c *crudMenuPriceUseCase) CreateMenu(menuFood *model.Menu_Price) error {
	return c.repo.Create(menuFood)
}

func NewCrudMenuPriceUseCase(repo repo.MenuPriceRepository) CrudMenuPriceUseCase {
	return &crudMenuPriceUseCase{
		repo:repo,
	}
}
