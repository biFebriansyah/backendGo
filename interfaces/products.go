package interfaces

import (
	"github.com/biFebriansyah/backintro/database/orm/models"
	"github.com/biFebriansyah/backintro/libs"
)

type ProductRepo interface {
	FindAll() (*models.Products, error)
	Save(data *models.Product) (*models.Product, error)
	GetById(uid string) (*models.Products, error)
	DeleteByid(uid string) (*models.Products, error)
}

type ProductService interface {
	GetAll() *libs.Response
	Add(data *models.Product) *libs.Response
	GetProdWithId(uid string) *libs.Response
	RemoveByUid(uid string) *libs.Response
}
