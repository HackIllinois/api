package service

import (
	"errors"
	"time"

	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/common/utils"
	"github.com/HackIllinois/api/services/event/config"
	"github.com/HackIllinois/api/services/event/models"
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

var db database.Database

func Initialize() error {
	if db != nil {
		db.Close()
		db = nil
	}

	var err error
	db, err = database.InitDatabase(config.EVENT_DB_HOST, config.EVENT_DB_NAME)

	if err != nil {
		return err
	}

	validate = validator.New()

	return nil
}

/*
	Returns the event with the given id
*/
func GetEvent(id string) (*models.Event, error) {
	query := database.QuerySelector{
		"id": id,
	}

	var event models.Event
	err := db.FindOne("events", query, &event)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

/*
	Deletes the event with the given id.
	Removes the event from event trackers and every user's tracker.
	Returns the event that was deleted.
*/
func DeleteEvent(id string) (*models.Event, error) {

	// Gets event to be able to return it later

	event, err := GetEvent(id)

	if err != nil {
		return nil, err
	}

	query := database.QuerySelector{
		"id": id,
	}

	// Remove event from events database

	err = db.RemoveOne("events", query)

	if err != nil {
		return nil, err
	}

	// Remove from event trackers database

	event_selector := database.QuerySelector{
		"eventid": id,
	}

	err = db.RemoveOne("eventtrackers", event_selector)

	if err != nil {
		return nil, err
	}

	// Find all elements, and remove `id` from the Events slice
	// All the updates are individually atomic

	update_expression := database.QuerySelector{
		"$pull": database.QuerySelector{
			"events": id,
		},
	}

	_, err = db.UpdateAll("usertrackers", nil, &update_expression)

	return event, err
}

/*
	Returns all the events
*/
func GetAllEvents() (*models.EventList, error) {
	events := []models.Event{}
	// nil implies there are no filters on the query, therefore everything in the "events" collection is returned.
	err := db.FindAll("events", nil, &events)

	if err != nil {
		return nil, err
	}

	event_list := models.EventList{
		Events: events,
	}

	return &event_list, nil
}

/*
	Returns all the events
*/
func GetFilteredEvents(parameters map[string][]string) (*models.EventList, error) {
	query, err := database.CreateFilterQuery(parameters, models.Event{})

	if err != nil {
		return nil, err
	}

	events := []models.Event{}
	filtered_events := models.EventList{Events: events}
	err = db.FindAll("events", query, &filtered_events.Events)

	if err != nil {
		return nil, err
	}

	return &filtered_events, nil
}

/*
	Creates an event with the given id
*/
func CreateEvent(id string, event models.Event) error {
	err := validate.Struct(event)

	if err != nil {
		return err
	}

	_, err = GetEvent(id)

	if err != database.ErrNotFound {
		if err != nil {
			return err
		}
		return errors.New("Event already exists")
	}

	err = db.Insert("events", &event)

	if err != nil {
		return err
	}

	event_tracker := models.EventTracker{
		EventID: id,
		Users:   []string{},
	}

	err = db.Insert("eventtrackers", &event_tracker)

	return err
}

/*
	Updates the event with the given id
*/
func UpdateEvent(id string, event models.Event) error {
	err := validate.Struct(event)

	if err != nil {
		return err
	}

	selector := database.QuerySelector{
		"id": id,
	}

	err = db.Update("events", selector, &event)

	return err
}

/*
	Returns the event tracker for the specified event
*/
func GetEventTracker(event_id string) (*models.EventTracker, error) {
	query := database.QuerySelector{
		"eventid": event_id,
	}

	var tracker models.EventTracker
	err := db.FindOne("eventtrackers", query, &tracker)

	if err != nil {
		return nil, err
	}

	return &tracker, nil
}

/*
	Returns the user tracker for the specified user
*/
func GetUserTracker(user_id string) (*models.UserTracker, error) {
	query := database.QuerySelector{
		"userid": user_id,
	}

	var tracker models.UserTracker
	err := db.FindOne("usertrackers", query, &tracker)

	if err != nil {
		if err == database.ErrNotFound {
			return &models.UserTracker{
				UserID: user_id,
				Events: []string{},
			}, nil
		}
		return nil, err
	}

	return &tracker, nil
}

/*
	Returns true is the user has already been marked as attending
	the specified event, false otherwise
*/
func IsUserAttendingEvent(event_id string, user_id string) (bool, error) {
	tracker, err := GetEventTracker(event_id)

	if err != nil {
		return false, err
	}

	for _, id := range tracker.Users {
		if user_id == id {
			return true, nil
		}
	}

	return false, nil
}

/*
	Marks the specified user as attending the specified event
	The user must not already marked as attending for this to return successfully
*/
func MarkUserAsAttendingEvent(event_id string, user_id string) error {
	is_attending, err := IsUserAttendingEvent(event_id, user_id)

	if err != nil {
		return err
	}

	if is_attending {
		return errors.New("User has already been marked as attending")
	}

	if config.EVENT_CHECKIN_TIME_RESTRICTED {
		is_event_active, err := IsEventActive(event_id)

		if err != nil {
			return err
		}

		if !is_event_active {
			return errors.New("People cannot be checked-in for the event at this time.")
		}
	}

	event_selector := database.QuerySelector{
		"eventid": event_id,
	}

	event_modifier := database.QuerySelector{
		"$addToSet": database.QuerySelector{
			"users": user_id,
		},
	}

	err = db.Update("eventtrackers", event_selector, &event_modifier)

	if err != nil {
		return err
	}

	user_selector := database.QuerySelector{
		"userid": user_id,
	}

	user_modifier := database.QuerySelector{
		"$addToSet": database.QuerySelector{
			"events": event_id,
		},
	}

	err = db.Update("usertrackers", user_selector, &user_modifier)

	if err == database.ErrNotFound {
		user_tracker := models.UserTracker{
			UserID: user_id,
			Events: []string{event_id},
		}
		err = db.Insert("usertrackers", &user_tracker)
	}

	return err
}

const PreEventCheckinIntervalInMinutes = 15
const PreEventCheckinIntervalInSeconds = PreEventCheckinIntervalInMinutes * 60

/*
	Check if an event is active, i.e., that check-ins are allowed for the event at the current time.
	Returns true if the current time is between `PreEventCheckinIntervalInMinutes` number of minutes before the event, and the end of event.
*/
func IsEventActive(event_id string) (bool, error) {
	event, err := GetEvent(event_id)

	if err != nil {
		return false, err
	}

	start_time := event.StartTime
	end_time := event.EndTime
	current_time := time.Now().Unix()

	if current_time < start_time {
		return start_time-current_time <= PreEventCheckinIntervalInSeconds, nil
	} else {
		return current_time < end_time, nil
	}
}

/*
	Returns the event favorites for the user with the given id
*/
func GetEventFavorites(id string) (*models.EventFavorites, error) {
	query := database.QuerySelector{
		"id": id,
	}

	var event_favorites models.EventFavorites
	err := db.FindOne("favorites", query, &event_favorites)

	if err != nil {
		if err == database.ErrNotFound {
			err = db.Insert("favorites", &models.EventFavorites{
				ID:     id,
				Events: []string{},
			})

			if err != nil {
				return nil, err
			}

			err = db.FindOne("favorites", query, &event_favorites)

			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return &event_favorites, nil
}

/*
	Adds the given event to the favorites for the user with the given id
*/
func AddEventFavorite(id string, event string) error {
	selector := database.QuerySelector{
		"id": id,
	}

	_, err := GetEvent(event)

	if err != nil {
		return errors.New("Could not find event with the given id.")
	}

	event_favorites, err := GetEventFavorites(id)

	if err != nil {
		return err
	}

	if !utils.ContainsString(event_favorites.Events, event) {
		event_favorites.Events = append(event_favorites.Events, event)
	}

	err = db.Update("favorites", selector, event_favorites)

	return err
}

/*
	Removes the given event to the favorites for the user with the given id
*/
func RemoveEventFavorite(id string, event string) error {
	selector := database.QuerySelector{
		"id": id,
	}

	event_favorites, err := GetEventFavorites(id)

	if err != nil {
		return err
	}

	event_favorites.Events, err = utils.RemoveString(event_favorites.Events, event)

	if err != nil {
		return errors.New("User's event favorites does not have specified event")
	}

	err = db.Update("favorites", selector, event_favorites)

	return err
}

/*
	Returns all event stats
*/
func GetStats() (map[string]interface{}, error) {
	query := database.QuerySelector{}

	var trackers []models.EventTracker
	err := db.FindAll("eventtrackers", query, &trackers)

	if err != nil {
		return nil, err
	}

	stats := make(map[string]interface{})

	for _, tracker := range trackers {
		stats[tracker.EventID] = len(tracker.Users)
	}

	return stats, nil
}
