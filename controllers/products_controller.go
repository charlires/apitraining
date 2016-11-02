package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type ProductsController struct {
	MyValue int
}

func NewProductsController(myValue int) *ProductsController {
	return &ProductsController{
		MyValue: myValue,
	}
}

func (ph ProductsController) List(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "List of products!")
}

func (ph ProductsController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, fmt.Sprintf("The ID requested was: %s", id))
}
