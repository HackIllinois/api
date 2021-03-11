package tests

import (
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
	profile := models.Profile{
		ID:          "testid",
		FirstName:   "testfirstname",
		LastName:    "testlastname",
		Points:      0,
		Timezone:    "America/Chicago",
		Description: "Hi",
		Discord:     "testdiscordusername",
		AvatarUrl:   "https://imgs.smoothradio.com/images/191589?crop=16_9&width=660&relax=1&signature=Rz93ikqcAz7BcX6SKiEC94zJnqo=",
		TeamStatus:  "Looking For Team",
		Interests:   []string{"testinterest1", "testinterest2"},
	}

	err := db.Insert("profiles", &profile)

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

	profile := models.Profile{
		ID:          "testid2",
		FirstName:   "testfirstname2",
		LastName:    "testlastname2",
		Points:      340,
		Timezone:    "America/New York",
		Description: "Hello",
		Discord:     "testdiscordusername2",
		AvatarUrl:   "https://yt3.ggpht.com/ytc/AAUvwniHNhQyp4hWj3nrADnils-6N3jNREP8rWKGDTp0Lg=s900-c-k-c0x00ffffff-no-rj",
		TeamStatus:  "Found Team",
		Interests:   []string{"testinterest2"},
	}

	err := db.Insert("profiles", &profile)

	if err != nil {
		t.Fatal(err)
	}

	actual_profile_list, err := service.GetAllProfiles()

	if err != nil {
		t.Fatal(err)
	}

	expected_profile_list := models.ProfileList{
		Profiles: []models.Profile{
			{
				ID:          "testid",
				FirstName:   "testfirstname",
				LastName:    "testlastname",
				Points:      0,
				Timezone:    "America/Chicago",
				Description: "Hi",
				Discord:     "testdiscordusername",
				AvatarUrl:   "https://imgs.smoothradio.com/images/191589?crop=16_9&width=660&relax=1&signature=Rz93ikqcAz7BcX6SKiEC94zJnqo=",
				TeamStatus:  "Looking For Team",
				Interests:   []string{"testinterest1", "testinterest2"},
			},
			{
				ID:          "testid2",
				FirstName:   "testfirstname2",
				LastName:    "testlastname2",
				Points:      340,
				Timezone:    "America/New York",
				Description: "Hello",
				Discord:     "testdiscordusername2",
				AvatarUrl:   "https://yt3.ggpht.com/ytc/AAUvwniHNhQyp4hWj3nrADnils-6N3jNREP8rWKGDTp0Lg=s900-c-k-c0x00ffffff-no-rj",
				TeamStatus:  "Found Team",
				Interests:   []string{"testinterest2"},
			},
		},
	}

	if !reflect.DeepEqual(actual_profile_list, &expected_profile_list) {
		t.Errorf("Wrong profile list. Expected %v, got %v", expected_profile_list, actual_profile_list)
	}

	db.RemoveAll("profiles", nil)

	actual_profile_list, err = service.GetAllProfiles()

	if err != nil {
		t.Fatal(err)
	}

	expected_profile_list = models.ProfileList{
		Profiles: []models.Profile{},
	}

	if !reflect.DeepEqual(actual_profile_list, &expected_profile_list) {
		t.Errorf("Wrong profile list. Expected %v, got %v", expected_profile_list, actual_profile_list)
	}

	CleanupTestDB(t)

}

/*
	Service level test for getting profile from db
*/
func TestGetProfileService(t *testing.T) {
	SetupTestDB(t)

	profile, err := service.GetProfile("testid")

	if err != nil {
		t.Fatal(err)
	}

	expected_profile := models.Profile{
		ID:          "testid",
		FirstName:   "testfirstname",
		LastName:    "testlastname",
		Points:      0,
		Timezone:    "America/Chicago",
		Description: "Hi",
		Discord:     "testdiscordusername",
		AvatarUrl:   "https://imgs.smoothradio.com/images/191589?crop=16_9&width=660&relax=1&signature=Rz93ikqcAz7BcX6SKiEC94zJnqo=",
		TeamStatus:  "Looking For Team",
		Interests:   []string{"testinterest1", "testinterest2"},
	}

	if !reflect.DeepEqual(profile, &expected_profile) {
		t.Errorf("Wrong profile info. Expected %v, got %v", &expected_profile, profile)
	}

	CleanupTestDB(t)
}

