package tests

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/HackIllinois/api/services/event/models"
	"go.mongodb.org/mongo-driver/bson"
)

func TestUpdateEventNormal(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	res := client.Database(events_db_name).Collection("events").FindOne(context.Background(), bson.M{"id": "testeventid12345"})

	event := models.Event{}

	res.Decode(&event)

	event.Description = "It's a new description!"
	var received_event models.Event

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

	res := client.Database(events_db_name).Collection("events").FindOne(context.Background(), bson.M{"id": "testeventid12345"})

	event := models.Event{}

	res.Decode(&event)

	event.Description = "It's a new description!"
	var received_event models.Event

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

	res := client.Database(events_db_name).Collection("events").FindOne(context.Background(), bson.M{"id": "testeventid12345"})

	event := models.Event{}

	res.Decode(&event)

	event.ID = "adifferentidthatsnotbeenadded"
	event.Description = "It's a new description!"
	var received_event models.Event

	response, err := staff_client.New().Put("/event/").BodyJSON(event).ReceiveSuccess(&received_event)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusNotFound {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}
}
