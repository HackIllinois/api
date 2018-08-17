package main

import (
	"github.com/ReflectionsProjections/api/common/middleware"
	"github.com/ReflectionsProjections/api/services/stat/config"
	"github.com/ReflectionsProjections/api/services/stat/controller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	controller.SetupController(router.PathPrefix("/stat"))

	router.Use(middleware.ErrorMiddleware)
	router.Use(middleware.ContentTypeMiddleware)
	log.Fatal(http.ListenAndServe(config.STAT_PORT, router))
}
