package interfaces

import (
	"github.com/biFebriansyah/backintro/database/orm/models"
	"github.com/biFebriansyah/backintro/libs"
)

type UsersRepo interface {
	FindAll() (*models.Users, error)
	FindByUsername(username string) (*models.User, error)
	FindById(uid string) (*models.User, error)
	UserExsist(username string) bool
	Add(data *models.User) (*models.User, error)
}

type UsersService interface {
	FindAll() *libs.Response
	FindByUsername(username string) *libs.Response
	GetByUUID(uid string) *libs.Response
	Add(data *models.User) *libs.Response
}
