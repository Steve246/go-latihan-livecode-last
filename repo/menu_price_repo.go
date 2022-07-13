package repo

import (
	"go_livecode_persiapan/model"

	"gorm.io/gorm"
)

type MenuPriceRepository interface {
	Create(menuFood *model.Menu_Price) error
	// Delete(menuFood *model.Menu) error
	Delete(id string) error
	Update(menuFood *model.Menu_Price, id string) error
}


type menuPriceRepository struct {
	db *gorm.DB
}

func (c *menuPriceRepository) Update(menuFood *model.Menu_Price, id string) error {

	var menu model.Menu_Price

	result := c.db.Model(menuFood).First(&menu, "id = ?", id).Updates(menuFood)

	if err := result.Error; err != nil {
		return err
	}
	return nil

}

func (c *menuPriceRepository) Delete(id string) error {
	// result := c.db.Delete(menuFood).Error

	var menu []model.Menu_Price

	// var menuFood *model.Menu

	result := c.db.Where("id = ?", id).Delete(&menu).Error
	return result
}

func (c *menuPriceRepository) Create(menuFood *model.Menu_Price) error {
	err := c.db.Create(menuFood).Error

	if err != nil {
		return err
	}
	return nil
}


func NewMenuPriceRepository(db *gorm.DB) MenuPriceRepository {
	repo := new(menuPriceRepository)
	repo.db = db
	return repo 
}