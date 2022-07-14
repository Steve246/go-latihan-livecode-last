package repo

import (
	"go_livecode_persiapan/model"

	"gorm.io/gorm"
)

type DiscountRepository interface {
	Create(discount *model.Discount) error
	// Delete(menuFood *model.Menu) error
	Delete(id string) error
	Update(discount *model.Discount, id string) error
}

type discountRepository struct {
	db *gorm.DB
}

func (c *discountRepository) Update(discount *model.Discount, id string) error {

	var discountModel model.Discount

	result := c.db.Model(discount).First(&discountModel, "id = ?", id).Updates(discount)

	if err := result.Error; err != nil {
		return err
	}
	return nil

}

func (c *discountRepository) Delete(id string) error {
	var table []model.Discount

	// var menuFood *model.Menu

	result := c.db.Where("id = ?", id).Delete(&table).Error
	return result

}

func (c *discountRepository) Create(discount *model.Discount) error {
	err := c.db.Create(discount).Error

	if err != nil {
		return err
	}
	return nil

}


func NewDiscountRepository(db *gorm.DB) DiscountRepository {
	repo := new(discountRepository)
	repo.db = db
	return repo 
}