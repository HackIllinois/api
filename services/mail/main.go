package main

import (
	"github.com/ethan-lord/api/common/middleware"
	"github.com/ethan-lord/api/services/mail/config"
	"github.com/ethan-lord/api/services/mail/controller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	controller.SetupController(router.PathPrefix("/mail"))

	router.Use(middleware.ErrorMiddleware)
	router.Use(middleware.ContentTypeMiddleware)
	log.Fatal(http.ListenAndServe(config.MAIL_PORT, router))
}
