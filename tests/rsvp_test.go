package tests

import (
	"github.com/HackIllinois/api-rsvp/config"
	"github.com/HackIllinois/api-rsvp/database"
	"github.com/HackIllinois/api-rsvp/models"
	"github.com/HackIllinois/api-rsvp/service"
	"reflect"
	"testing"
)

/*
	Initialize database with test user info
*/
func SetupTestDB(t *testing.T) {
	rsvp := models.UserRsvp{
		ID:          "testid",
		IsAttending: true,
	}

	err := database.Insert("rsvps", &rsvp)

	if err != nil {
		t.Fatal(err)
	}
}

/*
	Drop test database
*/
func CleanupTestDB(t *testing.T) {
	session := database.GetSession()
	defer session.Close()

	err := session.DB(config.RSVP_DB_NAME).DropDatabase()

	if err != nil {
		t.Fatal(err)
	}
}

/*
	Service level test for getting user rsvp from database
*/
func TestGetUserRsvpService(t *testing.T) {
	SetupTestDB(t)

	rsvp, err := service.GetUserRsvp("testid")

	if err != nil {
		t.Fatal(err)
	}

	expected_rsvp := models.UserRsvp{
		ID:          "testid",
		IsAttending: true,
	}

	if !reflect.DeepEqual(rsvp, &expected_rsvp) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_rsvp, rsvp)
	}

	CleanupTestDB(t)
}

/*
	Service level test for creating user rsvp in the database
*/
func TestCreateUserRsvpService(t *testing.T) {
	SetupTestDB(t)

	new_rsvp := models.UserRsvp{
		ID:          "testid2",
		IsAttending: false,
	}

	err := service.CreateUserRsvp("testid2", new_rsvp)

	if err != nil {
		t.Fatal(err)
	}

	rsvp, err := service.GetUserRsvp("testid2")

	if err != nil {
		t.Fatal(err)
	}

	expected_rsvp := models.UserRsvp{
		ID:          "testid2",
		IsAttending: false,
	}

	if !reflect.DeepEqual(rsvp, &expected_rsvp) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_rsvp, rsvp)
	}

	CleanupTestDB(t)
}

/*
	Service level test for updating user rsvp in the database
*/
func TestUpdateUserRsvpService(t *testing.T) {
	SetupTestDB(t)

	rsvp := models.UserRsvp{
		ID:          "testid",
		IsAttending: false,
	}

	err := service.UpdateUserRsvp("testid", rsvp)

	if err != nil {
		t.Fatal(err)
	}

	updated_rsvp, err := service.GetUserRsvp("testid")

	if err != nil {
		t.Fatal(err)
	}

	expected_rsvp := models.UserRsvp{
		ID:          "testid",
		IsAttending: false,
	}

	if !reflect.DeepEqual(updated_rsvp, &expected_rsvp) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_rsvp, updated_rsvp)
	}

	CleanupTestDB(t)
}
