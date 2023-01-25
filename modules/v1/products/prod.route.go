package products

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/product").Subrouter()

	repo := NewRepo(db)
	ctrl := NewCtrl(repo)

	// http://localhost:3001/prodcut/update

	route.HandleFunc("/", ctrl.Get).Methods("GET")
	route.HandleFunc("/", ctrl.Save).Methods("POST")
}
