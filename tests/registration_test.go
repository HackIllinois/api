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
	user_registration := base_registration
	err := database.Insert("attendees", &user_registration)

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

	expected_registration := base_registration

	if !reflect.DeepEqual(user_registration, &expected_registration) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_registration, user_registration)
	}

	CleanupTestDB(t)
}

/*
	Service level test for creating user registration in the database
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
		t.Errorf("Wrong user info. Expected %v, got %v", expected_registration, user_registration)
	}

	CleanupTestDB(t)
}

/*
	Service level test for updating user registration in the database
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
		t.Errorf("Wrong user info. Expected %v, got %v", expected_registration, user_registration)
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
