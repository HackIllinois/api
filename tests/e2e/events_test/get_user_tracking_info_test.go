package tests

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/HackIllinois/api/services/event/models"
)

func TestGetUserTrackingInfoNormal(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	recieved_tracked_events := models.UserTracker{}
	userid := TEST_USER_ID
	response, err := staff_client.New().Get(fmt.Sprintf("/event/track/user/%s/", userid)).ReceiveSuccess(&recieved_tracked_events)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_user_tracking := models.UserTracker{
		UserID: userid,
		Events: []string{
			TEST_EVENT_2_ID,
		},
	}

	if !reflect.DeepEqual(recieved_tracked_events, expected_user_tracking) {
		t.Fatalf("Wrong result received. Expected %v, got %v", expected_user_tracking, recieved_tracked_events)
	}
}

func TestGetUserTrackingInfoNone(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	recieved_tracked_events := models.UserTracker{}
	userid := "anotheruser"
	response, err := staff_client.New().Get(fmt.Sprintf("/event/track/user/%s/", userid)).ReceiveSuccess(&recieved_tracked_events)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_user_tracking := models.UserTracker{
		UserID: userid,
		Events: []string{},
	}

	if !reflect.DeepEqual(recieved_tracked_events, expected_user_tracking) {
		t.Fatalf("Wrong result received. Expected %v, got %v", expected_user_tracking, recieved_tracked_events)
	}
}
