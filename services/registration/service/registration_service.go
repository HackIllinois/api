package service

import (
	"errors"
	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/services/registration/config"
	"github.com/HackIllinois/api/services/registration/models"
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

/*
	Returns the registrations associated with the given parameters
*/
func GetFilteredUserRegistrations(parameters map[string][]string) (*models.FilteredRegistrations, error) {
	query := make(map[string]interface{})
	for key, values := range parameters {
		if len(values) == 1 {
			key = strings.ToLower(key)
			value_list := strings.Split(values[0], ",")

			correctly_typed_value_list := make([]interface{}, len(value_list))
			for i, value := range value_list {
				correctly_typed_value, err := AssignValueType(key, value)
				if err == nil {
					correctly_typed_value_list[i] = correctly_typed_value
				} else {
					return nil, err
				}
			}
			query[key] = bson.M{"$in": correctly_typed_value_list}
		} else {
			return nil, errors.New("Multiple usage of key " + key)
		}
	}

	var filtered_registrations models.FilteredRegistrations
	err := db.FindAll("attendees", query, &filtered_registrations.Registrations)
	if err != nil {
		return nil, err
	}

	return &filtered_registrations, nil
}

func AssignValueType(key, value string) (interface{}, error) {
	int_keys := []string{"age", "graduationyear"}
	if Contains(int_keys, key) {
		return strconv.Atoi(value)
	}

	bool_keys := []string{"isnovice", "isprivate"}
	if Contains(bool_keys, key) {
		return strconv.ParseBool(value)
	}

	return value, nil
}

func Contains(slice []string, str string) bool {
	for _, value := range slice {
		if value == str {
			return true
		}
	}
	return false
}

/*
	Returns the registration associated with the given mentor id
*/
func GetMentorRegistration(id string) (*models.MentorRegistration, error) {
	query := bson.M{"id": id}

	var mentor_registration models.MentorRegistration
	err := db.FindOne("mentors", query, &mentor_registration)

	if err != nil {
		return nil, err
	}

	return &mentor_registration, nil
}

/*
	Creates the registration associated with the given mentor id
*/
func CreateMentorRegistration(id string, mentor_registration models.MentorRegistration) error {
	err := validate.Struct(mentor_registration)

	if err != nil {
		return err
	}

	_, err = GetMentorRegistration(id)

	if err != mgo.ErrNotFound {
		if err != nil {
			return err
		}
		return errors.New("Registration already exists")
	}

	err = db.Insert("mentors", &mentor_registration)

	return err
}

/*
	Updates the registration associated with the given mentor id
*/
func UpdateMentorRegistration(id string, mentor_registration models.MentorRegistration) error {
	err := validate.Struct(mentor_registration)

	if err != nil {
		return err
	}

	selector := bson.M{"id": id}

	err = db.Update("mentors", selector, &mentor_registration)

	return err
}
