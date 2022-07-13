package usecase

import (
	"go_livecode_persiapan/model"
	"go_livecode_persiapan/repo"
)

type CrudMenuUseCase interface {
	CreateMenu(menuFood *model.Menu) error
	// DeleteMenu(menuFood *model.Menu) error
	DeleteMenu(id string) error
	UpdateMenu(menuFood *model.Menu, id string) error
}

type crudMenuUseCase struct {
	repo repo.MenuRepository
}

func (c *crudMenuUseCase) UpdateMenu(menuFood *model.Menu, id string) error {
	return c.repo.Update(menuFood, id)
}

func (c *crudMenuUseCase) DeleteMenu(id string) error{
	return c.repo.Delete(id)
}

func (c *crudMenuUseCase) CreateMenu(menuFood *model.Menu) error {
	return c.repo.Create(menuFood)
}

func NewCrudMenuUsecase(repo repo.MenuRepository) CrudMenuUseCase {
	return &crudMenuUseCase{
		repo:repo,
	}
}