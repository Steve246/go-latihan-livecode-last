package repo

import (
	"go_livecode_persiapan/model"

	"gorm.io/gorm"
)

type TableRepository interface {
	Create(table *model.Table) error
	// Delete(menuFood *model.Menu) error
	Delete(id string) error
	Update(table *model.Table, id string) error
}

type tableRepository struct {
	db *gorm.DB
}

func (c *tableRepository) Update(table *model.Table, id string) error {

	var tableModel model.Table

	result := c.db.Model(table).First(&tableModel, "id = ?", id).Updates(table)

	if err := result.Error; err != nil {
		return err
	}
	return nil

}


func (c *tableRepository) Delete(id string) error {
	var table []model.Table

	// var menuFood *model.Menu

	result := c.db.Where("id = ?", id).Delete(&table).Error
	return result

}

func (c *tableRepository) Create(table *model.Table) error {
	err := c.db.Create(table).Error

	if err != nil {
		return err
	}
	return nil

}


func NewTableRepository(db *gorm.DB) TableRepository {
	repo := new(tableRepository)
	repo.db = db
	return repo 
}