package tests

import (
	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/services/auth/config"
	"github.com/HackIllinois/api/services/auth/models"
	"github.com/HackIllinois/api/services/auth/service"
	"reflect"
	"testing"
)

var db database.MongoDatabase

func init() {
	db_connection, err := database.InitMongoDatabase(config.AUTH_DB_HOST, config.AUTH_DB_NAME)

	if err != nil {
		panic(err)
	}

	db = db_connection
}

/*
	Initialize roles db with a test user
*/
func SetupTestDB(t *testing.T) {
	err := db.Insert("roles", &models.UserRoles{
		ID:    "testid",
		Roles: []string{"User"},
	})

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

	err := session.DB(config.AUTH_DB_NAME).DropDatabase()

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
