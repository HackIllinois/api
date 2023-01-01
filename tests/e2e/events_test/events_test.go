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
var checkin_db_name string

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
	checkin_db_name, err = cfg.Get("CHECKIN_DB_NAME")
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
	client.Database(checkin_db_name).Drop(context.Background())

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

		db.Collection("usertrackers").Indexes().CreateOne(context.Background(), mongo.IndexModel{
			Keys:    bson.M{"userid": 1},
			Options: options.Index().SetUnique(true),
		})

		db.Collection("favorites").Indexes().CreateOne(context.Background(), mongo.IndexModel{
			Keys:    bson.M{"id": 1},
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
		StartTime:   current_unix_time + 60000,
		EndTime:     current_unix_time + 120000,
		Sponsor:     "",
		EventType:   "FOOD",
		Locations: []event_models.EventLocation{
			{
				Description: "testlocationdescription2",
				Tags:        []string{"SIEBEL3", "ECEB2"},
				Latitude:    123.456,
				Longitude:   123.456,
			},
		},
		Points: 0,
	}

	client.Database(events_db_name).Collection("events").InsertOne(context.Background(), event1)
	client.Database(events_db_name).Collection("events").InsertOne(context.Background(), event2)

	event_tracker1 := event_models.EventTracker{
		EventID: "testeventid12345",
		Users:   []string{},
	}
	event_tracker2 := event_models.EventTracker{
		EventID: "testeventid67890",
		Users: []string{
			"localadmin",
		},
	}

	client.Database(events_db_name).Collection("eventtrackers").InsertOne(context.Background(), event_tracker1)
	client.Database(events_db_name).Collection("eventtrackers").InsertOne(context.Background(), event_tracker2)

	user_tracker := event_models.UserTracker{
		UserID: "localadmin",
		Events: []string{
			"testeventid67890",
		},
	}

	client.Database(events_db_name).Collection("usertrackers").InsertOne(context.Background(), user_tracker)

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
	client.Database(events_db_name).Collection("usertrackers").DeleteMany(context.Background(), bson.D{})
	client.Database(events_db_name).Collection("eventcodes").DeleteMany(context.Background(), bson.D{})
	client.Database(events_db_name).Collection("favorites").DeleteMany(context.Background(), bson.D{})
}

func TestStaffActions(t *testing.T) {
	defer ClearEvents()
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