/*
	Service level test for creating a profile in the db
*/
func TestCreateProfileService(t *testing.T) {
	SetupTestDB(t)

	new_profile := models.Profile{
		ID:          "testid2",
		FirstName:   "testfirstname2",
		LastName:    "testlastname2",
		Points:      340,
		Timezone:    "America/New York",
		Description: "Hello",
		Discord:     "testdiscordusername2",
		AvatarUrl:   "https://yt3.ggpht.com/ytc/AAUvwniHNhQyp4hWj3nrADnils-6N3jNREP8rWKGDTp0Lg=s900-c-k-c0x00ffffff-no-rj",
		TeamStatus:  "Found Team",
		Interests:   []string{"testinterest2"},
	}

	err := service.CreateProfile("testid2", new_profile)

	if err != nil {
		t.Fatal(err)
	}

	profile, err := service.GetProfile("testid2")

	if err != nil {
		t.Fatal(err)
	}

	expected_profile := models.Profile{
		ID:          "testid2",
		FirstName:   "testfirstname2",
		LastName:    "testlastname2",
		Points:      340,
		Timezone:    "America/New York",
		Description: "Hello",
		Discord:     "testdiscordusername2",
		AvatarUrl:   "https://yt3.ggpht.com/ytc/AAUvwniHNhQyp4hWj3nrADnils-6N3jNREP8rWKGDTp0Lg=s900-c-k-c0x00ffffff-no-rj",
		TeamStatus:  "Found Team",
		Interests:   []string{"testinterest2"},
	}

	if !reflect.DeepEqual(profile, &expected_profile) {
		t.Errorf("Wrong profile info. Expected %v, got %v", expected_profile, profile)
	}

	CleanupTestDB(t)
}

/*
	Service level test for deleting a profile in the db
*/
func TestDeleteProfileService(t *testing.T) {
	SetupTestDB(t)

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

	CleanupTestDB(t)
}

/*
	Service level test for updating a profile in the db
*/
func TestUpdateProfileService(t *testing.T) {
	SetupTestDB(t)

	profile := models.Profile{
		ID:          "testid",
		FirstName:   "testfirstname2",
		LastName:    "testlastname2",
		Points:      340,
		Timezone:    "America/New York",
		Description: "Hello",
		Discord:     "testdiscordusername2",
		AvatarUrl:   "https://yt3.ggpht.com/ytc/AAUvwniHNhQyp4hWj3nrADnils-6N3jNREP8rWKGDTp0Lg=s900-c-k-c0x00ffffff-no-rj",
		TeamStatus:  "Found Team",
		Interests:   []string{"testinterest2"},
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
		ID:          "testid",
		FirstName:   "testfirstname2",
		LastName:    "testlastname2",
		Points:      340,
		Timezone:    "America/New York",
		Description: "Hello",
		Discord:     "testdiscordusername2",
		AvatarUrl:   "https://yt3.ggpht.com/ytc/AAUvwniHNhQyp4hWj3nrADnils-6N3jNREP8rWKGDTp0Lg=s900-c-k-c0x00ffffff-no-rj",
		TeamStatus:  "Found Team",
		Interests:   []string{"testinterest2"},
	}

	if !reflect.DeepEqual(updated_profile, &expected_profile) {
		t.Errorf("Wrong profile info. Expected %v, got %v", expected_profile, updated_profile)
	}

	CleanupTestDB(t)
}

