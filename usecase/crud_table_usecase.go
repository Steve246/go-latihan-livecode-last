package usecase

import (
	"go_livecode_persiapan/model"
	"go_livecode_persiapan/repo"
)

type CrudTableUseCase interface {
	CreateTable(table *model.Table) error
	DeleteTable(id string) error
	UpdateTable(table *model.Table, id string) error
}


type crudTableUseCase struct {
	repo repo.TableRepository
}

func (c *crudTableUseCase) 	CreateTable(table *model.Table) error{

	return c.repo.Create(table)

 }

func (c *crudTableUseCase) DeleteTable(id string) error {
	return c.repo.Delete(id)
}

func (c *crudTableUseCase) 	UpdateTable(table *model.Table, id string) error {
	return c.repo.Update(table, id)

}


func NewCrudTableUseCase(repo repo.TableRepository) CrudTableUseCase {
	return &crudTableUseCase{
		repo:repo,
	}
}