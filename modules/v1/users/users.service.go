package users

import (
	"github.com/biFebriansyah/backintro/database/orm/models"
	"github.com/biFebriansyah/backintro/interfaces"
	"github.com/biFebriansyah/backintro/libs"
)

// main logika

type user_service struct {
	repo interfaces.UsersRepoIF
}

func NewService(repo interfaces.UsersRepoIF) *user_service {
	return &user_service{repo}
}

func (s *user_service) GetAll() *libs.Response {
	data, err := s.repo.FindAll()
	if err != nil {
		return libs.NewRes(err.Error(), 400, true)
	}

	return libs.NewRes(data, 200, false)
}

func (s *user_service) GetByUid(uid string) *libs.Response {
	data, err := s.repo.FindById(uid)
	if err != nil {
		return libs.NewRes(err.Error(), 400, true)
	}

	return libs.NewRes(data, 200, false)
}

func (s *user_service) GetByUsername(username string) *libs.Response {
	data, err := s.repo.FindByUsername(username)
	if err != nil {
		return libs.NewRes(err.Error(), 400, true)
	}

	return libs.NewRes(data, 200, false)
}

func (s *user_service) Add(data *models.User) *libs.Response {
	HasPass, err := libs.HashPassword(data.Password)
	if err != nil {
		return libs.NewRes(err.Error(), 400, true)
	}

	data.Password = HasPass
	data, err = s.repo.Save(data)
	if err != nil {
		return libs.NewRes(err.Error(), 400, true)
	}

	return libs.NewRes(data, 200, false)
}
