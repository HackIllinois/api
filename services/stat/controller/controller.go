package controller

import (
	"encoding/json"
	"github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/services/stat/service"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"net/http"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.Handle("/{name}/", alice.New().ThenFunc(GetStat)).Methods("GET")
	router.Handle("/", alice.New().ThenFunc(GetAllStat)).Methods("GET")

}

/*
	Endpoint to retrieve stats for a specified service
*/
func GetStat(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]

	stat, err := service.GetAggregatedStats(name)

	if err != nil {
		panic(errors.InternalError(err.Error(), err.Error()))
	}

	json.NewEncoder(w).Encode(stat)
}

/*
	Endpoint to retrieve stats for all services
*/
func GetAllStat(w http.ResponseWriter, r *http.Request) {
	all_stat, err := service.GetAllAggregatedStats()

	if err != nil {
		panic(errors.InternalError(err.Error(), err.Error()))
	}

	json.NewEncoder(w).Encode(all_stat)
}
