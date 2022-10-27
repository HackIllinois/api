package tests

import (
	"testing"

	profile_models "github.com/HackIllinois/api/services/profile/models"
)

func PostProfileTest(t *testing.T) {
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

	if response.StatusCode != 200 {
		t.Errorf("Request returned HTTP error %d", response.StatusCode)
	}
}
