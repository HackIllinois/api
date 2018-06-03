package controller

import (
	"encoding/json"
	"github.com/HackIllinois/api-commons/errors"
	"github.com/HackIllinois/api-decision/models"
	"github.com/HackIllinois/api-decision/service"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"net/http"
	"time"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.Handle("/", alice.New().ThenFunc(GetCurrentDecision)).Methods("GET")
	router.Handle("/", alice.New().ThenFunc(UpdateDecision)).Methods("POST")
	router.Handle("/filter/", alice.New().ThenFunc(GetFilteredDecisions)).Methods("GET")
	router.Handle("/{id}/", alice.New().ThenFunc(GetDecision)).Methods("GET")
}

/*
	Endpoint to get the decision for the current user
*/
func GetCurrentDecision(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	decision, err := service.GetDecision(id)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(decision)
}

/*
	Endpoint to update the decision for the specified user
*/
func UpdateDecision(w http.ResponseWriter, r *http.Request) {
	var decision models.Decision
	json.NewDecoder(r.Body).Decode(&decision)

	if decision.ID == "" {
		panic(errors.UnprocessableError("Must provide id parameter"))
	}

	reviewer_id := r.Header.Get("HackIllinois-Identity")
	decision.Reviewer = reviewer_id
	decision.Timestamp = time.Now().Unix()

	err := service.UpdateDecision(decision.ID, decision)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	updated_decision, err := service.GetDecision(decision.ID)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(updated_decision)
}

/*
	Endpoint to get decisions based on a filter
*/
func GetFilteredDecisions(w http.ResponseWriter, r *http.Request) {
	
}

/*
	Endpoint to get the decision for the specified user
*/
func GetDecision(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	decision, err := service.GetDecision(id)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(decision)
}