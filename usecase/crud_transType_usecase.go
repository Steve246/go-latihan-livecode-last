package usecase

import (
	"go_livecode_persiapan/model"
	"go_livecode_persiapan/repo"
)

type CrudTransTypeUseCase interface {
	CreateTransType(trans *model.Trans_Type) error
	// DeleteMenu(menuFood *model.Menu) error
	DeleteTransType(id string) error
	UpdateTransType(trans *model.Trans_Type, id string) error
}

type crudTransTypeUseCase struct {
	repo repo.TransTypeRepository
}

func (c *crudTransTypeUseCase) UpdateTransType(trans *model.Trans_Type, id string) error {
	return c.repo.Update(trans, id)

}

func (c *crudTransTypeUseCase) DeleteTransType(id string) error {
	return c.repo.Delete(id)

}

func (c *crudTransTypeUseCase) CreateTransType(trans *model.Trans_Type) error {
	return c.repo.Create(trans)
}

func NewCrudTransTypeUseCase(repo repo.TransTypeRepository) CrudTransTypeUseCase {
	return &crudTransTypeUseCase{
		repo:repo,
	}
}