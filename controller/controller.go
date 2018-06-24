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

	// Masks the decision if not finalized
	if !decision.Finalized {
		decision.Status = "PENDING"
	}

	json.NewEncoder(w).Encode(decision)
}

/*
	Endpoint to update the decision for the specified user.
	If the existing decision is finalized, an error is reported.
*/
func UpdateDecision(w http.ResponseWriter, r *http.Request) {
	var decision models.Decision
	json.NewDecoder(r.Body).Decode(&decision)

	if decision.ID == "" {
		panic(errors.UnprocessableError("Must provide id parameter."))
	}
// Only affects the response, not status of the actual decision.
	existing_decision_history, err := service.GetDecision(decision.ID)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	if existing_decision_history.Finalized {
		panic(errors.UnprocessableError("Cannot modify finalized decisions."))
	}

	reviewer_id := r.Header.Get("HackIllinois-Identity")
	decision.Reviewer = reviewer_id
	decision.Timestamp = time.Now().Unix()
	// Finalized is always false, unless explicitly set to true via the appropriate endpoint.
	decision.Finalized = false

	err = service.UpdateDecision(decision.ID, decision)

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
	var decision_finalized models.DecisionFinalized
	json.NewDecoder(r.Body).Decode(&decision_finalized)
	
	id := decision_finalized.ID

	if id == "" {
		panic(errors.UnprocessableError("Must provide id parameter to retrieve current decision	"))
	}

	// Assuming we are working on the current user's decision 
	existing_decision_history, err := service.GetDecision(id)

	// If the decision is NOT already finalized, set it to what was provided in the request body
	if !existing_decision_history.Finalized {
		var latest_decision models.Decision
		latest_decision.Finalized = decision_finalized.Finalized
		latest_decision.ID = id
		latest_decision.Status = existing_decision_history.Status
		latest_decision.Wave = existing_decision_history.Wave
		latest_decision.Reviewer = r.Header.Get("HackIllinois-Identity")
		latest_decision.Timestamp = time.Now().Unix()

		err := service.UpdateDecision(id, latest_decision)
		
		if err != nil {
			panic(errors.UnprocessableError("Error updating the decision, in an attempt to finalize it."))
		}
	} else {
		panic(errors.UnprocessableError("Decision already finalized."))
	}

	updated_decision, err := service.GetDecision(id)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(updated_decision)
}