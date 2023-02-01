package interfaces

import (
	"github.com/biFebriansyah/backintro/database/orm/models"
	help "github.com/biFebriansyah/backintro/libs"
)

type CartRepo interface {
	FindByUserId(id string) (*models.Cart, error)
	All() (*models.Cart, error)
	Save(usersId string, items *models.CartItem) (*models.Cart, error)
}

type CartService interface {
	Get() *help.Response
	GetByUserId(id string) *help.Response
	Add(usersId string, items *models.CartItem) (*help.Response, error)
}
