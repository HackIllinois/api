package main

import (
	"github.com/pattyjogal/api/common/middleware"
	"github.com/pattyjogal/api/services/mail/config"
	"github.com/pattyjogal/api/services/mail/controller"
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
