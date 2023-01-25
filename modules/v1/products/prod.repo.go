package products

import (
	"github.com/biFebriansyah/backintro/database/orm/model"
	"gorm.io/gorm"
)

// tujuan untuk query db

type prod_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *prod_repo {
	return &prod_repo{db}
}

func (r *prod_repo) FindAll() (*model.Products, error) {

	var data model.Products
	result := r.db.Find(&data)

	if result.Error != nil {
		return nil, result.Error
	}

	return &data, nil
}

func (r *prod_repo) Add(data *model.Product) (*model.Product, error) {

	result := r.db.Create(data)
	if result.Error != nil {
		return nil, result.Error
	}

	return data, nil
}
