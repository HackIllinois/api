package controller

import (
	"encoding/json"
	"net/http"

	"github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/services/profile/models"
	"github.com/HackIllinois/api/services/profile/service"
	"github.com/gorilla/mux"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.HandleFunc("/", GetProfile).Methods("GET")
	router.HandleFunc("/", CreateProfile).Methods("POST")
	router.HandleFunc("/", UpdateProfile).Methods("PUT")
	router.HandleFunc("/", DeleteProfile).Methods("DELETE")

	router.HandleFunc("/list/", GetAllProfiles).Methods("GET")
}

/*
	GetProfile is the endpoint to get the profile for the current user
*/
func GetProfile(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	user_profile, err := service.GetProfile(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get current user's profile."))
		return
	}

	json.NewEncoder(w).Encode(user_profile)
}

/*
	CreateProfile is the endpoint to create the profile for the current user.
*/
func CreateProfile(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	if id == "" {
		errors.WriteError(w, r, errors.MalformedRequestError("Must provide id in request.", "Must provide id in request."))
		return
	}

	var profile models.Profile
	json.NewDecoder(r.Body).Decode(&profile)
	err := service.CreateProfile(id, profile)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not create new profile."))
		return
	}

	created_profile, err := service.GetProfile(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get created profile."))
		return
	}

	json.NewEncoder(w).Encode(created_profile)
}

/*
	UpdateProfile is the endpoint to update the profile for the current user
*/
func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	if id == "" {
		errors.WriteError(w, r, errors.MalformedRequestError("Must provide id in request.", "Must provide id in request."))
		return
	}

	var profile models.Profile
	json.NewDecoder(r.Body).Decode(&profile)

	err := service.UpdateProfile(id, profile)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not update the profile."))
		return
	}

	updated_profile, err := service.GetProfile(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get updated profile details."))
		return
	}

	json.NewEncoder(w).Encode(updated_profile)
}

/*
	DeleteProfile is the endpoint to delete the profile for the current user
*/
func DeleteProfile(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")
	if id == "" {
		errors.WriteError(w, r, errors.MalformedRequestError("Must provide id in request.", "Must provide id in request."))
		return
	}

	deleted_profile, err := service.DeleteProfile(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not delete the profile."))
		return
	}

	json.NewEncoder(w).Encode(deleted_profile)
}

/*
	GetAllProfiles is the endpoint to get all active user profiles
*/
func GetAllProfiles(w http.ResponseWriter, r *http.Request) {

	user_profile_list, err := service.GetAllProfiles()

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not list all user profiles."))
		return
	}

	json.NewEncoder(w).Encode(user_profile_list)
}
