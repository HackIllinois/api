package apiserver

import (
	"encoding/json"
	"fmt"
	"github.com/HackIllinois/api/common/config"
	"github.com/HackIllinois/api/common/middleware"
	"github.com/gorilla/mux"
	"github.com/thoas/stats"
	"net/http"
	"time"
)

func StartServer(address string, router *mux.Router, name string, initialize func() error) error {
	err := initialize()

	if err != nil {
		return err
	}

	router.Use(middleware.ContentTypeMiddleware)

	stats_middleware := stats.New()
	router.Use(stats_middleware.Handler)

	router.HandleFunc(fmt.Sprintf("/%s/internal/healthstats/", name), GetHealthStats(stats_middleware)).Methods("GET")
	router.HandleFunc(fmt.Sprintf("/%s/internal/reload/", name), Reload(initialize)).Methods("GET")

	server := &http.Server{
		Handler:      router,
		Addr:         address,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	return server.ListenAndServe()
}

/*
	Endpoint which returns health stats
	Returns HTTP200 when the service is healthy
	Returns HTTP503 when the service is unhealthy
*/
func GetHealthStats(stats_middleware *stats.Stats) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		health_stats := stats_middleware.Data()

		w.Header().Set("Content-Type", "application/json")

		// Set http response code based on service health
		// Used for easy health checks
		if IsHealthy(health_stats) {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusServiceUnavailable)
		}

		json.NewEncoder(w).Encode(health_stats)
	}
}

/*
	Reinitializes the service causing configuration variables to be reread
*/
func Reload(initialize func() error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := config.Initialize()

		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		err = initialize()

		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
