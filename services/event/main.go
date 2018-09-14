package main

import (
	"github.com/pattyjogal/api/common/middleware"
	"github.com/pattyjogal/api/services/event/config"
	"github.com/pattyjogal/api/services/event/controller"
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
