package tests

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/services/profile/config"
	"github.com/HackIllinois/api/services/profile/models"
	"github.com/HackIllinois/api/services/profile/service"
)

var db database.Database

func TestMain(m *testing.M) {
	err := config.Initialize()

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)

	}

	err = service.Initialize()

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	db, err = database.InitDatabase(config.PROFILE_DB_HOST, config.PROFILE_DB_NAME)

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	return_code := m.Run()

	os.Exit(return_code)
}

var TestTime = time.Now().Unix()

/*
	Initialize db with a test profile
*/
func SetupTestDB(t *testing.T) {
	profile_id := "testid"

	profile := models.Profile{
		ID:        profile_id,
		FirstName: "testfirstname",
		LastName:  "testlastname",
		Points:    0,
		IsVirtual: false,
	}

	id_map := models.IdMap{
		UserID:    "testuserid",
		ProfileID: profile_id,
	}

	err := db.Insert("profileids", &id_map)

	if err != nil {
		t.Fatal(err)
	}

	err = db.Insert("profiles", &profile)

	if err != nil {
		t.Fatal(err)
	}

	attendance_tracker := models.AttendanceTracker{
		ID:     profile_id,
		Events: []string{},
	}

	err = db.Insert("profileattendance", &attendance_tracker)

	if err != nil {
		t.Fatal(err)
	}

	profile_favorites := models.ProfileFavorites{
		ID:       profile_id,
		Profiles: []string{},
	}

	err = db.Insert("profilefavorites", &profile_favorites)

	if err != nil {
		t.Fatal(err)
	}
}

/*
	Drop test db
*/
func CleanupTestDB(t *testing.T) {
	err := db.DropDatabase()

	if err != nil {
		t.Fatal(err)
	}
}

/*
	Service level test for getting all profiles from db
*/
func TestGetAllProfilesService(t *testing.T) {
	SetupTestDB(t)
	defer CleanupTestDB(t)

	profile := models.Profile{
		ID:        "testid2",
		FirstName: "testfirstname2",
		LastName:  "testlastname2",
		Points:    340,
		IsVirtual: true,
	}

	err := db.Insert("profiles", &profile)

	if err != nil {
		t.Fatal(err)
	}

	parameters := map[string][]string{}

	actual_profile_list, err := service.GetFilteredProfiles(parameters)

	if err != nil {
		t.Fatal(err)
	}

	expected_profile_list := models.ProfileList{
		Profiles: []models.Profile{
			{
				ID:        "testid",
				FirstName: "testfirstname",
				LastName:  "testlastname",
				Points:    0,
				IsVirtual: false,
			},
			{
				ID:        "testid2",
				FirstName: "testfirstname2",
				LastName:  "testlastname2",
				Points:    340,
				IsVirtual: true,
			},
		},
	}

	if !reflect.DeepEqual(actual_profile_list, &expected_profile_list) {
		t.Errorf("Wrong profile list. Expected %v, got %v", expected_profile_list, actual_profile_list)
	}

	db.RemoveAll("profiles", nil)

	actual_profile_list, err = service.GetFilteredProfiles(parameters)

	if err != nil {
		t.Fatal(err)
	}

	expected_profile_list = models.ProfileList{
		Profiles: []models.Profile{},
	}

	if !reflect.DeepEqual(actual_profile_list, &expected_profile_list) {
		t.Errorf("Wrong profile list. Expected %v, got %v", expected_profile_list, actual_profile_list)
	}
}

/*
	Service level test for getting profile from db
*/
func TestGetProfileService(t *testing.T) {
	SetupTestDB(t)
	defer CleanupTestDB(t)

	profile, err := service.GetProfile("testid")

	if err != nil {
		t.Fatal(err)
	}

	expected_profile := models.Profile{
		ID:        "testid",
		FirstName: "testfirstname",
		LastName:  "testlastname",
		Points:    0,
		IsVirtual: false,
	}

	if !reflect.DeepEqual(profile, &expected_profile) {
		t.Errorf("Wrong profile info. Expected %v, got %v", &expected_profile, profile)
	}
}

