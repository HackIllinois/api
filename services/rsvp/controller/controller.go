package controller

import (
	"encoding/json"
	"net/http"

	"github.com/HackIllinois/api/common/datastore"
	"github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/common/metrics"
	"github.com/HackIllinois/api/services/rsvp/config"
	"github.com/HackIllinois/api/services/rsvp/service"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.Handle("/internal/metrics/", promhttp.Handler()).Methods("GET")

	metrics.RegisterHandler("/filter/", GetFilteredRsvps, "GET", router)
	metrics.RegisterHandler("/{id}/", GetUserRsvp, "GET", router)
	metrics.RegisterHandler("/", GetCurrentUserRsvp, "GET", router)
	metrics.RegisterHandler("/", CreateCurrentUserRsvp, "POST", router)
	metrics.RegisterHandler("/", UpdateCurrentUserRsvp, "PUT", router)

	metrics.RegisterHandler("/internal/stats/", GetStats, "GET", router)
}

/*
	Endpoint to get the rsvp for a specified user
*/
func GetUserRsvp(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	rsvp, err := service.GetUserRsvp(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Cannot get user's RSVP status."))
		return
	}

	json.NewEncoder(w).Encode(rsvp)
}

/*
	Endpoint to get the rsvp for the current user
*/
func GetCurrentUserRsvp(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	rsvp, err := service.GetUserRsvp(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Cannot get user's RSVP status."))
		return
	}

	json.NewEncoder(w).Encode(rsvp)
}

/*
	Endpoint to create the rsvp for the current user.
	On successful creation, sends the user a confirmation mail.
*/
func CreateCurrentUserRsvp(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	if id == "" {
		errors.WriteError(w, r, errors.MalformedRequestError("Must provide id in request.", "Must provide id in request."))
		return
	}

	isAccepted, isActive, err := service.IsApplicantAcceptedAndActive(id)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not determine status of applicant decision, which is needed to create an RSVP for the user."))
		return
	}

	if !isAccepted {
		errors.WriteError(w, r, errors.AttributeMismatchError("Applicant not accepted.", "Applicant must be accepted to RSVP."))
		return
	}

	if !isActive {
		errors.WriteError(w, r, errors.AttributeMismatchError("Applicant decision has expired.", "Applicant decision has expired."))
		return
	}

	rsvp := datastore.NewDataStore(config.RSVP_DEFINITION)
	err = json.NewDecoder(r.Body).Decode(&rsvp)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not decode user rsvp information. Failure in JSON validation or incorrect rsvp definition."))
		return
	}

	rsvp.Data["id"] = id

	registration_data, err := service.GetRegistrationData(id)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not retrieve registration data."))
		return
	}

	rsvp.Data["registrationData"] = registration_data

	err = service.CreateUserRsvp(id, rsvp)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not create an RSVP for the user."))
		return
	}

	isAttending, ok := rsvp.Data["isAttending"].(bool)

	if !ok {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Failure in parsing user rsvp"))
		return
	}

	if isAttending {
		err = service.AddAttendeeRole(id)

		if err != nil {
			errors.WriteError(w, r, errors.AuthorizationError(err.Error(), "Could not add Attendee role to applicant."))
			return
		}
	}

	updated_rsvp, err := service.GetUserRsvp(id)
	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get user's RSVP."))
		return
	}

	if isAttending {
		mail_template := "rsvp_confirmation"
		err = service.SendUserMail(id, mail_template)

		if err != nil {
			errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not send user RSVP confirmation mail."))
			return
		}
	}

	json.NewEncoder(w).Encode(updated_rsvp)
}

/*
	Endpoint to update the rsvp for the current user.
	On successful update, sends the user a confirmation mail.
*/
func UpdateCurrentUserRsvp(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	if id == "" {
		errors.WriteError(w, r, errors.MalformedRequestError("Must provide id in request.", "Must provide id in the request."))
		return
	}

	isAccepted, isActive, err := service.IsApplicantAcceptedAndActive(id)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not determine if applicant was accepted and/or decision expiration status."))
		return
	}

	if !isAccepted {
		errors.WriteError(w, r, errors.AttributeMismatchError("Applicant must be accepted to modify RSVP.", "Applicant must be accepted to modify RSVP."))
		return
	}

	if !isActive {
		errors.WriteError(w, r, errors.AttributeMismatchError("Cannot modify RSVP, applicant decision has expired.", "Cannot modify RSVP, applicant decision has expired."))
		return
	}

	original_rsvp, err := service.GetUserRsvp(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get user's RSVP status."))
		return
	}

	rsvp := datastore.NewDataStore(config.RSVP_DEFINITION)
	err = json.NewDecoder(r.Body).Decode(&rsvp)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not decode user rsvp information. Failure in JSON validation or incorrect rsvp definition."))
		return
	}

	rsvp.Data["id"] = id

	registration_data, err := service.GetRegistrationData(id)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not retrieve registration data."))
		return
	}

	rsvp.Data["registrationData"] = registration_data

	err = service.UpdateUserRsvp(id, rsvp)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not update user RSVP."))
		return
	}

	wasAttending, ok := original_rsvp.Data["isAttending"].(bool)

	if !ok {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Failure in parsing user rsvp"))
		return
	}

	isAttending, ok := rsvp.Data["isAttending"].(bool)

	if !ok {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Failure in parsing user rsvp"))
		return
	}

	if !wasAttending && isAttending {
		err = service.AddAttendeeRole(id)

		if err != nil {
			errors.WriteError(w, r, errors.AuthorizationError(err.Error(), "Could not add Attendee role to user."))
			return
		}
	} else if wasAttending && !isAttending {
		err = service.RemoveAttendeeRole(id)

		if err != nil {
			errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not remove Attendee role from user."))
			return
		}
	}

	updated_rsvp, err := service.GetUserRsvp(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get updated RSVP for user."))
		return
	}

	mail_template := "rsvp_update"
	err = service.SendUserMail(id, mail_template)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not send user confirmation mail for RSVP update."))
		return
	}

	json.NewEncoder(w).Encode(updated_rsvp)
}

/*
	Endpoint to get rsvps based on filters
*/
func GetFilteredRsvps(w http.ResponseWriter, r *http.Request) {
	parameters := r.URL.Query()
	rsvps, err := service.GetFilteredRsvps(parameters)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not fetch filtered list of rsvps."))
		return
	}

	json.NewEncoder(w).Encode(rsvps)
}

/*
	Endpoint to get rsvp stats
*/
func GetStats(w http.ResponseWriter, r *http.Request) {
	stats, err := service.GetStats()

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not get RSVP service statistics."))
		return
	}

	json.NewEncoder(w).Encode(stats)
}
