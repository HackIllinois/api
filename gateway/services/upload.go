package services

import (
	"github.com/HackIllinois/api/gateway/config"
	"github.com/HackIllinois/api/gateway/middleware"
	"github.com/HackIllinois/api/gateway/models"
	"github.com/arbor-dev/arbor"
	"github.com/justinas/alice"
	"net/http"
)

var UploadURL = config.UPLOAD_SERVICE

const UploadFormat string = "RAW"
const InfoFormat string = "JSON"

var UploadRoutes = arbor.RouteCollection{
	arbor.Route{
		"GetCurrentUploadInfo",
		"GET",
		"/upload/resume/",
		alice.New(middleware.AuthMiddleware([]string{models.UserRole}), middleware.IdentificationMiddleware).ThenFunc(GetCurrentUploadInfo).ServeHTTP,
	},
	arbor.Route{
		"UpdateCurrentUploadInfo",
		"PUT",
		"/upload/resume/",
		alice.New(middleware.AuthMiddleware([]string{models.UserRole}), middleware.IdentificationMiddleware).ThenFunc(UpdateCurrentUploadInfo).ServeHTTP,
	},
	arbor.Route{
		"GetUploadInfo",
		"GET",
		"/upload/resume/{id}/",
		alice.New(middleware.AuthMiddleware([]string{models.AdminRole, models.StaffRole}), middleware.IdentificationMiddleware).ThenFunc(GetUploadInfo).ServeHTTP,
	},
}

func GetCurrentUploadInfo(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, UploadURL+r.URL.String(), InfoFormat, "", r)
}

func UpdateCurrentUploadInfo(w http.ResponseWriter, r *http.Request) {
	arbor.PUT(w, UploadURL+r.URL.String(), UploadFormat, "", r)
}

func GetUploadInfo(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, UploadURL+r.URL.String(), InfoFormat, "", r)
}
