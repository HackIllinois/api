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

	expected_registration := models.UserRegistration{
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
	End to end test for creating current user registration
*/
func TestCreateCurrentUserRegistrationEndpoint(t *testing.T) {
	SetupTestDB(t)

	req_body := models.UserRegistration{
		FirstName: "testfirstname2",
		LastName:  "testlastname2",
	}

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

	expected_registration := models.UserRegistration{
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
	End to end test for updating current user registration
*/
func TestUpdateCurrentUserRegistrationEndpoint(t *testing.T) {
	SetupTestDB(t)

	req_body := models.UserRegistration{
		FirstName: "testupdatedfirstname",
		LastName:  "testupdatedlastname",
	}

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

	expected_registration := models.UserRegistration{
		ID:        "testid",
		FirstName: "testupdatedfirstname",
		LastName:  "testupdatedlastname",
	}

	if !reflect.DeepEqual(user_registration, expected_registration) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_registration, user_registration)
	}

	CleanupTestDB(t)
}
