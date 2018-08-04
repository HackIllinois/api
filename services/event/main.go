package main

import (
	"github.com/HackIllinois/api-commons/middleware"
	"github.com/HackIllinois/api-event/config"
	"github.com/HackIllinois/api-event/controller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	controller.SetupController(router.PathPrefix("/event"))

	router.Use(middleware.ErrorMiddleware)
	router.Use(middleware.ContentTypeMiddleware)
	log.Fatal(http.ListenAndServe(config.EVENT_PORT, router))
}
