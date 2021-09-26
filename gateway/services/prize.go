package services

import (
	"net/http"

	"github.com/HackIllinois/api/gateway/config"
	"github.com/HackIllinois/api/gateway/middleware"
	"github.com/HackIllinois/api/gateway/models"
	"github.com/arbor-dev/arbor"
	"github.com/justinas/alice"
)

const PrizeFormat string = "JSON"

var PrizeRoutes = arbor.RouteCollection{
	arbor.Route{
		"GetPrize",
		"GET",
		"/prize/",
		alice.New(middleware.AuthMiddleware([]models.Role{models.AdminRole, models.StaffRole}), middleware.IdentificationMiddleware).ThenFunc(GetPrize).ServeHTTP,
	},
	arbor.Route{
		"CreatePrize",
		"POST",
		"/prize/",
		alice.New(middleware.AuthMiddleware([]models.Role{models.AdminRole, models.StaffRole}), middleware.IdentificationMiddleware).ThenFunc(CreatePrize).ServeHTTP,
	},
	arbor.Route{
		"UpdatePrize",
		"PUT",
		"/prize/",
		alice.New(middleware.AuthMiddleware([]models.Role{models.AdminRole, models.StaffRole}), middleware.IdentificationMiddleware).ThenFunc(UpdatePrize).ServeHTTP,
	},
	arbor.Route{
		"DeletePrize",
		"DELETE",
		"/prize/",
		alice.New(middleware.AuthMiddleware([]models.Role{models.AdminRole, models.StaffRole}), middleware.IdentificationMiddleware).ThenFunc(DeletePrize).ServeHTTP,
	},
}

func GetPrize(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.PRIZE_SERVICE+r.URL.String(), PrizeFormat, "", r)
}

func CreatePrize(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, config.PRIZE_SERVICE+r.URL.String(), PrizeFormat, "", r)
}

func UpdatePrize(w http.ResponseWriter, r *http.Request) {
	arbor.PUT(w, config.PRIZE_SERVICE+r.URL.String(), PrizeFormat, "", r)
}

func DeletePrize(w http.ResponseWriter, r *http.Request) {
	arbor.DELETE(w, config.PRIZE_SERVICE+r.URL.String(), PrizeFormat, "", r)
}
