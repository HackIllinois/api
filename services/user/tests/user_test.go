package tests

import (
	"fmt"
	"net/url"
	"os"
	"reflect"
	"testing"

	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/services/user/config"
	"github.com/HackIllinois/api/services/user/models"
	"github.com/HackIllinois/api/services/user/service"
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

	db, err = database.InitDatabase(config.USER_DB_HOST, config.USER_DB_NAME)

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	return_code := m.Run()

	os.Exit(return_code)
}

/*
	Initialize database with test user info
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

func SetupFilterTestDB(t *testing.T) {
	err := db.Insert("info", &models.UserInfo{
		ID:        "testid1",
		FirstName: "Alex",
		Username:  "testusername",
		Email:     "testemail@domain.com",
	})

	if err != nil {
		t.Fatal(err)
	}

	err = db.Insert("info", &models.UserInfo{
		ID:        "testid2",
		FirstName: "Charlie",
		Username:  "testusername",
		Email:     "testemail@domain.com",
	})

	err = db.Insert("info", &models.UserInfo{
		ID:        "testid3",
		FirstName: "Bobby",
		Username:  "testusername",
		Email:     "testemail@domain.com",
	})

	err = db.Insert("info", &models.UserInfo{
		ID:        "testid4",
		FirstName: "Bobby",
		LastName:  "Adamson",
		Username:  "test-two-parameter-filter",
	})

	err = db.Insert("info", &models.UserInfo{
		ID:        "testid5",
		FirstName: "Bobby",
		LastName:  "Zulu",
		Username:  "test-two-parameter-filter",
	})
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
		"username": {"testusername"},
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

func TestGetFilteredUserInfoWithSortingService(t *testing.T) {
	SetupFilterTestDB(t)

	user_info_1, err := service.GetUserInfo("testid1")
	if err != nil {
		t.Fatal(err)
	}
	user_info_2, err := service.GetUserInfo("testid2")
	user_info_3, err := service.GetUserInfo("testid3")

	// Sort by first name and expect: Alex, Bobby, Charlie

	parameters := map[string][]string{
		"username": {"testusername"},
		"sortby":   {"FiRsTNAmE"},
	}

	filtered_info, err := service.GetFilteredUserInfo(parameters)
	if err != nil {
		t.Fatal(err)
	}

	expected_info := &models.FilteredUsers{
		[]models.UserInfo{
			*user_info_1, // Alex
			*user_info_3, // Bobby
			*user_info_2, // Charlie
		},
	}

	if !reflect.DeepEqual(filtered_info, expected_info) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_info, filtered_info)
	}

	// Reverse the sort and expect: Charlie, Bobby, Alex.

	parameters = map[string][]string{
		"username": {"testusername"},
		"sortby":   {"-FiRsTNAmE"},
	}

	filtered_info, err = service.GetFilteredUserInfo(parameters)
	if err != nil {
		t.Fatal(err)
	}

	expected_info = &models.FilteredUsers{
		[]models.UserInfo{
			*user_info_2, // Charlie
			*user_info_3, // Bobby
			*user_info_1, // Alex
		},
	}

	if !reflect.DeepEqual(filtered_info, expected_info) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_info, filtered_info)
	}

	// Sort by two parameters and expect: Bobby Adamson, Bobby Zulu

	user_info_1, err = service.GetUserInfo("testid5")
	user_info_2, err = service.GetUserInfo("testid4")
	parameters = map[string][]string{
		"username": {"test-two-parameter-filter"},
		"sortby":   {"firstName,lastName"},
	}

	filtered_info, err = service.GetFilteredUserInfo(parameters)
	if err != nil {
		t.Fatal(err)
	}

	expected_info = &models.FilteredUsers{
		[]models.UserInfo{
			*user_info_2, // Bobby Adamson
			*user_info_1, // Bobby Zulu
		},
	}

	if !reflect.DeepEqual(filtered_info, expected_info) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_info, filtered_info)
	}

	CleanupTestDB(t)
}

/*
	Service level test for generating QR code URI
*/
func TestGetQrInfo(t *testing.T) {
	SetupTestDB(t)

	actual_uri, err := service.GetQrInfo("testid")

	if err != nil {
		t.Fatal(err)
	}

	parsed_uri, err := url.Parse(actual_uri)

	if err != nil {
		t.Fatal(err)
	}

	actual_query_params, err := url.ParseQuery(parsed_uri.RawQuery)

	if err != nil {
		t.Fatal(err)
	}

	expected_query_params := url.Values{
		"userId": []string{"testid"},
	}

	if !reflect.DeepEqual(expected_query_params, actual_query_params) {
		t.Errorf("Wrong QR code URI. Expected %v, got %v", expected_query_params, actual_query_params)
	}

	CleanupTestDB(t)
}
