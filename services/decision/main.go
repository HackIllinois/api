package main

import (
	"github.com/ethan-lord/api/common/middleware"
	"github.com/ethan-lord/api/services/decision/config"
	"github.com/ethan-lord/api/services/decision/controller"
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
