package tests

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/HackIllinois/api/services/event/models"
)

func TestGetAllPublicEventsNormal(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	received_events := models.EventList[models.EventPublic]{}
	response, err := public_client.New().Get("/event/").ReceiveSuccess(&received_events)
	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_events := models.EventList[models.EventPublic]{
		Events: []models.EventPublic{
			{
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
		},
	}

	if !reflect.DeepEqual(received_events, expected_events) {
		t.Fatalf("Wrong event info. Expected %v, got %v", expected_events, received_events)
	}
}

func TestGetAllEventsNormal(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	received_events := models.EventList[models.EventDB]{}
	response, err := staff_client.New().Get("/event/").ReceiveSuccess(&received_events)
	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_events := models.EventList[models.EventDB]{
		Events: []models.EventDB{
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
		},
	}

	if !reflect.DeepEqual(received_events, expected_events) {
		t.Fatalf("Wrong event info. Expected %v, got %v", expected_events, received_events)
	}
}

func TestGetAllEventsNone(t *testing.T) {
	received_events := models.EventList[models.EventDB]{}
	response, err := staff_client.New().Get("/event/").ReceiveSuccess(&received_events)
	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_events := models.EventList[models.EventDB]{
		Events: []models.EventDB{},
	}

	if !reflect.DeepEqual(received_events, expected_events) {
		t.Fatalf("Wrong event info. Expected %v, got %v", expected_events, received_events)
	}
}
