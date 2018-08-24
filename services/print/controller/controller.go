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
	json.NewDecoder(r.Body).Decode(&user_checkin)

	resp, err = service.PublishPrintJob(print_job.ID, print_job.Email)
	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

  // TODO log stuff when extensive logging is implemented
	json.NewEncoder(w).Encode({})
}
