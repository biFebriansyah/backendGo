package products

import (
	"net/http"

	"github.com/biFebriansyah/backintro/database/orm/models"
	"github.com/biFebriansyah/backintro/interfaces"
	"github.com/biFebriansyah/backintro/libs"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

type prod_ctrl struct {
	svc interfaces.ProductService
}

func NewCtrl(reps interfaces.ProductService) *prod_ctrl {
	return &prod_ctrl{svc: reps}
}

func (re *prod_ctrl) GetByid(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)["id"]
	result := re.svc.GetProdWithId(vars)
	result.Send(w)
}

func (re *prod_ctrl) Delete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)["id"]
	result := re.svc.RemoveByUid(vars)
	result.Send(w)
}

func (re *prod_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	result := re.svc.GetAll()
	result.Send(w)
}

func (re *prod_ctrl) Add(w http.ResponseWriter, r *http.Request) {

	var datas models.Product
	var decoder = schema.NewDecoder()

	err := r.ParseForm()
	if err != nil {
		libs.NewRes(err, 500, true)
		return
	}

	uploads := r.Context().Value("file")
	if uploads != nil {
		datas.Image = uploads.(string)
	}

	err = decoder.Decode(&datas, r.PostForm)
	if err != nil {
		libs.NewRes(err, 500, true)
		return
	}

	re.svc.Add(&datas).Send(w)

}
