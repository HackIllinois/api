package tests

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/services/event/models"
	"go.mongodb.org/mongo-driver/bson"
)

func TestUpdateEventNormal(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	res := client.Database(events_db_name).Collection("events").FindOne(context.Background(), bson.M{"id": TEST_EVENT_1_ID})

	event := models.EventDB{}

	res.Decode(&event)

	event.Description = "It's a new description!"
	var received_event models.EventDB

	response, err := staff_client.New().Put("/event/").BodyJSON(event).ReceiveSuccess(&received_event)
	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	if !reflect.DeepEqual(received_event, event) {
		t.Fatalf("Wrong event info. Expected %v, got %v", event, received_event)
	}
}

func TestUpdateEventForbidden(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	res := client.Database(events_db_name).Collection("events").FindOne(context.Background(), bson.M{"id": TEST_EVENT_1_ID})

	event := models.EventDB{}

	res.Decode(&event)

	event.Description = "It's a new description!"
	event.IsPrivate = false
	var received_event models.EventDB

	response, err := user_client.New().Put("/event/").BodyJSON(event).ReceiveSuccess(&received_event)
	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusForbidden {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}
}

func TestUpdateEventNotFound(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	res := client.Database(events_db_name).Collection("events").FindOne(context.Background(), bson.M{"id": TEST_EVENT_1_ID})

	event := models.EventDB{}

	res.Decode(&event)

	event.ID = "adifferentidthatsnotbeenadded"
	event.Description = "It's a new description!"

	api_err := errors.ApiError{}
	response, err := staff_client.New().Put("/event/").BodyJSON(event).Receive(nil, &api_err)
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
		Message:  "Could not update the event.",
		RawError: "Error: NOT_FOUND",
	}

	if !reflect.DeepEqual(expected_error, api_err) {
		t.Fatalf("Wrong error response received. Expected %v, got %v", expected_error, api_err)
	}
}
