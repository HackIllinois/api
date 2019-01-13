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
		panic(errors.DatabaseError(err.Error(), "Could not get current user's decision."))
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
		panic(errors.MalformedRequestError("Must provide id parameter in request.", "Must provide id parameter in request."))
	}

	has_decision, err := service.HasDecision(decision.ID)

	if err != nil {
		panic(errors.DatabaseError(err.Error(), "Could not determine user's decision."))
	}

	if has_decision {
		existing_decision_history, err := service.GetDecision(decision.ID)

		if err != nil {
			panic(errors.DatabaseError(err.Error(), "Could not get current user's existing decision history."))
		}

		if existing_decision_history.Finalized {
			panic(errors.AttributeMismatchError("Cannot modify finalized decisions.", "Cannot modify finalized decisions."))
		}
	}

	decision.Reviewer = r.Header.Get("HackIllinois-Identity")
	decision.Timestamp = time.Now().Unix()
	// Finalized is always false, unless explicitly set to true via the appropriate endpoint.
	decision.Finalized = false

	err = service.UpdateDecision(decision.ID, decision)

	if err != nil {
		panic(errors.InternalError(err.Error(), "Could not update decision."))
	}

	updated_decision, err := service.GetDecision(decision.ID)

	if err != nil {
		panic(errors.DatabaseError(err.Error(), "Could not fetch updated decision."))
	}

	json.NewEncoder(w).Encode(updated_decision)
}

/*
	Finalizes / unfinalizes the decision associated with the provided ID.
	Finalized decisions are blocked from further review, unless unfinalized.
*/
func FinalizeDecision(w http.ResponseWriter, r *http.Request) {
	var decision_finalized models.DecisionFinalized
	json.NewDecoder(r.Body).Decode(&decision_finalized)

	id := decision_finalized.ID

	if id == "" {
		panic(errors.MalformedRequestError("Must provide id parameter to retrieve current decision.", "Must provide id parameter to retrieve current decision."))
	}

	// Assuming we are working on the specified user's decision
	existing_decision_history, err := service.GetDecision(id)

	// It is an error to finalize a finalized decision, or unfinalize an unfinalized decision.
	if existing_decision_history.Finalized == decision_finalized.Finalized {
		panic(errors.AttributeMismatchError("Superfluous request. Existing decision already at desired state of finalization.", "Superfluous request. Existing decision already at desired state of finalization."))
	}

	var latest_decision models.Decision
	latest_decision.Finalized = decision_finalized.Finalized
	latest_decision.ID = id
	latest_decision.Status = existing_decision_history.Status
	latest_decision.Wave = existing_decision_history.Wave
	latest_decision.Reviewer = r.Header.Get("HackIllinois-Identity")
	latest_decision.Timestamp = time.Now().Unix()

	err = service.UpdateDecision(id, latest_decision)

	if err != nil {
		panic(errors.InternalError(err.Error(), "Error updating the decision, in an attempt to alter its finalized status."))
	}

	updated_decision, err := service.GetDecision(id)

	if err != nil {
		panic(errors.DatabaseError(err.Error(), "Could not fetch updated decision."))
	}

	if updated_decision.Finalized {
		err = service.AddUserToMailList(id, updated_decision)

		if err != nil {
			panic(errors.InternalError(err.Error(), "Could not add user to mail list."))
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
		panic(errors.DatabaseError(err.Error(), "Could not retrieve filtered decisions."))
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
		panic(errors.DatabaseError(err.Error(), "Could not get decision for the specified user."))
	}

	json.NewEncoder(w).Encode(decision)
}

/*
	Endpoint to get decision stats
*/
func GetStats(w http.ResponseWriter, r *http.Request) {
	stats, err := service.GetStats()

	if err != nil {
		panic(errors.InternalError(err.Error(), "Could not get decision service statistics."))
	}

	json.NewEncoder(w).Encode(stats)
}
