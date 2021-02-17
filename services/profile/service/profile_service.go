package service

import (
	"errors"

	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/services/profile/config"
	"github.com/HackIllinois/api/services/profile/models"
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
	db, err = database.InitDatabase(config.PROFILE_DB_HOST, config.PROFILE_DB_NAME)

	if err != nil {
		return err
	}

	validate = validator.New()

	return nil
}

/*
	Returns the profile with the given id
*/
func GetProfile(id string) (*models.Profile, error) {
	query := database.QuerySelector{
		"id": id,
	}

	var profile models.Profile
	err := db.FindOne("profiles", query, &profile)

	if err != nil {
		return nil, err
	}

	return &profile, nil
}

/*
	Deletes the profile with the given id.
	Removes the profile from profile trackers and every user's tracker.
	Returns the profile that was deleted.
*/
func DeleteProfile(id string) (*models.Profile, error) {

	// Gets profile to be able to return it later

	profile, err := GetProfile(id)

	if err != nil {
		return nil, err
	}

	query := database.QuerySelector{
		"id": id,
	}

	// Remove profile from profile database

	err = db.RemoveOne("profiles", query)

	if err != nil {
		return nil, err
	}

	return profile, err
}

/*
	Creates a profile with the given id
*/
func CreateProfile(id string, profile models.Profile) error {
	profile.ID = id
	err := validate.Struct(profile)

	if err != nil {
		return err
	}

	_, err = GetProfile(id)

	if err != database.ErrNotFound {
		if err != nil {
			return err
		}
		return errors.New("Profile already exists")
	}

	err = db.Insert("profiles", &profile)

	return err
}

/*
	Updates the profile with the given id
*/
func UpdateProfile(id string, profile models.Profile) error {
	profile.ID = id
	err := validate.Struct(profile)

	if err != nil {
		return err
	}

	selector := database.QuerySelector{
		"id": id,
	}

	err = db.Update("profiles", selector, &profile)

	return err
}

/*
	Returns the list of all accessible profiles
*/
func GetAllProfiles() (*models.ProfileList, error) {
	profiles := []models.Profile{}

	err := db.FindAll("profiles", nil, &profiles)

	if err != nil {
		return nil, err
	}

	profile_list := models.ProfileList{
		Profiles: profiles,
	}

	return &profile_list, nil
}
