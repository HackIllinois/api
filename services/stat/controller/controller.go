package controller

import (
	"encoding/json"
	"github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/services/stat/service"
	"github.com/gorilla/mux"
	"net/http"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.HandleFunc("/{name}/", GetStat).Methods("GET")
	router.HandleFunc("/", GetAllStat).Methods("GET")

}

/*
	Endpoint to retrieve stats for a specified service
*/
func GetStat(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]

	stat, err := service.GetAggregatedStats(name)

	if err != nil {
		panic(errors.InternalError(err.Error(), "Failed to get statistics for service "+name+"."))
	}

	json.NewEncoder(w).Encode(stat)
}

/*
	Endpoint to retrieve stats for all services
*/
func GetAllStat(w http.ResponseWriter, r *http.Request) {
	all_stat, err := service.GetAllAggregatedStats()

	if err != nil {
		panic(errors.InternalError(err.Error(), "Failed to aggregate statistics."))
	}

	json.NewEncoder(w).Encode(all_stat)
}
