package services

import (
	"net/http"

	"github.com/HackIllinois/api/common/authtoken"
	"github.com/HackIllinois/api/gateway/config"
	"github.com/HackIllinois/api/gateway/middleware"
	"github.com/arbor-dev/arbor"
	"github.com/justinas/alice"
)

const ProjectFormat string = "JSON"

var ProjectRoutes = arbor.RouteCollection{
	arbor.Route{
		"GetProjectFavorites",
		"GET",
		"/project/favorite/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.UserRole}), middleware.IdentificationMiddleware).ThenFunc(GetProjectFavorites).ServeHTTP,
	},
	arbor.Route{
		"AddProjectFavorite",
		"POST",
		"/project/favorite/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.UserRole}), middleware.IdentificationMiddleware).ThenFunc(AddProjectFavorite).ServeHTTP,
	},
	arbor.Route{
		"RemoveProjectFavorite",
		"DELETE",
		"/project/favorite/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.UserRole}), middleware.IdentificationMiddleware).ThenFunc(RemoveProjectFavorite).ServeHTTP,
	},
	arbor.Route{
		"GetFilteredProjects",
		"GET",
		"/project/filter/",
		alice.New(middleware.IdentificationMiddleware).ThenFunc(GetFilteredProjects).ServeHTTP,
	},
	arbor.Route{
		"GetProject",
		"GET",
		"/project/{name}/",
		alice.New(middleware.IdentificationMiddleware).ThenFunc(GetProject).ServeHTTP,
	},
	arbor.Route{
		"DeleteProject",
		"DELETE",
		"/project/{name}/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole}), middleware.IdentificationMiddleware).ThenFunc(DeleteProject).ServeHTTP,
	},
	arbor.Route{
		"GetAllProjects",
		"GET",
		"/project/",
		alice.New(middleware.IdentificationMiddleware).ThenFunc(GetProject).ServeHTTP,
	},
	arbor.Route{
		"CreateProject",
		"POST",
		"/project/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole}), middleware.IdentificationMiddleware).ThenFunc(CreateProject).ServeHTTP,
	},
	arbor.Route{
		"UpdateProject",
		"PUT",
		"/project/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole}), middleware.IdentificationMiddleware).ThenFunc(UpdateProject).ServeHTTP,
	},
}

func GetProject(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.PROJECT_SERVICE+r.URL.String(), ProjectFormat, "", r)
}

func GetFilteredProjects(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.PROJECT_SERVICE+r.URL.String(), ProjectFormat, "", r)
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	arbor.DELETE(w, config.PROJECT_SERVICE+r.URL.String(), ProjectFormat, "", r)
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, config.PROJECT_SERVICE+r.URL.String(), ProjectFormat, "", r)
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	arbor.PUT(w, config.PROJECT_SERVICE+r.URL.String(), ProjectFormat, "", r)
}

func GetProjectFavorites(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.PROJECT_SERVICE+r.URL.String(), ProjectFormat, "", r)
}

func AddProjectFavorite(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, config.PROJECT_SERVICE+r.URL.String(), ProjectFormat, "", r)
}

func RemoveProjectFavorite(w http.ResponseWriter, r *http.Request) {
	arbor.DELETE(w, config.PROJECT_SERVICE+r.URL.String(), ProjectFormat, "", r)
}
