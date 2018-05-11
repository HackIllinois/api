package tests

import (
	"testing"
	"reflect"
	"bytes"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"github.com/HackIllinois/api-auth/controller"
	"github.com/HackIllinois/api-auth/service"
	"github.com/HackIllinois/api-auth/database"
	"github.com/HackIllinois/api-auth/models"
	"github.com/HackIllinois/api-auth/config"
)

/*
	Initialize roles database with a test user
*/
func SetupTestDB(t *testing.T) {
	err := database.Insert("roles", &models.UserRoles {
		ID: "testid",
		Roles: []string{"User"},
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

	err := session.DB(config.AUTH_DB_NAME).DropDatabase()


	if err != nil {
		t.Fatal(err)
	}
}

/*
	End to end test for GET /auth/roles/?id=ID
*/
func TestGetRolesEndpoint(t *testing.T) {
	SetupTestDB(t)

	req, err := http.NewRequest("GET", "/auth/roles/?id=testid", nil)

	if err != nil {
		t.Fatal(err)
	}

	res_recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetRoles)

	handler.ServeHTTP(res_recorder, req)

	if res_recorder.Code != http.StatusOK {
		t.Errorf("Wrong status code. Expected %v Got %v", http.StatusOK, res_recorder.Code)
	}

	var user_roles models.UserRoles
	json.NewDecoder(res_recorder.Body).Decode(&user_roles)

	if user_roles.ID != "testid" {
		t.Errorf("Wrong user id. Expected %v, got %v", "testid", user_roles.ID)
	}

	expected_roles := []string{"User"}
	if !reflect.DeepEqual(user_roles.Roles, expected_roles) {
		t.Errorf("Wrong user roles. Expected %v, got %v", expected_roles, user_roles.Roles)
	}

	CleanupTestDB(t)
}

/*
	End to end test for PUT /auth/roles/
*/
func TestPutRolesEndpoint(t *testing.T) {
	SetupTestDB(t)

	req_body := models.UserRoles {
		ID: "testid",
		Roles: []string{"User", "Admin"},
	}

	body := bytes.Buffer{}
	json.NewEncoder(&body).Encode(&req_body)

	req, err := http.NewRequest("PUT", "/auth/roles/", &body)

	if err != nil {
		t.Fatal(err)
	}

	res_recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.SetRoles)

	handler.ServeHTTP(res_recorder, req)

	if res_recorder.Code != http.StatusOK {
		t.Errorf("Wrong status code. Expected %v Got %v", http.StatusOK, res_recorder.Code)
	}

	var user_roles models.UserRoles
	json.NewDecoder(res_recorder.Body).Decode(&user_roles)

	if user_roles.ID != "testid" {
		t.Errorf("Wrong user id. Expected %v, got %v", "testid", user_roles.ID)
	}

	expected_roles := []string{"User", "Admin"}
	if !reflect.DeepEqual(user_roles.Roles, expected_roles) {
		t.Errorf("Wrong user roles. Expected %v, got %v", expected_roles, user_roles.Roles)
	}

	CleanupTestDB(t)
}

/*
	Service level test for getting a user's roles from the database
*/
func TestGetRolesService(t *testing.T) {
	SetupTestDB(t)

	expected_roles := []string{"User"}
	roles, err := service.GetUserRoles("testid", false)

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(roles, expected_roles) {
		t.Errorf("Wrong user roles. Expected %v, got %v", expected_roles, roles)
	}

	roles, err = service.GetUserRoles("testid2", true)

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(roles, expected_roles) {
		t.Errorf("Wrong user roles. Expected %v, got %v", expected_roles, roles)
	}

	CleanupTestDB(t)
}

/*
	Service level test for setting a user's roles in the database
*/
func TestPutRolesService(t *testing.T) {
	SetupTestDB(t)

	updated_roles := []string{"User", "Admin"}
	err := service.SetUserRoles("testid", updated_roles)

	if err != nil {
		t.Fatal(err)
	}

	roles, err := service.GetUserRoles("testid", false)

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(roles, updated_roles) {
		t.Errorf("Wrong user roles. Expected %v, got %v", updated_roles, roles)
	}

	CleanupTestDB(t)
}
