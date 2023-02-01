package users

import (
	"errors"

	"github.com/biFebriansyah/backintro/database/orm/models"
	"gorm.io/gorm"
)

type user_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *user_repo {
	return &user_repo{db}
}

func (r *user_repo) FindAll() (*models.Users, error) {

	var data models.Users

	result := r.db.Find(&data)

	if result.Error != nil {
		return nil, errors.New("gagal mengambil data")
	}

	return &data, nil
}

func (r *user_repo) FindById(uid string) (*models.User, error) {

	var data models.User
	var field = []string{"user_id", "username", "email"}

	result := r.db.Select(field).First(&data, "user_id = ?", uid)

	if result.Error != nil {
		return nil, errors.New("gagal mengambil data")
	}

	return &data, nil
}

func (r *user_repo) FindByUsername(username string) (*models.User, error) {

	var data models.User

	result := r.db.First(&data, "username = ?", username)

	if result.Error != nil {
		return nil, errors.New("gagal mengambil data")
	}

	return &data, nil
}

func (r *user_repo) UserExsist(username string) bool {
	var data models.User
	result := r.db.First(&data, "username = ?", username)

	if result.Error != nil {
		return false
	}

	return true
}

func (r *user_repo) Save(data *models.User) (*models.User, error) {

	result := r.db.Create(data)

	if result.Error != nil {
		return nil, errors.New("gagal mengambil data")
	}

	data.Password = ""

	return data, nil
}
