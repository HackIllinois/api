package tests

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/HackIllinois/api/services/user/models"
	user_models "github.com/HackIllinois/api/services/user/models"
)

var user_info_1 user_models.UserInfo
var user_info_2 user_models.UserInfo
var user_info_3 user_models.UserInfo
var user_info_4 user_models.UserInfo

func AddTestDataToDatabase() {
	user_info_1 = user_models.UserInfo{
		ID:        "12345",
		Username:  "johnsmith",
		FirstName: "John",
		LastName:  "Smith",
		Email:     "john.smith@aol.com",
	}
	user_info_2 = user_models.UserInfo{
		ID:        "12345",
		Username:  "hello",
		FirstName: "John",
		LastName:  "Smith",
		Email:     "Hi.Hey@aol.com",
	}

	user_info_3 = user_models.UserInfo{
		ID:        "54321",
		Username:  "james",
		FirstName: "James",
		LastName:  "Names",
		Email:     "James.Names@aol.com",
	}
	user_info_4 = user_models.UserInfo{
		ID:        "54321",
		Username:  "A",
		FirstName: "Aaa",
		LastName:  "Bbbbb",
		Email:     "a.b@aol.com",
	}
	client.Database(user_db_name).Collection("info").InsertOne(context.Background(), user_info_1)
	client.Database(user_db_name).Collection("info").InsertOne(context.Background(), user_info_2)
	client.Database(user_db_name).Collection("info").InsertOne(context.Background(), user_info_3)
	client.Database(user_db_name).Collection("info").InsertOne(context.Background(), user_info_4)
}

func ClearDatabase() {
	client.Database(user_db_name).Drop(context.Background())
}

func TestGetFilterByID(t *testing.T) {
	AddTestDataToDatabase()
	defer ClearDatabase()

	endpoint_address := "/user/filter/?id=12345"

	recieved_profile := user_models.FilteredUsers{}

	response, err := admin_client.New().Get(endpoint_address).ReceiveSuccess(&recieved_profile)

	expected_info := models.FilteredUsers{
		[]models.UserInfo{
			user_info_1,
			user_info_2,
		},
	}

	if err != nil {
		t.Errorf("Unable to make request")
	}
	if response.StatusCode != http.StatusOK {
		t.Errorf("Request returned HTTP error %d", response.StatusCode)
	}
	if !reflect.DeepEqual(recieved_profile, expected_info) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_info, recieved_profile)
	}
}

func TestGetFilterByUserName(t *testing.T) {
	AddTestDataToDatabase()
	defer ClearDatabase()

	endpoint_address := "/user/filter/?username=johnsmith"

	recieved_profile := user_models.FilteredUsers{}

	response, err := admin_client.New().Get(endpoint_address).ReceiveSuccess(&recieved_profile)

	expected_info := models.FilteredUsers{
		[]models.UserInfo{
			user_info_1,
		},
	}

	if err != nil {
		t.Errorf("Unable to make request")
	}
	if response.StatusCode != http.StatusOK {
		t.Errorf("Request returned HTTP error %d", response.StatusCode)
	}
	if !reflect.DeepEqual(recieved_profile, expected_info) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_info, recieved_profile)
	}
}

func TestGetFilterByFirstName(t *testing.T) {
	AddTestDataToDatabase()
	defer ClearDatabase()

	endpoint_address := "/user/filter/?firstName=John"

	recieved_profile := user_models.FilteredUsers{}

	response, err := admin_client.New().Get(endpoint_address).ReceiveSuccess(&recieved_profile)

	expected_info := models.FilteredUsers{
		[]models.UserInfo{
			user_info_1,
			user_info_2,
		},
	}

	if err != nil {
		t.Errorf("Unable to make request")
	}
	if response.StatusCode != http.StatusOK {
		t.Errorf("Request returned HTTP error %d", response.StatusCode)
	}
	if !reflect.DeepEqual(recieved_profile, expected_info) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_info, recieved_profile)
	}
}

func TestGetFilterByLastName(t *testing.T) {
	AddTestDataToDatabase()
	defer ClearDatabase()

	endpoint_address := "/user/filter/?lastName=Smith"

	recieved_profile := user_models.FilteredUsers{}

	response, err := admin_client.New().Get(endpoint_address).ReceiveSuccess(&recieved_profile)

	expected_info := models.FilteredUsers{
		[]models.UserInfo{
			user_info_1,
			user_info_2,
		},
	}

	if err != nil {
		t.Errorf("Unable to make request")
	}
	if response.StatusCode != http.StatusOK {
		t.Errorf("Request returned HTTP error %d", response.StatusCode)
	}
	if !reflect.DeepEqual(recieved_profile, expected_info) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_info, recieved_profile)
	}
}

func TestGetFilterByEmail(t *testing.T) {
	AddTestDataToDatabase()
	defer ClearDatabase()

	endpoint_address := "/user/filter/?email=James.Names@aol.com"

	recieved_profile := user_models.FilteredUsers{}

	response, err := admin_client.New().Get(endpoint_address).ReceiveSuccess(&recieved_profile)

	expected_info := models.FilteredUsers{
		[]models.UserInfo{
			user_info_3,
		},
	}

	if err != nil {
		t.Errorf("Unable to make request")
	}
	if response.StatusCode != http.StatusOK {
		t.Errorf("Request returned HTTP error %d", response.StatusCode)
	}
	if !reflect.DeepEqual(recieved_profile, expected_info) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_info, recieved_profile)
	}
}

func TestGetFilterLimit(t *testing.T) {
	AddTestDataToDatabase()
	defer ClearDatabase()

	endpoint_address := "/user/filter/?firstName=John&p=1&limit=1"

	recieved_profile := user_models.FilteredUsers{}

	response, err := admin_client.New().Get(endpoint_address).ReceiveSuccess(&recieved_profile)

	expected_info := models.FilteredUsers{
		[]models.UserInfo{
			user_info_1,
		},
	}

	if err != nil {
		t.Errorf("Unable to make request")
	}
	if response.StatusCode != http.StatusOK {
		t.Errorf("Request returned HTTP error %d", response.StatusCode)
	}
	if !reflect.DeepEqual(recieved_profile, expected_info) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_info, recieved_profile)
	}
}

func TestGetFilterOrder(t *testing.T) {
	AddTestDataToDatabase()
	defer ClearDatabase()

	endpoint_address := "/user/filter/?id=54321&sortby=firstName,lastName"

	recieved_profile := user_models.FilteredUsers{}

	response, err := admin_client.New().Get(endpoint_address).ReceiveSuccess(&recieved_profile)

	expected_info := models.FilteredUsers{
		[]models.UserInfo{
			user_info_4,
			user_info_3,
		},
	}

	if err != nil {
		t.Errorf("Unable to make request")
	}
	if response.StatusCode != http.StatusOK {
		t.Errorf("Request returned HTTP error %d", response.StatusCode)
	}
	if !reflect.DeepEqual(recieved_profile, expected_info) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_info, recieved_profile)
	}
}
