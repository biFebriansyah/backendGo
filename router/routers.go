package router

import (
	"github.com/biFebriansyah/backintro/database/orm"
	"github.com/biFebriansyah/backintro/modules/v1/auth"
	"github.com/biFebriansyah/backintro/modules/v1/products"
	"github.com/biFebriansyah/backintro/modules/v1/users"
	"github.com/gorilla/mux"
)

func NewApp() (*mux.Router, error) {
	mainRoute := mux.NewRouter()

	db, err := orm.NewDb()
	if err != nil {
		return nil, err
	}

	products.New(mainRoute, db)
	users.New(mainRoute, db)
	auth.New(mainRoute, db)

	return mainRoute, nil
}
