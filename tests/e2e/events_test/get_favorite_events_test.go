package tests

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/HackIllinois/api/services/event/models"
)

func TestGetFavoriteEventsNormal(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	favorite_events := models.EventFavorites{
		ID: TEST_USER_ID,
		Events: []string{
			TEST_EVENT_1_ID,
			TEST_EVENT_2_ID,
		},
	}

	client.Database(events_db_name).Collection("favorites").InsertOne(context.Background(), favorite_events)

	recieved_favorites := models.EventFavorites{}
	response, err := user_client.New().Get("/event/favorite/").ReceiveSuccess(&recieved_favorites)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	if !reflect.DeepEqual(recieved_favorites, favorite_events) {
		t.Fatalf("Wrong result received from database. Expected %v, got %v", favorite_events, recieved_favorites)
	}
}

func TestGetFavoriteEventsNone(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	favorite_events := models.EventFavorites{
		ID:     TEST_USER_ID,
		Events: []string{},
	}

	recieved_favorites := models.EventFavorites{}
	response, err := user_client.New().Get("/event/favorite/").ReceiveSuccess(&recieved_favorites)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	if !reflect.DeepEqual(recieved_favorites, favorite_events) {
		t.Fatalf("Wrong result received from database. Expected %v, got %v", favorite_events, recieved_favorites)
	}
}
