package tests

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/HackIllinois/api/services/event/models"
	"go.mongodb.org/mongo-driver/bson"
)

func TestDeleteNormal(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	event_id := "testeventid12345"
	received_event := models.Event{}
	response, err := staff_client.New().Delete(fmt.Sprintf("/event/%s/", event_id)).ReceiveSuccess(&received_event)

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

	cursor, _ := client.Database(events_db_name).Collection("events").Find(context.Background(), bson.D{})
	res := []models.Event{}
	err = cursor.All(context.TODO(), &res)

	if err != nil {
		t.Fatalf("Test failed due to unexpected error: %v", err)
		return
	}

	expected_events := []models.Event{
		{
			ID:          "testeventid67890",
			Name:        "testevent2",
			Description: "testdescription2",
			StartTime:   current_unix_time,
			EndTime:     current_unix_time + 60000,
			Sponsor:     "testsponsor2",
			EventType:   "WORKSHOP",
			Locations: []models.EventLocation{
				{
					Description: "testlocationdescription2",
					Tags:        []string{"SIEBEL3", "ECEB2"},
					Latitude:    123.456,
					Longitude:   123.456,
				},
			},
			Points: 100,
		},
	}

	if !reflect.DeepEqual(res, expected_events) {
		t.Fatalf("Database contained wrong event info. Expected %v, got %v", expected_events, res)
	}
}

func TestDeleteNotExist(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	event_id := "nonsense_eventid"
	received_event := models.Event{}
	response, err := staff_client.New().Delete(fmt.Sprintf("/event/%s/", event_id)).ReceiveSuccess(&received_event)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusInternalServerError {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	cursor, _ := client.Database(events_db_name).Collection("events").Find(context.Background(), bson.D{})
	res := []models.Event{}
	err = cursor.All(context.TODO(), &res)

	if err != nil {
		t.Fatalf("Test failed due to unexpected error: %v", err)
		return
	}

	expected_events := []models.Event{
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
		{
			ID:          "testeventid67890",
			Name:        "testevent2",
			Description: "testdescription2",
			StartTime:   current_unix_time,
			EndTime:     current_unix_time + 60000,
			Sponsor:     "testsponsor2",
			EventType:   "WORKSHOP",
			Locations: []models.EventLocation{
				{
					Description: "testlocationdescription2",
					Tags:        []string{"SIEBEL3", "ECEB2"},
					Latitude:    123.456,
					Longitude:   123.456,
				},
			},
			Points: 100,
		},
	}

	if !reflect.DeepEqual(res, expected_events) {
		t.Fatalf("Database contained wrong event info. Expected %v, got %v", expected_events, res)
	}
}

func TestDeleteForbidden(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	event_id := "testeventid12345"
	received_event := models.Event{}
	response, err := user_client.New().Delete(fmt.Sprintf("/event/%s/", event_id)).ReceiveSuccess(&received_event)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusForbidden {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	cursor, _ := client.Database(events_db_name).Collection("events").Find(context.Background(), bson.D{})
	res := []models.Event{}
	err = cursor.All(context.TODO(), &res)

	if err != nil {
		t.Fatalf("Test failed due to unexpected error: %v", err)
		return
	}

	expected_events := []models.Event{
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
		{
			ID:          "testeventid67890",
			Name:        "testevent2",
			Description: "testdescription2",
			StartTime:   current_unix_time,
			EndTime:     current_unix_time + 60000,
			Sponsor:     "testsponsor2",
			EventType:   "WORKSHOP",
			Locations: []models.EventLocation{
				{
					Description: "testlocationdescription2",
					Tags:        []string{"SIEBEL3", "ECEB2"},
					Latitude:    123.456,
					Longitude:   123.456,
				},
			},
			Points: 100,
		},
	}

	if !reflect.DeepEqual(res, expected_events) {
		t.Fatalf("Database contained wrong event info. Expected %v, got %v", expected_events, res)
	}
}