/*
	Service level test for creating a profile in the db
*/
func TestCreateProfileService(t *testing.T) {
	SetupTestDB(t)
	defer CleanupTestDB(t)

	new_profile := models.Profile{
		ID:        "testid2",
		FirstName: "testfirstname2",
		LastName:  "testlastname2",
		Points:    340,
		IsVirtual: true,
	}

	err := service.CreateProfile("testuserid2", "testid2", new_profile)

	if err != nil {
		t.Fatal(err)
	}

	profile, err := service.GetProfile("testid2")

	if err != nil {
		t.Fatal(err)
	}

	expected_profile := models.Profile{
		ID:        "testid2",
		FirstName: "testfirstname2",
		LastName:  "testlastname2",
		Points:    340,
		IsVirtual: true,
	}

	if !reflect.DeepEqual(profile, &expected_profile) {
		t.Errorf("Wrong profile info. Expected %v, got %v", expected_profile, profile)
	}

	// Test that id mapping was inserted correctly
	profile_id1, err := service.GetProfileIdFromUserId("testuserid2")

	if err != nil {
		t.Fatal(err)
	}

	if profile_id1 != "testid2" {
		t.Errorf("Wrong profile mapping found for user %s. Expected %s but got %s", "testuserid2", "testid2", profile_id1)
	}
}

/*
	Service level test for deleting a profile in the db
*/
func TestDeleteProfileService(t *testing.T) {
	SetupTestDB(t)
	defer CleanupTestDB(t)

	profile_id := "testid"

	// Try to delete the profile

	_, err := service.DeleteProfile(profile_id)

	if err != nil {
		t.Fatal(err)
	}

	// Try to find the profile in the profiles db
	profile, err := service.GetProfile(profile_id)

	if err == nil {
		t.Errorf("Found profile %v in profiles database.", profile)
	}
}

/*
	Service level test for updating a profile in the db
*/
func TestUpdateProfileService(t *testing.T) {
	SetupTestDB(t)
	defer CleanupTestDB(t)

	profile := models.Profile{
		ID:        "testid",
		FirstName: "testfirstname2",
		LastName:  "testlastname2",
		Points:    340,
		IsVirtual: true,
	}

	err := service.UpdateProfile("testid", profile)

	if err != nil {
		t.Fatal(err)
	}

	updated_profile, err := service.GetProfile("testid")

	if err != nil {
		t.Fatal(err)
	}

	expected_profile := models.Profile{
		ID:        "testid",
		FirstName: "testfirstname2",
		LastName:  "testlastname2",
		Points:    340,
		IsVirtual: true,
	}

	if !reflect.DeepEqual(updated_profile, &expected_profile) {
		t.Errorf("Wrong profile info. Expected %v, got %v", expected_profile, updated_profile)
	}
}

