package main

import (
	"github.com/HackIllinois/api-commons/middleware"
	"github.com/HackIllinois/api-rsvp/config"
	"github.com/HackIllinois/api-rsvp/controller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	controller.SetupController(router.PathPrefix("/rsvp"))

	router.Use(middleware.ErrorMiddleware)
	router.Use(middleware.ContentTypeMiddleware)
	log.Fatal(http.ListenAndServe(config.RSVP_PORT, router))
}
