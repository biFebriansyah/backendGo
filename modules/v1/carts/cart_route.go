package carts

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/carts").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	// route.HandleFunc("", middleware.Handle(ctrl.GetAll, middleware.AuthWithRole("user"))).Methods("GET")
	route.HandleFunc("/{id}", ctrl.GetByUserId).Methods("GET")
	// route.HandleFunc("", middleware.Handle(ctrl.Create, middleware.AuthWithRole("user"))).Methods("POST")
}
