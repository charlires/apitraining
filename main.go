package main

import (
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/rhaseven7h/apitraining/controllers"
)

func main() {
	bindAddress := "0.0.0.0:9099"

	logger := logrus.New()

	m := mux.NewRouter()

	productsController := controllers.NewProductsController(45)
	m.HandleFunc("/products", productsController.List).Methods("GET")
	m.HandleFunc("/products/{id}", productsController.Get).Methods("GET")

	servicesController := controllers.NewServicesController("ooyala!", logger)
	m.HandleFunc("/services", servicesController.List)
	m.HandleFunc("/services/{id}", servicesController.Get)

	fmt.Printf("Server started, listening at %s\n", bindAddress)
	http.ListenAndServe(bindAddress, m)
}
