package controller

import (
	"encoding/json"
	"net/http"

	common_config "github.com/HackIllinois/api/common/config"
	"github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/common/metrics"
	"github.com/HackIllinois/api/common/utils"
	"github.com/HackIllinois/api/services/checkin/models"
	"github.com/HackIllinois/api/services/checkin/service"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var validate *validator.Validate

func SetupController(route *mux.Route) {
	validate = validator.New()

	router := route.Subrouter()

	router.Handle("/internal/metrics/", promhttp.Handler()).Methods("GET")

	metrics.RegisterHandler("/", CreateUserCheckin, "POST", router)
	metrics.RegisterHandler("/", UpdateUserCheckin, "PUT", router)
	metrics.RegisterHandler("/", GetCurrentUserCheckin, "GET", router)
	metrics.RegisterHandler("/list/", GetAllCheckedInUsers, "GET", router)
	metrics.RegisterHandler("/{id}/", GetUserCheckin, "GET", router)
	metrics.RegisterHandler("/internal/stats/", GetStats, "GET", router)
}

/*
Endpoint to get a specified user's checkin
*/
func GetUserCheckin(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	user_checkin, err := service.GetUserCheckin(id)
	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get specified user's check-in details."))
		return
	}

	json.NewEncoder(w).Encode(user_checkin)
}

/*
Endpoint to get the current user's checkin
*/
func GetCurrentUserCheckin(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	user_checkin, err := service.GetUserCheckin(id)
	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get current user's check-in details."))
		return
	}

	json.NewEncoder(w).Encode(user_checkin)
}

/*
Endpoint to set the specified user's checkin
*/
func CreateUserCheckin(w http.ResponseWriter, r *http.Request) {
	var user_checkin models.UserCheckin
	json.NewDecoder(r.Body).Decode(&user_checkin)

	err := validate.Struct(user_checkin)
	if err != nil {
		errors.WriteError(w, r, errors.AttributeMismatchError(err.Error(), "No user token was provided."))
		return
	}

	id, err := utils.FetchIdFromSignedUserToken(common_config.TOKEN_SECRET, user_checkin.UserToken)
	if err != nil {
		errors.WriteError(w, r, errors.AuthorizationError(err.Error(), "Bad user token."))
		return
	}

	user_checkin.ID = id

	can_user_checkin, err := service.CanUserCheckin(user_checkin.ID, user_checkin.Override)

	// Ignore the error caused when a user hasn't been accepted (no RSVP status)
	if err != nil && err.Error() != "Rsvp service failed to return status" {
		if err.Error() == "User is not registered." {
			errors.WriteError(w, r, errors.AttributeMismatchError("User is not registered.", "User is not registered."))
		} else {
			errors.WriteError(w, r, errors.InternalError(err.Error(), "Unable to determine user's check-in permissions."))
		}
		return
	}

	if !can_user_checkin {
		errors.WriteError(
			w,
			r,
			errors.AttributeMismatchError(
				"Reasons for not being able to check-in include: no RSVP, no staff override (in case of no RSVP), or check-ins are not allowed at this time.",
				"Attendee has not RSVPed.",
			),
		)
		return
	}

	is_rsvped, err := service.IsAttendeeRsvped(user_checkin.ID)

	// Ignore the error caused when a user hasn't been accepted
	if err != nil && err.Error() != "Rsvp service failed to return status" {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not retrieve rsvp status."))
		return
	}

	if is_rsvped {
		rsvp_data, err := service.GetRsvpData(user_checkin.ID)
		if err != nil {
			errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not retrieve rsvp data."))
			return
		}

		user_checkin.RsvpData = rsvp_data
	}

	err = service.CreateUserCheckin(user_checkin.ID, user_checkin)

	if err != nil {
		if err.Error() == "Checkin already exists" {
			errors.WriteError(
				w,
				r,
				errors.AttributeMismatchError("User has already checked in.", "User has already checked in."),
			)
		} else {
			errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not create user check-in."))
		}
		return
	}

	updated_checkin, err := service.GetUserCheckin(user_checkin.ID)
	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get recently created check-in information."))
		return
	}

	if updated_checkin.Override {
		err = service.AddAttendeeRole(updated_checkin.ID)

		if err != nil {
			errors.WriteError(w, r, errors.AuthorizationError(err.Error(), "Could not add attendee role to user."))
			return
		}
	}

	json.NewEncoder(w).Encode(updated_checkin)
}

/*
Endpoint to update the specified user's checkin
*/
func UpdateUserCheckin(w http.ResponseWriter, r *http.Request) {
	var user_checkin models.UserCheckin
	json.NewDecoder(r.Body).Decode(&user_checkin)

	err := validate.Struct(user_checkin)
	if err != nil {
		errors.WriteError(w, r, errors.AttributeMismatchError(err.Error(), "No user token was provided."))
		return
	}

	id, err := utils.FetchIdFromSignedUserToken(common_config.TOKEN_SECRET, user_checkin.UserToken)
	if err != nil {
		errors.WriteError(w, r, errors.AuthorizationError(err.Error(), "Bad user token."))
		return
	}

	user_checkin.ID = id

	rsvp_data, err := service.GetRsvpData(user_checkin.ID)
	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not retrieve rsvp data."))
		return
	}

	user_checkin.RsvpData = rsvp_data

	err = service.UpdateUserCheckin(user_checkin.ID, user_checkin)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not update user check-in information."))
		return
	}

	updated_checkin, err := service.GetUserCheckin(user_checkin.ID)
	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not fetch updated check-in information."))
		return
	}

	if updated_checkin.Override {
		err = service.AddAttendeeRole(updated_checkin.ID)

		if err != nil {
			errors.WriteError(w, r, errors.AuthorizationError(err.Error(), "Could not add attendee role."))
			return
		}
	}

	json.NewEncoder(w).Encode(updated_checkin)
}

/*
Endpoint to get all checked in user IDs
*/
func GetAllCheckedInUsers(w http.ResponseWriter, r *http.Request) {
	checked_in_users, err := service.GetAllCheckedInUsers()
	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get all checked-in users."))
		return
	}

	json.NewEncoder(w).Encode(checked_in_users)
}

/*
Endpoint to get checkin stats
*/
func GetStats(w http.ResponseWriter, r *http.Request) {
	stats, err := service.GetStats()
	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not get check-in service statistics."))
		return
	}

	json.NewEncoder(w).Encode(stats)
}
