package tests

import (
	"fmt"
	"net/url"
	"os"
	"reflect"
	"testing"

	common_config "github.com/HackIllinois/api/common/config"
	"github.com/HackIllinois/api/common/database"
	user_config "github.com/HackIllinois/api/services/user/config"
	"github.com/HackIllinois/api/services/user/models"
	"github.com/HackIllinois/api/services/user/service"
	"github.com/golang-jwt/jwt/v4"
)

var db database.Database

func TestMain(m *testing.M) {
	err := user_config.Initialize()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)

	}

	err = service.Initialize()

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	db, err = database.InitDatabase(user_config.USER_DB_HOST, user_config.USER_DB_NAME)

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
	}, nil)

	if err != nil {
		t.Fatal(err)
	}

	err = db.Insert("info", &models.UserInfo{
		ID:       "testid2",
		Username: "testusername",
		Email:    "testemail@domain.com",
	}, nil)

	if err != nil {
		t.Fatal(err)
	}
}

/*
Initialize db for sortby filter tests
*/
func SetupFilterTestDB(t *testing.T) {
	err := db.Insert("info", &models.UserInfo{
		ID:        "testid1",
		FirstName: "Alex",
		Username:  "testusername",
		Email:     "testemail@domain.com",
	}, nil)

	if err != nil {
		t.Fatal(err)
	}

	err = db.Insert("info", &models.UserInfo{
		ID:        "testid2",
		FirstName: "Charlie",
		Username:  "testusername",
		Email:     "testemail@domain.com",
	}, nil)

	err = db.Insert("info", &models.UserInfo{
		ID:        "testid3",
		FirstName: "Bobby",
		Username:  "testusername",
		Email:     "testemail@domain.com",
	}, nil)

	err = db.Insert("info", &models.UserInfo{
		ID:        "testid4",
		FirstName: "Bobby",
		LastName:  "Adamson",
		Username:  "test-two-parameter-filter",
	}, nil)

	err = db.Insert("info", &models.UserInfo{
		ID:        "testid5",
		FirstName: "Bobby",
		LastName:  "Zulu",
		Username:  "test-two-parameter-filter",
	}, nil)
}

/*
Initialize db for pagination, filter tests
*/
func SetupPaginationDB(t *testing.T) {
	err := db.Insert("info", &models.UserInfo{
		ID:       "testid",
		Username: "testusername",
		Email:    "testemail@domain.com",
	}, nil)

	if err != nil {
		t.Fatal(err)
	}

	err = db.Insert("info", &models.UserInfo{
		ID:       "testid2",
		Username: "testusername",
		Email:    "testemail@domain.com",
	}, nil)

	err = db.Insert("info", &models.UserInfo{
		ID:       "testid3",
		Username: "testusername",
		Email:    "testemail@domain.com",
	}, nil)

	err = db.Insert("info", &models.UserInfo{
		ID:       "testid4",
		Username: "testusername",
		Email:    "testemail@domain.com",
	}, nil)

	err = db.Insert("info", &models.UserInfo{
		ID:       "testid5",
		Username: "testusername",
		Email:    "testemail@domain.com",
	}, nil)
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
Service level test for getting user info from db
*/
func TestGetUserInfoService(t *testing.T) {
	SetupTestDB(t)

	user_info, err := service.GetUserInfo("testid", nil)

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
	}, nil)

	if err != nil {
		t.Fatal(err)
	}

	user_info, err := service.GetUserInfo("testid2", nil)

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

	user_info_1, err := service.GetUserInfo("testid", nil)
	if err != nil {
		t.Fatal(err)
	}

	user_info_2, err := service.GetUserInfo("testid2", nil)
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

/*
Test Sortby parameter
*/
func TestGetFilteredUserInfoWithSortingService(t *testing.T) {
	SetupFilterTestDB(t)

	user_info_1, err := service.GetUserInfo("testid1", nil)
	if err != nil {
		t.Fatal(err)
	}
	user_info_2, err := service.GetUserInfo("testid2", nil)
	user_info_3, err := service.GetUserInfo("testid3", nil)

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

	user_info_1, err = service.GetUserInfo("testid5", nil)
	user_info_2, err = service.GetUserInfo("testid4", nil)
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
Test Pagination parameter
*/
func TestGetFilteredUserInfoServicePagination(t *testing.T) {
	SetupPaginationDB(t)

	user_info_1, err := service.GetUserInfo("testid", nil)
	if err != nil {
		t.Fatal(err)
	}

	user_info_2, err := service.GetUserInfo("testid2", nil)
	if err != nil {
		t.Fatal(err)
	}

	user_info_3, err := service.GetUserInfo("testid3", nil)
	if err != nil {
		t.Fatal(err)
	}

	user_info_4, err := service.GetUserInfo("testid4", nil)
	if err != nil {
		t.Fatal(err)
	}

	user_info_5, err := service.GetUserInfo("testid5", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Filter to page two, expect user_info_2
	parameters := map[string][]string{
		"username": {"testusername"},
		"p":        {"2"},
		"limit":    {"1"},
	}
	filtered_info, err := service.GetFilteredUserInfo(parameters)
	if err != nil {
		t.Fatal(err)
	}

	expected_info := &models.FilteredUsers{
		[]models.UserInfo{
			*user_info_2,
		},
	}

	if !reflect.DeepEqual(filtered_info, expected_info) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_info, filtered_info)
	}

	// Filter to page two with two users per page

	// page_1: user_info_1, user_info_2
	parameters = map[string][]string{
		"username": {"testusername"},
		"p":        {"1"},
		"limit":    {"2"},
	}
	filtered_info, err = service.GetFilteredUserInfo(parameters)
	if err != nil {
		t.Fatal(err)
	}

	expected_info = &models.FilteredUsers{
		[]models.UserInfo{
			*user_info_1,
			*user_info_2,
		},
	}

	if !reflect.DeepEqual(filtered_info, expected_info) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_info, filtered_info)
	}

	// page 2: user_info_3, user_info_4
	parameters = map[string][]string{
		"username": {"testusername"},
		"p":        {"2"},
		"limit":    {"2"},
	}
	filtered_info, err = service.GetFilteredUserInfo(parameters)
	if err != nil {
		t.Fatal(err)
	}

	expected_info = &models.FilteredUsers{
		[]models.UserInfo{
			*user_info_3,
			*user_info_4,
		},
	}

	if !reflect.DeepEqual(filtered_info, expected_info) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_info, filtered_info)
	}

	// page 3: user_info_5
	parameters = map[string][]string{
		"username": {"testusername"},
		"p":        {"3"},
		"limit":    {"2"},
	}
	filtered_info, err = service.GetFilteredUserInfo(parameters)
	if err != nil {
		t.Fatal(err)
	}

	expected_info = &models.FilteredUsers{
		[]models.UserInfo{
			*user_info_5,
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

	signed_token := actual_query_params.Get("userToken")

	token, err := jwt.Parse(signed_token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(common_config.TOKEN_SECRET), nil
	})

	if err != nil {
		t.Fatal(err)
	}

	actual_user_id := token.Claims.(jwt.MapClaims)["userId"]

	expected_user_id := "testid"

	if !reflect.DeepEqual(actual_user_id, expected_user_id) {
		t.Errorf("Wrong QR code URI. Expected %v, got %v", expected_user_id, actual_user_id)
	}

	CleanupTestDB(t)
}
