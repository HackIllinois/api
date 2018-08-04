package controller

import (
	"encoding/json"
	"net/http"

	"github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/services/event/models"
	"github.com/HackIllinois/api/services/event/service"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.Handle("/{name}/", alice.New().ThenFunc(GetEvent)).Methods("GET")
	router.Handle("/{name}/", alice.New().ThenFunc(DeleteEvent)).Methods("DELETE")
	router.Handle("/", alice.New().ThenFunc(CreateEvent)).Methods("POST")
	router.Handle("/", alice.New().ThenFunc(UpdateEvent)).Methods("PUT")
	router.Handle("/", alice.New().ThenFunc(GetAllEvents)).Methods("GET")

	router.Handle("/track/", alice.New().ThenFunc(MarkUserAsAttendingEvent)).Methods("POST")
	router.Handle("/track/event/{name}/", alice.New().ThenFunc(GetEventTrackingInfo)).Methods("GET")
	router.Handle("/track/user/{id}/", alice.New().ThenFunc(GetUserTrackingInfo)).Methods("GET")
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
	Endpoint to delete an event with the specified name.
	It removes the event from the event trackers, and every user's tracker.
	On successful deletion, it returns the event that was deleted.
*/
func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]

	event, err := service.DeleteEvent(name)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(event)
}

/*
	Endpoint to get all events
*/
func GetAllEvents(w http.ResponseWriter, r *http.Request) {
	event_list, err := service.GetAllEvents()

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(event_list)
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

/*
	Endpoint to get tracking info by event
*/
func GetEventTrackingInfo(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]

	tracker, err := service.GetEventTracker(name)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(tracker)
}

/*
	Endpoint to get tracking info by user
*/
func GetUserTrackingInfo(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	tracker, err := service.GetUserTracker(id)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(tracker)
}

/*
	Mark a user as attending an event
*/
func MarkUserAsAttendingEvent(w http.ResponseWriter, r *http.Request) {
	var tracking_info models.TrackingInfo
	json.NewDecoder(r.Body).Decode(&tracking_info)

	is_checkedin, err := service.IsUserCheckedIn(tracking_info.UserID)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	if !is_checkedin {
		panic(errors.UnprocessableError("User must be checked in to attend event"))
	}

	err = service.MarkUserAsAttendingEvent(tracking_info.EventName, tracking_info.UserID)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	event_tracker, err := service.GetEventTracker(tracking_info.EventName)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	user_tracker, err := service.GetUserTracker(tracking_info.UserID)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	tracking_status := &models.TrackingStatus{
		EventTracker: *event_tracker,
		UserTracker:  *user_tracker,
	}

	json.NewEncoder(w).Encode(tracking_status)
}
