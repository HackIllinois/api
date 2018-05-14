package tests

import (
	"bytes"
	"encoding/json"
	"github.com/HackIllinois/api-registration/config"
	"github.com/HackIllinois/api-registration/controller"
	"github.com/HackIllinois/api-registration/database"
	"github.com/HackIllinois/api-registration/models"
	"github.com/HackIllinois/api-registration/service"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
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

/*
	End to end test for getting user registration
*/
func TestGetUserRegistrationEndpoint(t *testing.T) {
	SetupTestDB(t)

	req, err := http.NewRequest("GET", "/registration/testid/", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "testid"})

	if err != nil {
		t.Fatal(err)
	}

	res_recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetUserRegistration)

	handler.ServeHTTP(res_recorder, req)

	if res_recorder.Code != http.StatusOK {
		t.Errorf("Wrong status code. Expected %v Got %v", http.StatusOK, res_recorder.Code)
	}

	var user_registration models.UserRegistration
	json.NewDecoder(res_recorder.Body).Decode(&user_registration)

	expected_registration := base_registration

	if !reflect.DeepEqual(user_registration, expected_registration) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_registration, user_registration)
	}

	CleanupTestDB(t)
}

/*
	End to end test for creating current user registration
*/
func TestCreateCurrentUserRegistrationEndpoint(t *testing.T) {
	SetupTestDB(t)

	req_body := base_registration
	req_body.FirstName = "first2"
	req_body.LastName = "last2"

	body := bytes.Buffer{}
	json.NewEncoder(&body).Encode(&req_body)

	req, err := http.NewRequest("POST", "/registration/", &body)
	req.Header.Set("HackIllinois-Identity", "testid2")

	if err != nil {
		t.Fatal(err)
	}

	res_recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.CreateCurrentUserRegistration)

	handler.ServeHTTP(res_recorder, req)

	if res_recorder.Code != http.StatusOK {
		t.Errorf("Wrong status code. Expected %v Got %v", http.StatusOK, res_recorder.Code)
	}

	var user_registration models.UserRegistration
	json.NewDecoder(res_recorder.Body).Decode(&user_registration)

	expected_registration := base_registration
	expected_registration.ID = "testid2"
	expected_registration.FirstName = "first2"
	expected_registration.LastName = "last2"

	if !reflect.DeepEqual(user_registration, expected_registration) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_registration, user_registration)
	}

	CleanupTestDB(t)
}

/*
	End to end test for updating current user registration
*/
func TestUpdateCurrentUserRegistrationEndpoint(t *testing.T) {
	SetupTestDB(t)

	req_body := base_registration
	req_body.FirstName = "firstupdate"
	req_body.LastName = "lastupdate"

	body := bytes.Buffer{}
	json.NewEncoder(&body).Encode(&req_body)

	req, err := http.NewRequest("PUT", "/registration/", &body)
	req.Header.Set("HackIllinois-Identity", "testid")

	if err != nil {
		t.Fatal(err)
	}

	res_recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.UpdateCurrentUserRegistration)

	handler.ServeHTTP(res_recorder, req)

	if res_recorder.Code != http.StatusOK {
		t.Errorf("Wrong status code. Expected %v Got %v", http.StatusOK, res_recorder.Code)
	}

	var user_registration models.UserRegistration
	json.NewDecoder(res_recorder.Body).Decode(&user_registration)

	expected_registration := base_registration
	expected_registration.ID = "testid"
	expected_registration.FirstName = "firstupdate"
	expected_registration.LastName = "lastupdate"

	if !reflect.DeepEqual(user_registration, expected_registration) {
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
