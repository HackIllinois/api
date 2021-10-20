package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/common/utils"
	"github.com/HackIllinois/api/services/decision/config"
	"github.com/HackIllinois/api/services/decision/models"
	"github.com/HackIllinois/api/services/decision/service"
	"github.com/gorilla/mux"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.HandleFunc("/", GetCurrentDecision).Methods("GET")
	router.HandleFunc("/", UpdateDecision).Methods("POST")
	router.HandleFunc("/finalize/", FinalizeDecision).Methods("POST")
	router.HandleFunc("/finalize/batch/", FinalizeDecisionBatch).Methods("POST")
	router.HandleFunc("/filter/", GetFilteredDecisions).Methods("GET")
	router.HandleFunc("/batch/", UpdateDecisionBatch).Methods("POST")
	router.HandleFunc("/{id}/", GetDecision).Methods("GET")

	router.HandleFunc("/internal/stats/", GetStats).Methods("GET")
}

/*
	Endpoint to get the decision for the current user
*/
func GetCurrentDecision(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	decision, err := service.GetDecision(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get current user's decision."))
		return
	}

	decision_view := models.DecisionView{
		ID: decision.ID,
	}

	// Expose the decision if it has been finalized
	if decision.Finalized {
		decision_view.Status = decision.Status

		// Expose the expiration only for finalized ACCEPTED decisions
		if decision.Status == "ACCEPTED" {
			decision_view.ExpiresAt = decision.ExpiresAt
		}
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
		errors.WriteError(w, r, errors.MalformedRequestError("Must provide id parameter in request.", "Must provide id parameter in request."))
		return
	}

	has_decision, err := service.HasDecision(decision.ID)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not determine user's decision."))
		return
	}

	if has_decision {
		existing_decision_history, err := service.GetDecision(decision.ID)

		if err != nil {
			errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get current user's existing decision history."))
			return
		}

		if existing_decision_history.Finalized {
			errors.WriteError(w, r, errors.AttributeMismatchError("Cannot modify finalized decisions.", "Cannot modify finalized decisions."))
			return
		}
	}

	decision.Reviewer = r.Header.Get("HackIllinois-Identity")
	decision.Timestamp = time.Now().Unix()
	decision.ExpiresAt = decision.Timestamp + utils.HoursToUnixSeconds(config.DECISION_EXPIRATION_HOURS)
	// Finalized is always false, unless explicitly set to true via the appropriate endpoint.
	decision.Finalized = false

	err = service.UpdateDecision(decision.ID, decision)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not update decision."))
		return
	}

	updated_decision, err := service.GetDecision(decision.ID)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not fetch updated decision."))
		return
	}

	json.NewEncoder(w).Encode(updated_decision)
}

/*
	Endpoint to update the decisions for specified users.
	If any of the existing decisions are already finalized, an error is reported.
*/
func UpdateDecisionBatch(w http.ResponseWriter, r *http.Request) {
	var decisions models.Decisions
	json.NewDecoder(r.Body).Decode(&decisions)

	var updated_decisions models.FilteredDecisions

	// Check if any of the decisions have been finalized already. If so, fail early.
	for _, decision := range decisions.Decisions {
		if decision.ID == "" {
			errors.WriteError(w, r, errors.MalformedRequestError("Must provide id parameter in request.", "Must provide id parameter in request."))
			return
		}

		has_decision, err := service.HasDecision(decision.ID)

		if err != nil {
			errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not determine user's decision."))
			return
		}

		if has_decision {
			existing_decision_history, err := service.GetDecision(decision.ID)

			if err != nil {
				errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get current user's existing decision history."))
				return
			}

			if existing_decision_history.Finalized {
				errors.WriteError(w, r, errors.AttributeMismatchError("Cannot modify finalized decision for user ID "+decision.ID, "Cannot modify finalized decision for user ID "+decision.ID))
				return
			}
		}
	}

	for _, decision := range decisions.Decisions {

		decision.Reviewer = r.Header.Get("HackIllinois-Identity")
		decision.Timestamp = time.Now().Unix()
		decision.ExpiresAt = decision.Timestamp + utils.HoursToUnixSeconds(config.DECISION_EXPIRATION_HOURS)
		// Finalized is always false, unless explicitly set to true via the appropriate endpoint.
		decision.Finalized = false

		err := service.UpdateDecision(decision.ID, decision)

		if err != nil {
			errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not update decision."))
			return
		}

		updated_decision, err := service.GetDecision(decision.ID)

		if err != nil {
			errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not fetch updated decision."))
			return
		}

		updated_decisions.Decisions = append(updated_decisions.Decisions, *updated_decision)
	}

	json.NewEncoder(w).Encode(updated_decisions)
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
		errors.WriteError(w, r, errors.MalformedRequestError("Must provide id parameter to retrieve current decision.", "Must provide id parameter to retrieve current decision."))
		return
	}

	// Assuming we are working on the specified user's decision
	existing_decision_history, err := service.GetDecision(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get current user's existing decision history."))
		return
	}

	// It is an error to finalize a finalized decision, or unfinalize an unfinalized decision.
	if existing_decision_history.Finalized == decision_finalized.Finalized {
		errors.WriteError(w, r, errors.AttributeMismatchError("Superfluous request. Existing decision already at desired state of finalization.", "Superfluous request. Existing decision already at desired state of finalization."))
		return
	}

	var latest_decision models.Decision
	latest_decision.Finalized = decision_finalized.Finalized
	latest_decision.ID = id
	latest_decision.Status = existing_decision_history.Status
	latest_decision.Wave = existing_decision_history.Wave
	latest_decision.Reviewer = r.Header.Get("HackIllinois-Identity")
	latest_decision.Timestamp = time.Now().Unix()
	latest_decision.ExpiresAt = latest_decision.Timestamp + utils.HoursToUnixSeconds(config.DECISION_EXPIRATION_HOURS)

	err = service.UpdateDecision(id, latest_decision)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Error updating the decision, in an attempt to alter its finalized status."))
		return
	}

	updated_decision, err := service.GetDecision(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not fetch updated decision."))
		return
	}

	if updated_decision.Finalized {
		err = service.AddUserToMailList(id, updated_decision)

		if err != nil {
			errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not add user to mail list."))
			return
		}
	} else {
		err = service.RemoveUserFromMailList(id, updated_decision)

		if err != nil {
			errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not remove user from mail list."))
			return
		}
	}

	json.NewEncoder(w).Encode(updated_decision)
}

