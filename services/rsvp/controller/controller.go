package controller

import (
	"encoding/json"
	"github.com/HackIllinois/api/common/datastore"
	"github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/services/rsvp/config"
	"github.com/HackIllinois/api/services/rsvp/service"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"net/http"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.Handle("/{id}/", alice.New().ThenFunc(GetUserRsvp)).Methods("GET")
	router.Handle("/", alice.New().ThenFunc(GetCurrentUserRsvp)).Methods("GET")
	router.Handle("/", alice.New().ThenFunc(CreateCurrentUserRsvp)).Methods("POST")
	router.Handle("/", alice.New().ThenFunc(UpdateCurrentUserRsvp)).Methods("PUT")

	router.Handle("/internal/stats/", alice.New().ThenFunc(GetStats)).Methods("GET")
}

/*
	Endpoint to get the rsvp for a specified user
*/
func GetUserRsvp(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	rsvp, err := service.GetUserRsvp(id)

	if err != nil {
		panic(errors.DatabaseError(err.Error(), "Cannot get user's RSVP status."))
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
		panic(errors.DatabaseError(err.Error(), "Cannot get user's RSVP status."))
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
		panic(errors.MalformedRequestError("Must provide id in request.", "Must provide id in request."))
	}

	isAccepted, isActive, err := service.IsApplicantAcceptedAndActive(id)

	if err != nil {
		panic(errors.InternalError(err.Error(), "Could not determine status of applicant decision, which is needed to create an RSVP for the user."))
	}

	if !isAccepted {
		panic(errors.AttributeMismatchError("Applicant not accepted.", "Applicant must be accepted to RSVP."))
	}

	if !isActive {
		panic(errors.AttributeMismatchError("Applicant decision has expired.", "Applicant decision has expired."))
	}

	rsvp := datastore.NewDataStore(config.RSVP_DEFINITION)
	err = json.NewDecoder(r.Body).Decode(&rsvp)

	if err != nil {
		panic(errors.InternalError(err.Error(), "Could not decode user rsvp information. Failure in JSON validation or incorrect rsvp definition."))
	}

	rsvp.Data["id"] = id

	err = service.CreateUserRsvp(id, rsvp)

	if err != nil {
		panic(errors.InternalError(err.Error(), "Could not create an RSVP for the user."))
	}

	isAttending, ok := rsvp.Data["isAttending"].(bool)

	if !ok {
		panic(errors.InternalError(err.Error(), "Failure in parsing user rsvp"))
	}

	if isAttending {
		err = service.AddAttendeeRole(id)

		if err != nil {
			panic(errors.AuthorizationError(err.Error(), "Could not add Attendee role to applicant."))
		}
	}

	updated_rsvp, err := service.GetUserRsvp(id)

	if err != nil {
		panic(errors.DatabaseError(err.Error(), "Could not get user's RSVP."))
	}

	mail_template := "rsvp_confirmation"
	err = service.SendUserMail(id, mail_template)

	if err != nil {
		panic(errors.InternalError(err.Error(), "Could not send user RSVP confirmation mail."))
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
		panic(errors.MalformedRequestError("Must provide id in request.", "Must provide id in the request."))
	}

	isAccepted, isActive, err := service.IsApplicantAcceptedAndActive(id)

	if err != nil {
		panic(errors.InternalError(err.Error(), "Could not determine if applicant was accepted and/or decision expiration status."))
	}

	if !isAccepted {
		panic(errors.AttributeMismatchError("Applicant must be accepted to modify RSVP.", "Applicant must be accepted to modify RSVP."))
	}

	if !isActive {
		panic(errors.AttributeMismatchError("Cannot modify RSVP, applicant decision has expired.", "Cannot modify RSVP, applicant decision has expired."))
	}

	original_rsvp, err := service.GetUserRsvp(id)

	if err != nil {
		panic(errors.DatabaseError(err.Error(), "Could not get user's RSVP status."))
	}

	rsvp := datastore.NewDataStore(config.RSVP_DEFINITION)
	err = json.NewDecoder(r.Body).Decode(&rsvp)

	if err != nil {
		panic(errors.InternalError(err.Error(), "Could not decode user rsvp information. Failure in JSON validation or incorrect rsvp definition."))
	}

	rsvp.Data["id"] = id

	err = service.UpdateUserRsvp(id, rsvp)

	if err != nil {
		panic(errors.DatabaseError(err.Error(), "Could not update user RSVP."))
	}

	wasAttending, ok := original_rsvp.Data["isAttending"].(bool)

	if !ok {
		panic(errors.InternalError(err.Error(), "Failure in parsing user rsvp"))
	}

	isAttending, ok := rsvp.Data["isAttending"].(bool)

	if !ok {
		panic(errors.InternalError(err.Error(), "Failure in parsing user rsvp"))
	}

	if !wasAttending && isAttending {
		err = service.AddAttendeeRole(id)

		if err != nil {
			panic(errors.AuthorizationError(err.Error(), "Could not add Attendee role to user."))
		}
	} else if wasAttending && !isAttending {
		err = service.RemoveAttendeeRole(id)

		if err != nil {
			panic(errors.InternalError(err.Error(), "Could not remove Attendee role from user."))
		}
	}

	updated_rsvp, err := service.GetUserRsvp(id)

	if err != nil {
		panic(errors.DatabaseError(err.Error(), "Could not get updated RSVP for user."))
	}

	mail_template := "rsvp_update"
	err = service.SendUserMail(id, mail_template)

	if err != nil {
		panic(errors.InternalError(err.Error(), "Could not send user confirmation mail for RSVP update."))
	}

	json.NewEncoder(w).Encode(updated_rsvp)
}

/*
	Endpoint to get rsvp stats
*/
func GetStats(w http.ResponseWriter, r *http.Request) {
	stats, err := service.GetStats()

	if err != nil {
		panic(errors.InternalError(err.Error(), "Could not get RSVP service statistics."))
	}

	json.NewEncoder(w).Encode(stats)
}
