package tests

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/HackIllinois/api/services/event/models"
)

func TestGetEventTrackingInfoNone(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	recieved_tracking_users := models.EventTracker{}
	eventid := "testeventid12345"
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

	recieved_tracking_users := models.EventTracker{}
	eventid := "invalideventid"
	response, err := staff_client.New().Get(fmt.Sprintf("/event/track/event/%s/", eventid)).ReceiveSuccess(&recieved_tracking_users)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusNotFound {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}
}

func TestGetEventTrackingInfoNormal(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	recieved_tracking_users := models.EventTracker{}
	eventid := "testeventid67889"
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
			"localadmin",
		},
	}

	if !reflect.DeepEqual(recieved_tracking_users, expected_event_tracking) {
		t.Fatalf("Wrong result received. Expected %v, got %v", expected_event_tracking, recieved_tracking_users)
	}
}
