package controller

import (
	"encoding/json"
	"github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/services/stat/models"
	"github.com/HackIllinois/api/services/stat/service"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"net/http"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.Handle("/service/{name}/", alice.New().ThenFunc(GetService)).Methods("GET")
	router.Handle("/service/", alice.New().ThenFunc(RegisterService)).Methods("POST")

	router.Handle("/{name}/", alice.New().ThenFunc(GetStat)).Methods("GET")
	router.Handle("/", alice.New().ThenFunc(GetAllStat)).Methods("GET")

}

/*
	Endpoint to get the service with the specified name
*/
func GetService(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]

	api_service, err := service.GetService(name)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(api_service)
}

/*
	Endpoint to register a service
*/
func RegisterService(w http.ResponseWriter, r *http.Request) {
	var api_service models.Service
	json.NewDecoder(r.Body).Decode(&api_service)

	if api_service.Name == "" {
		panic(errors.UnprocessableError("Service must have a name"))
	}

	err := service.RegisterService(api_service.Name, api_service)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	updated_api_service, err := service.GetService(api_service.Name)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(updated_api_service)
}

/*
	Endpoint to retreive stats for a specified service
*/
func GetStat(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]

	stat, err := service.GetAggregatedStats(name)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(stat)
}

/*
	Endpoint to retreive stats for all services
*/
func GetAllStat(w http.ResponseWriter, r *http.Request) {
	all_stat, err := service.GetAllAggregatedStats()

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(all_stat)
}
