package controller

import (
	"encoding/json"
	"net/http"

	"github.com/HackIllinois/api/common/authtoken"
	"github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/common/metrics"
	"github.com/HackIllinois/api/common/utils"
	"github.com/HackIllinois/api/services/profile/config"
	"github.com/HackIllinois/api/services/profile/models"
	"github.com/HackIllinois/api/services/profile/service"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.Handle("/internal/metrics/", promhttp.Handler()).Methods("GET")

	metrics.RegisterHandler("/", GetProfile, "GET", router)
	metrics.RegisterHandler("/", CreateProfile, "POST", router)
	metrics.RegisterHandler("/", UpdateProfile, "PUT", router)
	metrics.RegisterHandler("/", DeleteProfile, "DELETE", router)

	metrics.RegisterHandler("/list/", GetFilteredProfiles, "GET", router)
	metrics.RegisterHandler("/leaderboard/", GetProfileLeaderboard, "GET", router)
	metrics.RegisterHandler("/search/", GetValidFilteredProfiles, "GET", router)

	metrics.RegisterHandler("/event/checkin/", RedeemEvent, "POST", router)
	metrics.RegisterHandler("/points/award/", AwardPoints, "POST", router)

	metrics.RegisterHandler("/favorite/", GetProfileFavorites, "GET", router)
	metrics.RegisterHandler("/favorite/", AddProfileFavorite, "POST", router)
	metrics.RegisterHandler("/favorite/", RemoveProfileFavorite, "DELETE", router)

	metrics.RegisterHandler("/{id}/", GetProfileById, "GET", router)

	metrics.RegisterHandler("/tier/threshold/", GetTierThresholds, "GET", router)
}

/*
	GetProfile is the endpoint to get the profile for the current user
*/
func GetProfile(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	profile_id, err := service.GetProfileIdFromUserId(id)
	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get profile id associated with the user"))
		return
	}

	user_profile, err := service.GetProfile(profile_id)
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
	profile_id := mux.Vars(r)["id"]

	user_profile, err := service.GetProfile(profile_id)
	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get profile for profile id "+profile_id))
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

	profile_id, err := service.GetProfileIdFromUserId(id)

	if err == nil {
		errors.WriteError(w, r, errors.DatabaseError("", "User already has a profile with profile id "+profile_id))
		return
	}

	profile_id = utils.GenerateUniqueID()

	var profile models.Profile
	json.NewDecoder(r.Body).Decode(&profile)

	profile.Points = 0
	profile.FoodWave = 0

	err = service.CreateProfile(id, profile_id, profile)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not create new profile."))
		return
	}

	created_profile, err := service.GetProfile(profile_id)
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

	profile_id, err := service.GetProfileIdFromUserId(id)
	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get profile id associated with the user"))
		return
	}

	var profile models.Profile
	json.NewDecoder(r.Body).Decode(&profile)

	old_profile, err := service.GetProfile(profile_id)
	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get profile associated with this profile id."))
		return
	}

	is_staff := authtoken.IsRequestFromStaffOrHigher(r)

	if !is_staff {
		profile.Points = old_profile.Points
	}

	if !is_staff || profile.FoodWave == 0 {
		profile.FoodWave = old_profile.FoodWave
	}

	err = service.UpdateProfile(profile_id, profile)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not update the profile."))
		return
	}

	updated_profile, err := service.GetProfile(profile_id)
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
	// id := r.Header.Get("HackIllinois-Identity")

	// if id == "" {
	// 	errors.WriteError(w, r, errors.MalformedRequestError("Must provide id in request.", "Must provide id in request."))
	// 	return
	// }

	// profile_id, err := service.GetProfileIdFromUserId(id)

	// if err != nil {
	// 	errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get profile id associated with the user"))
	// 	return
	// }

	// deleted_profile, err := service.DeleteProfile(profile_id)

	// if err != nil {
	// 	errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not delete the profile."))
	// 	return
	// }

	// json.NewEncoder(w).Encode(deleted_profile)
	errors.WriteError(w, r, errors.InternalError("Endpoint temporarily disabled.", "Endpoint temporarily disabled."))
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
	Filters the profiles by TeamStatus and Interests. Additionally filters out profiles that have the TeamStatus "NOT_LOOKING".
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
	RedeemEvent checks the appropriate table to check whether the given event id has already
	been redeemed by the specified user. If the event is not in the table, add it to the array.
