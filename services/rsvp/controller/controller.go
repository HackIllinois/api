package controller

import (
	"encoding/json"
	"github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/services/rsvp/models"
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
		panic(errors.MalformedRequestError("Must provide ID in request.", "Must provide ID in request."))
	}

	isAccepted, err := service.IsApplicantAccepted(id)

	if err != nil {
		panic(errors.InternalError(err.Error(), err.Error()))
	}

	if !isAccepted {
		panic(errors.AttributeMismatchError(err.Error(), "Applicant must be accepted to RSVP."))
	}

	var rsvp models.UserRsvp
	json.NewDecoder(r.Body).Decode(&rsvp)

	rsvp.ID = id

	err = service.CreateUserRsvp(id, rsvp)

	if err != nil {
		panic(errors.InternalError(err.Error(), err.Error()))
	}

	if rsvp.IsAttending {
		err = service.AddAttendeeRole(id)

		if err != nil {
			panic(errors.InternalError(err.Error(), err.Error()))
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
		panic(errors.MalformedRequestError("Must provide ID in request.", "Must provide ID in the request."))
	}

	original_rsvp, err := service.GetUserRsvp(id)

	if err != nil {
		panic(errors.DatabaseError(err.Error(), "Could not get user's RSVP status."))
	}

	var rsvp models.UserRsvp
	json.NewDecoder(r.Body).Decode(&rsvp)

	rsvp.ID = id

	err = service.UpdateUserRsvp(id, rsvp)

	if err != nil {
		panic(errors.DatabaseError(err.Error(), "Could not update user RSVP."))
	}

	if !original_rsvp.IsAttending && rsvp.IsAttending {
		err = service.AddAttendeeRole(id)

		if err != nil {
			panic(errors.InternalError(err.Error(), err.Error()))
		}
	} else if original_rsvp.IsAttending && !rsvp.IsAttending {
		err = service.RemoveAttendeeRole(id)

		if err != nil {
			panic(errors.InternalError(err.Error(), err.Error()))
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
