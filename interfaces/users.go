package interfaces

import (
	"github.com/biFebriansyah/backintro/database/orm/models"
	"github.com/biFebriansyah/backintro/libs"
)

type UsersRepoIF interface {
	FindAll() (*models.Users, error)
	FindById(uid string) (*models.User, error)
	UserExsist(username string) bool
	FindByUsername(username string) (*models.User, error)
	Save(data *models.User) (*models.User, error)
}

type UsersServiceIF interface {
	GetAll() *libs.Response
	GetByUid(uid string) *libs.Response
	GetByUsername(username string) *libs.Response
	Add(data *models.User) *libs.Response
}
