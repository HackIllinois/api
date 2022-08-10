// Staff tries to make calls to events endpoints
package tests

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/HackIllinois/api/common/configloader"
	event_models "github.com/HackIllinois/api/services/event/models"
	profile_models "github.com/HackIllinois/api/services/profile/models"
	"github.com/HackIllinois/api/tests/common"
	"github.com/dghubble/sling"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var staff_client *sling.Sling
var public_client *sling.Sling
var user_client *sling.Sling
var client *mongo.Client

var events_db_name string
var profile_db_name string

var current_unix_time = time.Now().Unix()

func TestMain(m *testing.M) {

	cfg, err := configloader.Load(os.Getenv("HI_CONFIG"))

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	staff_client = common.GetSlingClient("Staff")
	public_client = common.GetSlingClient("")
	user_client = common.GetSlingClient("User")

	client = common.GetLocalMongoSession()

	events_db_name, err = cfg.Get("EVENT_DB_NAME")
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}
	profile_db_name, err = cfg.Get("PROFILE_DB_NAME")
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}
	ResetDatabase()

	return_code := m.Run()
	ResetDatabase()
	os.Exit(return_code)
}

func ResetDatabase() {
	client.Database(events_db_name).Drop(context.Background())
	client.Database(profile_db_name).Drop(context.Background())

	{
		// establishes unique id indexes to prevent duplicate documents
		db := client.Database(events_db_name)

		db.Collection("events").Indexes().CreateOne(context.Background(), mongo.IndexModel{
			Keys:    bson.M{"id": 1},
			Options: options.Index().SetUnique(true),
		})

		db.Collection("eventcodes").Indexes().CreateOne(context.Background(), mongo.IndexModel{
			Keys:    bson.M{"id": 1},
			Options: options.Index().SetUnique(true),
		})

		db.Collection("eventtrackers").Indexes().CreateOne(context.Background(), mongo.IndexModel{
			Keys:    bson.M{"eventid": 1},
			Options: options.Index().SetUnique(true),
		})
	}
}

func CreateEvents() {
	event1 := event_models.Event{
		ID:          "testeventid12345",
		Name:        "testevent1",
		Description: "testdescription1",
		StartTime:   current_unix_time,
		EndTime:     current_unix_time + 60000,
		Sponsor:     "testsponsor1",
		EventType:   "WORKSHOP",
		Locations: []event_models.EventLocation{
			{
				Description: "testlocationdescription1",
				Tags:        []string{"SIEBEL3", "ECEB2"},
				Latitude:    123.456,
				Longitude:   123.456,
			},
		},
		Points: 50,
	}

	event2 := event_models.Event{
		ID:          "testeventid67890",
		Name:        "testevent2",
		Description: "testdescription2",
		StartTime:   current_unix_time,
		EndTime:     current_unix_time + 60000,
		Sponsor:     "testsponsor2",
		EventType:   "WORKSHOP",
		Locations: []event_models.EventLocation{
			{
				Description: "testlocationdescription2",
				Tags:        []string{"SIEBEL3", "ECEB2"},
				Latitude:    123.456,
				Longitude:   123.456,
			},
		},
		Points: 100,
	}

	client.Database(events_db_name).Collection("events").InsertOne(context.Background(), event1)
	client.Database(events_db_name).Collection("events").InsertOne(context.Background(), event2)

	event_tracker1 := event_models.EventTracker{
		EventID: "testeventid12345",
		Users:   []string{},
	}
	event_tracker2 := event_models.EventTracker{
		EventID: "testeventid67890",
		Users:   []string{},
	}

	client.Database(events_db_name).Collection("eventtrackers").InsertOne(context.Background(), event_tracker1)
	client.Database(events_db_name).Collection("eventtrackers").InsertOne(context.Background(), event_tracker2)

	event_code1 := event_models.EventCode{
		ID:         "testeventid12345",
		Code:       "123456",
		Expiration: current_unix_time + 60000,
	}
	event_code2 := event_models.EventCode{
		ID:         "testeventid67890",
		Code:       "abcdef",
		Expiration: current_unix_time - 60000,
	}

	client.Database(events_db_name).Collection("eventcodes").InsertOne(context.Background(), event_code1)
	client.Database(events_db_name).Collection("eventcodes").InsertOne(context.Background(), event_code2)
}

