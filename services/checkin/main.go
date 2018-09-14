package main

import (
	"github.com/pattyjogal/api/common/middleware"
	"github.com/pattyjogal/api/services/checkin/config"
	"github.com/pattyjogal/api/services/checkin/controller"
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
