package controller

import (
	"encoding/json"
	"github.com/HackIllinois/api-commons/errors"
	"github.com/HackIllinois/api-event/models"
	"github.com/HackIllinois/api-event/service"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"net/http"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.Handle("/{name}/", alice.New().ThenFunc(GetEvent)).Methods("GET")
	router.Handle("/", alice.New().ThenFunc(CreateEvent)).Methods("POST")
	router.Handle("/", alice.New().ThenFunc(UpdateEvent)).Methods("PUT")
}

/*
	Endpoint to get the event with the specified name
*/
func GetEvent(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]

	event, err := service.GetEvent(name)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(event)
}

/*
	Endpoint to create an event
*/
func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event models.Event
	json.NewDecoder(r.Body).Decode(&event)

	err := service.CreateEvent(event.Name, event)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	updated_event, err := service.GetEvent(event.Name)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(updated_event)
}

/*
	Endpoint to update an event
*/
func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	var event models.Event
	json.NewDecoder(r.Body).Decode(&event)

	err := service.UpdateEvent(event.Name, event)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	updated_event, err := service.GetEvent(event.Name)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(updated_event)
}
