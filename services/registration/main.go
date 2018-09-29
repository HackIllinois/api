package main

import (
	"github.com/ethan-lord/api/common/middleware"
	"github.com/ethan-lord/api/services/registration/config"
	"github.com/ethan-lord/api/services/registration/controller"
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
