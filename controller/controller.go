package controller

import (
	"encoding/json"
	"github.com/HackIllinois/api-commons/errors"
	"github.com/HackIllinois/api-registration/models"
	"github.com/HackIllinois/api-registration/service"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"net/http"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.Handle("/{id}/", alice.New().ThenFunc(GetUserRegistration)).Methods("GET")
	router.Handle("/", alice.New().ThenFunc(GetCurrentUserRegistration)).Methods("GET")
	router.Handle("/", alice.New().ThenFunc(CreateCurrentUserRegistration)).Methods("POST")
	router.Handle("/", alice.New().ThenFunc(UpdateCurrentUserRegistration)).Methods("PUT")
}

/*
	Endpoint to get the registration for a specified user
*/
func GetUserRegistration(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	user_registration, err := service.GetUserRegistration(id)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(user_registration)
}

/*
	Endpoint to get the registration for the current user
*/
func GetCurrentUserRegistration(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	user_registration, err := service.GetUserRegistration(id)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(user_registration)
}

/*
	Endpoint to create the registration for the current user
*/
func CreateCurrentUserRegistration(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	if id == "" {
		panic(errors.UnprocessableError("Must provide id"))
	}

	var user_registration models.UserRegistration
	json.NewDecoder(r.Body).Decode(&user_registration)

	user_registration.ID = id

	user_info, err := service.GetUserInfo(id)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	user_registration.GitHub = user_info.Username
	user_registration.Email = user_info.Email

	err = service.AddAttendeeRole(id)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	err = service.CreateUserRegistration(id, user_registration)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	updated_registration, err := service.GetUserRegistration(id)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(updated_registration)
}

/*
	Endpoint to update the registration for the current user
*/
func UpdateCurrentUserRegistration(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	if id == "" {
		panic(errors.UnprocessableError("Must provide id"))
	}

	var user_registration models.UserRegistration
	json.NewDecoder(r.Body).Decode(&user_registration)

	user_registration.ID = id

	user_info, err := service.GetUserInfo(id)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	user_registration.GitHub = user_info.Username
	user_registration.Email = user_info.Email

	err = service.UpdateUserRegistration(id, user_registration)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	updated_registration, err := service.GetUserRegistration(id)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(updated_registration)
}
