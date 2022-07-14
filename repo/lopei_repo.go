package repo

import (
	"go_livecode_persiapan/model"
)

// type LopeiRepository interface {
// 	RetrieveById(id int32) (model.Bill, error)
// }

// //retrieve by id
// // update tablenya

// type lopeiRepository struct {
// 	db *gorm.DB
// }

// func (p *lopeiRepository) RetrieveById(id int32) (model.Bill, error) {

// 	var bill model.Bill

// 	result:= p.db.Unscoped().Where("id = ?", id).First(&bill)

// 	if err := result.Error; err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return bill, nil
// 		} else {
// 			return bill, err
// 		}
// 	}
// 	return bill, nil

// }

// func NewLopeiRepository(db *gorm.DB) LopeiRepository {
// 	repo := new(lopeiRepository)
// 	repo.db = db
// 	return repo
// }

type LopeiRepository interface {
	RetrieveById(id int32) (model.CustomerStruk, error)
}

type lopeiRepository struct {
	db []model.CustomerStruk
}

func (p *lopeiRepository) RetrieveById(id int32) (model.CustomerStruk, error) {
	for _, cust := range p.db {
		if cust.LopeiId == id {
			return cust, nil
		}
	}
	return model.CustomerStruk{}, nil
}

func NewLopeiRepository() LopeiRepository {
	repo := new(lopeiRepository)
	repo.db = []model.CustomerStruk{
		{LopeiId: 1, Balance: 5000},
		{LopeiId: 2, Balance: 4000},
		{LopeiId: 3, Balance: 2000},
	}
	return repo
}

