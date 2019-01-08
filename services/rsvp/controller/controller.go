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
		panic(errors.UnprocessableError(err.Error()))
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
		panic(errors.UnprocessableError(err.Error()))
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
		panic(errors.UnprocessableError("Must provide id"))
	}

	isAccepted, isActive, err := service.IsApplicantAcceptedAndActive(id)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	if !isAccepted {
		panic(errors.UnprocessableError("Applicant must be accepted to rsvp"))
	}

	if !isActive {
		panic(errors.UnprocessableError("Applicant decision has expired"))
	}

	rsvp := datastore.NewDataStore(config.RSVP_DEFINITION)
	json.NewDecoder(r.Body).Decode(&rsvp)

	rsvp.Data["id"] = id

	err = service.CreateUserRsvp(id, rsvp)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	isAttending, ok := rsvp.Data["isAttending"].(bool)

	if !ok {
		panic(errors.UnprocessableError("Failure in parsing user rsvp"))
	}

	if isAttending {
		err = service.AddAttendeeRole(id)

		if err != nil {
			panic(errors.UnprocessableError(err.Error()))
		}
	}

	updated_rsvp, err := service.GetUserRsvp(id)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	mail_template := "rsvp_confirmation"
	err = service.SendUserMail(id, mail_template)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
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
		panic(errors.UnprocessableError("Must provide id"))
	}

	isAccepted, isActive, err := service.IsApplicantAcceptedAndActive(id)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	if !isAccepted {
		panic(errors.UnprocessableError("Applicant must be accepted to modify rsvp"))
	}

	if !isActive {
		panic(errors.UnprocessableError("Cannot modify rsvp, applicant decision has expired"))
	}

	original_rsvp, err := service.GetUserRsvp(id)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	rsvp := datastore.NewDataStore(config.RSVP_DEFINITION)
	json.NewDecoder(r.Body).Decode(&rsvp)

	rsvp.Data["id"] = id

	err = service.UpdateUserRsvp(id, rsvp)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	wasAttending, ok := original_rsvp.Data["isAttending"].(bool)

	if !ok {
		panic(errors.UnprocessableError("Failure in parsing user rsvp"))
	}

	isAttending, ok := rsvp.Data["isAttending"].(bool)

	if !ok {
		panic(errors.UnprocessableError("Failure in parsing user rsvp"))
	}

	if !wasAttending && isAttending {
		err = service.AddAttendeeRole(id)

		if err != nil {
			panic(errors.UnprocessableError(err.Error()))
		}
	} else if wasAttending && !isAttending {
		err = service.RemoveAttendeeRole(id)

		if err != nil {
			panic(errors.UnprocessableError(err.Error()))
		}
	}

	updated_rsvp, err := service.GetUserRsvp(id)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	mail_template := "rsvp_update"
	err = service.SendUserMail(id, mail_template)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(updated_rsvp)
}

/*
	Endpoint to get rsvp stats
*/
func GetStats(w http.ResponseWriter, r *http.Request) {
	stats, err := service.GetStats()

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(stats)
}
