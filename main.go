package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/hackillinois/api-auth/middleware"
)

func main() {
	router := mux.NewRouter()
	SetupController(router.PathPrefix("/auth"))

	router.Use(middleware.ErrorMiddleware)
	router.Use(middleware.ContentTypeMiddleware)
	log.Fatal(http.ListenAndServe(":8002", router))
}
