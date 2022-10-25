package tests

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	profile_models "github.com/HackIllinois/api/services/profile/models"
)

func GetProfileTest(t *testing.T) {
	profile_info := profile_models.Profile{
		ID:        "12345",
		FirstName: "John",
		LastName:  "Smith",
		Points:    5,
		Timezone:  "CST",
		Discord:   "discord",
		AvatarUrl: "url",
	}
	client.Database(profile_db_name).Collection("profile").InsertOne(context.Background(), profile_info)

	endpoint_address := fmt.Sprintf("/profile/%s/", "12345")

	recieved_profile := profile_models.Profile{}
	response, err := admin_client.New().Get(endpoint_address).ReceiveSuccess(&recieved_profile)

	if err != nil {
		t.Errorf("Unable to make request")
	}
	if response.StatusCode != 200 {
		t.Errorf("Request returned HTTP error %d", response.StatusCode)
	}
	if !reflect.DeepEqual(recieved_profile, profile_info) {
		t.Errorf("Wrong event info. Expected %v, got %v", profile_info, recieved_profile)
	}
}