func TestGetFilteredProfiles(t *testing.T) {
	SetupTestDB(t)
	defer CleanupTestDB(t)

	profile := models.Profile{
		ID:        "testid2",
		FirstName: "testfirstname2",
		LastName:  "testlastname2",
		Points:    340,
		IsVirtual: true,
	}

	err := db.Insert("profiles", &profile)

	profile = models.Profile{
		ID:        "testid3",
		FirstName: "testfirstname3",
		LastName:  "testlastname3",
		Points:    342,
		IsVirtual: true,
	}
	err = db.Insert("profiles", &profile)

	parameters := map[string][]string{
		"IsVirtual": {"true"},
		"limit":     {"0"},
	}

	filtered_profile_list, err := service.GetFilteredProfiles(parameters)

	if err != nil {
		t.Fatal(err)
	}

	expected_filtered_profile_list := models.ProfileList{
		Profiles: []models.Profile{
			{
				ID:        "testid2",
				FirstName: "testfirstname2",
				LastName:  "testlastname2",
				Points:    340,
				IsVirtual: true,
			},
			{
				ID:        "testid3",
				FirstName: "testfirstname3",
				LastName:  "testlastname3",
				Points:    342,
				IsVirtual: true,
			},
		},
	}

	if !reflect.DeepEqual(filtered_profile_list, &expected_filtered_profile_list) {
		t.Errorf("Wrong profile list. Expected %v, got %v", expected_filtered_profile_list, filtered_profile_list)
	}

	// Add a limit and test that
	parameters = map[string][]string{
		"IsVirtual": {"true"},
		"limit":     {"1"},
	}

	filtered_profile_list, err = service.GetFilteredProfiles(parameters)

	expected_filtered_profile_list = models.ProfileList{
		Profiles: []models.Profile{
			{
				ID:        "testid2",
				FirstName: "testfirstname2",
				LastName:  "testlastname2",
				Points:    340,
				IsVirtual: true,
			},
		},
	}

	if !reflect.DeepEqual(filtered_profile_list, &expected_filtered_profile_list) {
		t.Errorf("Wrong profile list. Expected %v, got %v", expected_filtered_profile_list, filtered_profile_list)
	}

	// Remove filter by virtual status
	parameters = map[string][]string{
		"limit": {"0"},
	}

	filtered_profile_list, err = service.GetFilteredProfiles(parameters)

	expected_filtered_profile_list = models.ProfileList{
		Profiles: []models.Profile{
			{
				ID:        "testid",
				FirstName: "testfirstname",
				LastName:  "testlastname",
				Points:    0,
				IsVirtual: false,
			},
			{
				ID:        "testid2",
				FirstName: "testfirstname2",
				LastName:  "testlastname2",
				Points:    340,
				IsVirtual: true,
			},
			{
				ID:        "testid3",
				FirstName: "testfirstname3",
				LastName:  "testlastname3",
				Points:    342,
				IsVirtual: true,
			},
		},
	}

	if !reflect.DeepEqual(filtered_profile_list, &expected_filtered_profile_list) {
		t.Errorf("Wrong profile list. Expected %v, got %v", expected_filtered_profile_list, filtered_profile_list)
	}
}

func TestGetProfileLeaderboard(t *testing.T) {
	SetupTestDB(t)
	defer CleanupTestDB(t)

	profile := models.Profile{
		ID:        "testid2",
		FirstName: "testfirstname2",
		LastName:  "testlastname2",
		Points:    340,
		IsVirtual: true,
	}

	err := db.Insert("profiles", &profile)

	if err != nil {
		t.Fatal(err)
	}

	parameters := map[string][]string{}

	leaderboard, err := service.GetProfileLeaderboard(parameters)

	if err != nil {
		t.Fatal(err)
	}

	expected_leaderboard := models.LeaderboardEntryList{
		LeaderboardEntries: []models.LeaderboardEntry{
			{
				ID:        "testid2",
				FirstName: "testfirstname2",
				LastName:  "testlastname2",
				Points:    340,
			},
			{
				ID:        "testid",
				FirstName: "testfirstname",
				LastName:  "testlastname",
				Points:    0,
			},
		},
	}

	if !reflect.DeepEqual(leaderboard, &expected_leaderboard) {
		t.Errorf("Wrong profile info. Expected %v, got %v", expected_leaderboard, leaderboard)
	}

	// Insert another profile and test
	profile = models.Profile{
		ID:        "testid3",
		FirstName: "testfirstname3",
		LastName:  "testlastname3",
		Points:    999,
		IsVirtual: true,
	}

	err = db.Insert("profiles", &profile)

	if err != nil {
		t.Fatal(err)
	}

	parameters = map[string][]string{
		"limit": {"0"},
	}

	leaderboard, err = service.GetProfileLeaderboard(parameters)

	if err != nil {
		t.Fatal(err)
	}

	expected_leaderboard = models.LeaderboardEntryList{
		LeaderboardEntries: []models.LeaderboardEntry{
			{
				ID:        "testid3",
				FirstName: "testfirstname3",
				LastName:  "testlastname3",
				Points:    999,
			},
			{
				ID:        "testid2",
				FirstName: "testfirstname2",
				LastName:  "testlastname2",
				Points:    340,
			},
			{
				ID:        "testid",
				FirstName: "testfirstname",
				LastName:  "testlastname",
				Points:    0,
			},
		},
	}
	if !reflect.DeepEqual(leaderboard, &expected_leaderboard) {
		t.Errorf("Wrong profile info. Expected %v, got %v", expected_leaderboard, leaderboard)
	}

	// Add a limit and test again
	parameters = map[string][]string{
		"limit": {"2"}, // Get the top two
	}

	leaderboard, err = service.GetProfileLeaderboard(parameters)

	if err != nil {
		t.Fatal(err)
	}

	expected_leaderboard = models.LeaderboardEntryList{
		LeaderboardEntries: []models.LeaderboardEntry{
			{
				ID:        "testid3",
				FirstName: "testfirstname3",
				LastName:  "testlastname3",
				Points:    999,
			},
			{
				ID:        "testid2",
				FirstName: "testfirstname2",
				LastName:  "testlastname2",
				Points:    340,
			},
		},
	}
	if !reflect.DeepEqual(leaderboard, &expected_leaderboard) {
		t.Errorf("Wrong profile info. Expected %v, got %v", expected_leaderboard, leaderboard)
	}
}

