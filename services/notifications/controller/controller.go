package controller

import (
	"encoding/json"
	"github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/services/notifications/service"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"net/http"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.Handle("/topics/", alice.New().ThenFunc(GetAllTopics)).Methods("GET")
}

/*
	Endpoint to get all SNS Topics
*/
func GetAllTopics(w http.ResponseWriter, r *http.Request) {
	topics, err := service.GetAllTopics()

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(topics)
}
