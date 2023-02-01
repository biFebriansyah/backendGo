package interfaces

import (
	"github.com/biFebriansyah/backintro/database/orm/models"
	"github.com/biFebriansyah/backintro/libs"
)

type Auth_ServiceIF interface {
	Login(body *models.User) *libs.Response
}
