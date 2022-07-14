package repo

import (
	"go_livecode_persiapan/model"

	"gorm.io/gorm"
)

type TransacDetailRepository interface {
	Create(billDetail *model.Bill_Detail) error
	// Delete(menuFood *model.Menu) error
	Delete(id string) error
	Update(billDetail *model.Bill_Detail, id string) error
}

type transacDetailRepository struct {
	db *gorm.DB
}

func(c *transacDetailRepository) Update(billDetail *model.Bill_Detail, id string) error {
	var billDetailModel model.Bill_Detail

	result := c.db.Model(billDetail).First(&billDetailModel, "id = ?", id).Updates(billDetail)

	if err := result.Error; err != nil {
		return err
	}
	return nil
	

}

func(c *transacDetailRepository) Delete(id string) error {

	// result := c.db.Delete(menuFood).Error

	var billDetailModel []model.Bill_Detail

	// var menuFood *model.Menu

	result := c.db.Where("id = ?", id).Delete(&billDetailModel).Error
	return result

}

func(c *transacDetailRepository) Create(billDetail *model.Bill_Detail) error {
	err := c.db.Create(billDetail).Error

	if err != nil {
		return err
	}
	return nil

}

func NewTransacDetailRepository(db *gorm.DB) TransacDetailRepository {
	repo := new(transacDetailRepository)
	repo.db = db
	return repo 
}
