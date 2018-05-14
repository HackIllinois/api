package tests

import (
	"bytes"
	"encoding/json"
	"github.com/HackIllinois/api-auth/config"
	"github.com/HackIllinois/api-auth/controller"
	"github.com/HackIllinois/api-auth/database"
	"github.com/HackIllinois/api-auth/models"
	"github.com/HackIllinois/api-auth/service"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

/*
	Initialize roles database with a test user
*/
func SetupTestDB(t *testing.T) {
	err := database.Insert("roles", &models.UserRoles{
		ID:    "testid",
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