func TestGetValidFilteredProfiles(t *testing.T) {
	SetupTestDB(t)
	defer CleanupTestDB(t)

	profile := models.Profile{
		ID:        "testid2",
		FirstName: "testfirstname2",
		LastName:  "testlastname2",
		Points:    340,
		IsVirtual: false,
	}

	err := db.Insert("profiles", &profile)

	profile = models.Profile{
		ID:        "testid3",
		FirstName: "testfirstname3",
		LastName:  "testlastname3",
		Points:    342,
		IsVirtual: true,
	}
	err = db.Insert("profiles", &profile)

	parameters := map[string][]string{
		"IsVirtual": {"true"},
		"limit":     {"0"},
	}

	filtered_profile_list, err := service.GetValidFilteredProfiles(parameters)

	if err != nil {
		t.Fatal(err)
	}

	expected_filtered_profile_list := models.ProfileList{
		Profiles: []models.Profile{
			{
				ID:        "testid3",
				FirstName: "testfirstname3",
				LastName:  "testlastname3",
				Points:    342,
				IsVirtual: true,
			},
		},
	}

	if !reflect.DeepEqual(filtered_profile_list, &expected_filtered_profile_list) {
		t.Errorf("Wrong profile list. Expected %v, got %v", expected_filtered_profile_list, filtered_profile_list)
	}

	profile = models.Profile{
		ID:        "testid4",
		FirstName: "testfirstname4",
		LastName:  "testlastname4",
		Points:    342,
		IsVirtual: false,
	}
	err = db.Insert("profiles", &profile)

	// Remove the virtual status filter. Now every profile should show up.

	parameters = map[string][]string{
		"limit": {"0"},
	}

	filtered_profile_list, err = service.GetValidFilteredProfiles(parameters)

	if err != nil {
		t.Fatal(err)
	}

	expected_filtered_profile_list = models.ProfileList{
		Profiles: []models.Profile{
			{
				ID:        "testid",
				FirstName: "testfirstname",
				LastName:  "testlastname",
				Points:    0,
				IsVirtual: false,
			},
			{
				ID:        "testid2",
				FirstName: "testfirstname2",
				LastName:  "testlastname2",
				Points:    340,
				IsVirtual: false,
			},
			{
				ID:        "testid3",
				FirstName: "testfirstname3",
				LastName:  "testlastname3",
				Points:    342,
				IsVirtual: true,
			},
			{
				ID:        "testid4",
				FirstName: "testfirstname4",
				LastName:  "testlastname4",
				Points:    342,
				IsVirtual: false,
			},
		},
	}
	if !reflect.DeepEqual(filtered_profile_list, &expected_filtered_profile_list) {
		t.Errorf("Wrong profile list. Expected %v, got %v", expected_filtered_profile_list, filtered_profile_list)
	}
}

