package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

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
	router.HandleFunc("/search/", GetFilteredProfiles).Methods("GET")
	router.HandleFunc("/leaderboard/", GetProfileLeaderboard).Methods("GET")
	router.HandleFunc("/checkin/{id}/", IsEventRedeemed).Methods("GET")
	router.HandleFunc("/award/{points}/", AwardPoints).Methods("GET")
	router.HandleFunc("/{id}/", GetProfileById).Methods("GET")
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
	GetProfileById is used to get a profile for a provided id.
*/
func GetProfileById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	user_profile, err := service.GetProfile(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get profile for id "+id+"."))
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

/*
	GetProfileLeaderboard is the endpoint used to return a list of profiles, sorted by the amount of points they have (descending).
*/
func GetProfileLeaderboard(w http.ResponseWriter, r *http.Request) {
	parameters := r.URL.Query()

	user_profile_list, err := service.GetProfileLeaderboard(parameters)
	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get the profile leaderboard."))
		return
	}

	json.NewEncoder(w).Encode(user_profile_list)
}

/*
	Filters the profiles by TeamStatus and Interests
*/
func GetFilteredProfiles(w http.ResponseWriter, r *http.Request) {
	parameters := r.URL.Query()

	filtered_profile_list, err := service.GetFilteredProfiles(parameters)
	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get the filtered profiles."))
		return
	}

	json.NewEncoder(w).Encode(filtered_profile_list)
}

/*
	Filters the profiles by TeamStatus and Interests. Additionally filters out profiles that have the TeamStatus "Not Looking".
*/
func GetValidFilteredProfiles(w http.ResponseWriter, r *http.Request) {
	parameters := r.URL.Query()

	filtered_profile_list, err := service.GetValidFilteredProfiles(parameters)
	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get the valid filtered profiles."))
		return
	}

	json.NewEncoder(w).Encode(filtered_profile_list)
}

/*
	IsEventRedeemed checks the appropriate table to check whether the given event id has already been redeemed. If the event is not in the table, it add to the array.
*/
func IsEventRedeemed(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")
	event_id := mux.Vars(r)["id"]

	redemption_status, err := service.EventRedemptionStatus(id, event_id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not check if event was redeemed for id "+id+"."))
		return
	}
	if !redemption_status {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Event with id "+id+" was already redeemed."))
		return
	}

	json.NewEncoder(w).Encode(redemption_status)
}

/*
	AwardPoints gives the specified number of points to the current user.
*/
func AwardPoints(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")
	points, err := strconv.Atoi(mux.Vars(r)["points"])

	user_profile, err := service.GetProfile(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get profile for id "+id+" when trying to award points."))
		return
	}

	user_profile.Points += points

	err = service.UpdateProfile(id, *user_profile)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not update the profile when trying to award points."))
		return
	}

	updated_profile, err := service.GetProfile(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get updated profile details after awarding points."))
		return
	}

	json.NewEncoder(w).Encode(updated_profile)
}
