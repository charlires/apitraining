package main

import (
	"fmt"
	"net/http"
	"regexp"
)

type ProductsHandler struct{}

func (ph ProductsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "In products!")
}

type ServicesHandler struct{}

func (ph ServicesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "In services!")
}

type GenericHandler struct{}

func (gh GenericHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pathstr := r.URL.Path

	ph := &ProductsHandler{}
	sh := &ServicesHandler{}

	ok, err := regexp.MatchString("product", pathstr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error matching routes at products: %s", err.Error())
		return
	}

	if ok {
		ph.ServeHTTP(w, r)
		return
	}

	ok, err = regexp.MatchString("service", pathstr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error matching routes at services: %s", err.Error())
		return
	}

	if ok {
		sh.ServeHTTP(w, r)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "no route found for: %s", pathstr)
}

func main() {
	bindAddress := "0.0.0.0:8080"
	gh := &GenericHandler{}
	fmt.Printf("Server started, listening at %s\n", bindAddress)
	http.ListenAndServe(bindAddress, gh)
}
