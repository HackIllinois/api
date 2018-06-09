package main

import (
	"github.com/HackIllinois/api-upload/config"
	"github.com/HackIllinois/api-upload/controller"
	"github.com/HackIllinois/api-commons/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	controller.SetupController(router.PathPrefix("/upload"))

	router.Use(middleware.ErrorMiddleware)
	router.Use(middleware.ContentTypeMiddleware)
	log.Fatal(http.ListenAndServe(config.UPLOAD_PORT, router))
}
