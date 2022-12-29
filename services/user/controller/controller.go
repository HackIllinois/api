package controller

import (
	"encoding/json"
	"net/http"

	"github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/common/metrics"
	"github.com/HackIllinois/api/services/user/models"
	"github.com/HackIllinois/api/services/user/service"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.Handle("/internal/metrics/", promhttp.Handler()).Methods("GET")

	metrics.RegisterHandler("/", GetCurrentUserInfo, "GET", router)
	metrics.RegisterHandler("/", SetUserInfo, "POST", router)
	metrics.RegisterHandler("/filter/", GetFilteredUserInfo, "GET", router)

	metrics.RegisterHandler("/qr/", GetCurrentQrCodeInfo, "GET", router)
	metrics.RegisterHandler("/qr/{id}/", GetQrCodeInfo, "GET", router)

	metrics.RegisterHandler("/{id}/", GetUserInfo, "GET", router)

	metrics.RegisterHandler("/internal/stats/", GetStats, "GET", router)
}

/*
	Endpoint to get the info for the current user
*/
func GetCurrentUserInfo(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	user_info, err := service.GetUserInfo(id, nil)

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

	updated_info, err := service.UpsertUserInfo(user_info.ID, user_info)

	if err != nil {
		errors.WriteError(w, r, *err)
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

	user_info, err := service.GetUserInfo(id, nil)

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
