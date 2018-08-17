package main

import (
	"github.com/ReflectionsProjections/api/common/middleware"
	"github.com/ReflectionsProjections/api/services/mail/config"
	"github.com/ReflectionsProjections/api/services/mail/controller"
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
