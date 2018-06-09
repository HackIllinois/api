package tests

import (
	"github.com/HackIllinois/api-commons/database"
	"github.com/HackIllinois/api-registration/config"
	"github.com/HackIllinois/api-registration/models"
	"github.com/HackIllinois/api-registration/service"
	"reflect"
	"testing"
)

var db database.MongoDatabase

func init() {
	db_connection, err := database.InitMongoDatabase(config.REGISTRATION_DB_HOST, config.REGISTRATION_DB_NAME)

	if err != nil {
		panic(err)
	}

	db = db_connection
}

/*
	Initialize db with test user info
*/
func SetupTestDB(t *testing.T) {
	user_registration := base_registration
	err := db.Insert("attendees", &user_registration)

	if err != nil {
		t.Fatal(err)
	}
}

/*
	Drop test db
*/
func CleanupTestDB(t *testing.T) {
	session := db.GetSession()
	defer session.Close()

	err := session.DB(config.REGISTRATION_DB_NAME).DropDatabase()

	if err != nil {
		t.Fatal(err)
	}
}

/*
	Service level test for getting user registration from db
*/
func TestGetUserRegistrationService(t *testing.T) {
	SetupTestDB(t)

	user_registration, err := service.GetUserRegistration("testid")

	if err != nil {
		t.Fatal(err)
	}

	expected_registration := base_registration

	if !reflect.DeepEqual(user_registration, &expected_registration) {
		t.Errorf("Wrong user info.\nExpected %v\ngot %v\n", expected_registration, user_registration)
	}

	CleanupTestDB(t)
}

/*
	Service level test for creating user registration in the db
*/
func TestCreateUserRegistrationService(t *testing.T) {
	SetupTestDB(t)

	new_registration := base_registration
	new_registration.ID = "testid2"
	new_registration.FirstName = "first2"
	new_registration.LastName = "last2"
	err := service.CreateUserRegistration("testid2", new_registration)

	if err != nil {
		t.Fatal(err)
	}

	user_registration, err := service.GetUserRegistration("testid2")

	if err != nil {
		t.Fatal(err)
	}

	expected_registration := base_registration
	expected_registration.ID = "testid2"
	expected_registration.FirstName = "first2"
	expected_registration.LastName = "last2"

	if !reflect.DeepEqual(user_registration, &expected_registration) {
		t.Errorf("Wrong user info.\nExpected %v\ngot %v\n", expected_registration, user_registration)
	}

	CleanupTestDB(t)
}

/*
	Service level test for updating user registration in the db
*/
func TestUpdateUserRegistrationService(t *testing.T) {
	SetupTestDB(t)

	updated_registration := base_registration
	updated_registration.ID = "testid"
	updated_registration.FirstName = "first2"
	updated_registration.LastName = "last2"
	err := service.UpdateUserRegistration("testid", updated_registration)

	if err != nil {
		t.Fatal(err)
	}

	user_registration, err := service.GetUserRegistration("testid")

	if err != nil {
		t.Fatal(err)
	}

	expected_registration := base_registration
	expected_registration.ID = "testid"
	expected_registration.FirstName = "first2"
	expected_registration.LastName = "last2"

	if !reflect.DeepEqual(user_registration, &expected_registration) {
		t.Errorf("Wrong user info.\nExpected %v\ngot %v\n", expected_registration, user_registration)
	}

	CleanupTestDB(t)
}

/*
	Service level test for filtering user registrations in the db
*/
func TestGetFilteredUserRegistrationsService(t *testing.T) {
	SetupTestDB(t)

	registration_1 := base_registration

	registration_2 := base_registration
	registration_2.ID = "testid2"
	err := service.CreateUserRegistration(registration_2.ID, registration_2)
	if err != nil {
		t.Fatal(err)
	}

	// Test capitalized keys
	parameters := map[string][]string{
		"ID": []string{"testid"},
	}
	user_registrations, err := service.GetFilteredUserRegistrations(parameters)
	if err != nil {
		t.Fatal(err)
	}

	expected_registrations := []models.UserRegistration{
		registration_1,
	}
	if !reflect.DeepEqual(user_registrations, &expected_registrations) {
		t.Errorf("Wrong user info.\nExpected %v\ngot %v\n", expected_registrations, user_registrations)
	}

	// Test multiple values
	parameters = map[string][]string{
		"id": []string{"testid,testid2"},
	}
	user_registrations, err = service.GetFilteredUserRegistrations(parameters)
	if err != nil {
		t.Fatal(err)
	}

	expected_registrations = []models.UserRegistration{
		registration_1,
		registration_2,
	}
	if !reflect.DeepEqual(user_registrations, &expected_registrations) {
		t.Errorf("Wrong user info.\nExpected %v\ngot %v\n", expected_registrations, user_registrations)
	}

	// Test type casting
	parameters = map[string][]string{
		"firstName": []string{"first"},
		"age": []string{"20"},
		"isNovice": []string{"true"},
	}
	user_registrations, err = service.GetFilteredUserRegistrations(parameters)
	if err != nil {
		t.Fatal(err)
	}

	expected_registrations = []models.UserRegistration{
		registration_1,
		registration_2,
	}
	if !reflect.DeepEqual(user_registrations, &expected_registrations) {
		t.Errorf("Wrong user info.\nExpected %v\ngot %v\n", expected_registrations, user_registrations)
	}

	CleanupTestDB(t)
}

var base_registration models.UserRegistration = models.UserRegistration{
	ID:                   "testid",
	FirstName:            "first",
	LastName:             "last",
	Email:                "test@gmail.com",
	ShirtSize:            "M",
	Diet:                 "NONE",
	Age:                  20,
	GraduationYear:       2019,
	Transportation:       "BUS",
	School:               "University of Illinois at Urbana-Champaign",
	Major:                "Computer Science",
	Gender:               "MALE",
	ProfessionalInterest: "INTERNSHIP",
	GitHub:               "githubusername",
	Linkedin:             "linkedinusername",
	Interests:            "things",
	IsNovice:             true,
	IsPrivate:            false,
	PhoneNumber:          "555-287-2903",
	LongForms: []models.UserLongForm{
		models.UserLongForm{
			Response: "longformresponse",
		},
	},
	ExtraInfos: []models.UserExtraInfo{
		models.UserExtraInfo{
			Response: "extrainforesponse",
		},
	},
	OsContributors: []models.UserOsContributor{
		models.UserOsContributor{
			Name:        "contributorname",
			ContactInfo: "contact@test.com",
		},
	},
	Collaborators: []models.UserCollaborator{
		models.UserCollaborator{
			Github: "collaboratorgithub",
		},
	},
}
