package controller

import (
	"encoding/json"
	"github.com/HackIllinois/api-upload/service"
	"github.com/HackIllinois/api-commons/errors"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"net/http"
	"io/ioutil"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.Handle("/resume/{id}/", alice.New().ThenFunc(GetUserResume)).Methods("GET")
	router.Handle("/resume/", alice.New().ThenFunc(UpdateUserResume)).Methods("PUT")
	router.Handle("/resume/", alice.New().ThenFunc(GetCurrentUserResume)).Methods("GET")
}

/*
	Endpoint to get a specified user's resume
*/
func GetUserResume(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	resume, err := service.GetUserResumeLink(id)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(resume)
}

/*
	Endpoint to get the current user's resume
*/
func GetCurrentUserResume(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	resume, err := service.GetUserResumeLink(id)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(resume)
}

/*
	Endpoint to update the specified user's resume
*/
func UpdateUserResume(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	file_buffer, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	err = service.UpdateUserResume(id, file_buffer)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	resume, err := service.GetUserResumeLink(id)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(resume)
}
