package service

import (
	"errors"
	"strconv"
	"strings"

	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/services/registration/config"
	"github.com/HackIllinois/api/services/registration/models"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

var db database.Database

func Initialize() error {
	if db != nil {
		db.Close()
		db = nil
	}

	var err error
	db, err = database.InitDatabase(config.REGISTRATION_DB_HOST, config.REGISTRATION_DB_NAME)

	if err != nil {
		return err
	}

	validate = validator.New()

	return nil
}

/*
	Returns the registration associated with the given user id
*/
func GetUserRegistration(id string) (*models.UserRegistration, error) {
	query := database.QuerySelector{"id": id}

	var user_registration models.UserRegistration
	err := db.FindOne("attendees", query, &user_registration, nil)

	if err != nil {
		return nil, err
	}

	return &user_registration, nil
}

/*
	Creates the registration associated with the given user id
*/
func CreateUserRegistration(id string, user_registration models.UserRegistration) error {
	err := user_registration.Validate()

	if err != nil {
		return err
	}

	_, err = GetUserRegistration(id)

	if err != database.ErrNotFound {
		if err != nil {
			return err
		}
		return errors.New("Registration already exists.")
	}

	err = db.Insert("attendees", &user_registration, nil)

	return err
}

/*
	Updates the registration associated with the given user id
*/
func UpdateUserRegistration(id string, user_registration models.UserRegistration) error {
	err := user_registration.Validate()

	if err != nil {
		return err
	}

	selector := database.QuerySelector{"id": id}

	err = db.Replace("attendees", selector, &user_registration, false, nil)

	return err
}

/*
	Returns db search query based on given parameters
*/
func getFilterQuery(parameters map[string][]string) (map[string]interface{}, error) {
	query := make(map[string]interface{})
	for key, values := range parameters {
		if len(values) == 1 {
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
			query[key] = database.QuerySelector{"$in": correctly_typed_value_list}
		} else {
			return nil, errors.New("Multiple usage of key " + key)
		}
	}

	return query, nil
}

/*
	Returns the user registrations associated with the given parameters
*/
func GetFilteredUserRegistrations(parameters map[string][]string) (*models.FilteredUserRegistrations, error) {
	query, err := getFilterQuery(parameters)
	if err != nil {
		return nil, err
	}

	var filtered_registrations models.FilteredUserRegistrations
	err = db.FindAll("attendees", query, &filtered_registrations.Registrations, nil)
	if err != nil {
		return nil, err
	}

	return &filtered_registrations, nil
}

func AssignValueType(key, value string) (interface{}, error) {
	int_keys := []string{"age", "graduationYear"}
	if Contains(int_keys, key) {
		return strconv.Atoi(value)
	}

	bool_keys := []string{"isNovice", "isPrivate"}
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
	query := database.QuerySelector{"id": id}

	var mentor_registration models.MentorRegistration
	err := db.FindOne("mentors", query, &mentor_registration, nil)

	if err != nil {
		return nil, err
	}

	return &mentor_registration, nil
}

/*
	Creates the registration associated with the given mentor id
*/
func CreateMentorRegistration(id string, mentor_registration models.MentorRegistration) error {
	err := mentor_registration.Validate()

	if err != nil {
		return err
	}

	_, err = GetMentorRegistration(id)

	if err != database.ErrNotFound {
		if err != nil {
			return err
		}
		return errors.New("Registration already exists")
	}

	err = db.Insert("mentors", &mentor_registration, nil)

	return err
}

/*
	Updates the registration associated with the given mentor id
*/
func UpdateMentorRegistration(id string, mentor_registration models.MentorRegistration) error {
	err := mentor_registration.Validate()

	if err != nil {
		return err
	}

	selector := database.QuerySelector{"id": id}

	err = db.Replace("mentors", selector, &mentor_registration, false, nil)

	return err
}

/*
	Returns the mentor registrations associated with the given parameters
*/
func GetFilteredMentorRegistrations(parameters map[string][]string) (*models.FilteredMentorRegistrations, error) {
	query, err := getFilterQuery(parameters)
	if err != nil {
		return nil, err
	}

	var filtered_registrations models.FilteredMentorRegistrations
	err = db.FindAll("mentors", query, &filtered_registrations.Registrations, nil)
	if err != nil {
		return nil, err
	}

	return &filtered_registrations, nil
}

/*
	Returns all registration stats
*/
func GetStats() (map[string]interface{}, error) {
	attendee_stats, err := db.GetStats("attendees", config.REGISTRATION_STAT_FIELDS, nil)

	if err != nil {
		return nil, err
	}

	mentor_stats, err := db.GetStats("mentors", []string{}, nil)

	if err != nil {
		return nil, err
	}

	stats := make(map[string]interface{})
	stats["attendees"] = attendee_stats
	stats["mentors"] = mentor_stats

	return stats, nil
}
