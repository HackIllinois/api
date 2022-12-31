package tests

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/HackIllinois/api/services/event/models"
	"go.mongodb.org/mongo-driver/bson"
)

func TestAddFavoriteEventNormal(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	req := models.EventFavoriteModification{
		EventID: "testeventid12345",
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
		ID: "localadmin",
		Events: []string{
			"testeventid12345",
		},
	}

	if !reflect.DeepEqual(received_res, expected_res) {
		t.Fatalf("Wrong result received. Expected %v, got %v", expected_res, received_res)
		return
	}

	res := client.Database(events_db_name).Collection("favorites").FindOne(context.Background(), bson.M{"id": "localadmin"})

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

	response, err := user_client.New().Post("/event/favorite/").BodyJSON(req).ReceiveSuccess(nil)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusInternalServerError {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	cursor, err := client.Database(events_db_name).Collection("favorites").Find(context.Background(), bson.M{"id": "localadmin"})

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
