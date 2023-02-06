package service

import (
	"errors"

	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/common/utils"
	"github.com/HackIllinois/api/services/checkin/config"
	"github.com/HackIllinois/api/services/checkin/models"
)

var db database.Database

func Initialize() error {
	if db != nil {
		db.Close()
		db = nil
	}

	var err error
	db, err = database.InitDatabase(config.CHECKIN_DB_HOST, config.CHECKIN_DB_NAME)

	if err != nil {
		return err
	}

	return nil
}

/*
	Returns the checkin associated with the given user id
*/
func GetUserCheckin(id string) (*models.UserCheckin, error) {
	query := database.QuerySelector{
		"id": id,
	}

	var user_checkin models.UserCheckin
	err := db.FindOne("checkins", query, &user_checkin, nil)

	if err != nil {
		return nil, err
	}

	return &user_checkin, nil
}

/*
	Create the checkin associated with the given user id
*/
func CreateUserCheckin(id string, user_checkin models.UserCheckin) error {
	_, err := GetUserCheckin(id)

	if err != database.ErrNotFound {
		if err != nil {
			return err
		}
		return errors.New("Checkin already exists")
	}

	err = db.Insert("checkins", &user_checkin, nil)

	return err
}

/*
	Update the checkin associated with the given user id
*/
func UpdateUserCheckin(id string, user_checkin models.UserCheckin) error {
	selector := database.QuerySelector{
		"id": id,
	}

	err := db.Replace("checkins", selector, &user_checkin, false, nil)

	return err
}

/*
	Returns true, nil if a user with specified ID is allowed to checkin, and false, nil if not allowed.
	Sponsors, mentors, and those with staff overrides do not need an RSVP to check-in.
*/
func CanUserCheckin(id string, user_has_override bool) (bool, error) {
	is_user_registered, err := IsUserRegistered(id)

	if err != nil {
		return false, err
	}

	if !is_user_registered {
		return false, errors.New("User is not registered.")
	}

	if user_has_override {
		return true, nil
	}

	user_roles, err := GetRoles(id)

	if err != nil {
		return false, err
	}

	is_sponsor_or_mentor := utils.ContainsString(user_roles.Roles, models.SponsorRole) || utils.ContainsString(user_roles.Roles, models.MentorRole)

	if is_sponsor_or_mentor {
		return true, nil
	}

	// We do not want to call the below service function if the above condition is met, as it results
	// in a 400 (Bad Request) / error if the user's RSVP info cannot be found.
	// Therefore, we do not combine the conditions, and return as early as possible.
	return IsAttendeeRsvped(id)
}

/*
	Returns a list of all checked in user IDs
*/
func GetAllCheckedInUsers() (*models.CheckinList, error) {
	query := database.QuerySelector{
		"hascheckedin": true,
	}

	var check_ins []models.UserCheckin
	err := db.FindAll("checkins", query, &check_ins, nil)

	if err != nil {
		return nil, err
	}

	var checkin_list models.CheckinList
	for _, check_in := range check_ins {
		checkin_list.CheckedInUsers = append(checkin_list.CheckedInUsers, check_in.ID)
	}

	return &checkin_list, nil
}

/*
	Returns all checkin stats
*/
func GetStats() (map[string]interface{}, error) {
	return db.GetStats("checkins", []string{"override", "hascheckedin", "haspickedupswag"}, nil)
}
