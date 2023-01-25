package products

import (
	"encoding/json"
	"net/http"

	"github.com/biFebriansyah/backintro/database/orm/model"
)

type prod_ctrl struct {
	repo *prod_repo
}

func NewCtrl(repo *prod_repo) *prod_ctrl {
	return &prod_ctrl{repo}
}

func (c *prod_ctrl) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	respone, err := c.repo.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(respone)
}

func (c *prod_ctrl) Save(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var data model.Product
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	respone, err := c.repo.Add(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(respone)
}
