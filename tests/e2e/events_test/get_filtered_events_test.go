package tests

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/services/event/models"
)

func AddAdditionalEvents() {
	event3 := models.Event{
		ID:          "testeventid1337",
		Name:        "eventhax0r",
		Description: "A hax0r event",
		StartTime:   current_unix_time + 120000,
		EndTime:     current_unix_time + 150000,
		Sponsor:     "testsponsor1",
		EventType:   "MINIEVENT",
		Locations: []models.EventLocation{
			{
				Description: "testlocationdescription3",
				Tags:        []string{"SIEBEL3", "ECEB2"},
				Latitude:    123.456,
				Longitude:   123.456,
			},
		},
		Points: 1337,
	}

	event4 := models.Event{
		ID:          "testeventid4",
		Name:        "eventdinner",
		Description: "Get your dinner!",
		StartTime:   current_unix_time + 240000,
		EndTime:     current_unix_time + 300000,
		Sponsor:     "",
		EventType:   "FOOD",
		Locations: []models.EventLocation{
			{
				Description: "testlocationdescription4",
				Tags:        []string{"SIEBEL3", "ECEB2"},
				Latitude:    123.456,
				Longitude:   123.456,
			},
		},
		Points: 0,
	}

	client.Database(events_db_name).Collection("events").InsertOne(context.Background(), event3)
	client.Database(events_db_name).Collection("events").InsertOne(context.Background(), event4)
}

func TestGetFilteredEventsNormal(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	received_events := models.EventList{}
	response, err := public_client.New().Get("/event/filter/?name=testevent1").ReceiveSuccess(&received_events)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_events := models.EventList{
		Events: []models.Event{
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
		},
	}

	if !reflect.DeepEqual(received_events, expected_events) {
		t.Fatalf("Wrong event info. Expected %v, got %v", expected_events, received_events)
	}
}

