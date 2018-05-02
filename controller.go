package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"./models"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.Handle("/github/", alice.New().ThenFunc(LoginGithub)).Methods("POST")
}

func LoginGithub(w http.ResponseWriter, r *http.Request) {
	// TODO: Login github here
	var token models.Token
	json.NewEncoder(w).Encode(token)
}
