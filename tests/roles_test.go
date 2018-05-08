package tests

import (
	"testing"
	"reflect"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"github.com/HackIllinois/api-auth/controller"
	"github.com/HackIllinois/api-auth/database"
	"github.com/HackIllinois/api-auth/models"
	"github.com/HackIllinois/api-auth/config"
)

func SetupTestDB(t *testing.T) {
	err := database.Insert("roles", &models.UserRoles {
		ID: "testid",
		Roles: []string{"User"},
	})

	if err != nil {
		t.Fatal(err)
	}
}

func CleanupTestDB(t *testing.T) {
	session := database.GetSession()
	defer session.Close()

	err := session.DB(config.AUTH_DB_NAME).DropDatabase()


	if err != nil {
		t.Fatal(err)
	}
}

func TestGetRoles(t *testing.T) {
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
		t.Errorf("Wrong user id. Expected %v, got %v", expected_roles, user_roles.Roles)
	}

	CleanupTestDB(t)
}
