package service

import (
	"errors"
	"github.com/HackIllinois/api-commons/database"
	"github.com/HackIllinois/api-event/config"
	"github.com/HackIllinois/api-event/models"
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
		"userId": user_id,
	}

	var tracker models.UserTracker
	err := db.FindOne("usertrackers", query, &tracker)

	if err != nil {
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
	The user must be checkedin and not already marked as attending
	for this to return successfully
*/
func MarkUserAsAttendingEvent(event_name string, user_id string) error {
	is_attending, err := IsUserAttendingEvent(event_name, user_id)

	if err != nil {
		return err
	}

	if is_attending {
		return errors.New("User has already been marked as attending")
	}

	is_checkedin, err := IsUserCheckedIn(user_id)

	if err != nil {
		return err
	}

	if !is_checkedin {
		return errors.New("User must be checked in to attend event")
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

	return err
}
