package tests

import (
	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/services/user/config"
	"github.com/HackIllinois/api/services/user/models"
	"github.com/HackIllinois/api/services/user/service"
	"reflect"
	"testing"
)

var db database.Database

func init() {
	db_connection, err := database.InitDatabase(config.USER_DB_HOST, config.USER_DB_NAME)

	if err != nil {
		panic(err)
	}

	db = db_connection
}

/*
	Initialize databse with test user info
*/
func SetupTestDB(t *testing.T) {
	err := db.Insert("info", &models.UserInfo{
		ID:       "testid",
		Username: "testusername",
		Email:    "testemail@domain.com",
	})

	if err != nil {
		t.Fatal(err)
	}

	err = db.Insert("info", &models.UserInfo{
		ID:       "testid2",
		Username: "testusername",
		Email:    "testemail@domain.com",
	})

	if err != nil {
		t.Fatal(err)
	}
}

/*
	Drop test db
*/
func CleanupTestDB(t *testing.T) {
	err := db.DropDatabase()

	if err != nil {
		t.Fatal(err)
	}
}

/*
	Service level test for getting user info from db
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
	Service level test for setting user info in the db
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
	Service level test for getting filtered user info from db
*/
func TestGetFilteredUserInfoService(t *testing.T) {
	SetupTestDB(t)

	user_info_1, err := service.GetUserInfo("testid")
	if err != nil {
		t.Fatal(err)
	}

	user_info_2, err := service.GetUserInfo("testid2")
	if err != nil {
		t.Fatal(err)
	}

	parameters := map[string][]string{
		"username": []string{"testusername"},
	}
	filtered_info, err := service.GetFilteredUserInfo(parameters)
	if err != nil {
		t.Fatal(err)
	}

	expected_info := &models.FilteredUsers{
		[]models.UserInfo{
			*user_info_1,
			*user_info_2,
		},
	}

	if !reflect.DeepEqual(filtered_info, expected_info) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_info, filtered_info)
	}

	CleanupTestDB(t)
}
