package services

import (
	"net/http"

	"github.com/HackIllinois/api/common/authtoken"
	"github.com/HackIllinois/api/gateway/config"
	"github.com/HackIllinois/api/gateway/middleware"
	"github.com/arbor-dev/arbor"
	"github.com/justinas/alice"
)

const UploadFormat string = "RAW"
const InfoFormat string = "JSON"

var UploadRoutes = arbor.RouteCollection{
	arbor.Route{
		"GetCurrentUploadInfo",
		"GET",
		"/upload/resume/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.UserRole}), middleware.IdentificationMiddleware).ThenFunc(GetCurrentUploadInfo).ServeHTTP,
	},
	arbor.Route{
		"UpdateCurrentUploadInfo",
		"GET",
		"/upload/resume/upload/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.UserRole}), middleware.IdentificationMiddleware).ThenFunc(UpdateCurrentUploadInfo).ServeHTTP,
	},
	arbor.Route{
		"GetUploadInfo",
		"GET",
		"/upload/resume/{id}/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole, authtoken.StaffRole}), middleware.IdentificationMiddleware).ThenFunc(GetUploadInfo).ServeHTTP,
	},
	arbor.Route{
		"GetCurrentUploadInfo",
		"GET",
		"/upload/photo/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.UserRole}), middleware.IdentificationMiddleware).ThenFunc(GetCurrentUploadInfo).ServeHTTP,
	},
	arbor.Route{
		"UpdateCurrentUploadInfo",
		"GET",
		"/upload/photo/upload/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.UserRole}), middleware.IdentificationMiddleware).ThenFunc(UpdateCurrentUploadInfo).ServeHTTP,
	},
	arbor.Route{
		"GetUploadInfo",
		"GET",
		"/upload/photo/{id}/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole, authtoken.StaffRole}), middleware.IdentificationMiddleware).ThenFunc(GetUploadInfo).ServeHTTP,
	},
	arbor.Route{
		"CreateBlob",
		"POST",
		"/upload/blobstore/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole, authtoken.StaffRole, authtoken.BlobstoreRole}), middleware.IdentificationMiddleware).ThenFunc(CreateBlob).ServeHTTP,
	},
	arbor.Route{
		"UpdateBlob",
		"PUT",
		"/upload/blobstore/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole, authtoken.StaffRole, authtoken.BlobstoreRole}), middleware.IdentificationMiddleware).ThenFunc(UpdateBlob).ServeHTTP,
	},
	arbor.Route{
		"GetBlob",
		"GET",
		"/upload/blobstore/{id}/",
		alice.New(middleware.IdentificationMiddleware).ThenFunc(GetBlob).ServeHTTP,
	},
	arbor.Route{
		"DeleteBlob",
		"DELETE",
		"/upload/blobstore/{id}/",
		alice.New(middleware.AuthMiddleware([]authtoken.Role{authtoken.AdminRole}), middleware.IdentificationMiddleware).ThenFunc(DeleteBlob).ServeHTTP,
	},
}

func GetCurrentUploadInfo(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.UPLOAD_SERVICE+r.URL.String(), InfoFormat, "", r)
}

func UpdateCurrentUploadInfo(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.UPLOAD_SERVICE+r.URL.String(), UploadFormat, "", r)
}

func GetUploadInfo(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.UPLOAD_SERVICE+r.URL.String(), InfoFormat, "", r)
}

func CreateBlob(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, config.UPLOAD_SERVICE+r.URL.String(), InfoFormat, "", r)
}

func UpdateBlob(w http.ResponseWriter, r *http.Request) {
	arbor.PUT(w, config.UPLOAD_SERVICE+r.URL.String(), InfoFormat, "", r)
}

func GetBlob(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.UPLOAD_SERVICE+r.URL.String(), InfoFormat, "", r)
}

func DeleteBlob(w http.ResponseWriter, r *http.Request) {
	arbor.DELETE(w, config.UPLOAD_SERVICE+r.URL.String(), InfoFormat, "", r)
}
