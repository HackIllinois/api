package tests

import (
	"net/http"
	"testing"

	profile_models "github.com/HackIllinois/api/services/profile/models"
)

func TestPostProfile(t *testing.T) {
	profile_info := profile_models.Profile{
		ID:        "12345",
		FirstName: "John",
		LastName:  "Smith",
		Points:    5,
		Timezone:  "CST",
		Discord:   "discord",
		AvatarUrl: "url",
	}

	recieved_profile := profile_models.Profile{}
	response, err := admin_client.New().Post("/profile/").BodyJSON(profile_info).ReceiveSuccess(&recieved_profile)

	if err != nil {
		t.Errorf("Unable to make request")
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("Request returned HTTP error %d", response.StatusCode)
	}
}

func TestUnauthenticatedPostProfile(t *testing.T) {
	profile_info := profile_models.Profile{
		ID:        "12345",
		FirstName: "John",
		LastName:  "Smith",
		Points:    5,
		Timezone:  "CST",
		Discord:   "discord",
		AvatarUrl: "url",
	}

	recieved_profile := profile_models.Profile{}
	response, _ := unauthenticated_client.New().Post("/profile/").BodyJSON(profile_info).ReceiveSuccess(&recieved_profile)

	if response.StatusCode != http.StatusForbidden {
		t.Errorf("Unauthenticated attendee able to access endpoint that requires authentication")
	}
}

func TestBadPostProfile(t *testing.T) {
	profile_info := profile_models.Profile{
		ID:        "12345",
		FirstName: "John",
	}

	recieved_profile := profile_models.Profile{}
	response, _ := admin_client.New().Post("/profile/").BodyJSON(profile_info).ReceiveSuccess(&recieved_profile)

	if response.StatusCode != http.StatusInternalServerError {
		t.Errorf("Profile with not enough fields can be inserted")
	}
}
