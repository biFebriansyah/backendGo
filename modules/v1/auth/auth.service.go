package auth

import (
	"github.com/biFebriansyah/backintro/database/orm/models"
	"github.com/biFebriansyah/backintro/interfaces"
	"github.com/biFebriansyah/backintro/libs"
)

type auth_service struct {
	repo interfaces.UsersRepoIF
}

type tokenRes struct {
	Token string `json:"token"`
}

func NewService(repo interfaces.UsersRepoIF) *auth_service {
	return &auth_service{repo}
}

func (s *auth_service) Login(body *models.User) *libs.Response {
	user, err := s.repo.FindByUsername(body.Username)
	if err != nil {
		return libs.NewRes("Username Belum terdaftar", 401, true)
	}

	if libs.CheckPassword(body.Password, user.Password) {
		return libs.NewRes("password salah", 401, true)
	}

	jwt := libs.NewToken(user.UserId, user.Role)
	token, err := jwt.Create()
	if err != nil {
		return libs.NewRes(err, 500, true)
	}

	return libs.NewRes(tokenRes{Token: token}, 200, false)

}
