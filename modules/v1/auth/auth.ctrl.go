package auth

import (
	"encoding/json"
	"net/http"

	"github.com/biFebriansyah/backintro/database/orm/models"
	"github.com/biFebriansyah/backintro/interfaces"
	"github.com/biFebriansyah/backintro/libs"
)

type auth_ctrl struct {
	repo interfaces.Auth_ServiceIF
}

func NewCtrl(repo interfaces.Auth_ServiceIF) *auth_ctrl {
	return &auth_ctrl{repo}
}

func (c *auth_ctrl) Signin(w http.ResponseWriter, r *http.Request) {
	var data models.User
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		libs.NewRes(err.Error(), 500, true)
		return
	}

	c.repo.Login(&data).Send(w)

}