func FinalizeDecisionBatch(w http.ResponseWriter, r *http.Request) {
	var decisions_finalized models.DecisionsFinalized // decisions which need to be finalized
	json.NewDecoder(r.Body).Decode(&decisions_finalized)

	var decisions models.FilteredDecisions // decisions which have been finalized

	// Check if any of the IDs are missing, or if any of the requests are superfluous. If so, fail early.
	for _, decision_finalized := range decisions_finalized.DecisionsFinalized {
		if decision_finalized.ID == "" {
			errors.WriteError(w, r, errors.MalformedRequestError("Must provide id parameter in request.", "Must provide id parameter in request."))
			return
		}

		// Assuming we are working on the specified user's decision
		existing_decision_history, err := service.GetDecision(decision_finalized.ID)

		if err != nil {
			errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get current user's existing decision history."))
			return
		}

		// It is an error to finalize a finalized decision, or unfinalize an unfinalized decision.
		if existing_decision_history.Finalized == decision_finalized.Finalized {
			errors.WriteError(w, r, errors.AttributeMismatchError("Superfluous request. Existing decision already at desired state of finalization.", "Superfluous request. Existing decision already at desired state of finalization."))
			return
		}
	}

	for _, decision_finalized := range decisions_finalized.DecisionsFinalized {
		id := decision_finalized.ID

		existing_decision_history, err := service.GetDecision(id)

		if err != nil {
			errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get current user's existing decision history."))
			return
		}

		var latest_decision models.Decision
		latest_decision.Finalized = decision_finalized.Finalized
		latest_decision.ID = decision_finalized.ID
		latest_decision.Status = existing_decision_history.Status
		latest_decision.Wave = existing_decision_history.Wave
		latest_decision.Reviewer = r.Header.Get("HackIllinois-Identity")
		latest_decision.Timestamp = time.Now().Unix()
		latest_decision.ExpiresAt = latest_decision.Timestamp + utils.HoursToUnixSeconds(config.DECISION_EXPIRATION_HOURS)

		err = service.UpdateDecision(id, latest_decision)

		if err != nil {
			errors.WriteError(w, r, errors.InternalError(err.Error(), "Error updating the decision, in an attempt to alter its finalized status."))
			return
		}

		updated_decision, err := service.GetDecision(id)

		if err != nil {
			errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not fetch updated decision."))
			return
		}

		if updated_decision.Finalized {
			err = service.AddUserToMailList(id, updated_decision)

			if err != nil {
				errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not add user to mail list."))
				return
			}
		} else {
			err = service.RemoveUserFromMailList(id, updated_decision)

			if err != nil {
				errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not remove user from mail list."))
				return
			}
		}

		decisions.Decisions = append(decisions.Decisions, *updated_decision)
	}

	json.NewEncoder(w).Encode(decisions)
}

/*
	Endpoint to get decisions based on a filter
*/
func GetFilteredDecisions(w http.ResponseWriter, r *http.Request) {
	parameters := r.URL.Query()
	decisions, err := service.GetFilteredDecisions(parameters)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not retrieve filtered decisions."))
		return
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
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get decision for the specified user."))
		return
	}

	json.NewEncoder(w).Encode(decision)
}

/*
	Endpoint to get decision stats
*/
func GetStats(w http.ResponseWriter, r *http.Request) {
	stats, err := service.GetStats()

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not get decision service statistics."))
		return
	}

	json.NewEncoder(w).Encode(stats)
}
