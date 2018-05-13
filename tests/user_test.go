package tests

import (
	"bytes"
	"encoding/json"
	"github.com/HackIllinois/api-user/config"
	"github.com/HackIllinois/api-user/controller"
	"github.com/HackIllinois/api-user/database"
	"github.com/HackIllinois/api-user/models"
	"github.com/HackIllinois/api-user/service"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

/*
	Initialize databse with test user info
*/
func SetupTestDB(t *testing.T) {
	err := database.Insert("info", &models.UserInfo{
		ID:       "testid",
		Username: "testusername",
		Email:    "testemail@domain.com",
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

	err := session.DB(config.USER_DB_NAME).DropDatabase()

	if err != nil {
		t.Fatal(err)
	}
}

/*
	Service level test for getting user info from database
*/
func TestGetUserInfoService(t *testing.T) {
	SetupTestDB(t)

	user_info, err := service.GetUserInfo("testid")

	if err != nil {
		t.Fatal(err)
	}

	expected_info := &models.UserInfo{
		ID:       "testid",
		Username: "testusername",
		Email:    "testemail@domain.com",
	}

	if !reflect.DeepEqual(user_info, expected_info) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_info, user_info)
	}

	CleanupTestDB(t)
}

/*
	Service level test for setting user info in the database
*/
func TestSetUserInfoService(t *testing.T) {
	SetupTestDB(t)

	err := service.SetUserInfo("testid2", models.UserInfo{
		ID:       "testid2",
		Username: "testusername2",
		Email:    "testemail2@domain.com",
	})

	if err != nil {
		t.Fatal(err)
	}

	user_info, err := service.GetUserInfo("testid2")

	if err != nil {
		t.Fatal(err)
	}

	expected_info := &models.UserInfo{
		ID:       "testid2",
		Username: "testusername2",
		Email:    "testemail2@domain.com",
	}

	if !reflect.DeepEqual(user_info, expected_info) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_info, user_info)
	}

	CleanupTestDB(t)
}

/*
	End to end test for getting user info
*/
func TestGetUserInfoEndpoint(t *testing.T) {
	SetupTestDB(t)

	req, err := http.NewRequest("GET", "/user/testid/", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "testid"})

	if err != nil {
		t.Fatal(err)
	}

	res_recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetUserInfo)

	handler.ServeHTTP(res_recorder, req)

	if res_recorder.Code != http.StatusOK {
		t.Errorf("Wrong status code. Expected %v Got %v", http.StatusOK, res_recorder.Code)
	}

	var user_info models.UserInfo
	json.NewDecoder(res_recorder.Body).Decode(&user_info)

	expected_info := models.UserInfo{
		ID:       "testid",
		Username: "testusername",
		Email:    "testemail@domain.com",
	}

	if !reflect.DeepEqual(user_info, expected_info) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_info, user_info)
	}

	CleanupTestDB(t)
}

/*
	End to end test for setting user info
*/
func TestSetUserInfoEndpoint(t *testing.T) {
	SetupTestDB(t)

	req_body := models.UserInfo{
		ID:       "testid2",
		Username: "testusername2",
		Email:    "testemail2@domain.com",
	}

	body := bytes.Buffer{}
	json.NewEncoder(&body).Encode(&req_body)

	req, err := http.NewRequest("POST", "/user/", &body)

	if err != nil {
		t.Fatal(err)
	}

	res_recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.SetUserInfo)

	handler.ServeHTTP(res_recorder, req)

	if res_recorder.Code != http.StatusOK {
		t.Errorf("Wrong status code. Expected %v Got %v", http.StatusOK, res_recorder.Code)
	}

	var user_info models.UserInfo
	json.NewDecoder(res_recorder.Body).Decode(&user_info)

	expected_info := models.UserInfo{
		ID:       "testid2",
		Username: "testusername2",
		Email:    "testemail2@domain.com",
	}

	if !reflect.DeepEqual(user_info, expected_info) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_info, user_info)
	}

	CleanupTestDB(t)
}
