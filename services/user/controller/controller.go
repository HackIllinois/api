package controller

import (
	"encoding/json"
	"github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/services/user/models"
	"github.com/HackIllinois/api/services/user/service"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"net/http"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.Handle("/", alice.New().ThenFunc(GetCurrentUserInfo)).Methods("GET")
	router.Handle("/", alice.New().ThenFunc(SetUserInfo)).Methods("POST")
	router.Handle("/filter/", alice.New().ThenFunc(GetFilteredUserInfo)).Methods("GET")

	router.Handle("/qr/", alice.New().ThenFunc(GetCurrentQrCodeInfo)).Methods("GET")
	router.Handle("/qr/{id}/", alice.New().ThenFunc(GetQrCodeInfo)).Methods("GET")

	router.Handle("/{id}/", alice.New().ThenFunc(GetUserInfo)).Methods("GET")

	router.Handle("/internal/stats/", alice.New().ThenFunc(GetStats)).Methods("GET")
}

/*
	Endpoint to get the info for the current user
*/
func GetCurrentUserInfo(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	user_info, err := service.GetUserInfo(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not fetch user info by ID."))
		return
	}

	json.NewEncoder(w).Encode(user_info)
}

/*
	Endpoint to set the info for a specified user
*/
func SetUserInfo(w http.ResponseWriter, r *http.Request) {
	var user_info models.UserInfo
	json.NewDecoder(r.Body).Decode(&user_info)

	if user_info.ID == "" {
		errors.WriteError(w, r, errors.MalformedRequestError("Must provide user id in request.", "Must provide user id in request."))
		return
	}

	err := service.SetUserInfo(user_info.ID, user_info)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not upsert user info."))
		return
	}

	updated_info, err := service.GetUserInfo(user_info.ID)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not fetch user info by ID."))
		return
	}

	json.NewEncoder(w).Encode(updated_info)
}

/*
	Endpoint to get user info based on filters
*/
func GetFilteredUserInfo(w http.ResponseWriter, r *http.Request) {
	parameters := r.URL.Query()
	user_info, err := service.GetFilteredUserInfo(parameters)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not fetch filtered list of users."))
		return
	}

	json.NewEncoder(w).Encode(user_info)
}

/*
	Endpoint to get the info for a specified user
*/
func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	user_info, err := service.GetUserInfo(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not fetch user information by user id."))
		return
	}

	json.NewEncoder(w).Encode(user_info)
}

/*
	Endpoint to get the string to be embedded into the current user's QR code
*/
func GetCurrentQrCodeInfo(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	uri, err := service.GetQrInfo(id)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not parse QR code URI."))
		return
	}

	qr_info_container := models.QrInfoContainer{
		ID:     id,
		QrInfo: uri,
	}

	json.NewEncoder(w).Encode(qr_info_container)
}

/*
	Endpoint to get the string to be embedded into the specified user's QR code
*/
func GetQrCodeInfo(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	uri, err := service.GetQrInfo(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not parse QR code URI."))
		return
	}

	qr_info_container := models.QrInfoContainer{
		ID:     id,
		QrInfo: uri,
	}

	json.NewEncoder(w).Encode(qr_info_container)
}

/*
	Endpoint to get user stats
*/
func GetStats(w http.ResponseWriter, r *http.Request) {
	stats, err := service.GetStats()

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not retrieve user service statistics."))
		return
	}

	json.NewEncoder(w).Encode(stats)
}
