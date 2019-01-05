package apiserver

import (
	"encoding/json"
	"fmt"
	"github.com/HackIllinois/api/common/middleware"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/thoas/stats"
	"net/http"
	"time"
)

func StartServer(address string, router *mux.Router, name string) error {
	router.Use(middleware.ErrorMiddleware)
	router.Use(middleware.ContentTypeMiddleware)

	stats_middleware := stats.New()
	router.Use(stats_middleware.Handler)

	router.Handle(fmt.Sprintf("/%s/internal/healthstats/", name), alice.New().ThenFunc(GetHealthStats(stats_middleware))).Methods("GET")

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
