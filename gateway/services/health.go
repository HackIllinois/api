package services

import (
	"encoding/json"
	"fmt"
	"github.com/HackIllinois/api/common/apirequest"
	"github.com/HackIllinois/api/gateway/config"
	"github.com/HackIllinois/api/gateway/middleware"
	"github.com/HackIllinois/api/gateway/models"
	"github.com/arbor-dev/arbor"
	"github.com/justinas/alice"
	"net/http"
)

var ServiceLocations = map[string]string{
	"auth":          config.AUTH_SERVICE,
	"user":          config.USER_SERVICE,
	"registration":  config.REGISTRATION_SERVICE,
	"decision":      config.DECISION_SERVICE,
	"rsvp":          config.RSVP_SERVICE,
	"checkin":       config.CHECKIN_SERVICE,
	"upload":        config.UPLOAD_SERVICE,
	"mail":          config.MAIL_SERVICE,
	"event":         config.EVENT_SERVICE,
	"stat":          config.STAT_SERVICE,
	"notifications": config.NOTIFICATIONS_SERVICE,
}

var HealthRoutes = arbor.RouteCollection{
	arbor.Route{
		"Health Check",
		"GET",
		"/health/",
		alice.New(middleware.AuthMiddleware([]models.Role{models.AdminRole}), middleware.IdentificationMiddleware).ThenFunc(GetHealthChecks).ServeHTTP,
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
