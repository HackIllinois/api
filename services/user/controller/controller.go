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

	router.Handle("/{id}/", alice.New().ThenFunc(GetUserInfo)).Methods("GET")
}

/*
	Endpoint to get the info for the current user
*/
func GetCurrentUserInfo(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	user_info, err := service.GetUserInfo(id)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
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
		panic(errors.UnprocessableError("Must provide id parameter"))
	}

	err := service.SetUserInfo(user_info.ID, user_info)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	updated_info, err := service.GetUserInfo(user_info.ID)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
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
		panic(errors.UnprocessableError(err.Error()))
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
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(user_info)
}
