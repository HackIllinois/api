package main

import (
	"github.com/HackIllinois/api-commons/middleware"
	"github.com/HackIllinois/api-decision/config"
	"github.com/HackIllinois/api-decision/controller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	controller.SetupController(router.PathPrefix("/decision"))

	router.Use(middleware.ErrorMiddleware)
	router.Use(middleware.ContentTypeMiddleware)
	log.Fatal(http.ListenAndServe(config.DECISION_PORT, router))
}