func TestGetFilteredProfiles(t *testing.T) {
	SetupTestDB(t)

	profile := models.Profile{
		ID:          "testid2",
		FirstName:   "testfirstname2",
		LastName:    "testlastname2",
		Points:      340,
		Timezone:    "America/New York",
		Description: "Hello",
		Discord:     "testdiscordusername2",
		AvatarUrl:   "https://yt3.ggpht.com/ytc/AAUvwniHNhQyp4hWj3nrADnils-6N3jNREP8rWKGDTp0Lg=s900-c-k-c0x00ffffff-no-rj",
		TeamStatus:  "Found Team",
		Interests:   []string{"Cpp", "Machine Learning", "Additional Interest"},
	}

	err := db.Insert("profiles", &profile)

	profile = models.Profile{
		ID:          "testid3",
		FirstName:   "testfirstname3",
		LastName:    "testlastname3",
		Points:      342,
		Timezone:    "America/New York",
		Description: "Hello",
		Discord:     "testdiscordusername3",
		AvatarUrl:   "https://yt3.ggpht.com/ytc/AAUvwniHNhQyp4hWj3nrADnils-6N3jNREP8rWKGDTp0Lg=s900-c-k-c0x00ffffff-no-rj",
		TeamStatus:  "Found Team",
		Interests:   []string{"Cpp", "Machine Learning"},
	}
	err = db.Insert("profiles", &profile)

	parameters := map[string][]string{
		"teamStatus": {"Found Team"},
		"interests":  {"Cpp,Machine Learning"},
		"limit":      {"0"},
	}

	filtered_profile_list, err := service.GetFilteredProfiles(parameters)

	if err != nil {
		t.Fatal(err)
	}

	expected_filtered_profile_list := models.ProfileList{
		Profiles: []models.Profile{
			{
				ID:          "testid2",
				FirstName:   "testfirstname2",
				LastName:    "testlastname2",
				Points:      340,
				Timezone:    "America/New York",
				Description: "Hello",
				Discord:     "testdiscordusername2",
				AvatarUrl:   "https://yt3.ggpht.com/ytc/AAUvwniHNhQyp4hWj3nrADnils-6N3jNREP8rWKGDTp0Lg=s900-c-k-c0x00ffffff-no-rj",
				TeamStatus:  "Found Team",
				Interests:   []string{"Cpp", "Machine Learning", "Additional Interest"},
			},
			{
				ID:          "testid3",
				FirstName:   "testfirstname3",
				LastName:    "testlastname3",
				Points:      342,
				Timezone:    "America/New York",
				Description: "Hello",
				Discord:     "testdiscordusername3",
				AvatarUrl:   "https://yt3.ggpht.com/ytc/AAUvwniHNhQyp4hWj3nrADnils-6N3jNREP8rWKGDTp0Lg=s900-c-k-c0x00ffffff-no-rj",
				TeamStatus:  "Found Team",
				Interests:   []string{"Cpp", "Machine Learning"},
			},
		},
	}

	if !reflect.DeepEqual(filtered_profile_list, &expected_filtered_profile_list) {
		t.Errorf("Wrong profile list. Expected %v, got %v", expected_filtered_profile_list, filtered_profile_list)
	}

	// Add a limit and test that
	parameters = map[string][]string{
		"teamStatus": {"Found Team"},
		"interests":  {"Cpp,Machine Learning"},
		"limit":      {"1"},
	}

	filtered_profile_list, err = service.GetFilteredProfiles(parameters)

	expected_filtered_profile_list = models.ProfileList{
		Profiles: []models.Profile{
			{
				ID:          "testid2",
				FirstName:   "testfirstname2",
				LastName:    "testlastname2",
				Points:      340,
				Timezone:    "America/New York",
				Description: "Hello",
				Discord:     "testdiscordusername2",
				AvatarUrl:   "https://yt3.ggpht.com/ytc/AAUvwniHNhQyp4hWj3nrADnils-6N3jNREP8rWKGDTp0Lg=s900-c-k-c0x00ffffff-no-rj",
				TeamStatus:  "Found Team",
				Interests:   []string{"Cpp", "Machine Learning", "Additional Interest"},
			},
		},
	}

	if !reflect.DeepEqual(filtered_profile_list, &expected_filtered_profile_list) {
		t.Errorf("Wrong profile list. Expected %v, got %v", expected_filtered_profile_list, filtered_profile_list)
	}

	// Change the interests to be off by one
	parameters = map[string][]string{
		"teamStatus": {"Found Team"},
		"interests":  {"Cpp,Machine Learning,Additional Interest"},
		"limit":      {"0"},
	}

	filtered_profile_list, err = service.GetFilteredProfiles(parameters)

	expected_filtered_profile_list = models.ProfileList{
		Profiles: []models.Profile{
			{
				ID:          "testid2",
				FirstName:   "testfirstname2",
				LastName:    "testlastname2",
				Points:      340,
				Timezone:    "America/New York",
				Description: "Hello",
				Discord:     "testdiscordusername2",
				AvatarUrl:   "https://yt3.ggpht.com/ytc/AAUvwniHNhQyp4hWj3nrADnils-6N3jNREP8rWKGDTp0Lg=s900-c-k-c0x00ffffff-no-rj",
				TeamStatus:  "Found Team",
				Interests:   []string{"Cpp", "Machine Learning", "Additional Interest"},
			},
		},
	}

	if !reflect.DeepEqual(filtered_profile_list, &expected_filtered_profile_list) {
		t.Errorf("Wrong profile list. Expected %v, got %v", expected_filtered_profile_list, filtered_profile_list)
	}

	// Remove filter by interests
	parameters = map[string][]string{
		"teamStatus": {"Found Team"},
		"limit":      {"0"},
	}

	filtered_profile_list, err = service.GetFilteredProfiles(parameters)

	expected_filtered_profile_list = models.ProfileList{
		Profiles: []models.Profile{
			{
				ID:          "testid2",
				FirstName:   "testfirstname2",
				LastName:    "testlastname2",
				Points:      340,
				Timezone:    "America/New York",
				Description: "Hello",
				Discord:     "testdiscordusername2",
				AvatarUrl:   "https://yt3.ggpht.com/ytc/AAUvwniHNhQyp4hWj3nrADnils-6N3jNREP8rWKGDTp0Lg=s900-c-k-c0x00ffffff-no-rj",
				TeamStatus:  "Found Team",
				Interests:   []string{"Cpp", "Machine Learning", "Additional Interest"},
			},
			{
				ID:          "testid3",
				FirstName:   "testfirstname3",
				LastName:    "testlastname3",
				Points:      342,
				Timezone:    "America/New York",
				Description: "Hello",
				Discord:     "testdiscordusername3",
				AvatarUrl:   "https://yt3.ggpht.com/ytc/AAUvwniHNhQyp4hWj3nrADnils-6N3jNREP8rWKGDTp0Lg=s900-c-k-c0x00ffffff-no-rj",
				TeamStatus:  "Found Team",
				Interests:   []string{"Cpp", "Machine Learning"},
			},
		},
	}

	if !reflect.DeepEqual(filtered_profile_list, &expected_filtered_profile_list) {
		t.Errorf("Wrong profile list. Expected %v, got %v", expected_filtered_profile_list, filtered_profile_list)
	}

	// Remove filter by teamStatus
	parameters = map[string][]string{
		"interests": {"Cpp,Machine Learning,Additional Interest"},
		"limit":     {"0"},
	}

	filtered_profile_list, err = service.GetFilteredProfiles(parameters)

	expected_filtered_profile_list = models.ProfileList{
		Profiles: []models.Profile{
			{
				ID:          "testid2",
				FirstName:   "testfirstname2",
				LastName:    "testlastname2",
				Points:      340,
				Timezone:    "America/New York",
				Description: "Hello",
				Discord:     "testdiscordusername2",
				AvatarUrl:   "https://yt3.ggpht.com/ytc/AAUvwniHNhQyp4hWj3nrADnils-6N3jNREP8rWKGDTp0Lg=s900-c-k-c0x00ffffff-no-rj",
				TeamStatus:  "Found Team",
				Interests:   []string{"Cpp", "Machine Learning", "Additional Interest"},
			},
		},
	}

	if !reflect.DeepEqual(filtered_profile_list, &expected_filtered_profile_list) {
		t.Errorf("Wrong profile list. Expected %v, got %v", expected_filtered_profile_list, filtered_profile_list)
	}

	CleanupTestDB(t)
}

