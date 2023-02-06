package tests

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/services/checkin/config"
	"github.com/HackIllinois/api/services/checkin/models"
	"github.com/HackIllinois/api/services/checkin/service"
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

	db, err = database.InitDatabase(config.CHECKIN_DB_HOST, config.CHECKIN_DB_NAME)

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	return_code := m.Run()

	os.Exit(return_code)
}

/*
	Initialize db with test checkin info
*/
func SetupTestDB(t *testing.T) {
	checkin := models.UserCheckin{
		ID:              "testid",
		HasCheckedIn:    true,
		HasPickedUpSwag: true,
		RsvpData:        map[string]interface{}{},
	}

	err := db.Insert("checkins", &checkin, nil)

	if err != nil {
		t.Fatal(err)
	}
}

/*
	Drop test db
*/
func CleanupTestDB(t *testing.T) {
	err := db.DropDatabase(nil)

	if err != nil {
		t.Fatal(err)
	}
}

/*
	Service level test for getting user checkin from db
*/
func TestGetUserCheckinService(t *testing.T) {
	SetupTestDB(t)

	checkin, err := service.GetUserCheckin("testid")

	if err != nil {
		t.Fatal(err)
	}

	expected_checkin := models.UserCheckin{
		ID:              "testid",
		HasCheckedIn:    true,
		HasPickedUpSwag: true,
		RsvpData:        map[string]interface{}{},
	}

	if !reflect.DeepEqual(checkin, &expected_checkin) {
		t.Errorf("Wrong user info. Expected %v, got %v", &expected_checkin, checkin)
	}

	CleanupTestDB(t)
}

/*
	Service level test for creating user checkin in the db
*/
func TestCreateUserCheckinService(t *testing.T) {
	SetupTestDB(t)

	new_checkin := models.UserCheckin{
		ID:              "testid2",
		HasCheckedIn:    true,
		HasPickedUpSwag: false,
		RsvpData:        map[string]interface{}{},
	}

	err := service.CreateUserCheckin("testid2", new_checkin)

	if err != nil {
		t.Fatal(err)
	}

	checkin, err := service.GetUserCheckin("testid2")

	if err != nil {
		t.Fatal(err)
	}

	expected_checkin := models.UserCheckin{
		ID:              "testid2",
		HasCheckedIn:    true,
		HasPickedUpSwag: false,
		RsvpData:        map[string]interface{}{},
	}

	if !reflect.DeepEqual(checkin, &expected_checkin) {
		t.Errorf("Wrong user info. Expected %v, got %v", &expected_checkin, checkin)
	}

	CleanupTestDB(t)
}

/*
	Service level test for updating user checkin in the db
*/
func TestUpdateUserCheckinService(t *testing.T) {
	SetupTestDB(t)

	checkin := models.UserCheckin{
		ID:              "testid",
		HasCheckedIn:    true,
		HasPickedUpSwag: false,
		RsvpData:        map[string]interface{}{},
	}

	err := service.UpdateUserCheckin("testid", checkin)

	if err != nil {
		t.Fatal(err)
	}

	updated_checkin, err := service.GetUserCheckin("testid")

	if err != nil {
		t.Fatal(err)
	}

	expected_checkin := models.UserCheckin{
		ID:              "testid",
		HasCheckedIn:    true,
		HasPickedUpSwag: false,
		RsvpData:        map[string]interface{}{},
	}

	if !reflect.DeepEqual(updated_checkin, &expected_checkin) {
		t.Errorf("Wrong user info. Expected %v, got %v", &expected_checkin, updated_checkin)
	}

	CleanupTestDB(t)
}

/*
	Service level test for getting list of all checked in users
*/
func TestGetAllCheckedInUsersService(t *testing.T) {
	SetupTestDB(t)

	new_checkin := models.UserCheckin{
		ID:              "testid2",
		HasCheckedIn:    false,
		HasPickedUpSwag: false,
		RsvpData:        map[string]interface{}{},
	}

	err := service.CreateUserCheckin("testid2", new_checkin)

	if err != nil {
		t.Fatal(err)
	}

	new_checkin = models.UserCheckin{
		ID:              "testid3",
		HasCheckedIn:    true,
		HasPickedUpSwag: false,
		RsvpData:        map[string]interface{}{},
	}

	err = service.CreateUserCheckin("testid3", new_checkin)

	if err != nil {
		t.Fatal(err)
	}

	checkin_list, err := service.GetAllCheckedInUsers()

	if err != nil {
		t.Fatal(err)
	}

	expected_checkin_list := models.CheckinList{
		CheckedInUsers: []string{"testid", "testid3"},
	}

	if !reflect.DeepEqual(checkin_list, &expected_checkin_list) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_checkin_list, checkin_list)
	}

	CleanupTestDB(t)
}
