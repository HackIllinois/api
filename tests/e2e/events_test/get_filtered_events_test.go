package tests

import (
	"net/http"
	"reflect"
	"testing"

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

	response, err := public_client.New().Get("/event/filter/?nonsensefield=trydecipheringthis!").ReceiveSuccess(nil)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusUnprocessableEntity {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}
}
