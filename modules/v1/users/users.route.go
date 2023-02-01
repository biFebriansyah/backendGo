package users

import (
	"github.com/biFebriansyah/backintro/middleware"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/users").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	// midle.auth(middle.cahce(ctrl.GetAll)) ->

	// midle.auth(middle.cahce(ctrl.GetAll))

	route.HandleFunc("", middleware.Handle(ctrl.GetById, middleware.AuthMidle("user"))).Methods("GET")
	route.HandleFunc("", ctrl.Add).Methods("POST")
}
