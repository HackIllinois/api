package tests

import (
	"github.com/HackIllinois/api-registration/config"
	"github.com/HackIllinois/api-registration/database"
	"github.com/HackIllinois/api-registration/models"
	"github.com/HackIllinois/api-registration/service"
	"reflect"
	"testing"
)

/*
	Initialize database with test user info
*/
func SetupTestDB(t *testing.T) {
	err := database.Insert("attendees", &models.UserRegistration{
		ID:        "testid",
		FirstName: "testfirstname",
		LastName:  "testlastname",
	})

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

	err := session.DB(config.REGISTRATION_DB_NAME).DropDatabase()

	if err != nil {
		t.Fatal(err)
	}
}

/*
	Service level test for getting user registration from database
*/
func TestGetUserRegistrationService(t *testing.T) {
	SetupTestDB(t)

	user_registration, err := service.GetUserRegistration("testid")

	if err != nil {
		t.Fatal(err)
	}

	expected_registration := &models.UserRegistration{
		ID:        "testid",
		FirstName: "testfirstname",
		LastName:  "testlastname",
	}

	if !reflect.DeepEqual(user_registration, expected_registration) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_registration, user_registration)
	}

	CleanupTestDB(t)
}

/*
	Service level test for creating user registration in the database
*/
func TestCreateUserRegistrationService(t *testing.T) {
	SetupTestDB(t)

	err := service.CreateUserRegistration("testid2", models.UserRegistration{
		ID:        "testid2",
		FirstName: "testfirstname2",
		LastName:  "testlastname2",
	})

	if err != nil {
		t.Fatal(err)
	}

	user_registration, err := service.GetUserRegistration("testid2")

	if err != nil {
		t.Fatal(err)
	}

	expected_registration := &models.UserRegistration{
		ID:        "testid2",
		FirstName: "testfirstname2",
		LastName:  "testlastname2",
	}

	if !reflect.DeepEqual(user_registration, expected_registration) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_registration, user_registration)
	}

	CleanupTestDB(t)
}

/*
	Service level test for updating user registration in the database
*/
func TestUpdateUserRegistrationService(t *testing.T) {
	SetupTestDB(t)

	err := service.UpdateUserRegistration("testid", models.UserRegistration{
		ID:        "testid",
		FirstName: "testupdatedfirstname",
		LastName:  "testupdatedlastname",
	})

	if err != nil {
		t.Fatal(err)
	}

	user_registration, err := service.GetUserRegistration("testid")

	if err != nil {
		t.Fatal(err)
	}

	expected_registration := &models.UserRegistration{
		ID:        "testid",
		FirstName: "testupdatedfirstname",
		LastName:  "testupdatedlastname",
	}

	if !reflect.DeepEqual(user_registration, expected_registration) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_registration, user_registration)
	}

	CleanupTestDB(t)
}
