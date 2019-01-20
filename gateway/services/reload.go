package services

import (
	"encoding/json"
	"fmt"
	"github.com/HackIllinois/api/common/apirequest"
	"github.com/HackIllinois/api/gateway/middleware"
	"github.com/HackIllinois/api/gateway/models"
	"github.com/arbor-dev/arbor"
	"github.com/justinas/alice"
	"net/http"
)

var ReloadRoutes = arbor.RouteCollection{
	arbor.Route{
		"Reload",
		"GET",
		"/reload/",
		alice.New(middleware.AuthMiddleware([]models.Role{models.AdminRole}), middleware.IdentificationMiddleware).ThenFunc(Reload).ServeHTTP,
	},
}

func Reload(w http.ResponseWriter, r *http.Request) {
	reload_success := []string{}
	reload_failed := []string{}

	for service_name, service_location := range ServiceLocations {
		status, err := apirequest.Get(fmt.Sprintf("%s/%s/internal/reload/", service_location, service_name), nil)

		if err != nil {
			reload_failed = append(reload_failed, service_name)
			continue
		}

		if status == http.StatusOK {
			reload_success = append(reload_success, service_name)
		} else {
			reload_failed = append(reload_failed, service_name)
		}
	}

	reload_info := map[string][]string{
		"success": reload_success,
		"failed":  reload_failed,
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(reload_info)
}
