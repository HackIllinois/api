package services

import (
	"net/http"

	"github.com/HackIllinois/api/gateway/config"
	"github.com/HackIllinois/api/gateway/middleware"
	"github.com/HackIllinois/api/gateway/models"
	"github.com/arbor-dev/arbor"
	"github.com/justinas/alice"
)

const RecognitionFormat string = "JSON"

var RecognitionRoutes = arbor.RouteCollection{
	arbor.Route{
		"GetAllRecognitions",
		"GET",
		"/recognition/",
		alice.New(middleware.IdentificationMiddleware).ThenFunc(GetRecognition).ServeHTTP,
	},
	arbor.Route{
		"CreateRecognition",
		"POST",
		"/recognition/",
		alice.New(middleware.AuthMiddleware([]models.Role{models.AdminRole}), middleware.IdentificationMiddleware).ThenFunc(CreateRecognition).ServeHTTP,
	},
}

func GetRecognition(w http.ResponseWriter, r *http.Request) {
	arbor.GET(w, config.RECOGNITION_SERVICE+r.URL.String(), RecognitionFormat, "", r)
}

func CreateRecognition(w http.ResponseWriter, r *http.Request) {
	arbor.POST(w, config.RECOGNITION_SERVICE+r.URL.String(), RecognitionFormat, "", r)
}
