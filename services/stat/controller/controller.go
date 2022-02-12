package controller

import (
	"encoding/json"
	"net/http"

	"github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/common/metrics"
	"github.com/HackIllinois/api/services/stat/service"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.Handle("/internal/metrics/", promhttp.Handler()).Methods("GET")

	metrics.RegisterHandler("/{name}/", GetStat, "GET", router)
	metrics.RegisterHandler("/", GetAllStat, "GET", router)
}

/*
	Endpoint to retrieve stats for a specified service
*/
func GetStat(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]

	stat, err := service.GetAggregatedStats(name)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Failed to get statistics for service "+name+"."))
		return
	}

	json.NewEncoder(w).Encode(stat)
}

/*
	Endpoint to retrieve stats for all services
*/
func GetAllStat(w http.ResponseWriter, r *http.Request) {
	all_stat, err := service.GetAllAggregatedStats()

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Failed to aggregate statistics."))
		return
	}

	json.NewEncoder(w).Encode(all_stat)
}
