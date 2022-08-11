package service

import (
	"errors"

	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/services/rsvp/config"
	"github.com/HackIllinois/api/services/rsvp/models"
)

var db database.Database

func Initialize() error {
	if db != nil {
		db.Close()
		db = nil
	}

	var err error
	db, err = database.InitDatabase(config.RSVP_DB_HOST, config.RSVP_DB_NAME)

	if err != nil {
		return err
	}

	return nil
}

/*
	Returns the rsvp associated with the given user id
*/
func GetUserRsvp(id string) (*models.UserRsvp, error) {
	query := database.QuerySelector{
		"id": id,
	}

	var rsvp models.UserRsvp
	err := db.FindOne("rsvps", query, &rsvp, nil)

	if err != nil {
		return nil, err
	}

	return &rsvp, nil
}

/*
	Creates the rsvp associated with the given user id
*/
func CreateUserRsvp(id string, rsvp models.UserRsvp) error {
	isAttending, ok := rsvp.Data["isAttending"].(bool)

	if !ok {
		return errors.New("Type mismatch in rsvp data")
	}

	if isAttending {
		err := rsvp.Validate()

		if err != nil {
			return err
		}
	}

	_, err := GetUserRsvp(id)

	if err != database.ErrNotFound {
		if err != nil {
			return err
		}
		return errors.New("RSVP already exists.")
	}

	err = db.Insert("rsvps", &rsvp, nil)

	return err
}

/*
	Updates the rsvp associated with the given user id
*/
func UpdateUserRsvp(id string, rsvp models.UserRsvp) error {
	isAttending, ok := rsvp.Data["isAttending"].(bool)

	if !ok {
		return errors.New("Type mismatch in rsvp data")
	}

	if isAttending {
		err := rsvp.Validate()

		if err != nil {
			return err
		}
	}

	selector := database.QuerySelector{
		"id": id,
	}

	err := db.Replace("rsvps", selector, &rsvp, false, nil)

	return err
}

/*
	Returns the rsvps associated with the given parameters
*/
func GetFilteredRsvps(parameters map[string][]string) (*models.FilteredRsvps, error) {
	query, err := database.CreateFilterQuery(parameters, models.UserRsvp{})

	if err != nil {
		return nil, err
	}

	var filtered_rsvps models.FilteredRsvps
	err = db.FindAll("rsvps", query, &filtered_rsvps.Rsvps, nil)

	if err != nil {
		return nil, err
	}

	return &filtered_rsvps, nil
}

/*
	Returns all rsvp stats
*/
func GetStats() (map[string]interface{}, error) {
	return db.GetStats("rsvps", config.RSVP_STAT_FIELDS, nil)
}
