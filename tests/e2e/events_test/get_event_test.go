package tests

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/HackIllinois/api/services/event/models"
)

func TestGetEventNormal(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	event_id := "testeventid12345"
	received_event := models.Event{}
	response, err := public_client.New().Get(fmt.Sprintf("/event/%s/", event_id)).ReceiveSuccess(&received_event)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_event := models.Event{
		ID:          "testeventid12345",
		Name:        "testevent1",
		Description: "testdescription1",
		StartTime:   current_unix_time,
		EndTime:     current_unix_time + 60000,
		Sponsor:     "testsponsor1",
		EventType:   "WORKSHOP",
		Locations: []models.EventLocation{
			{
				Description: "testlocationdescription1",
				Tags:        []string{"SIEBEL3", "ECEB2"},
				Latitude:    123.456,
				Longitude:   123.456,
			},
		},
		Points: 50,
	}

	if !reflect.DeepEqual(received_event, expected_event) {
		t.Fatalf("Wrong event info. Expected %v, got %v", expected_event, received_event)
	}
}

func TestGetEventNotExist(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	event_id := "nonsense_eventid"
	received_event := models.Event{}
	response, err := public_client.New().Get(fmt.Sprintf("/event/%s/", event_id)).ReceiveSuccess(&received_event)

	if err != nil {
		t.Error("Unable to make request")
	}
	if response.StatusCode != http.StatusInternalServerError {
		t.Errorf("Request returned HTTP error %d", response.StatusCode)
	}
}