func CreateProfile() {
	profile := profile_models.Profile{
		ID:        "theadminprofile",
		FirstName: "HackIllinois",
		LastName:  "Admin",
		Points:    0,
	}

	client.Database(profile_db_name).Collection("profiles").InsertOne(context.Background(), profile)

	userid_profileid := profile_models.IdMap{
		UserID:    "localadmin",
		ProfileID: "theadminprofile",
	}
	client.Database(profile_db_name).Collection("profileids").InsertOne(context.Background(), userid_profileid)
}

func ClearProfiles() {
	client.Database(profile_db_name).Collection("profiles").DeleteMany(context.Background(), bson.D{})
	client.Database(profile_db_name).Collection("profileids").DeleteMany(context.Background(), bson.D{})
	client.Database(profile_db_name).Collection("profileattendance").DeleteMany(context.Background(), bson.D{})
}

func ClearEvents() {
	client.Database(events_db_name).Collection("events").DeleteMany(context.Background(), bson.D{})
	client.Database(events_db_name).Collection("eventtrackers").DeleteMany(context.Background(), bson.D{})
	client.Database(events_db_name).Collection("eventcodes").DeleteMany(context.Background(), bson.D{})
	client.Database(events_db_name).Collection("favorites").DeleteMany(context.Background(), bson.D{})
}

func TestGetEventNormal(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	event_id := "testeventid12345"
	received_event := event_models.Event{}
	response, err := public_client.New().Get(fmt.Sprintf("/event/%s/", event_id)).ReceiveSuccess(&received_event)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_event := event_models.Event{
		ID:          "testeventid12345",
		Name:        "testevent1",
		Description: "testdescription1",
		StartTime:   current_unix_time,
		EndTime:     current_unix_time + 60000,
		Sponsor:     "testsponsor1",
		EventType:   "WORKSHOP",
		Locations: []event_models.EventLocation{
			{
				Description: "testlocationdescription1",
				Tags:        []string{"SIEBEL3", "ECEB2"},
				Latitude:    123.456,
				Longitude:   123.456,
			},
		},
		Points: 50,
	}

	if !reflect.DeepEqual(received_event, expected_event) {
		t.Fatalf("Wrong event info. Expected %v, got %v", expected_event, received_event)
	}
}

func TestGetEventNotExist(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	event_id := "nonsense_eventid"
	received_event := event_models.Event{}
	response, err := public_client.New().Get(fmt.Sprintf("/event/%s/", event_id)).ReceiveSuccess(&received_event)

	if err != nil {
		t.Error("Unable to make request")
	}
	if response.StatusCode != http.StatusInternalServerError {
		t.Errorf("Request returned HTTP error %d", response.StatusCode)
	}
}

func TestDeleteNormal(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	event_id := "testeventid12345"
	received_event := event_models.Event{}
	response, err := staff_client.New().Delete(fmt.Sprintf("/event/%s/", event_id)).ReceiveSuccess(&received_event)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_event := event_models.Event{
		ID:          "testeventid12345",
		Name:        "testevent1",
		Description: "testdescription1",
		StartTime:   current_unix_time,
		EndTime:     current_unix_time + 60000,
		Sponsor:     "testsponsor1",
		EventType:   "WORKSHOP",
		Locations: []event_models.EventLocation{
			{
				Description: "testlocationdescription1",
				Tags:        []string{"SIEBEL3", "ECEB2"},
				Latitude:    123.456,
				Longitude:   123.456,
			},
		},
		Points: 50,
	}

	if !reflect.DeepEqual(received_event, expected_event) {
		t.Fatalf("Wrong event info. Expected %v, got %v", expected_event, received_event)
	}

	response, err = staff_client.New().Get(fmt.Sprintf("/event/%s/", event_id)).ReceiveSuccess(nil)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusInternalServerError {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}
}

