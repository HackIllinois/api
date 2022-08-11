package service

import (
	"errors"
	"strconv"

	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/common/utils"
	"github.com/HackIllinois/api/services/profile/config"
	"github.com/HackIllinois/api/services/profile/models"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
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
	Returns the profile id associated with the given user id
*/
func GetProfileIdFromUserId(id string) (string, error) {
	query := database.QuerySelector{
		"userid": id,
	}

	var id_map models.IdMap
	err := db.FindOne("profileids", query, &id_map, nil)

	// Returns error if no mapping was found
	if err != nil {
		return "", err
	}

	return id_map.ProfileID, nil
}

/*
	Returns the profile with the given id
*/
func GetProfile(profile_id string) (*models.Profile, error) {
	query := database.QuerySelector{
		"id": profile_id,
	}

	var profile models.Profile
	err := db.FindOne("profiles", query, &profile, nil)

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
func DeleteProfile(profile_id string) (*models.Profile, error) {
	// Gets profile to be able to return it later
	profile, err := GetProfile(profile_id)

	if err != nil {
		return nil, err
	}

	// Remove user id to profile id mapping
	query := database.QuerySelector{
		"profileid": profile_id,
	}

	err = db.RemoveOne("profileids", query, nil)

	if err != nil {
		return nil, err
	}

	// Remove profile from profile database
	query = database.QuerySelector{
		"id": profile_id,
	}

	err = db.RemoveOne("profiles", query, nil)

	if err != nil {
		return nil, err
	}

	err = db.RemoveOne("profileattendance", query, nil)

	if err != nil {
		return nil, err
	}

	err = db.RemoveOne("profilefavorites", query, nil)

	if err != nil {
		return nil, err
	}

	return profile, err
}

/*
	Creates a profile with the given id
*/
func CreateProfile(id string, profile_id string, profile models.Profile) error {
	profile.ID = profile_id
	err := validate.Struct(profile)

	if err != nil {
		return err
	}

	_, err = GetProfile(profile_id)

	if err != database.ErrNotFound {
		if err != nil {
			return err
		}
		return errors.New("Profile already exists")
	}

	// TODO: Look into mongodb multi-document transactions
	// Create user id to profile id mapping
	var id_map models.IdMap

	id_map.UserID = id
	id_map.ProfileID = profile_id

	err = db.Insert("profileids", &id_map, nil)

	if err != nil {
		return err
	}

	err = db.Insert("profiles", &profile, nil)

	if err != nil {
		return err
	}

	attendance_tracker := models.AttendanceTracker{
		ID:     profile_id,
		Events: []string{},
	}

	err = db.Insert("profileattendance", &attendance_tracker, nil)

	if err != nil {
		return err
	}

	profile_favorites := models.ProfileFavorites{
		ID:       profile_id,
		Profiles: []string{},
	}

	err = db.Insert("profilefavorites", &profile_favorites, nil)

	if err != nil {
		return err
	}

	return err
}

/*
	Updates the profile with the given id
*/
func UpdateProfile(profile_id string, profile models.Profile) error {
	profile.ID = profile_id
	err := validate.Struct(profile)

	if err != nil {
		return err
	}

	selector := database.QuerySelector{
		"id": profile_id,
	}

	err = db.Replace("profiles", selector, &profile, false, nil)

	return err
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

	sort_field := bson.D{
		{
			"points",
			-1,
		},
	}

	err = db.FindAllSorted("profiles", nil, sort_field, &leaderboard_entries, nil)

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
	err = db.FindAll("profiles", query, &profiles, nil)

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

/*
	Returns a list of profiles filtered upon teamStatus and interests. Will be limited to only include the first "limit" results.
	Will also remove profiles with a TeamStatus set to "NOT_LOOKING"
*/
func GetValidFilteredProfiles(parameters map[string][]string) (*models.ProfileList, error) {
	filtered_profile_list, err := GetFilteredProfiles(parameters)

	if err != nil {
		return nil, errors.New("Could not get filtered profiles")
	}

	return filtered_profile_list, nil
}

/*
  Redeems the event with `event_id` for the user with profile id `id`
*/
func RedeemEvent(profile_id string, event_id string) (*models.RedeemEventResponse, error) {
	var redemption_status models.RedeemEventResponse
	redemption_status.Status = "Success"

	selector := database.QuerySelector{
		"id": profile_id,
	}

	var attended_events models.AttendanceTracker
	err := db.FindOne("profileattendance", selector, &attended_events, nil)

	if err != nil {
		if err == database.ErrNotFound {
			attended_events = models.AttendanceTracker{
				ID:     profile_id,
				Events: []string{},
			}
			err = db.Insert("profileattendance", &attended_events, nil)

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

	err = db.Replace("profileattendance", selector, attended_events, false, nil)

	return &redemption_status, err
}

/*
	Returns the profile favorites for the user with the given id
*/
func GetProfileFavorites(profile_id string) (*models.ProfileFavorites, error) {
	query := database.QuerySelector{
		"id": profile_id,
	}

	var profile_favorites models.ProfileFavorites
	err := db.FindOne("profilefavorites", query, &profile_favorites, nil)

	if err != nil {
		return nil, err
	}

	return &profile_favorites, nil
}

/*
	Adds the given profile to the favorites for the user with the given id
*/
func AddProfileFavorite(profile_id string, profile string) error {
	if profile_id == profile {
		return errors.New("User's profile matches the specified profile.")
	}

	selector := database.QuerySelector{
		"id": profile_id,
	}

	_, err := GetProfile(profile)

	if err != nil {
		return errors.New("Could not find profile with the given id.")
	}

	profile_favorites, err := GetProfileFavorites(profile_id)

	if err != nil {
		return err
	}

	if !utils.ContainsString(profile_favorites.Profiles, profile) {
		profile_favorites.Profiles = append(profile_favorites.Profiles, profile)
	}

	err = db.Replace("profilefavorites", selector, profile_favorites, false, nil)

	return err
}

/*
	Removes the given profile from the favorites for the user with the given id
*/
func RemoveProfileFavorite(profile_id string, profile string) error {
	selector := database.QuerySelector{
		"id": profile_id,
	}

	profile_favorites, err := GetProfileFavorites(profile_id)

	if err != nil {
		return err
	}

	profile_favorites.Profiles, err = utils.RemoveString(profile_favorites.Profiles, profile)

	if err != nil {
		return errors.New("User's profile favorites does not have specified profile")
	}

	err = db.Replace("profilefavorites", selector, profile_favorites, false, nil)

	return err
}
