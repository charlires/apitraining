package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rhaseven7h/apitraining/loggingadapter"
	//"github.com/Sirupsen/logrus"
)

type ServicesHandler struct {
	MyServiceID string
	Logger loggingadapter.OOLogger
}

func NewServicesController(myServiceID string, logger loggingadapter.OOLogger) *ServicesHandler  {
	logger.
	WithField("input_id", myServiceID).
		Info("creating services controller")
	return &ServicesHandler{
		MyServiceID: myServiceID,
		Logger: logger,
	}
}

//func NewServicesControllerMock(myServiceID string, logger loggingadapter.OOLogger) *ServicesHandler {
//	logger.
//		WithField("input_id", myServiceID).
//		Info("creating services controller")
//	return &ServicesHandler{
//		MyServiceID: myServiceID,
//		Logger: logger,
//	}
//}

func (ph ServicesHandler) List(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "List of services!")
}

func (ph ServicesHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ph.Logger.
			WithField("id_received", vars["id"]).
			Error("error occurred with parameters")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, fmt.Sprintf("The ID requested was not an int: %s", vars["id"]))
		return
	}
	ph.Logger.
		WithField("id_received", id).
		Info("error occurred with parameters")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, fmt.Sprintf("The ID requested was: %d", id))
}
