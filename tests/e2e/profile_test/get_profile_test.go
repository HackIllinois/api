package tests

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	profile_models "github.com/HackIllinois/api/services/profile/models"
)

func TestGetProfile(t *testing.T) {
	profile_info := profile_models.Profile{
		ID:        "12345",
		FirstName: "John",
		LastName:  "Smith",
		Points:    5,
		Discord:   "discord",
		AvatarUrl: "url",
		FoodWave:  2,
	}
	client.Database(profile_db_name).Collection("profiles").InsertOne(context.Background(), profile_info)

	endpoint_address := fmt.Sprintf("/profile/%s/", "12345")

	received_profile := profile_models.Profile{}
	response, err := admin_client.New().Get(endpoint_address).ReceiveSuccess(&received_profile)
	if err != nil {
		t.Fatalf("Unable to make request")
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
	}
	if !reflect.DeepEqual(received_profile, profile_info) {
		t.Errorf("Wrong event info. Expected %v, got %v", profile_info, received_profile)
	}
}

func TestUnauthenticatedGetProfile(t *testing.T) {
	profile_info := profile_models.Profile{
		ID:        "12345",
		FirstName: "John",
		LastName:  "Smith",
		Points:    5,
		Discord:   "discord",
		AvatarUrl: "url",
		FoodWave:  1,
	}
	client.Database(profile_db_name).Collection("profiles").InsertOne(context.Background(), profile_info)

	endpoint_address := fmt.Sprintf("/profile/%s/", "12345")

	received_profile := profile_models.Profile{}
	response, err := unauthenticated_client.New().Get(endpoint_address).ReceiveSuccess(&received_profile)
	if err != nil {
		t.Fatalf("Unable to make request")
	}

	if response.StatusCode != http.StatusForbidden {
		t.Errorf("Unauthenticated attendee able to access endpoint that requires authentication")
	}
}

func TestNonExistantGetProfile(t *testing.T) {
	endpoint_address := fmt.Sprintf("/profile/%s/", "00000")

	recieved_profile := profile_models.Profile{}
	response, err := admin_client.New().Get(endpoint_address).ReceiveSuccess(&recieved_profile)
	if err != nil {
		t.Fatalf("Unable to make request")
	}

	if response.StatusCode != http.StatusInternalServerError { // change to http.StatusNotFound once we standardize response codes
		t.Errorf("Attendee able to access non-existent profile")
	}
}
