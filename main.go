package main

import (
	"github.com/HackIllinois/api-commons/middleware"
	"github.com/HackIllinois/api-registration/controller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	controller.SetupController(router.PathPrefix("/registration"))

	router.Use(middleware.ErrorMiddleware)
	router.Use(middleware.ContentTypeMiddleware)
	log.Fatal(http.ListenAndServe(":8004", router))
}
