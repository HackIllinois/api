package main

import (
	"github.com/ReflectionsProjections/api/common/middleware"
	"github.com/ReflectionsProjections/api/services/registration/config"
	"github.com/ReflectionsProjections/api/services/registration/controller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	controller.SetupController(router.PathPrefix("/registration"))

	router.Use(middleware.ErrorMiddleware)
	router.Use(middleware.ContentTypeMiddleware)
	log.Fatal(http.ListenAndServe(config.REGISTRATION_PORT, router))
}
