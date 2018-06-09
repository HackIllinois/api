package main

import (
	"github.com/HackIllinois/api-checkin/config"
	"github.com/HackIllinois/api-checkin/controller"
	"github.com/HackIllinois/api-commons/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	controller.SetupController(router.PathPrefix("/checkin"))

	router.Use(middleware.ErrorMiddleware)
	router.Use(middleware.ContentTypeMiddleware)
	log.Fatal(http.ListenAndServe(config.CHECKIN_PORT, router))
}
