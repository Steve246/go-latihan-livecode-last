package repo

import (
	"go_livecode_persiapan/model"

	"gorm.io/gorm"
)

type TransTypeRepository interface {
	Create(transType *model.Trans_Type) error
	// Delete(menuFood *model.Menu) error
	Delete(id string) error
	Update(transType *model.Trans_Type, id string) error
}

type transTypeRepository struct {
	db *gorm.DB
}

func (c *transTypeRepository) Update(transType *model.Trans_Type, id string) error {

	var trans model.Menu

	result := c.db.Model(transType).First(&trans, "id = ?", id).Updates(transType)

	if err := result.Error; err != nil {
		return err
	}
	return nil

}

func (c *transTypeRepository) Delete(id string) error {

	var trans []model.Trans_Type

	// var menuFood *model.Menu

	result := c.db.Where("id = ?", id).Delete(&trans).Error
	return result

}

func (c *transTypeRepository) Create(transType *model.Trans_Type) error {

	err := c.db.Create(transType).Error

	if err != nil {
		return err
	}
	return nil

}

func NewTransTypeRepository(db *gorm.DB) TransTypeRepository {
	repo := new(transTypeRepository)
	repo.db = db
	return repo 
}
