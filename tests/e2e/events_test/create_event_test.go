package tests

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/HackIllinois/api/services/event/models"
	"go.mongodb.org/mongo-driver/bson"
)

func TestCreateEventNormal(t *testing.T) {
	defer ClearEvents()
	event_info := models.Event{
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
	received_event := models.Event{}
	response, err := staff_client.New().Post("/event/").BodyJSON(event_info).ReceiveSuccess(&received_event)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	event_info.ID = received_event.ID

	if !reflect.DeepEqual(received_event, event_info) {
		t.Fatalf("Wrong event info. Expected %v, got %v", event_info, received_event)
	}

	cursor, _ := client.Database(events_db_name).Collection("events").Find(context.Background(), bson.D{})
	res := []models.Event{}
	err = cursor.All(context.TODO(), &res)

	if err != nil {
		t.Fatalf("Test failed due to unexpected error: %v", err)
		return
	}

	expected_res := []models.Event{
		event_info,
	}

	if !reflect.DeepEqual(res, expected_res) {
		t.Fatalf("Database contained wrong event info. Expected %v, got %v", expected_res, res)
	}
}

func TestCreateEventForbidden(t *testing.T) {
	defer ClearEvents()
	event_info := models.Event{
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

	response, err := user_client.New().Post("/event/").BodyJSON(event_info).Receive(nil, nil)

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

	expected_res := []models.Event{}

	if !reflect.DeepEqual(res, expected_res) {
		t.Fatalf("Database contained wrong event info. Expected %v, got %v", expected_res, res)
	}
}
