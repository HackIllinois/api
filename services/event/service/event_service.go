package service

import (
	"errors"

	"github.com/ReflectionsProjections/api/common/database"
	"github.com/ReflectionsProjections/api/services/event/config"
	"github.com/ReflectionsProjections/api/services/event/models"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

var db database.MongoDatabase

func init() {
	db_connection, err := database.InitMongoDatabase(config.EVENT_DB_HOST, config.EVENT_DB_NAME)

	if err != nil {
		panic(err)
	}

	db = db_connection
}

/*
	Returns the event with the given name
*/
func GetEvent(name string) (*models.Event, error) {
	query := bson.M{
		"name": name,
	}

	var event models.Event
	err := db.FindOne("events", query, &event)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

/*
	Deletes the event with the given name.
	Removes the event from event trackers and every user's tracker.
	Returns the event that was deleted.
*/
func DeleteEvent(name string) (*models.Event, error) {

	// Gets event to be able to return it later

	event, err := GetEvent(name)

	if err != nil {
		return nil, err
	}

	query := bson.M{
		"name": name,
	}

	// Remove event from events database

	err = db.RemoveOne("events", query)

	// Remove from event trackers database

	event_selector := bson.M{
		"eventname": name,
	}

	err = db.RemoveOne("eventtrackers", event_selector)

	if err != nil {
		return nil, err
	}

	// Find all elements, and remove `name` from the Events slice
	// All the updates are individually atomic

	update_expression := bson.M{
		"$pull": bson.M{
			"events": name,
		},
	}

	_, err = db.UpdateAll("usertrackers", nil, &update_expression)

	return event, err
}

/*
	Returns all the events
*/
func GetAllEvents() (*models.EventList, error) {
	var events []models.Event
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
	Creates an event with the given name
*/
func CreateEvent(name string, event models.Event) error {
	err := validate.Struct(event)

	if err != nil {
		return err
	}

	_, err = GetEvent(name)

	if err != mgo.ErrNotFound {
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
		EventName: name,
		Users:     []string{},
	}

	err = db.Insert("eventtrackers", &event_tracker)

	return err
}

/*
	Updates the event with the given name
*/
func UpdateEvent(name string, event models.Event) error {
	err := validate.Struct(event)

	if err != nil {
		return err
	}

	selector := bson.M{
		"name": name,
	}

	err = db.Update("events", selector, &event)

	return err
}

/*
	Returns the event tracker for the specified event
*/
func GetEventTracker(event_name string) (*models.EventTracker, error) {
	query := bson.M{
		"eventname": event_name,
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
	query := bson.M{
		"userid": user_id,
	}

	var tracker models.UserTracker
	err := db.FindOne("usertrackers", query, &tracker)

	if err != nil {
		if err == mgo.ErrNotFound {
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
func IsUserAttendingEvent(event_name string, user_id string) (bool, error) {
	tracker, err := GetEventTracker(event_name)

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
func MarkUserAsAttendingEvent(event_name string, user_id string) error {
	is_attending, err := IsUserAttendingEvent(event_name, user_id)

	if err != nil {
		return err
	}

	if is_attending {
		return errors.New("User has already been marked as attending")
	}

	event_selector := bson.M{
		"eventname": event_name,
	}

	event_modifier := bson.M{
		"$addToSet": bson.M{
			"users": user_id,
		},
	}

	err = db.Update("eventtrackers", event_selector, &event_modifier)

	if err != nil {
		return err
	}

	user_selector := bson.M{
		"userid": user_id,
	}

	user_modifier := bson.M{
		"$addToSet": bson.M{
			"events": event_name,
		},
	}

	err = db.Update("usertrackers", user_selector, &user_modifier)

	if err == mgo.ErrNotFound {
		user_tracker := models.UserTracker{
			UserID: user_id,
			Events: []string{event_name},
		}
		err = db.Insert("usertrackers", &user_tracker)
	}

	return err
}