*/
func RedeemEvent(w http.ResponseWriter, r *http.Request) {
	var request models.RedeemEventRequest
	json.NewDecoder(r.Body).Decode(&request)

	id := request.ID

	profile_id, err := service.GetProfileIdFromUserId(id)
	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get profile id associated with the user"))
		return
	}

	redemption_status, err := service.RedeemEvent(profile_id, request.EventID)
	if err != nil {
		errors.WriteError(
			w,
			r,
			errors.DatabaseError(
				err.Error(),
				"Could not check if event was redeemed for id "+request.ID+" and event id "+request.EventID+". "+redemption_status.Status,
			),
		)
		return
	}

	json.NewEncoder(w).Encode(redemption_status)
}

/*
	AwardPoints gives the specified number of points to the current user.
*/
func AwardPoints(w http.ResponseWriter, r *http.Request) {
	var request models.AwardPointsRequest
	json.NewDecoder(r.Body).Decode(&request)

	id := request.ID

	profile_id, err := service.GetProfileIdFromUserId(id)
	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get profile id associated with the user"))
		return
	}

	user_profile, err := service.GetProfile(profile_id)
	if err != nil {
		errors.WriteError(
			w,
			r,
			errors.DatabaseError(err.Error(), "Could not get profile for id "+request.ID+" when trying to award points."),
		)
		return
	}

	user_profile.Points += request.Points

	err = service.UpdateProfile(profile_id, *user_profile)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not update the profile when trying to award points."))
		return
	}

	updated_profile, err := service.GetProfile(profile_id)
	if err != nil {
		errors.WriteError(
			w,
			r,
			errors.DatabaseError(err.Error(), "Could not get updated profile details after awarding points."),
		)
		return
	}

	json.NewEncoder(w).Encode(updated_profile)
}

/*
	Endpoint to get the current user's profile favorites
*/
func GetProfileFavorites(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	if id == "" {
		errors.WriteError(w, r, errors.MalformedRequestError("Must provide id in request.", "Must provide id in request."))
		return
	}

	profile_id, err := service.GetProfileIdFromUserId(id)
	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get profile id associated with the user"))
		return
	}

	favorites, err := service.GetProfileFavorites(profile_id)
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

	if id == "" {
		errors.WriteError(w, r, errors.MalformedRequestError("Must provide id in request.", "Must provide id in request."))
		return
	}

	profile_id, err := service.GetProfileIdFromUserId(id)
	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get profile id associated with the user"))
		return
	}

	var profile_favorite_modification models.ProfileFavoriteModification
	json.NewDecoder(r.Body).Decode(&profile_favorite_modification)

	err = service.AddProfileFavorite(profile_id, profile_favorite_modification.ID)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not add a profile favorite for the current user."))
		return
	}

	favorites, err := service.GetProfileFavorites(profile_id)
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

	if id == "" {
		errors.WriteError(w, r, errors.MalformedRequestError("Must provide id in request.", "Must provide id in request."))
		return
	}

	profile_id, err := service.GetProfileIdFromUserId(id)
	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get profile id associated with the user"))
		return
	}

	var profile_favorite_modification models.ProfileFavoriteModification
	json.NewDecoder(r.Body).Decode(&profile_favorite_modification)

	err = service.RemoveProfileFavorite(profile_id, profile_favorite_modification.ID)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not remove a profile favorite for the current user."))
		return
	}

	favorites, err := service.GetProfileFavorites(profile_id)
	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get updated user profile favorites."))
		return
	}

	json.NewEncoder(w).Encode(favorites)
}

/*
	Returns the tier name to threshold mapping
*/
func GetTierThresholds(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(config.TIER_THRESHOLDS)
}
