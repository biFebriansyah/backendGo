package carts

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/biFebriansyah/backintro/database/orm/models"
	"github.com/biFebriansyah/backintro/interfaces"
	"github.com/biFebriansyah/backintro/libs"
	"github.com/gorilla/mux"
)

type cart_ctrl struct {
	repo interfaces.CartService
}

func NewCtrl(rep interfaces.CartService) *cart_ctrl {
	return &cart_ctrl{rep}
}

func (rep *cart_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	usersId := r.Context().Value("user")
	result := rep.repo.GetByUserId(usersId.(string))
	result.Send(w)
}

func (rep *cart_ctrl) GetByUserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	vars := mux.Vars(r)["id"]
	result := rep.repo.GetByUserId(vars)
	result.Send(w)
}

func (rep *cart_ctrl) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var data models.CartItem
	usersId := r.Context().Value("user")

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		libs.NewRes(err, 500, true).Send(w)
		return
	}

	result, err := rep.repo.Add(usersId.(string), &data)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	result.Send(w)
}
