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
		ID:        "testid",
		Name:      "testname",
		Email:     "testemail",
		Github:    "testgithub",
		Linkedin:  "testlinkedin",
		Interests: []string{"testinterest1", "testinterest2"},
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
func TestGetAllProjectsService(t *testing.T) {
	SetupTestDB(t)

	profile := models.Profile{
		ID:        "testid2",
		Name:      "testname2",
		Email:     "testemail2",
		Github:    "testgithub2",
		Linkedin:  "testlinkedin2",
		Interests: []string{"testinterest3", "testinterest4"},
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
				ID:        "testid",
				Name:      "testname",
				Email:     "testemail",
				Github:    "testgithub",
				Linkedin:  "testlinkedin",
				Interests: []string{"testinterest1", "testinterest2"},
			},
			{
				ID:        "testid2",
				Name:      "testname2",
				Email:     "testemail2",
				Github:    "testgithub2",
				Linkedin:  "testlinkedin2",
				Interests: []string{"testinterest3", "testinterest4"},
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
