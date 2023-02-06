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

func TestAddFavoriteEventNormal(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	req := models.EventFavoriteModification{
		EventID: TEST_EVENT_2_ID,
	}
	received_res := models.EventFavorites{}
	response, err := user_client.New().Post("/event/favorite/").BodyJSON(req).ReceiveSuccess(&received_res)
	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_res := models.EventFavorites{
		ID: TEST_USER_ID,
		Events: []string{
			TEST_EVENT_2_ID,
		},
	}

	if !reflect.DeepEqual(received_res, expected_res) {
		t.Fatalf("Wrong result received. Expected %v, got %v", expected_res, received_res)
		return
	}

	res := client.Database(events_db_name).Collection("favorites").FindOne(context.Background(), bson.M{"id": TEST_USER_ID})

	err = res.Decode(&received_res)

	if err != nil {
		t.Fatalf("Had trouble finding favorites in database: %v", err)
		return
	}

	if !reflect.DeepEqual(received_res, expected_res) {
		t.Fatalf("Wrong result received from database. Expected %v, got %v", expected_res, received_res)
	}
}

func TestAddFavoriteEventNotExist(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	req := models.EventFavoriteModification{
		EventID: "nonexistantid",
	}

	api_err := errors.ApiError{}
	response, err := user_client.New().Post("/event/favorite/").BodyJSON(req).Receive(nil, &api_err)
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
		Message:  "Could not add an event favorite for the current user.",
		RawError: "Could not find event with the given id.",
	}

	if !reflect.DeepEqual(api_err, expected_error) {
		t.Fatalf("Wrong reponse received. Expected %v, got %v", expected_error, api_err)
	}

	cursor, err := client.Database(events_db_name).
		Collection("favorites").
		Find(context.Background(), bson.M{"id": TEST_USER_ID})

	res := []models.EventFavorites{}
	err = cursor.All(context.TODO(), &res)

	if err != nil {
		t.Fatalf("Test failed due to unexpected error: %v", err)
		return
	}

	expected_res := []models.EventFavorites{}

	if !reflect.DeepEqual(res, expected_res) {
		t.Fatalf("Wrong result received from database. Expected %v, got %v", expected_res, res)
	}
}
