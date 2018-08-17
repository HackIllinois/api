package services

import (
	"github.com/ReflectionsProjections/api/gateway/config"
	"github.com/ReflectionsProjections/api/gateway/middleware"
	"github.com/arbor-dev/arbor"
	"github.com/justinas/alice"
	"net/http"
)

var UploadURL = config.UPLOAD_SERVICE

const UploadFormat string = "JSON"

var UploadRoutes = arbor.RouteCollection{
	arbor.Route{
		"GetCurrentUploadInfo",
		"GET",
		"/upload/resume/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"User"})).ThenFunc(GetCurrentUploadInfo).ServeHTTP,
	},
	arbor.Route{
		"UpdateCurrentUploadInfo",
		"PUT",
		"/upload/resume/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"User"})).ThenFunc(UpdateCurrentUploadInfo).ServeHTTP,
	},
	arbor.Route{
		"GetUploadInfo",
		"GET",
		"/upload/resume/{id}/",
		alice.New(middleware.IdentificationMiddleware, middleware.AuthMiddleware([]string{"Admin"})).ThenFunc(GetUploadInfo).ServeHTTP,
	},
}

func GetCurrentUploadInfo(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, UploadURL+r.URL.String(), UploadFormat, "", r)
}

func UpdateCurrentUploadInfo(w http.ResponseWriter, r *http.Request) {
	arbor.PUT(w, UploadURL+r.URL.String(), UploadFormat, "", r)
}

func GetUploadInfo(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, UploadURL+r.URL.String(), UploadFormat, "", r)
}
