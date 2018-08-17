package main

import (
	"github.com/ReflectionsProjections/api/common/middleware"
	"github.com/ReflectionsProjections/api/services/event/config"
	"github.com/ReflectionsProjections/api/services/event/controller"
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
