package tests

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/services/event/models"
)

func AddAdditionalEvents() {
	event3 := models.EventDB{
		EventPublic: models.EventPublic{
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
		IsPrivate:             false,
		DisplayOnStaffCheckin: false,
	}

	event4 := models.EventDB{
		EventPublic: models.EventPublic{
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
		IsPrivate:             false,
		DisplayOnStaffCheckin: true,
	}

	client.Database(events_db_name).Collection("events").InsertOne(context.Background(), event3)
	client.Database(events_db_name).Collection("events").InsertOne(context.Background(), event4)
}

func TestGetFilteredPublicEventsNormal(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	received_events := models.EventList[models.EventPublic]{}
	response, err := public_client.New().Get("/event/filter/?name=testevent2").ReceiveSuccess(&received_events)
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

func TestGetFilteredPrivateEventAsPublic(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	received_events := models.EventList[models.EventPublic]{}
	response, err := public_client.New().Get("/event/filter/?name=testevent1").ReceiveSuccess(&received_events)
	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_events := models.EventList[models.EventPublic]{
		Events: []models.EventPublic{},
	}

	if !reflect.DeepEqual(received_events, expected_events) {
		t.Fatalf("Wrong event info. Expected %v, got %v", expected_events, received_events)
	}
}

func TestGetFilteredAllEventsAsStaff(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	received_events := models.EventList[models.EventDB]{}
	response, err := staff_client.New().Get("/event/filter/?name=testevent1").ReceiveSuccess(&received_events)
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
		},
	}

	if !reflect.DeepEqual(received_events, expected_events) {
		t.Fatalf("Wrong event info. Expected %v, got %v", expected_events, received_events)
	}
}

func TestGetFilteredAllEventsNoFilter(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	received_events := models.EventList[models.EventDB]{}
	response, err := staff_client.New().Get("/event/filter/").ReceiveSuccess(&received_events)
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

func TestGetFilteredEventsNoFilter(t *testing.T) {
	CreateEvents()
	AddAdditionalEvents()
	defer ClearEvents()

	received_events := models.EventList[models.EventPublic]{}
	response, err := public_client.New().Get("/event/filter/").ReceiveSuccess(&received_events)
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

	received_events := models.EventList[models.EventPublic]{}
	response, err := public_client.New().Get("/event/filter/?name=testevent2").ReceiveSuccess(&received_events)
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

func TestGetFilteredEventsByDescription(t *testing.T) {
	CreateEvents()
	AddAdditionalEvents()
	defer ClearEvents()

	received_events := models.EventList[models.EventPublic]{}
	response, err := public_client.New().
		Get(fmt.Sprintf("/event/filter/?description=%s", url.QueryEscape("A hax0r event"))).
		ReceiveSuccess(&received_events)
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

func TestGetFilteredEventsByStartTime(t *testing.T) {
	CreateEvents()
	AddAdditionalEvents()
	defer ClearEvents()

	// This argument only does exact unix time which is kinda dumb
	// TODO: Change the filters to treat StartTime and EndTime as a range (if not provided, do not
	// include that bound)

	received_events := models.EventList[models.EventPublic]{}
	response, err := public_client.New().
		Get(fmt.Sprintf("/event/filter/?startTime=%d", current_unix_time+120000)).
		ReceiveSuccess(&received_events)
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

func TestGetFilteredEventsByEndTime(t *testing.T) {
	CreateEvents()
	AddAdditionalEvents()
	defer ClearEvents()

	received_events := models.EventList[models.EventPublic]{}
	response, err := public_client.New().
		Get(fmt.Sprintf("/event/filter/?endTime=%d", current_unix_time+150000)).
		ReceiveSuccess(&received_events)
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

func TestGetFilteredEventBySponsor(t *testing.T) {
	CreateEvents()
	AddAdditionalEvents()
	defer ClearEvents()

	received_events := models.EventList[models.EventPublic]{}
	response, err := public_client.New().Get("/event/filter/?sponsor=testsponsor1").ReceiveSuccess(&received_events)
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

	received_events := models.EventList[models.EventPublic]{}
	response, err := public_client.New().Get("/event/filter/?eventType=FOOD").ReceiveSuccess(&received_events)
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

	received_events := models.EventList[models.EventPublic]{}
	response, err := public_client.New().Get("/event/filter/?points=1337").ReceiveSuccess(&received_events)
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

func TestGetFilteredPrivateEventsAsPublic(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	received_events := models.EventList[models.EventDB]{}
	api_err := errors.ApiError{}
	response, err := public_client.New().Get("/event/filter/?isPrivate=true").Receive(&received_events, &api_err)
	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusInternalServerError {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_api_err := errors.ApiError{
		Status:   http.StatusInternalServerError,
		Type:     "DATABASE_ERROR",
		Message:  "Could not fetch filtered list of events.",
		RawError: "Invalid key isprivate",
	}
	if !reflect.DeepEqual(api_err, expected_api_err) {
		t.Fatalf("Wrong event info. Expected %v, got %v", expected_api_err, api_err)
	}
}

func TestGetFilteredDisplayStaffCheckinEventsAsStaff(t *testing.T) {
	CreateEvents()
	AddAdditionalEvents()
	defer ClearEvents()

	received_events := models.EventList[models.EventDB]{}
	response, err := staff_client.New().Get("/event/filter/?displayOnStaffCheckin=true").ReceiveSuccess(&received_events)
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
			{
				EventPublic: models.EventPublic{
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
				IsPrivate:             false,
				DisplayOnStaffCheckin: true,
			},
		},
	}

	if !reflect.DeepEqual(received_events, expected_events) {
		t.Fatalf("Wrong event info. Expected %v, got %v", expected_events, received_events)
	}
}

func TestGetFilteredPrivateEventsAsStaff(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	received_events := models.EventList[models.EventDB]{}
	response, err := staff_client.New().Get("/event/filter/?isPrivate=true").ReceiveSuccess(&received_events)
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

	received_events := models.EventList[models.EventPublic]{}
	response, err := public_client.New().
		Get("/event/filter/?sponsor=testsponsor1&eventType=MINIEVENT").
		ReceiveSuccess(&received_events)
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
