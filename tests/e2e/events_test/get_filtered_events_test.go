package tests

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/services/event/models"
)

func TestGetFilteredEventsNormal(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	received_events := models.EventList{}
	response, err := public_client.New().Get("/event/filter/?name=testevent1").ReceiveSuccess(&received_events)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_events := models.EventList{
		Events: []models.Event{
			{
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
			},
		},
	}

	if !reflect.DeepEqual(received_events, expected_events) {
		t.Fatalf("Wrong event info. Expected %v, got %v", expected_events, received_events)
	}
}

func TestGetFilteredEventsBadArgs(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	api_err := errors.ApiError{}
	response, err := public_client.New().Get("/event/filter/?nonsensefield=trydecipheringthis!").Receive(nil, &api_err)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusInternalServerError {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_error := errors.ApiError{
		Status:   http.StatusInternalServerError,
		Type:     "DATABASE_ERROR",
		Message:  "Could not fetch filtered list of events.",
		RawError: "Invalid key nonsensefield",
	}

	if !reflect.DeepEqual(expected_error, api_err) {
		t.Fatalf("Wrong error resonse received. Expected %v, got %v", expected_error, api_err)
	}
}
