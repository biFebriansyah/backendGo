package users

import (
	"encoding/json"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/biFebriansyah/backintro/database/orm/models"
	"github.com/biFebriansyah/backintro/interfaces"
	"github.com/biFebriansyah/backintro/libs"
)

// khusus untuk request dan respone
type user_ctrl struct {
	repo interfaces.UsersServiceIF
}

func NewCtrl(repo interfaces.UsersServiceIF) *user_ctrl {
	return &user_ctrl{repo}
}

func (re *user_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	result := re.repo.GetAll()
	result.Send(w)
	return

}

func (re *user_ctrl) GetById(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user")
	re.repo.GetByUid(user_id.(string)).Send(w)
}

func (re *user_ctrl) Add(w http.ResponseWriter, r *http.Request) {
	var datas models.User
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		libs.NewRes(err.Error(), 500, true).Send(w)
		return
	}

	_, err = govalidator.ValidateStruct(&datas)
	if err != nil {
		libs.NewRes(err, 500, true).Send(w)
		return
	}

	re.repo.Add(&datas).Send(w)

}
