package repo

import (
	"go_livecode_persiapan/model"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	Create(customer *model.Customer) error
	// Delete(menuFood *model.Menu) error
	Delete(id string) error
	Update(customer *model.Customer, id string) error
	
}

type customerRepository struct {
	db *gorm.DB
}

func (c *customerRepository) Update(customer *model.Customer, id string) error {
	var customerModel model.Customer

	result := c.db.Model(customer).First(&customerModel, "id = ?", id).Updates(customer)

	if err := result.Error; err != nil {
		return err
	}
	return nil

}


func (c *customerRepository) Delete(id string) error {
	var customer []model.Customer

	// var menuFood *model.Menu

	result := c.db.Where("id = ?", id).Delete(&customer).Error
	return result

}

func (c *customerRepository) Create(customer *model.Customer) error {
	err := c.db.Create(customer).Error

	if err != nil {
		return err
	}
	return nil

}


func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	repo := new(customerRepository)
	repo.db = db
	return repo 
}

