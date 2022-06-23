package tests

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/services/auth/config"
	"github.com/HackIllinois/api/services/auth/models"
	"github.com/HackIllinois/api/services/auth/service"
)

var db database.Database

func TestMain(m *testing.M) {
	err := config.Initialize()

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)

	}

	err = service.Initialize()

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	db, err = database.InitDatabase(config.AUTH_DB_HOST, config.AUTH_DB_NAME)

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	return_code := m.Run()

	os.Exit(return_code)
}

/*
	Initialize roles db with a test user
*/
func SetupTestDB(t *testing.T) {
	err := db.Insert("roles", &models.UserRoles{
		ID:    "testid",
		Roles: []string{"User"},
	}, nil)

	if err != nil {
		t.Fatal(err)
	}
}

/*
	Drop test db
*/
func CleanupTestDB(t *testing.T) {
	err := db.DropDatabase(nil)

	if err != nil {
		t.Fatal(err)
	}
}

/*
	Service level test for getting a user's roles from the db
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
	Service level test for adding a role to a user in the DB
*/
func TestAddRoleService(t *testing.T) {
	SetupTestDB(t)

	expected_roles := []string{"User", "Admin"}
	err := service.AddUserRole("testid", "Admin")

	if err != nil {
		t.Fatal(err)
	}

	roles, err := service.GetUserRoles("testid", false)

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(roles, expected_roles) {
		t.Errorf("Wrong user roles. Expected %v, got %v", expected_roles, roles)
	}

	// Test adding duplicate role
	err = service.AddUserRole("testid", "Admin")

	if err != nil {
		t.Fatal(err)
	}

	roles, err = service.GetUserRoles("testid", false)

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(roles, expected_roles) {
		t.Errorf("Wrong user roles. Expected %v, got %v", expected_roles, roles)
	}

	CleanupTestDB(t)
}

/*
	Service level test for removing a role from a user in the DB
*/
func TestRemoveRoleService(t *testing.T) {
	SetupTestDB(t)

	expected_roles := []string{}
	err := service.RemoveUserRole("testid", "User")

	if err != nil {
		t.Fatal(err)
	}

	roles, err := service.GetUserRoles("testid", false)

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(roles, expected_roles) {
		t.Errorf("Wrong user roles. Expected %v, got %v", expected_roles, roles)
	}

	// Ensure removing a user's role fails if they do not have that role
	err = service.RemoveUserRole("testid", "User")

	if err == nil {
		t.Errorf("Able to remove role \"User\" from a user that does not have the \"User\" role")
	}

	CleanupTestDB(t)
}

func TestGetUsersByRoleService(t *testing.T) {
	SetupTestDB(t)

	err := db.Insert("roles", &models.UserRoles{
		ID:    "testid2",
		Roles: []string{"Staff"},
	}, nil)

	if err != nil {
		t.Fatal(err)
	}

	userids, err := service.GetUsersByRole("Staff")

	if err != nil {
		t.Fatal(err)
	}

	expected_userids := []string{"testid2"}

	if !reflect.DeepEqual(userids, expected_userids) {
		t.Errorf("Wrong user ids. Expected %v, got %v", expected_userids, userids)
	}

	CleanupTestDB(t)
}
