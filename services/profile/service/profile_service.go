package service

import (
	"errors"
	"strconv"

	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/common/utils"
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

	if err != nil {
		return err
	}

	attendance_tracker := models.AttendanceTracker{
		ID:     id,
		Events: []string{},
	}

	err = db.Insert("profileattendance", &attendance_tracker)

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

/*
	Returns a list of "limit" profiles sorted decesending by points.
	If "limit" is not provided, this will return a list of all profiles.
*/
func GetProfileLeaderboard(parameters map[string][]string) (*models.LeaderboardEntryList, error) {
	limit_param, ok := parameters["limit"]

	if !ok {
		limit_param = []string{"0"}
	}

	limit, err := strconv.Atoi(limit_param[0])

	if err != nil {
		return nil, errors.New("Could not convert 'limit' to int.")
	}

	leaderboard_entries := []models.LeaderboardEntry{}

	sort_field := database.SortField{
		Name:     "points",
		Reversed: true,
	}

	err = db.FindAllSorted("profiles", nil, []database.SortField{sort_field}, &leaderboard_entries)

	if err != nil {
		return nil, err
	}

	if limit > 0 {
		limit = utils.Min(limit, len(leaderboard_entries))
		leaderboard_entries = leaderboard_entries[:limit]
	}

	leaderboard_entry_list := models.LeaderboardEntryList{
		LeaderboardEntries: leaderboard_entries,
	}

	return &leaderboard_entry_list, nil
}

/*
	Returns a list of profiles filtered upon teamStatus and interests. Will be limited to only include the first "limit" results.
*/
func GetFilteredProfiles(parameters map[string][]string) (*models.ProfileList, error) {
	limit_param, ok := parameters["limit"]

	if !ok {
		limit_param = []string{"0"}
	}

	limit, err := strconv.Atoi(limit_param[0])

	if err != nil {
		return nil, errors.New("Could not convert 'limit' to int.")
	}

	// Remove "limit" from parameters before querying db
	delete(parameters, "limit")

	query, err := database.CreateFilterQuery(parameters, models.Profile{})

	if err != nil {
		return nil, err
	}

	profiles := []models.Profile{}
	err = db.FindAll("profiles", query, &profiles)

	if err != nil {
		return nil, err
	}

	// TODO: add some kind of recommendation sort/metric here

	if limit > 0 {
		limit = utils.Min(limit, len(profiles))
		profiles = profiles[:limit]
	}

	profile_list := models.ProfileList{
		Profiles: profiles,
	}

	return &profile_list, nil
}

func GetValidFilteredProfiles(parameters map[string][]string) (*models.ProfileList, error) {
	parameters["teamStatusNot"] = append(parameters["teamStatusNot"], "NOT_LOOKING")
	filtered_profile_list, err := GetFilteredProfiles(parameters)

	if err != nil {
		return nil, errors.New("Could not get filtered profiles")
	}

	return filtered_profile_list, nil
}

func RedeemEvent(id string, event_id string) (*models.RedeemEventResponse, error) {
	var redemption_status models.RedeemEventResponse
	redemption_status.Status = "Success"

	selector := database.QuerySelector{
		"id": id,
	}

	var attended_events models.AttendanceTracker
	err := db.FindOne("profileattendance", selector, &attended_events)

	if err != nil {
		if err == database.ErrNotFound {
			err = db.Insert("profileattendance", &models.AttendanceTracker{
				ID:     id,
				Events: []string{},
			})

			if err != nil {
				redemption_status.Status = "Could not add tracker to db"
				return &redemption_status, err
			}
		} else {
			redemption_status.Status = "Could not access db"
			return &redemption_status, err
		}
	}

	if utils.ContainsString(attended_events.Events, event_id) {
		redemption_status.Status = "Event already redeemed"
		return &redemption_status, nil
	} else {
		attended_events.Events = append(attended_events.Events, event_id)
	}

	err = db.Update("profileattendance", selector, attended_events)

	return &redemption_status, err
}
