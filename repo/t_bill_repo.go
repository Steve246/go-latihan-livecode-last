package repo

import (
	"go_livecode_persiapan/model"

	"gorm.io/gorm"
)

type TransacRepository interface {
	Create(bill *model.Bill) error
	// Delete(menuFood *model.Menu) error
	Delete(id string) error
	Update(bill *model.Bill, id string) error
}

type transacRepository struct {
	db *gorm.DB
}

func(c *transacRepository) Update(bill *model.Bill, id string) error {

	var billModel model.Bill

	result := c.db.Model(billModel).First(&bill, "id = ?", id).Updates(billModel)

	if err := result.Error; err != nil {
		return err
	}
	return nil

}

func(c *transacRepository) Delete(id string) error {
	// result := c.db.Delete(menuFood).Error

	var billModel []model.Bill

	// var menuFood *model.Menu

	result := c.db.Where("id = ?", id).Delete(&billModel).Error
	return result

}

func(c *transacRepository) Create(bill *model.Bill) error {
	err := c.db.Create(bill).Error

	if err != nil {
		return err
	}
	return nil

}


func NewTransacRepository(db *gorm.DB) TransacRepository {
	repo := new(transacRepository)
	repo.db = db
	return repo 
}
