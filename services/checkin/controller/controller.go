package controller

import (
	"encoding/json"
	"github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/services/checkin/models"
	"github.com/HackIllinois/api/services/checkin/service"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"net/http"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.Handle("/", alice.New().ThenFunc(CreateUserCheckin)).Methods("POST")
	router.Handle("/", alice.New().ThenFunc(UpdateUserCheckin)).Methods("PUT")
	router.Handle("/", alice.New().ThenFunc(GetCurrentUserCheckin)).Methods("GET")
	router.Handle("/list/", alice.New().ThenFunc(GetAllCheckedInUsers)).Methods("GET")
	router.Handle("/{id}/", alice.New().ThenFunc(GetUserCheckin)).Methods("GET")

	router.Handle("/internal/stats/", alice.New().ThenFunc(GetStats)).Methods("GET")
}

/*
	Endpoint to get a specified user's checkin
*/
func GetUserCheckin(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	user_checkin, err := service.GetUserCheckin(id)

	if err != nil {
		panic(errors.DATABASE_ERROR("Could not get specified user's check-in details."))
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
		panic(errors.DATABASE_ERROR("Could not get current user's check-in details."))
	}

	json.NewEncoder(w).Encode(user_checkin)
}

/*
	Endpoint to set the specified user's checkin
*/
func CreateUserCheckin(w http.ResponseWriter, r *http.Request) {
	var user_checkin models.UserCheckin
	json.NewDecoder(r.Body).Decode(&user_checkin)

	can_user_checkin, err := service.CanUserCheckin(user_checkin.ID, user_checkin.Override)

	if err != nil {
		panic(errors.INTERNAL_ERROR("Unable to determine user's check-in permissions."))
	}

	if !can_user_checkin {
		panic(errors.ATTRIBUTE_MISMATCH_ERROR("Attendee must be RSVPed to check-in (or have a staff override)."))
	}

	err = service.CreateUserCheckin(user_checkin.ID, user_checkin)

	if err != nil {
		panic(errors.DATABASE_ERROR("Could not create user check-in."))
	}

	updated_checkin, err := service.GetUserCheckin(user_checkin.ID)

	if err != nil {
		panic(errors.DATABASE_ERROR("Could not get recently created check-in information."))
	}

	if updated_checkin.Override {
		err = service.AddAttendeeRole(updated_checkin.ID)

		if err != nil {
			panic(errors.INTERNAL_ERROR("Could not add attendee role to user.\n" + err.Error()))
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

	err := service.UpdateUserCheckin(user_checkin.ID, user_checkin)

	if err != nil {
		panic(errors.DATABASE_ERROR("Could not update user check-in information."))
	}

	updated_checkin, err := service.GetUserCheckin(user_checkin.ID)

	if err != nil {
		panic(errors.DATABASE_ERROR("Could not fetch updated check-in information."))
	}

	if updated_checkin.Override {
		err = service.AddAttendeeRole(updated_checkin.ID)

		if err != nil {
			panic(errors.INTERNAL_ERROR("Could not add attendee role.\n" + err.Error()))
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
		panic(errors.DATABASE_ERROR("Could not get all checked-in users."))
	}

	json.NewEncoder(w).Encode(checked_in_users)
}

/*
	Endpoint to get checkin stats
*/
func GetStats(w http.ResponseWriter, r *http.Request) {
	stats, err := service.GetStats()

	if err != nil {
		panic(errors.INTERNAL_ERROR("Could not get check-in service statistics."))
	}

	json.NewEncoder(w).Encode(stats)
}
