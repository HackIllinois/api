package tests

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/HackIllinois/api/services/event/models"
	"go.mongodb.org/mongo-driver/bson"
)

func TestDeleteFavoriteEventsNormal(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	favorite_events := models.EventFavorites{
		ID: "localadmin",
		Events: []string{
			"testeventid12345",
			"testeventid67890",
		},
	}

	client.Database(events_db_name).Collection("favorites").InsertOne(context.Background(), favorite_events)

	req := models.EventFavoriteModification{
		EventID: "testeventid67890",
	}
	recieved_favorites := models.EventFavorites{}
	response, err := user_client.New().Delete("/event/favorite/").BodyJSON(req).ReceiveSuccess(&recieved_favorites)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_favorites := models.EventFavorites{
		ID: "localadmin",
		Events: []string{
			"testeventid12345",
		},
	}

	if !reflect.DeepEqual(recieved_favorites, expected_favorites) {
		t.Fatalf("Wrong result received. Expected %v, got %v", expected_favorites, recieved_favorites)
	}

	res := client.Database(events_db_name).Collection("favorites").FindOne(context.Background(), bson.M{"id": "localadmin"})

	err = res.Decode(&recieved_favorites)

	if err != nil {
		t.Fatalf("Had trouble finding favorites in database: %v", err)
		return
	}

	if !reflect.DeepEqual(recieved_favorites, expected_favorites) {
		t.Fatalf("Wrong result received from database. Expected %v, got %v", expected_favorites, recieved_favorites)
	}
}

func TestDeleteFavoriteEventsNone(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	favorite_events := models.EventFavorites{
		ID: "localadmin",
		Events: []string{
			"testeventid12345",
			"testeventid67890",
		},
	}

	client.Database(events_db_name).Collection("favorites").InsertOne(context.Background(), favorite_events)

	req := models.EventFavoriteModification{
		EventID: "nonexistantevent",
	}
	recieved_favorites := models.EventFavorites{}
	response, err := user_client.New().Delete("/event/favorite/").BodyJSON(req).ReceiveSuccess(&recieved_favorites)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusInternalServerError {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	res := client.Database(events_db_name).Collection("favorites").FindOne(context.Background(), bson.M{"id": "localadmin"})

	err = res.Decode(&recieved_favorites)

	if err != nil {
		t.Fatalf("Had trouble finding favorites in database: %v", err)
		return
	}

	if !reflect.DeepEqual(recieved_favorites, favorite_events) {
		t.Fatalf("Wrong result received from database. Expected %v, got %v", favorite_events, recieved_favorites)
	}
}
