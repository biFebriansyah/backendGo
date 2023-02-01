package products

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// akan memangil semua method
// inisialisasi endpoint

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/products").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("", ctrl.GetAll).Methods("GET")
	route.HandleFunc("/{id}", ctrl.GetByid).Methods("GET")
	route.HandleFunc("/{id}", ctrl.Delete).Methods("DELETE")
	route.HandleFunc("", ctrl.Add).Methods("POST")
}
