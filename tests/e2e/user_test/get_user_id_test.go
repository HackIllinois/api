package tests

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	user_models "github.com/HackIllinois/api/services/user/models"
)

func TestUserProfile(t *testing.T) {
	user_info := user_models.UserInfo{
		ID:        "12345",
		Username:  "johnsmith",
		FirstName: "John",
		LastName:  "Smith",
		Email:     "john.smith@aol.com",
	}

	client.Database(user_db_name).Collection("info").InsertOne(context.Background(), user_info)

	endpoint_address := fmt.Sprintf("/user/%s/", "12345")

	recieved_user := user_models.UserInfo{}
	response, err := admin_client.New().Get(endpoint_address).ReceiveSuccess(&recieved_user)

	if err != nil {
		t.Errorf("Unable to make request")
	}
	if response.StatusCode != http.StatusOK {
		t.Errorf("Request returned HTTP error %d", response.StatusCode)
	}
	if !reflect.DeepEqual(recieved_user, user_info) {
		t.Errorf("Wrong user info. Expected %v, got %v", user_info, recieved_user)
	}
}
