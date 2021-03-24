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
	router.HandleFunc("/search/", GetFilteredProfiles).Methods("GET")
	router.HandleFunc("/leaderboard/", GetProfileLeaderboard).Methods("GET")

	router.HandleFunc("/favorite/", GetProfileFavorites).Methods("GET")
	router.HandleFunc("/favorite/add/", AddProfileFavorite).Methods("POST")
	router.HandleFunc("/favorite/remove/", RemoveProfileFavorite).Methods("POST")

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
	Endpoint to get the current user's profile favorites
*/
func GetProfileFavorites(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	favorites, err := service.GetProfileFavorites(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get user's profile favorites."))
		return
	}

	json.NewEncoder(w).Encode(favorites)
}

/*
	Endpoint to add a profile favorite for the current user
*/
func AddProfileFavorite(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	var profile_favorite_modification models.ProfileFavoriteModification
	json.NewDecoder(r.Body).Decode(&profile_favorite_modification)

	err := service.AddProfileFavorite(id, profile_favorite_modification.ProfileID)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not add a profile favorite for the current user."))
		return
	}

	favorites, err := service.GetProfileFavorites(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get updated user profile favorites."))
		return
	}

	json.NewEncoder(w).Encode(favorites)
}

/*
	Endpoint to remove a profile favorite for the current user
*/
func RemoveProfileFavorite(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	var profile_favorite_modification models.ProfileFavoriteModification
	json.NewDecoder(r.Body).Decode(&profile_favorite_modification)

	err := service.RemoveProfileFavorite(id, profile_favorite_modification.ProfileID)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not remove a profile favorite for the current user."))
		return
	}

	favorites, err := service.GetProfileFavorites(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get updated user profile favorites."))
		return
	}

	json.NewEncoder(w).Encode(favorites)
}
