package controller

import (
	"encoding/json"
	"net/http"

	"github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/common/utils"
	"github.com/HackIllinois/api/services/event/models"
	"github.com/HackIllinois/api/services/event/service"
	"github.com/gorilla/mux"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.HandleFunc("/favorite/", GetEventFavorites).Methods("GET")
	router.HandleFunc("/favorite/add/", AddEventFavorite).Methods("POST")
	router.HandleFunc("/favorite/remove/", RemoveEventFavorite).Methods("POST")

	router.HandleFunc("/filter/", GetFilteredEvents).Methods("GET")
	router.HandleFunc("/{id}/", GetEvent).Methods("GET")
	router.HandleFunc("/{id}/", DeleteEvent).Methods("DELETE")
	router.HandleFunc("/", CreateEvent).Methods("POST")
	router.HandleFunc("/", UpdateEvent).Methods("PUT")
	router.HandleFunc("/", GetAllEvents).Methods("GET")

	router.HandleFunc("/track/", MarkUserAsAttendingEvent).Methods("POST")
	router.HandleFunc("/track/event/{id}/", GetEventTrackingInfo).Methods("GET")
	router.HandleFunc("/track/user/{id}/", GetUserTrackingInfo).Methods("GET")

	router.HandleFunc("/internal/stats/", GetStats).Methods("GET")
}

/*
	Endpoint to get the event with the specified id
*/
func GetEvent(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	event, err := service.GetEvent(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not fetch the event details."))
		return
	}

	json.NewEncoder(w).Encode(event)
}

/*
	Endpoint to delete an event with the specified id.
	It removes the event from the event trackers, and every user's tracker.
	On successful deletion, it returns the event that was deleted.
*/
func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	event, err := service.DeleteEvent(id)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not delete either the event, event trackers, or user trackers, or an intermediary subroutine failed."))
		return
	}

	json.NewEncoder(w).Encode(event)
}

/*
	Endpoint to get all events
*/
func GetAllEvents(w http.ResponseWriter, r *http.Request) {
	event_list, err := service.GetAllEvents()

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get all events."))
		return
	}

	json.NewEncoder(w).Encode(event_list)
}

/*
	Endpoint to get events based on filters
*/
func GetFilteredEvents(w http.ResponseWriter, r *http.Request) {
	parameters := r.URL.Query()
	event, err := service.GetFilteredEvents(parameters)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not fetch filtered list of events."))
		return
	}

	json.NewEncoder(w).Encode(event)
}

/*
	Endpoint to create an event
*/
func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event models.Event
	json.NewDecoder(r.Body).Decode(&event)

	event.ID = utils.GenerateUniqueID()

	err := service.CreateEvent(event.ID, event)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not create new event."))
		return
	}

	updated_event, err := service.GetEvent(event.ID)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get updated event."))
		return
	}

	json.NewEncoder(w).Encode(updated_event)
}

/*
	Endpoint to update an event
*/
func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	var event models.Event
	json.NewDecoder(r.Body).Decode(&event)

	err := service.UpdateEvent(event.ID, event)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not update the event."))
		return
	}

	updated_event, err := service.GetEvent(event.ID)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get updated event details."))
		return
	}

	json.NewEncoder(w).Encode(updated_event)
}

/*
	Endpoint to get tracking info by event
*/
func GetEventTrackingInfo(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	tracker, err := service.GetEventTracker(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get event tracker."))
		return
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
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get user tracker."))
		return
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
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not determine check-in status of user."))
		return
	}

	if !is_checkedin {
		errors.WriteError(w, r, errors.AttributeMismatchError("User must be checked-in to attend event.", "User must be checked-in to attend event."))
		return
	}

	err = service.MarkUserAsAttendingEvent(tracking_info.EventID, tracking_info.UserID)

	if err != nil {
		if err.Error() == "User has already been marked as attending" {
			errors.WriteError(w, r, errors.AttributeMismatchError("User has already checked in.", "User has already checked in."))
		} else if err.Error() == "People cannot be checked-in for the event at this time." {
			errors.WriteError(w, r, errors.AttributeMismatchError("Event is not open for check-in at this time.", "Event is not open for check-in at this time."))
		} else {
			errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not mark user as attending the event."))
		}
		return
	}

	event_tracker, err := service.GetEventTracker(tracking_info.EventID)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get event trackers."))
		return
	}

	user_tracker, err := service.GetUserTracker(tracking_info.UserID)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get user trackers."))
		return
	}

	tracking_status := &models.TrackingStatus{
		EventTracker: *event_tracker,
		UserTracker:  *user_tracker,
	}

	json.NewEncoder(w).Encode(tracking_status)
}

/*
	Endpoint to get the current user's event favorites
*/
func GetEventFavorites(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	favorites, err := service.GetEventFavorites(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get user's event favourites."))
		return
	}

	json.NewEncoder(w).Encode(favorites)
}

/*
	Endpoint to add an event favorite for the current user
*/
func AddEventFavorite(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	var event_favorite_modification models.EventFavoriteModification
	json.NewDecoder(r.Body).Decode(&event_favorite_modification)

	err := service.AddEventFavorite(id, event_favorite_modification.EventID)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not add an event favorite for the current user."))
		return
	}

	favorites, err := service.GetEventFavorites(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get updated user event favorites."))
		return
	}

	json.NewEncoder(w).Encode(favorites)
}

/*
	Endpoint to remove an event favorite for the current user
*/
func RemoveEventFavorite(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	var event_favorite_modification models.EventFavoriteModification
	json.NewDecoder(r.Body).Decode(&event_favorite_modification)

	err := service.RemoveEventFavorite(id, event_favorite_modification.EventID)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not remove an event favorite for the current user."))
		return
	}

	favorites, err := service.GetEventFavorites(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not fetch updated event favourites for the user (post-removal)."))
		return
	}

	json.NewEncoder(w).Encode(favorites)
}

/*
	Endpoint to get event stats
*/
func GetStats(w http.ResponseWriter, r *http.Request) {
	stats, err := service.GetStats()

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not fetch event service statistics."))
		return
	}

	json.NewEncoder(w).Encode(stats)
}
