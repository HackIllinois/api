package tests

import (
	"github.com/ReflectionsProjections/api/common/database"
	"github.com/ReflectionsProjections/api/services/auth/config"
	"github.com/ReflectionsProjections/api/services/auth/models"
	"github.com/ReflectionsProjections/api/services/auth/service"
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
	Service level test for setting a user's roles in the db
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
