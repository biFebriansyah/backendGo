package carts

import (
	"github.com/biFebriansyah/backintro/database/orm/models"
	"github.com/biFebriansyah/backintro/interfaces"
	help "github.com/biFebriansyah/backintro/libs"
)

type cart_service struct {
	re interfaces.CartRepo
}

func NewService(rep interfaces.CartRepo) *cart_service {
	return &cart_service{rep}
}

func (r *cart_service) GetByUserId(id string) *help.Response {

	data, err := r.re.FindByUserId(id)
	if err != nil {
		return help.NewRes(err.Error(), 500, true)
	}

	result := help.NewRes(data, 200, false)
	return result
}

func (r *cart_service) Get() *help.Response {

	data, err := r.re.All()
	if err != nil {
		return help.NewRes(err.Error(), 500, true)
	}

	result := help.NewRes(data, 200, false)
	return result
}

func (r cart_service) Add(usersId string, items *models.CartItem) (*help.Response, error) {

	data, err := r.re.Save(usersId, items)
	if err != nil {
		return nil, err
	}

	result := help.NewRes(data, 200, false)
	return result, nil
}
