package tests

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/common/datastore"
	"github.com/HackIllinois/api/services/rsvp/config"
	"github.com/HackIllinois/api/services/rsvp/service"
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

	db, err = database.InitDatabase(config.RSVP_DB_HOST, config.RSVP_DB_NAME)

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	return_code := m.Run()

	os.Exit(return_code)
}

/*
	Initialize db with test user info
*/
func SetupTestDB(t *testing.T) {
	rsvp := getBaseUserRsvp()

	err := db.Insert("rsvps", &rsvp, nil)

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
	Service level test for getting user rsvp from db
*/
func TestGetUserRsvpService(t *testing.T) {
	SetupTestDB(t)

	rsvp, err := service.GetUserRsvp("testid")

	if err != nil {
		t.Fatal(err)
	}

	expected_rsvp := getBaseUserRsvp()

	if !reflect.DeepEqual(rsvp.Data["isAttending"], expected_rsvp.Data["isAttending"]) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_rsvp.Data["isAttending"], rsvp.Data["isAttending"])
	}

	CleanupTestDB(t)
}

/*
	Service level test for creating user rsvp in the db
*/
func TestCreateUserRsvpService(t *testing.T) {
	SetupTestDB(t)

	new_rsvp := getBaseUserRsvp()
	new_rsvp.Data["id"] = "testid2"
	new_rsvp.Data["isAttending"] = false

	err := service.CreateUserRsvp("testid2", new_rsvp)

	if err != nil {
		t.Fatal(err)
	}

	rsvp, err := service.GetUserRsvp("testid2")

	if err != nil {
		t.Fatal(err)
	}

	expected_rsvp := getBaseUserRsvp()
	expected_rsvp.Data["id"] = "testid2"
	expected_rsvp.Data["isAttending"] = false

	if !reflect.DeepEqual(rsvp.Data["isAttending"], expected_rsvp.Data["isAttending"]) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_rsvp.Data["isAttending"], rsvp.Data["isAttending"])
	}

	CleanupTestDB(t)
}

/*
	Service level test for updating user rsvp in the db
*/
func TestUpdateUserRsvpService(t *testing.T) {
	SetupTestDB(t)

	rsvp := getBaseUserRsvp()
	rsvp.Data["isAttending"] = false

	err := service.UpdateUserRsvp("testid", rsvp)

	if err != nil {
		t.Fatal(err)
	}

	updated_rsvp, err := service.GetUserRsvp("testid")

	if err != nil {
		t.Fatal(err)
	}

	expected_rsvp := getBaseUserRsvp()
	expected_rsvp.Data["isAttending"] = false

	if !reflect.DeepEqual(updated_rsvp.Data["isAttending"], expected_rsvp.Data["isAttending"]) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_rsvp.Data["isAttending"], updated_rsvp.Data["isAttending"])
	}

	CleanupTestDB(t)
}

/*
	Returns a basic user registration
*/
func getBaseUserRsvp() datastore.DataStore {
	base_user_rsvp := datastore.NewDataStore(config.RSVP_DEFINITION)
	json.Unmarshal([]byte(user_rsvp_data), &base_user_rsvp)
	return base_user_rsvp
}

var user_rsvp_data string = `
{
	"id": "testid",
	"isAttending": true
}
`
