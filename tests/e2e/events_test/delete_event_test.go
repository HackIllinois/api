package tests

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/services/event/models"
	"go.mongodb.org/mongo-driver/bson"
)

func TestDeleteEventNormal(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	event_id := TEST_EVENT_1_ID
	received_event := models.EventDB{}
	response, err := staff_client.New().Delete(fmt.Sprintf("/event/%s/", event_id)).ReceiveSuccess(&received_event)
	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_event := models.EventDB{
		EventPublic: models.EventPublic{
			ID:          TEST_EVENT_1_ID,
			Name:        "testevent1",
			Description: "testdescription1",
			StartTime:   current_unix_time,
			EndTime:     current_unix_time + 60000,
			Sponsor:     "testsponsor1",
			EventType:   "OTHER",
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
		IsPrivate:             true,
		DisplayOnStaffCheckin: false,
	}

	if !reflect.DeepEqual(received_event, expected_event) {
		t.Fatalf("Wrong event info. Expected %v, got %v", expected_event, received_event)
	}

	cursor, _ := client.Database(events_db_name).Collection("events").Find(context.Background(), bson.D{})
	res := []models.EventDB{}
	err = cursor.All(context.TODO(), &res)

	if err != nil {
		t.Fatalf("Test failed due to unexpected error: %v", err)
		return
	}

	expected_events := []models.EventDB{
		{
			EventPublic: models.EventPublic{
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
			},
			IsPrivate:             false,
			DisplayOnStaffCheckin: true,
		},
	}

	if !reflect.DeepEqual(res, expected_events) {
		t.Fatalf("Database contained wrong event info. Expected %v, got %v", expected_events, res)
	}
}

func TestDeleteEventNotExist(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	event_id := "nonsense_eventid"
	api_err := errors.ApiError{}
	response, err := staff_client.New().Delete(fmt.Sprintf("/event/%s/", event_id)).Receive(nil, &api_err)
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
		Type:     "INTERNAL_ERROR",
		Message:  "Could not delete either the event, event trackers, or user trackers, or an intermediary subroutine failed.",
		RawError: "Error: NOT_FOUND",
	}

	if !reflect.DeepEqual(api_err, expected_error) {
		t.Fatalf("Wrong error response received. Expected %v, got %v", expected_error, api_err)
	}

	cursor, _ := client.Database(events_db_name).Collection("events").Find(context.Background(), bson.D{})
	res := []models.EventDB{}
	err = cursor.All(context.TODO(), &res)

	if err != nil {
		t.Fatalf("Test failed due to unexpected error: %v", err)
		return
	}

	expected_events := []models.EventDB{
		{
			EventPublic: models.EventPublic{
				ID:          TEST_EVENT_1_ID,
				Name:        "testevent1",
				Description: "testdescription1",
				StartTime:   current_unix_time,
				EndTime:     current_unix_time + 60000,
				Sponsor:     "testsponsor1",
				EventType:   "OTHER",
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
			IsPrivate:             true,
			DisplayOnStaffCheckin: false,
		},
		{
			EventPublic: models.EventPublic{
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
			},
			IsPrivate:             false,
			DisplayOnStaffCheckin: true,
		},
	}

	if !reflect.DeepEqual(res, expected_events) {
		t.Fatalf("Database contained wrong event info. Expected %v, got %v", expected_events, res)
	}
}

func TestDeleteEventForbidden(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	event_id := TEST_EVENT_1_ID
	response, err := user_client.New().Delete(fmt.Sprintf("/event/%s/", event_id)).Receive(nil, nil)
	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusForbidden {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	cursor, _ := client.Database(events_db_name).Collection("events").Find(context.Background(), bson.D{})
	res := []models.EventDB{}
	err = cursor.All(context.TODO(), &res)

	if err != nil {
		t.Fatalf("Test failed due to unexpected error: %v", err)
		return
	}

	expected_events := []models.EventDB{
		{
			EventPublic: models.EventPublic{
				ID:          TEST_EVENT_1_ID,
				Name:        "testevent1",
				Description: "testdescription1",
				StartTime:   current_unix_time,
				EndTime:     current_unix_time + 60000,
				Sponsor:     "testsponsor1",
				EventType:   "OTHER",
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
			IsPrivate:             true,
			DisplayOnStaffCheckin: false,
		},
		{
			EventPublic: models.EventPublic{
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
			},
			IsPrivate:             false,
			DisplayOnStaffCheckin: true,
		},
	}

	if !reflect.DeepEqual(res, expected_events) {
		t.Fatalf("Database contained wrong event info. Expected %v, got %v", expected_events, res)
	}
}
