package main

import (
	"github.com/HackIllinois/api/common/middleware"
	"github.com/HackIllinois/api/services/user/config"
	"github.com/HackIllinois/api/services/user/controller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	controller.SetupController(router.PathPrefix("/user"))

	router.Use(middleware.ErrorMiddleware)
	router.Use(middleware.ContentTypeMiddleware)
	log.Fatal(http.ListenAndServe(config.USER_PORT, router))
}