func TestDeleteNotExist(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	event_id := "nonsense_eventid"
	received_event := event_models.Event{}
	response, err := staff_client.New().Delete(fmt.Sprintf("/event/%s/", event_id)).ReceiveSuccess(&received_event)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusInternalServerError {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	cursor, _ := client.Database(events_db_name).Collection("events").Find(context.Background(), bson.D{})
	res := []event_models.Event{}
	err = cursor.All(context.TODO(), &res)

	if err != nil {
		t.Fatalf("Test failed due to unexpected error: %v", err)
		return
	}

	expected_events := []event_models.Event{
		{
			ID:          "testeventid12345",
			Name:        "testevent1",
			Description: "testdescription1",
			StartTime:   current_unix_time,
			EndTime:     current_unix_time + 60000,
			Sponsor:     "testsponsor1",
			EventType:   "WORKSHOP",
			Locations: []event_models.EventLocation{
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
			StartTime:   current_unix_time,
			EndTime:     current_unix_time + 60000,
			Sponsor:     "testsponsor2",
			EventType:   "WORKSHOP",
			Locations: []event_models.EventLocation{
				{
					Description: "testlocationdescription2",
					Tags:        []string{"SIEBEL3", "ECEB2"},
					Latitude:    123.456,
					Longitude:   123.456,
				},
			},
			Points: 100,
		},
	}

	if !reflect.DeepEqual(res, expected_events) {
		t.Fatalf("Database contained wrong event info. Expected %v, got %v", expected_events, res)
	}
}

func TestDeleteForbidden(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	event_id := "testeventid12345"
	received_event := event_models.Event{}
	response, err := user_client.New().Delete(fmt.Sprintf("/event/%s/", event_id)).ReceiveSuccess(&received_event)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusForbidden {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	cursor, _ := client.Database(events_db_name).Collection("events").Find(context.Background(), bson.D{})
	res := []event_models.Event{}
	err = cursor.All(context.TODO(), &res)

	if err != nil {
		t.Fatalf("Test failed due to unexpected error: %v", err)
		return
	}

	expected_events := []event_models.Event{
		{
			ID:          "testeventid12345",
			Name:        "testevent1",
			Description: "testdescription1",
			StartTime:   current_unix_time,
			EndTime:     current_unix_time + 60000,
			Sponsor:     "testsponsor1",
			EventType:   "WORKSHOP",
			Locations: []event_models.EventLocation{
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
			StartTime:   current_unix_time,
			EndTime:     current_unix_time + 60000,
			Sponsor:     "testsponsor2",
			EventType:   "WORKSHOP",
			Locations: []event_models.EventLocation{
				{
					Description: "testlocationdescription2",
					Tags:        []string{"SIEBEL3", "ECEB2"},
					Latitude:    123.456,
					Longitude:   123.456,
				},
			},
			Points: 100,
		},
	}

	if !reflect.DeepEqual(res, expected_events) {
		t.Fatalf("Database contained wrong event info. Expected %v, got %v", expected_events, res)
	}
}

func TestGetAllEventsNormal(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	received_events := event_models.EventList{}
	response, err := public_client.New().Get("/event/").ReceiveSuccess(&received_events)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_events := event_models.EventList{
		Events: []event_models.Event{
			{
				ID:          "testeventid12345",
				Name:        "testevent1",
				Description: "testdescription1",
				StartTime:   current_unix_time,
				EndTime:     current_unix_time + 60000,
				Sponsor:     "testsponsor1",
				EventType:   "WORKSHOP",
				Locations: []event_models.EventLocation{
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
				StartTime:   current_unix_time,
				EndTime:     current_unix_time + 60000,
				Sponsor:     "testsponsor2",
				EventType:   "WORKSHOP",
				Locations: []event_models.EventLocation{
					{
						Description: "testlocationdescription2",
						Tags:        []string{"SIEBEL3", "ECEB2"},
						Latitude:    123.456,
						Longitude:   123.456,
					},
				},
				Points: 100,
			},
		},
	}

	if !reflect.DeepEqual(received_events, expected_events) {
		t.Fatalf("Wrong event info. Expected %v, got %v", expected_events, received_events)
	}
}

func TestGetAllEventsNone(t *testing.T) {
	received_events := event_models.EventList{}
	response, err := staff_client.New().Get("/event/").ReceiveSuccess(&received_events)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_events := event_models.EventList{
		Events: []event_models.Event{},
	}

	if !reflect.DeepEqual(received_events, expected_events) {
		t.Fatalf("Wrong event info. Expected %v, got %v", expected_events, received_events)
	}
}

func TestGetFilteredEventsNormal(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	received_events := event_models.EventList{}
	response, err := public_client.New().Get("/event/filter/?name=testevent1").ReceiveSuccess(&received_events)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_events := event_models.EventList{
		Events: []event_models.Event{
			{
				ID:          "testeventid12345",
				Name:        "testevent1",
				Description: "testdescription1",
				StartTime:   current_unix_time,
				EndTime:     current_unix_time + 60000,
				Sponsor:     "testsponsor1",
				EventType:   "WORKSHOP",
				Locations: []event_models.EventLocation{
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

func TestGetFilteredEventsBadArgs(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	response, err := public_client.New().Get("/event/filter/?nonsensefield=trydecipheringthis!").ReceiveSuccess(nil)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusInternalServerError {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}
}

func TestCreateEventNormal(t *testing.T) {
	defer ClearEvents()
	event_info := event_models.Event{
		Name:        "testevent1",
		Description: "testdescription1",
		StartTime:   current_unix_time,
		EndTime:     current_unix_time + 60000,
		Sponsor:     "testsponsor1",
		EventType:   "WORKSHOP",
		Locations: []event_models.EventLocation{
			{
				Description: "testlocationdescription1",
				Tags:        []string{"SIEBEL3", "ECEB2"},
				Latitude:    123.456,
				Longitude:   123.456,
			},
		},
		Points: 50,
	}
	received_event := event_models.Event{}
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
	res := []event_models.Event{}
	err = cursor.All(context.TODO(), &res)

	if err != nil {
		t.Fatalf("Test failed due to unexpected error: %v", err)
		return
	}

	expected_res := []event_models.Event{
		event_info,
	}

	if !reflect.DeepEqual(res, expected_res) {
		t.Fatalf("Database contained wrong event info. Expected %v, got %v", expected_res, res)
	}
}

func TestCreateEventForbidden(t *testing.T) {
	defer ClearEvents()
	event_info := event_models.Event{
		Name:        "testevent1",
		Description: "testdescription1",
		StartTime:   current_unix_time,
		EndTime:     current_unix_time + 60000,
		Sponsor:     "testsponsor1",
		EventType:   "WORKSHOP",
		Locations: []event_models.EventLocation{
			{
				Description: "testlocationdescription1",
				Tags:        []string{"SIEBEL3", "ECEB2"},
				Latitude:    123.456,
				Longitude:   123.456,
			},
		},
		Points: 50,
	}
	received_event := event_models.Event{}
	response, err := user_client.New().Post("/event/").BodyJSON(event_info).ReceiveSuccess(&received_event)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusForbidden {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	cursor, _ := client.Database(events_db_name).Collection("events").Find(context.Background(), bson.D{})
	res := []event_models.Event{}
	err = cursor.All(context.TODO(), &res)

	if err != nil {
		t.Fatalf("Test failed due to unexpected error: %v", err)
		return
	}

	expected_res := []event_models.Event{}

	if !reflect.DeepEqual(res, expected_res) {
		t.Fatalf("Database contained wrong event info. Expected %v, got %v", expected_res, res)
	}
}

func TestUpdateEventNormal(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	res := client.Database(events_db_name).Collection("events").FindOne(context.Background(), bson.M{"id": "testeventid12345"})

	event := event_models.Event{}

	res.Decode(&event)

	event.Description = "It's a new description!"
	var received_event event_models.Event

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

	event := event_models.Event{}

	res.Decode(&event)

	event.Description = "It's a new description!"
	var received_event event_models.Event

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

	event := event_models.Event{}

	res.Decode(&event)

	event.ID = "adifferentidthatsnotbeenadded"
	event.Description = "It's a new description!"
	var received_event event_models.Event

	response, err := staff_client.New().Put("/event/").BodyJSON(event).ReceiveSuccess(&received_event)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusInternalServerError {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}
}

func TestGetEventCodeNormal(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	recieved_code := event_models.EventCode{}
	id := "testeventid12345"
	response, err := staff_client.New().Get(fmt.Sprintf("/event/code/%s/", id)).ReceiveSuccess(&recieved_code)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_code := event_models.EventCode{
		ID:         "testeventid12345",
		Code:       "123456",
		Expiration: current_unix_time + 60000,
	}

	if !reflect.DeepEqual(recieved_code, expected_code) {
		t.Fatalf("Wrong event code info. Expected %v, got %v", expected_code, recieved_code)
	}
}

func TestGetEventCodeForbidden(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	id := "testeventid12345"
	response, err := user_client.New().Get(fmt.Sprintf("/event/code/%s/", id)).ReceiveSuccess(nil)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusForbidden {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}
}

func TestCheckinNormal(t *testing.T) {
	CreateEvents()
	defer ClearEvents()
	CreateProfile()
	defer ClearProfiles()

	req := event_models.CheckinRequest{
		Code: "123456",
	}
	received_res := event_models.CheckinResult{}
	response, err := staff_client.New().Post("/event/checkin/").BodyJSON(req).ReceiveSuccess(&received_res)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_res := event_models.CheckinResult{
		NewPoints:   50,
		TotalPoints: 50,
		Status:      "Success",
	}

	if !reflect.DeepEqual(received_res, expected_res) {
		t.Fatalf("Wrong result received. Expected %v, got %v", expected_res, received_res)
	}
}

func TestCheckinAddToExistingPoints(t *testing.T) {
	CreateEvents()
	defer ClearEvents()
	CreateProfile()
	defer ClearProfiles()

	client.Database(profile_db_name).Collection("profiles").UpdateOne(
		context.Background(),
		bson.M{"id": "theadminprofile"},
		bson.M{"$set": bson.M{
			"points": 19,
		}},
	)

	req := event_models.CheckinRequest{
		Code: "123456",
	}
	received_res := event_models.CheckinResult{}
	response, err := staff_client.New().Post("/event/checkin/").BodyJSON(req).ReceiveSuccess(&received_res)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_res := event_models.CheckinResult{
		NewPoints:   50,
		TotalPoints: 69,
		Status:      "Success",
	}

	if !reflect.DeepEqual(received_res, expected_res) {
		t.Fatalf("Wrong result received. Expected %v, got %v", expected_res, received_res)
	}
}

func TestCheckinInvalidCode(t *testing.T) {
	CreateEvents()
	defer ClearEvents()
	CreateProfile()
	defer ClearProfiles()

	req := event_models.CheckinRequest{
		Code: "wrongcode",
	}
	received_res := event_models.CheckinResult{}
	response, err := staff_client.New().Post("/event/checkin/").BodyJSON(req).ReceiveSuccess(&received_res)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_res := event_models.CheckinResult{
		NewPoints:   -1,
		TotalPoints: -1,
		Status:      "InvalidCode",
	}

	if !reflect.DeepEqual(received_res, expected_res) {
		t.Fatalf("Wrong result received. Expected %v, got %v", expected_res, received_res)
	}
}

func TestCheckinInvalidTime(t *testing.T) {
	CreateEvents()
	defer ClearEvents()
	CreateProfile()
	defer ClearProfiles()

	req := event_models.CheckinRequest{
		Code: "abcdef",
	}
	received_res := event_models.CheckinResult{}
	response, err := staff_client.New().Post("/event/checkin/").BodyJSON(req).ReceiveSuccess(&received_res)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_res := event_models.CheckinResult{
		NewPoints:   -1,
		TotalPoints: -1,
		Status:      "InvalidTime",
	}

	if !reflect.DeepEqual(received_res, expected_res) {
		t.Fatalf("Wrong result received. Expected %v, got %v", expected_res, received_res)
	}
}

func TestCheckinAlreadyCheckedIn(t *testing.T) {
	CreateEvents()
	defer ClearEvents()
	CreateProfile()
	defer ClearProfiles()

	client.Database(profile_db_name).Collection("profileattendance").UpdateOne(
		context.Background(),
		bson.M{"id": "theadminprofile"},
		bson.M{"$addToSet": bson.M{
			"events": "testeventid12345",
		}},
		options.Update().SetUpsert(true),
	)

	req := event_models.CheckinRequest{
		Code: "123456",
	}
	received_res := event_models.CheckinResult{}
	response, err := staff_client.New().Post("/event/checkin/").BodyJSON(req).ReceiveSuccess(&received_res)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_res := event_models.CheckinResult{
		NewPoints:   -1,
		TotalPoints: -1,
		Status:      "AlreadyCheckedIn",
	}

	if !reflect.DeepEqual(received_res, expected_res) {
		t.Fatalf("Wrong result received. Expected %v, got %v", expected_res, received_res)
	}
}

func TestAddFavoriteEventNormal(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	req := event_models.EventFavoriteModification{
		EventID: "testeventid12345",
	}
	received_res := event_models.EventFavorites{}
	response, err := user_client.New().Post("/event/favorite/").BodyJSON(req).ReceiveSuccess(&received_res)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_res := event_models.EventFavorites{
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

	req := event_models.EventFavoriteModification{
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

	res := []event_models.EventFavorites{}
	err = cursor.All(context.TODO(), &res)

	if err != nil {
		t.Fatalf("Test failed due to unexpected error: %v", err)
		return
	}

	expected_res := []event_models.EventFavorites{}

	if !reflect.DeepEqual(res, expected_res) {
		t.Fatalf("Wrong result received from database. Expected %v, got %v", expected_res, res)
	}
}

func TestStaffActions(t *testing.T) {
	// 1. Create event
	event_info := event_models.Event{
		Name:        "testname",
		Description: "testdescription2",
		StartTime:   534545,
		EndTime:     534545 + 60000,
		Sponsor:     "testsponsor",
		EventType:   "WORKSHOP",
		Locations: []event_models.EventLocation{
			{
				Description: "testlocationdescription",
				Tags:        []string{"SIEBEL3", "ECEB2"},
				Latitude:    123.456,
				Longitude:   123.456,
			},
		},
		Points: 100,
	}

	received_event := event_models.Event{}
	response, err := staff_client.New().Post("/event/").BodyJSON(event_info).ReceiveSuccess(&received_event)

	if err != nil {
		t.Errorf("Unable to make request")
	}
	if response.StatusCode != http.StatusOK {
		t.Errorf("Request returned HTTP error %d", response.StatusCode)
	}

	// 2. Update event
	event_id := received_event.ID
	event_info_updated := event_models.Event{
		ID:          event_id,
		Name:        "testname",
		Description: "testdescription2",
		StartTime:   534545,
		EndTime:     534545 + 60000,
		Sponsor:     "testsponsor",
		EventType:   "WORKSHOP",
		Locations: []event_models.EventLocation{
			{
				Description: "testlocationdescription",
				Tags:        []string{"SIEBEL3", "ECEB2"},
				Latitude:    123.456,
				Longitude:   123.456,
			},
		},
		Points: 200,
	}

	received_event = event_models.Event{}
	response, err = staff_client.New().Put("/event/").BodyJSON(event_info_updated).ReceiveSuccess(&received_event)

	if err != nil {
		t.Errorf("Unable to make request")
	}
	if response.StatusCode != http.StatusOK {
		t.Errorf("Request returned HTTP error %d", response.StatusCode)
	}
	if !reflect.DeepEqual(received_event, event_info_updated) {
		t.Errorf("Wrong event info. Expected %v, got %v", event_info_updated, received_event)
	}

	// 3. Fetch event
	endpoint_address := fmt.Sprintf("/event/%s/", event_id)

	received_event = event_models.Event{}
	response, err = staff_client.New().Get(endpoint_address).ReceiveSuccess(&received_event)

	if err != nil {
		t.Errorf("Unable to make request")
	}
	if response.StatusCode != http.StatusOK {
		t.Errorf("Request returned HTTP error %d", response.StatusCode)
	}
	if !reflect.DeepEqual(received_event, event_info_updated) {
		t.Errorf("Wrong event info. Expected %v, got %v", event_info_updated, received_event)
	}
}
