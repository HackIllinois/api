package service

import (
	"errors"
	"github.com/HackIllinois/api-commons/database"
	"github.com/HackIllinois/api-registration/config"
	"github.com/HackIllinois/api-registration/models"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"strconv"
	"strings"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

var db database.MongoDatabase

func init() {
	db_connection, err := database.InitMongoDatabase(config.REGISTRATION_DB_HOST, config.REGISTRATION_DB_NAME)

	if err != nil {
		panic(err)
	}

	db = db_connection
}

/*
	Returns the registration associated with the given user id
*/
func GetUserRegistration(id string) (*models.UserRegistration, error) {
	query := bson.M{"id": id}

	var user_registration models.UserRegistration
	err := db.FindOne("attendees", query, &user_registration)

	if err != nil {
		return nil, err
	}

	return &user_registration, nil
}

/*
	Creates the registration associated with the given user id
*/
func CreateUserRegistration(id string, user_registration models.UserRegistration) error {
	err := validate.Struct(user_registration)

	if err != nil {
		return err
	}

	_, err = GetUserRegistration(id)

	if err != mgo.ErrNotFound {
		if err != nil {
			return err
		}
		return errors.New("Registration already exists")
	}

	err = db.Insert("attendees", &user_registration)

	return err
}

/*
	Updates the registration associated with the given user id
*/
func UpdateUserRegistration(id string, user_registration models.UserRegistration) error {
	err := validate.Struct(user_registration)

	if err != nil {
		return err
	}

	selector := bson.M{"id": id}

	err = db.Update("attendees", selector, &user_registration)

	return err
}

func AssignValueType(key, value string) (interface{}, error) {
	if key == "age" || key == "graduationyear" {
		return strconv.Atoi(value)
	} else if key == "isnovice" || key == "isprivate" {
		return strconv.ParseBool(value)
	} else {
		return value, nil
	}
}

/*
	Returns the registrations associated with the given parameters
*/
func GetFilteredUserRegistrations(parameters map[string][]string) (*[]models.UserRegistration, error) {
	// Build query
	query := make(map[string]interface{})
	for key, values := range parameters {
		if len(values) == 1 {
			key = strings.ToLower(key)

			// Handle multiple comma separated values
			value_list := strings.Split(values[0], ",")

			// Assign correct type to values
			correctly_typed_value_list := make([]interface{}, len(value_list))
			for i, value := range value_list {
				correctly_typed_value, err := AssignValueType(key, value)
				if err == nil {
					correctly_typed_value_list[i] = correctly_typed_value
				} else {
					return nil, err
				}
			}

			// Update query
			query[key] = bson.M{"$in": correctly_typed_value_list}
		} else {
			return nil, errors.New("Multiple usage of key "+key)
		}
	}

	// Run query
	var user_registrations []models.UserRegistration
	err := db.FindAll("attendees", query, &user_registrations)
	if err != nil {
		return nil, err
	}

	return &user_registrations, nil
}
