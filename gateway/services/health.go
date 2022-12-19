package services

import (
	"encoding/json"
	"fmt"
	"github.com/HackIllinois/api/common/apirequest"
	"github.com/HackIllinois/api/common/authtoken"
	"github.com/HackIllinois/api/gateway/middleware"
	"github.com/arbor-dev/arbor"
	"github.com/justinas/alice"
	"net/http"
)

var HealthRoutes = arbor.RouteCollection{
	arbor.Route{
		"Health Check",
		"GET",
		"/health/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole}), middleware.IdentificationMiddleware).ThenFunc(GetHealthChecks).ServeHTTP,
	},
}

func GetHealthChecks(w http.ResponseWriter, r *http.Request) {
	healthy_services := []string{}
	unhealthy_services := []string{}

	service_health_stats := make(map[string]interface{})

	for service_name, service_location := range ServiceLocations {
		var health_stats map[string]interface{}
		status, err := apirequest.Get(fmt.Sprintf("%s/%s/internal/healthstats/", service_location, service_name), &health_stats)

		if err != nil {
			continue
		}

		service_health_stats[service_name] = health_stats

		if status == http.StatusOK {
			healthy_services = append(healthy_services, service_name)
		} else {
			unhealthy_services = append(unhealthy_services, service_name)
		}
	}

	health_info := map[string]interface{}{
		"healthyServices":   healthy_services,
		"unhealthyServices": unhealthy_services,
		"serviceHealthInfo": service_health_stats,
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(health_info)
}
