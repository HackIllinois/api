package controller

import (
	"encoding/json"
	"github.com/HackIllinois/api/services/print/errors"
	"github.com/HackIllinois/api/services/print/models"
	"github.com/HackIllinois/api/services/print/service"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"net/http"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.Handle("/", alice.New().ThenFunc(CreatePrintJob)).Methods("POST")
}

/*
	Endpoint to send a print job to the sns topic
*/
func CreatePrintJob(w http.ResponseWriter, r *http.Request) {
	var print_job models.PrintJob
	json.NewDecoder(r.Body).Decode(&print_job)

	resp, err := service.PublishPrintJob(print_job)
	if err != nil {
		panic(errors.PrintError(err.Error()))
	}

  // TODO log stuff when extensive logging is implemented
	w.WriteHeader(http.StatusOK)
}
