package main

import (
	"github.com/ethan-lord/api/common/middleware"
	"github.com/ethan-lord/api/services/upload/config"
	"github.com/ethan-lord/api/services/upload/controller"
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
