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
