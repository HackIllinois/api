package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/services/decision/models"
	"github.com/HackIllinois/api/services/decision/service"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.Handle("/", alice.New().ThenFunc(GetCurrentDecision)).Methods("GET")
	router.Handle("/", alice.New().ThenFunc(UpdateDecision)).Methods("POST")
	router.Handle("/finalize/", alice.New().ThenFunc(FinalizeDecision)).Methods("POST")
	router.Handle("/filter/", alice.New().ThenFunc(GetFilteredDecisions)).Methods("GET")
	router.Handle("/{id}/", alice.New().ThenFunc(GetDecision)).Methods("GET")

	router.Handle("/internal/stats/", alice.New().ThenFunc(GetStats)).Methods("GET")
}

/*
	Endpoint to get the decision for the current user
*/
func GetCurrentDecision(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	decision, err := service.GetDecision(id)

	if err != nil {
		panic(errors.DATABASE_ERROR("Could not get current user's decision."))
	}

	decision_view := models.DecisionView{
		ID:     decision.ID,
		Status: decision.Status,
	}

	// Masks the decision if not finalized
	if !decision.Finalized {
		decision_view.Status = "PENDING"
	}

	json.NewEncoder(w).Encode(decision_view)
}

/*
	Endpoint to update the decision for the specified user.
	If the existing decision is finalized, an error is reported.
*/
func UpdateDecision(w http.ResponseWriter, r *http.Request) {
	var decision models.Decision
	json.NewDecoder(r.Body).Decode(&decision)

	if decision.ID == "" {
		panic(errors.MALFORMED_REQUEST_ERROR("Must provide ID parameter in request."))
	}

	has_decision, err := service.HasDecision(decision.ID)

	if err != nil {
		panic(errors.DATABASE_ERROR("Could not determine user's decision."))
	}

	if has_decision {
		existing_decision_history, err := service.GetDecision(decision.ID)

		if err != nil {
			panic(errors.DATABASE_ERROR("Could not get current user's existing decision history."))
		}

		if existing_decision_history.Finalized {
			panic(errors.ATTRIBUTE_MISMATCH_ERROR("Cannot modify finalized decisions."))
		}
	}

	decision.Reviewer = r.Header.Get("HackIllinois-Identity")
	decision.Timestamp = time.Now().Unix()
	// Finalized is always false, unless explicitly set to true via the appropriate endpoint.
	decision.Finalized = false

	err = service.UpdateDecision(decision.ID, decision)

	if err != nil {
		panic(errors.INTERNAL_ERROR("Could not update decision."))
	}

	updated_decision, err := service.GetDecision(decision.ID)

	if err != nil {
		panic(errors.DATABASE_ERROR("Could not fetch updated decision."))
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
		panic(errors.MALFORMED_REQUEST_ERROR("Must provide id parameter to retrieve current decision."))
	}

	// Assuming we are working on the specified user's decision
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
			panic(errors.INTERNAL_ERROR("Error updating the decision, in an attempt to finalize it."))
		}
	} else {
		panic(errors.ATTRIBUTE_MISMATCH_ERROR("Decision already finalized."))
	}

	updated_decision, err := service.GetDecision(id)

	if err != nil {
		panic(errors.DATABASE_ERROR("Could not fetch updated decision."))
	}

	if updated_decision.Finalized {
		err = service.AddUserToMailList(id, updated_decision)

		if err != nil {
			panic(errors.INTERNAL_ERROR("Could not add user to mail list."))
		}
	}

	json.NewEncoder(w).Encode(updated_decision)
}

/*
	Endpoint to get decisions based on a filter
*/
func GetFilteredDecisions(w http.ResponseWriter, r *http.Request) {
	parameters := r.URL.Query()
	decisions, err := service.GetFilteredDecisions(parameters)

	if err != nil {
		panic(errors.DATABASE_ERROR("Could not retrieve filtered decisions."))
	}

	json.NewEncoder(w).Encode(decisions)
}

/*
	Endpoint to get the decision for the specified user
*/
func GetDecision(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	decision, err := service.GetDecision(id)

	if err != nil {
		panic(errors.DATABASE_ERROR("Could not get decision for the specified user."))
	}

	json.NewEncoder(w).Encode(decision)
}

/*
	Endpoint to get decision stats
*/
func GetStats(w http.ResponseWriter, r *http.Request) {
	stats, err := service.GetStats()

	if err != nil {
		panic(errors.INTERNAL_ERROR("Could not get decision service statistics."))
	}

	json.NewEncoder(w).Encode(stats)
}
