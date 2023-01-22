package tests

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/services/event/models"
)

func TestGetEventTrackingInfoNone(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	recieved_tracking_users := models.EventTracker{}
	eventid := TEST_EVENT_1_ID
	response, err := staff_client.New().Get(fmt.Sprintf("/event/track/event/%s/", eventid)).ReceiveSuccess(&recieved_tracking_users)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_event_tracking := models.EventTracker{
		EventID: eventid,
		Users:   []string{},
	}

	if !reflect.DeepEqual(recieved_tracking_users, expected_event_tracking) {
		t.Fatalf("Wrong result received. Expected %v, got %v", expected_event_tracking, recieved_tracking_users)
	}
}

func TestGetEventTrackingInfoNonexist(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	eventid := "invalideventid"
	api_err := errors.ApiError{}
	response, err := staff_client.New().Get(fmt.Sprintf("/event/track/event/%s/", eventid)).Receive(nil, &api_err)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusInternalServerError {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_error := errors.ApiError{
		Status:   500,
		Type:     "DATABASE_ERROR",
		Message:  "Could not get event tracker.",
		RawError: "Error: NOT_FOUND",
	}

	if !reflect.DeepEqual(expected_error, api_err) {
		t.Fatalf("Wrong error resonse recieved. Expected %v, got %v", expected_error, api_err)
	}
}

func TestGetEventTrackingInfoNormal(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	recieved_tracking_users := models.EventTracker{}
	eventid := TEST_EVENT_2_ID
	response, err := staff_client.New().Get(fmt.Sprintf("/event/track/event/%s/", eventid)).ReceiveSuccess(&recieved_tracking_users)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_event_tracking := models.EventTracker{
		EventID: eventid,
		Users: []string{
			TEST_USER_ID,
		},
	}

	if !reflect.DeepEqual(recieved_tracking_users, expected_event_tracking) {
		t.Fatalf("Wrong result received. Expected %v, got %v", expected_event_tracking, recieved_tracking_users)
	}
}