func TestProfileFavorites(t *testing.T) {
	SetupTestDB(t)
	defer CleanupTestDB(t)

	profile := models.Profile{
		ID:        "testid2",
		FirstName: "testfirstname2",
		LastName:  "testlastname2",
		Points:    340,
		IsVirtual: true,
	}

	err := db.Insert("profiles", &profile)

	if err != nil {
		t.Fatal(err)
	}

	profile_favorites, err := service.GetProfileFavorites("testid")

	expected_profile_favorites := models.ProfileFavorites{
		ID:       "testid",
		Profiles: []string{},
	}

	if !reflect.DeepEqual(profile_favorites, &expected_profile_favorites) {
		t.Errorf("Wrong favorite profile list. Expected %v, got %v", expected_profile_favorites, profile_favorites)
	}

	// Add a profile to the favorites
	err = service.AddProfileFavorite("testid", "testid2")
	if err != nil {
		t.Fatal(err)
	}

	profile_favorites, err = service.GetProfileFavorites("testid")
	if err != nil {
		t.Fatal(err)
	}

	expected_profile_favorites = models.ProfileFavorites{
		ID:       "testid",
		Profiles: []string{"testid2"},
	}

	if !reflect.DeepEqual(profile_favorites, &expected_profile_favorites) {
		t.Errorf("Wrong favorite profile list. Expected %v, got %v", expected_profile_favorites, profile_favorites)
	}

	// Favorite another (nonexistent) profile and make sure it fails.
	err = service.AddProfileFavorite("testid", "testid3")
	expected_err := errors.New("Could not find profile with the given id.")
	if !reflect.DeepEqual(err, expected_err) {
		t.Errorf("The service did not return the correct error. Expected %v, got %v", expected_err, err)
	}

	// Remove the (nonexistent) profile from the favorites and make sure it fails.
	err = service.RemoveProfileFavorite("testid", "testid3")
	expected_err = errors.New("User's profile favorites does not have specified profile")
	if !reflect.DeepEqual(err, expected_err) {
		t.Errorf("The service did not return the correct error. Expected %v, got %v", expected_err, err)
	}

	// Add yourself to the favorites and make sure it fails
	err = service.AddProfileFavorite("testid", "testid")
	expected_err = errors.New("User's profile matches the specified profile.")
	if !reflect.DeepEqual(err, expected_err) {
		t.Errorf("The service did not return the correct error. Expected %v, got %v", expected_err, err)
	}

	// Favorite another profile
	profile = models.Profile{
		ID:        "testid3",
		FirstName: "testfirstname3",
		LastName:  "testlastname3",
		Points:    342,
		IsVirtual: true,
	}
	err = db.Insert("profiles", &profile)
	if err != nil {
		t.Fatal(err)
	}

	err = service.AddProfileFavorite("testid", "testid3")
	if err != nil {
		t.Fatal(err)
	}

	profile_favorites, err = service.GetProfileFavorites("testid")
	expected_profile_favorites = models.ProfileFavorites{
		ID:       "testid",
		Profiles: []string{"testid2", "testid3"},
	}

	if !reflect.DeepEqual(profile_favorites, &expected_profile_favorites) {
		t.Errorf("Wrong favorite profile list. Expected %v, got %v", expected_profile_favorites, profile_favorites)
	}

	// Remove a favorite
	err = service.RemoveProfileFavorite("testid", "testid2")
	if err != nil {
		t.Fatal(err)
	}

	profile_favorites, err = service.GetProfileFavorites("testid")
	expected_profile_favorites = models.ProfileFavorites{
		ID:       "testid",
		Profiles: []string{"testid3"},
	}

	if !reflect.DeepEqual(profile_favorites, &expected_profile_favorites) {
		t.Errorf("Wrong favorite profile list. Expected %v, got %v", expected_profile_favorites, profile_favorites)
	}
}
