package main

import (
	"github.com/HackIllinois/api/common/middleware"
	"github.com/HackIllinois/api/services/print/config"
	"github.com/HackIllinois/api/services/print/controller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	controller.SetupController(router.PathPrefix("/print"))

	router.Use(middleware.ErrorMiddleware)
	router.Use(middleware.ContentTypeMiddleware)
	log.Fatal(http.ListenAndServe(config.PRINT_PORT, router))
}
