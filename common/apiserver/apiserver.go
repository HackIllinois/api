package apiserver

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/HackIllinois/api/common/middleware"
)

func StartServer(address string, router *mux.Router) error {
	router.Use(middleware.ErrorMiddleware)
	router.Use(middleware.ContentTypeMiddleware)

	return http.ListenAndServe(address, router)
}