func TestGetProfileLeaderboard(t *testing.T) {
	SetupTestDB(t)

	profile := models.Profile{
		ID:          "testid2",
		FirstName:   "testfirstname2",
		LastName:    "testlastname2",
		Points:      340,
		Timezone:    "America/New York",
		Description: "Hello",
		Discord:     "testdiscordusername2",
		AvatarUrl:   "https://yt3.ggpht.com/ytc/AAUvwniHNhQyp4hWj3nrADnils-6N3jNREP8rWKGDTp0Lg=s900-c-k-c0x00ffffff-no-rj",
		TeamStatus:  "Found Team",
		Interests:   []string{"Cpp", "Machine Learning", "Additional Interest"},
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

	expected_leaderboard := models.ProfileList{
		Profiles: []models.Profile{
			{
				ID:          "testid2",
				FirstName:   "testfirstname2",
				LastName:    "testlastname2",
				Points:      340,
				Timezone:    "America/New York",
				Description: "Hello",
				Discord:     "testdiscordusername2",
				AvatarUrl:   "https://yt3.ggpht.com/ytc/AAUvwniHNhQyp4hWj3nrADnils-6N3jNREP8rWKGDTp0Lg=s900-c-k-c0x00ffffff-no-rj",
				TeamStatus:  "Found Team",
				Interests:   []string{"Cpp", "Machine Learning", "Additional Interest"},
			},
			{
				ID:          "testid",
				FirstName:   "testfirstname",
				LastName:    "testlastname",
				Points:      0,
				Timezone:    "America/Chicago",
				Description: "Hi",
				Discord:     "testdiscordusername",
				AvatarUrl:   "https://imgs.smoothradio.com/images/191589?crop=16_9&width=660&relax=1&signature=Rz93ikqcAz7BcX6SKiEC94zJnqo=",
				TeamStatus:  "Looking For Team",
				Interests:   []string{"testinterest1", "testinterest2"},
			},
		},
	}

	if !reflect.DeepEqual(leaderboard, &expected_leaderboard) {
		t.Errorf("Wrong profile info. Expected %v, got %v", expected_leaderboard, leaderboard)
	}

	// Insert another profile and test
	profile = models.Profile{
		ID:          "testid3",
		FirstName:   "testfirstname3",
		LastName:    "testlastname3",
		Points:      999,
		Timezone:    "America/New York",
		Description: "Hello",
		Discord:     "testdiscordusername3",
		AvatarUrl:   "https://yt3.ggpht.com/ytc/AAUvwniHNhQyp4hWj3nrADnils-6N3jNREP8rWKGDTp0Lg=s900-c-k-c0x00ffffff-no-rj",
		TeamStatus:  "Found Team",
		Interests:   []string{"Cpp"},
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

	expected_leaderboard = models.ProfileList{
		Profiles: []models.Profile{
			{
				ID:          "testid3",
				FirstName:   "testfirstname3",
				LastName:    "testlastname3",
				Points:      999,
				Timezone:    "America/New York",
				Description: "Hello",
				Discord:     "testdiscordusername3",
				AvatarUrl:   "https://yt3.ggpht.com/ytc/AAUvwniHNhQyp4hWj3nrADnils-6N3jNREP8rWKGDTp0Lg=s900-c-k-c0x00ffffff-no-rj",
				TeamStatus:  "Found Team",
				Interests:   []string{"Cpp"},
			},
			{
				ID:          "testid2",
				FirstName:   "testfirstname2",
				LastName:    "testlastname2",
				Points:      340,
				Timezone:    "America/New York",
				Description: "Hello",
				Discord:     "testdiscordusername2",
				AvatarUrl:   "https://yt3.ggpht.com/ytc/AAUvwniHNhQyp4hWj3nrADnils-6N3jNREP8rWKGDTp0Lg=s900-c-k-c0x00ffffff-no-rj",
				TeamStatus:  "Found Team",
				Interests:   []string{"Cpp", "Machine Learning", "Additional Interest"},
			},
			{
				ID:          "testid",
				FirstName:   "testfirstname",
				LastName:    "testlastname",
				Points:      0,
				Timezone:    "America/Chicago",
				Description: "Hi",
				Discord:     "testdiscordusername",
				AvatarUrl:   "https://imgs.smoothradio.com/images/191589?crop=16_9&width=660&relax=1&signature=Rz93ikqcAz7BcX6SKiEC94zJnqo=",
				TeamStatus:  "Looking For Team",
				Interests:   []string{"testinterest1", "testinterest2"},
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

	expected_leaderboard = models.ProfileList{
		Profiles: []models.Profile{
			{
				ID:          "testid3",
				FirstName:   "testfirstname3",
				LastName:    "testlastname3",
				Points:      999,
				Timezone:    "America/New York",
				Description: "Hello",
				Discord:     "testdiscordusername3",
				AvatarUrl:   "https://yt3.ggpht.com/ytc/AAUvwniHNhQyp4hWj3nrADnils-6N3jNREP8rWKGDTp0Lg=s900-c-k-c0x00ffffff-no-rj",
				TeamStatus:  "Found Team",
				Interests:   []string{"Cpp"},
			},
			{
				ID:          "testid2",
				FirstName:   "testfirstname2",
				LastName:    "testlastname2",
				Points:      340,
				Timezone:    "America/New York",
				Description: "Hello",
				Discord:     "testdiscordusername2",
				AvatarUrl:   "https://yt3.ggpht.com/ytc/AAUvwniHNhQyp4hWj3nrADnils-6N3jNREP8rWKGDTp0Lg=s900-c-k-c0x00ffffff-no-rj",
				TeamStatus:  "Found Team",
				Interests:   []string{"Cpp", "Machine Learning", "Additional Interest"},
			},
		},
	}
	if !reflect.DeepEqual(leaderboard, &expected_leaderboard) {
		t.Errorf("Wrong profile info. Expected %v, got %v", expected_leaderboard, leaderboard)
	}

	CleanupTestDB(t)
}

func TestGetValidFilteredProfiles(t *testing.T) {
	SetupTestDB(t)

	profile := models.Profile{
		ID:          "testid2",
		FirstName:   "testfirstname2",
		LastName:    "testlastname2",
		Points:      340,
		Timezone:    "America/New York",
		Description: "Hello",
		Discord:     "testdiscordusername2",
		AvatarUrl:   "https://yt3.ggpht.com/ytc/AAUvwniHNhQyp4hWj3nrADnils-6N3jNREP8rWKGDTp0Lg=s900-c-k-c0x00ffffff-no-rj",
		TeamStatus:  "Not Looking",
		Interests:   []string{"Cpp", "Machine Learning", "Additional Interest"},
	}

	err := db.Insert("profiles", &profile)

	profile = models.Profile{
		ID:          "testid3",
		FirstName:   "testfirstname3",
		LastName:    "testlastname3",
		Points:      342,
		Timezone:    "America/New York",
		Description: "Hello",
		Discord:     "testdiscordusername3",
		AvatarUrl:   "https://yt3.ggpht.com/ytc/AAUvwniHNhQyp4hWj3nrADnils-6N3jNREP8rWKGDTp0Lg=s900-c-k-c0x00ffffff-no-rj",
		TeamStatus:  "Found Team",
		Interests:   []string{"Cpp", "Machine Learning"},
	}
	err = db.Insert("profiles", &profile)

	parameters := map[string][]string{
		"interests": {"Cpp,Machine Learning"},
		"limit":     {"0"},
	}

	filtered_profile_list, err := service.GetValidFilteredProfiles(parameters)

	if err != nil {
		t.Fatal(err)
	}

	expected_filtered_profile_list := models.ProfileList{
		Profiles: []models.Profile{
			{
				ID:          "testid3",
				FirstName:   "testfirstname3",
				LastName:    "testlastname3",
				Points:      342,
				Timezone:    "America/New York",
				Description: "Hello",
				Discord:     "testdiscordusername3",
				AvatarUrl:   "https://yt3.ggpht.com/ytc/AAUvwniHNhQyp4hWj3nrADnils-6N3jNREP8rWKGDTp0Lg=s900-c-k-c0x00ffffff-no-rj",
				TeamStatus:  "Found Team",
				Interests:   []string{"Cpp", "Machine Learning"},
			},
		},
	}

	profile = models.Profile{
		ID:          "testid4",
		FirstName:   "testfirstname3",
		LastName:    "testlastname3",
		Points:      342,
		Timezone:    "America/New York",
		Description: "Hello",
		Discord:     "testdiscordusername3",
		AvatarUrl:   "https://yt3.ggpht.com/ytc/AAUvwniHNhQyp4hWj3nrADnils-6N3jNREP8rWKGDTp0Lg=s900-c-k-c0x00ffffff-no-rj",
		TeamStatus:  "Looking For Team",
		Interests:   []string{},
	}
	err = db.Insert("profiles", &profile)

	// Remove the interests filter. Now every profile should show up except for those that are "Not Looking" for a team.

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
				ID:          "testid",
				FirstName:   "testfirstname",
				LastName:    "testlastname",
				Points:      0,
				Timezone:    "America/Chicago",
				Description: "Hi",
				Discord:     "testdiscordusername",
				AvatarUrl:   "https://imgs.smoothradio.com/images/191589?crop=16_9&width=660&relax=1&signature=Rz93ikqcAz7BcX6SKiEC94zJnqo=",
				TeamStatus:  "Looking For Team",
				Interests:   []string{"testinterest1", "testinterest2"},
			},
			{
				ID:          "testid3",
				FirstName:   "testfirstname3",
				LastName:    "testlastname3",
				Points:      342,
				Timezone:    "America/New York",
				Description: "Hello",
				Discord:     "testdiscordusername3",
				AvatarUrl:   "https://yt3.ggpht.com/ytc/AAUvwniHNhQyp4hWj3nrADnils-6N3jNREP8rWKGDTp0Lg=s900-c-k-c0x00ffffff-no-rj",
				TeamStatus:  "Found Team",
				Interests:   []string{"Cpp", "Machine Learning"},
			},
			{
				ID:          "testid4",
				FirstName:   "testfirstname3",
				LastName:    "testlastname3",
				Points:      342,
				Timezone:    "America/New York",
				Description: "Hello",
				Discord:     "testdiscordusername3",
				AvatarUrl:   "https://yt3.ggpht.com/ytc/AAUvwniHNhQyp4hWj3nrADnils-6N3jNREP8rWKGDTp0Lg=s900-c-k-c0x00ffffff-no-rj",
				TeamStatus:  "Looking For Team",
				Interests:   []string{},
			},
		},
	}
	if !reflect.DeepEqual(filtered_profile_list, &expected_filtered_profile_list) {
		t.Errorf("Wrong profile list. Expected %v, got %v", expected_filtered_profile_list, filtered_profile_list)
	}

	// Add a TeamStatus filter.

	parameters = map[string][]string{
		"teamStatus": {"Looking For Team"},
		"limit":      {"0"},
	}

	filtered_profile_list, err = service.GetValidFilteredProfiles(parameters)

	if err != nil {
		t.Fatal(err)
	}

	expected_filtered_profile_list = models.ProfileList{
		Profiles: []models.Profile{
			{
				ID:          "testid",
				FirstName:   "testfirstname",
				LastName:    "testlastname",
				Points:      0,
				Timezone:    "America/Chicago",
				Description: "Hi",
				Discord:     "testdiscordusername",
				AvatarUrl:   "https://imgs.smoothradio.com/images/191589?crop=16_9&width=660&relax=1&signature=Rz93ikqcAz7BcX6SKiEC94zJnqo=",
				TeamStatus:  "Looking For Team",
				Interests:   []string{"testinterest1", "testinterest2"},
			},
			{
				ID:          "testid4",
				FirstName:   "testfirstname3",
				LastName:    "testlastname3",
				Points:      342,
				Timezone:    "America/New York",
				Description: "Hello",
				Discord:     "testdiscordusername3",
				AvatarUrl:   "https://yt3.ggpht.com/ytc/AAUvwniHNhQyp4hWj3nrADnils-6N3jNREP8rWKGDTp0Lg=s900-c-k-c0x00ffffff-no-rj",
				TeamStatus:  "Looking For Team",
				Interests:   []string{},
			},
		},
	}

	if !reflect.DeepEqual(filtered_profile_list, &expected_filtered_profile_list) {
		t.Errorf("Wrong profile list. Expected %v, got %v", expected_filtered_profile_list, filtered_profile_list)
	}

	CleanupTestDB(t)
}
