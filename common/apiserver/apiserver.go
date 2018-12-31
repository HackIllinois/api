package apiserver

import (
	"time"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/HackIllinois/api/common/middleware"
)

func StartServer(address string, router *mux.Router, name string) error {
	router.Use(middleware.ErrorMiddleware)
	router.Use(middleware.ContentTypeMiddleware)

	server := &http.Server{
		Handler:      router,
		Addr:         address,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	return server.ListenAndServe()
}
