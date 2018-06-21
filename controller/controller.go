package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/HackIllinois/api-commons/errors"
	"github.com/HackIllinois/api-decision/models"
	"github.com/HackIllinois/api-decision/service"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.Handle("/{id}/", alice.New().ThenFunc(GetDecision)).Methods("GET")
	router.Handle("/", alice.New().ThenFunc(GetCurrentDecision)).Methods("GET")
	router.Handle("/", alice.New().ThenFunc(UpdateDecision)).Methods("POST")
	router.Handle("/finalize/", alice.New().ThenFunc(FinalizeDecision)).Methods("POST")
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
	Endpoint to update the decision for the specified user.
	If the existing decision is finalized, then it is not updated.
*/
func UpdateDecision(w http.ResponseWriter, r *http.Request) {
	var decision models.Decision
	json.NewDecoder(r.Body).Decode(&decision)

	if decision.ID == "" {
		panic(errors.UnprocessableError("Must provide id parameter"))
	}

	existing_decision, err := service.GetDecision(decision.ID)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	if existing_decision.Finalized {
		json.NewEncoder(w).Encode(existing_decision)
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
	Finalizes the decision associated with the provided ID.
	Finalized decisions are blocked from further review.
*/
func FinalizeDecision(w http.ResponseWriter, r *http.Request) {
	var existing_decision models.Decision
	json.NewDecoder(r.Body).Decode(&existing_decision)

	if existing_decision.ID == "" {
		panic(errors.UnprocessableError("Must provide id parameter"))
	}

	reviewer_id := r.Header.Get("HackIllinois-Identity")
	existing_decision.Reviewer = reviewer_id
	existing_decision.Finalized = true
	existing_decision.Timestamp = time.Now().Unix()

	err := service.UpdateDecision(existing_decision.ID, existing_decision)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	updated_decision, err := service.GetDecision(existing_decision.ID)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(updated_decision)
}
