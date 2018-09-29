package main

import (
	"github.com/ethan-lord/api/common/middleware"
	"github.com/ethan-lord/api/services/auth/config"
	"github.com/ethan-lord/api/services/auth/controller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	controller.SetupController(router.PathPrefix("/auth"))

	router.Use(middleware.ErrorMiddleware)
	router.Use(middleware.ContentTypeMiddleware)
	log.Fatal(http.ListenAndServe(config.AUTH_PORT, router))
}
