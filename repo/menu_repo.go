package repo

import (
	"go_livecode_persiapan/model"

	"gorm.io/gorm"
)

type MenuRepository interface {
	Create(menuFood *model.Menu) error
	// Delete(menuFood *model.Menu) error
	Delete(id string) error
	Update(menuFood *model.Menu, id string) error
}

type menuRepository struct {
	db *gorm.DB
}

func (c *menuRepository) Update(menuFood *model.Menu, id string) error {

	var menu model.Menu

	result := c.db.Model(menuFood).First(&menu, "id = ?", id).Updates(menuFood)

	if err := result.Error; err != nil {
		return err
	}
	return nil

}

func (c *menuRepository) Delete(id string) error {
	// result := c.db.Delete(menuFood).Error

	var menu []model.Menu

	// var menuFood *model.Menu

	result := c.db.Where("id = ?", id).Delete(&menu).Error
	return result
}

func (c *menuRepository) Create(menuFood *model.Menu) error {
	err := c.db.Create(menuFood).Error

	if err != nil {
		return err
	}
	return nil
}


func NewMenuRepository(db *gorm.DB) MenuRepository {
	repo := new(menuRepository)
	repo.db = db
	return repo 
}