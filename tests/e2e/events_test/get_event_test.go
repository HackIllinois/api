package tests

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/services/event/models"
)

func TestGetEventNormal(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	event_id := TEST_EVENT_2_ID
	received_event := models.EventPublic{}
	response, err := public_client.New().Get(fmt.Sprintf("/event/%s/", event_id)).ReceiveSuccess(&received_event)
	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_event := models.EventPublic{
		ID:          TEST_EVENT_2_ID,
		Name:        "testevent2",
		Description: "testdescription2",
		StartTime:   current_unix_time + 60000,
		EndTime:     current_unix_time + 120000,
		Sponsor:     "",
		EventType:   "FOOD",
		Locations: []models.EventLocation{
			{
				Description: "testlocationdescription2",
				Tags:        []string{"SIEBEL3", "ECEB2"},
				Latitude:    123.456,
				Longitude:   123.456,
			},
		},
		Points: 0,
	}

	if !reflect.DeepEqual(received_event, expected_event) {
		t.Fatalf("Wrong event info. Expected %v, got %v", expected_event, received_event)
	}
}

func TestGetEventNotExist(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	event_id := "nonsense_eventid"
	api_err := errors.ApiError{}
	response, err := public_client.New().Get(fmt.Sprintf("/event/%s/", event_id)).Receive(nil, &api_err)
	if err != nil {
		t.Error("Unable to make request")
	}
	if response.StatusCode != http.StatusInternalServerError {
		t.Errorf("Request returned HTTP error %d", response.StatusCode)
	}

	expected_error := errors.ApiError{
		Status:   http.StatusInternalServerError,
		Type:     "DATABASE_ERROR",
		Message:  "Could not fetch the event details.",
		RawError: "Error: NOT_FOUND",
	}

	if !reflect.DeepEqual(expected_error, api_err) {
		t.Fatalf("Wrong error response received. Expected %v, got %v", expected_error, api_err)
	}
}