func TestGetFilteredEventsNoFilter(t *testing.T) {
	CreateEvents()
	AddAdditionalEvents()
	defer ClearEvents()

	received_events := models.EventList{}
	response, err := public_client.New().Get("/event/filter/").ReceiveSuccess(&received_events)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_events := models.EventList{
		Events: []models.Event{
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
			{
				ID:          "testeventid1337",
				Name:        "eventhax0r",
				Description: "A hax0r event",
				StartTime:   current_unix_time + 120000,
				EndTime:     current_unix_time + 150000,
				Sponsor:     "testsponsor1",
				EventType:   "MINIEVENT",
				Locations: []models.EventLocation{
					{
						Description: "testlocationdescription3",
						Tags:        []string{"SIEBEL3", "ECEB2"},
						Latitude:    123.456,
						Longitude:   123.456,
					},
				},
				Points: 1337,
			},
			{
				ID:          "testeventid4",
				Name:        "eventdinner",
				Description: "Get your dinner!",
				StartTime:   current_unix_time + 240000,
				EndTime:     current_unix_time + 300000,
				Sponsor:     "",
				EventType:   "FOOD",
				Locations: []models.EventLocation{
					{
						Description: "testlocationdescription4",
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

func TestGetFilteredEventsByName(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	received_events := models.EventList{}
	response, err := public_client.New().Get("/event/filter/?name=testevent2").ReceiveSuccess(&received_events)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_events := models.EventList{
		Events: []models.Event{
			{
				ID:          "testeventid67890",
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

func TestGetFilteredEventsByDescription(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	received_events := models.EventList{}
	response, err := public_client.New().Get("/event/filter/?description=testdescription1").ReceiveSuccess(&received_events)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_events := models.EventList{
		Events: []models.Event{
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
		},
	}

	if !reflect.DeepEqual(received_events, expected_events) {
		t.Fatalf("Wrong event info. Expected %v, got %v", expected_events, received_events)
	}
}

func TestGetFilteredEventsByStartTime(t *testing.T) {
	CreateEvents()
	AddAdditionalEvents()
	defer ClearEvents()

	// This argument only does exact unix time which is kinda dumb
	// TODO: Change the filters to treat StartTime and EndTime as a range (if not provided, do not
	// include that bound)

	received_events := models.EventList{}
	response, err := public_client.New().Get(fmt.Sprintf("/event/filter/?startTime=%d", current_unix_time)).ReceiveSuccess(&received_events)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_events := models.EventList{
		Events: []models.Event{
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
		},
	}

	if !reflect.DeepEqual(received_events, expected_events) {
		t.Fatalf("Wrong event info. Expected %v, got %v", expected_events, received_events)
	}
}

func TestGetFilteredEventsByEndTime(t *testing.T) {
	CreateEvents()
	AddAdditionalEvents()
	defer ClearEvents()

	received_events := models.EventList{}
	response, err := public_client.New().Get(fmt.Sprintf("/event/filter/?endTime=%d", current_unix_time+60000)).ReceiveSuccess(&received_events)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_events := models.EventList{
		Events: []models.Event{
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
		},
	}

	if !reflect.DeepEqual(received_events, expected_events) {
		t.Fatalf("Wrong event info. Expected %v, got %v", expected_events, received_events)
	}
}

func TestGetFilteredEventBySponsor(t *testing.T) {
	CreateEvents()
	AddAdditionalEvents()
	defer ClearEvents()

	received_events := models.EventList{}
	response, err := public_client.New().Get("/event/filter/?sponsor=testsponsor1").ReceiveSuccess(&received_events)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_events := models.EventList{
		Events: []models.Event{
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
				ID:          "testeventid1337",
				Name:        "eventhax0r",
				Description: "A hax0r event",
				StartTime:   current_unix_time + 120000,
				EndTime:     current_unix_time + 150000,
				Sponsor:     "testsponsor1",
				EventType:   "MINIEVENT",
				Locations: []models.EventLocation{
					{
						Description: "testlocationdescription3",
						Tags:        []string{"SIEBEL3", "ECEB2"},
						Latitude:    123.456,
						Longitude:   123.456,
					},
				},
				Points: 1337,
			},
		},
	}

	if !reflect.DeepEqual(received_events, expected_events) {
		t.Fatalf("Wrong event info. Expected %v, got %v", expected_events, received_events)
	}
}

func TestGetFilteredEventByEventType(t *testing.T) {
	CreateEvents()
	AddAdditionalEvents()
	defer ClearEvents()

	received_events := models.EventList{}
	response, err := public_client.New().Get("/event/filter/?eventType=FOOD").ReceiveSuccess(&received_events)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_events := models.EventList{
		Events: []models.Event{
			{
				ID:          "testeventid67890",
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
			{
				ID:          "testeventid4",
				Name:        "eventdinner",
				Description: "Get your dinner!",
				StartTime:   current_unix_time + 240000,
				EndTime:     current_unix_time + 300000,
				Sponsor:     "",
				EventType:   "FOOD",
				Locations: []models.EventLocation{
					{
						Description: "testlocationdescription4",
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

func TestGetFilteredEventByPoints(t *testing.T) {
	CreateEvents()
	AddAdditionalEvents()
	defer ClearEvents()

	// TODO: Probably should make this some "range" idk

	received_events := models.EventList{}
	response, err := public_client.New().Get("/event/filter/?points=1337").ReceiveSuccess(&received_events)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_events := models.EventList{
		Events: []models.Event{
			{
				ID:          "testeventid1337",
				Name:        "eventhax0r",
				Description: "A hax0r event",
				StartTime:   current_unix_time + 120000,
				EndTime:     current_unix_time + 150000,
				Sponsor:     "testsponsor1",
				EventType:   "MINIEVENT",
				Locations: []models.EventLocation{
					{
						Description: "testlocationdescription3",
						Tags:        []string{"SIEBEL3", "ECEB2"},
						Latitude:    123.456,
						Longitude:   123.456,
					},
				},
				Points: 1337,
			},
		},
	}

	if !reflect.DeepEqual(received_events, expected_events) {
		t.Fatalf("Wrong event info. Expected %v, got %v", expected_events, received_events)
	}
}

func TestGetFilteredEventMultipleKeys(t *testing.T) {
	CreateEvents()
	AddAdditionalEvents()
	defer ClearEvents()

	received_events := models.EventList{}
	response, err := public_client.New().Get("/event/filter/?sponsor=testsponsor1&eventType=MINIEVENT").ReceiveSuccess(&received_events)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_events := models.EventList{
		Events: []models.Event{
			{
				ID:          "testeventid1337",
				Name:        "eventhax0r",
				Description: "A hax0r event",
				StartTime:   current_unix_time + 120000,
				EndTime:     current_unix_time + 150000,
				Sponsor:     "testsponsor1",
				EventType:   "MINIEVENT",
				Locations: []models.EventLocation{
					{
						Description: "testlocationdescription3",
						Tags:        []string{"SIEBEL3", "ECEB2"},
						Latitude:    123.456,
						Longitude:   123.456,
					},
				},
				Points: 1337,
			},
		},
	}

	if !reflect.DeepEqual(received_events, expected_events) {
		t.Fatalf("Wrong event info. Expected %v, got %v", expected_events, received_events)
	}
}

func TestGetFilteredEventsBadArgs(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	api_err := errors.ApiError{}
	response, err := public_client.New().Get("/event/filter/?nonsensefield=trydecipheringthis!").Receive(nil, &api_err)

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
		Message:  "Could not fetch filtered list of events.",
		RawError: "Invalid key nonsensefield",
	}

	if !reflect.DeepEqual(expected_error, api_err) {
		t.Fatalf("Wrong error resonse received. Expected %v, got %v", expected_error, api_err)
	}
}
