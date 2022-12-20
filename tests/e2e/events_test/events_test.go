// Staff tries to make calls to events endpoints
package tests

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/HackIllinois/api/common/configloader"
	event_models "github.com/HackIllinois/api/services/event/models"
	"github.com/HackIllinois/api/tests/common"
	"github.com/dghubble/sling"
	"go.mongodb.org/mongo-driver/mongo"
)

var staff_client *sling.Sling
var client *mongo.Client

func TestMain(m *testing.M) {

	cfg, err := configloader.Load(os.Getenv("HI_CONFIG"))

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	staff_client = common.GetSlingClient("Staff")

	client = common.GetLocalMongoSession()

	events_db_name, err := cfg.Get("EVENT_DB_NAME")
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}
	client.Database(events_db_name).Drop(context.Background())

	return_code := m.Run()
	os.Exit(return_code)
}

func TestStaffActions(t *testing.T) {
	// 1. Create event
	event_info := event_models.EventDB{
		EventPublic: event_models.EventPublic{
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
		},
		IsPrivate: false,
	}

	received_event := event_models.EventDB{}
	response, err := staff_client.New().Post("/event/").BodyJSON(event_info).ReceiveSuccess(&received_event)

	if err != nil {
		t.Errorf("Unable to make request")
	}
	if response.StatusCode != 200 {
		t.Errorf("Request returned HTTP error %d", response.StatusCode)
	}

	// 2. Update event
	event_id := received_event.ID
	event_info_updated := event_models.EventDB{
		EventPublic: event_models.EventPublic{
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
		},
		IsPrivate: false,
	}

	received_event = event_models.EventDB{}
	response, err = staff_client.New().Put("/event/").BodyJSON(event_info_updated).ReceiveSuccess(&received_event)

	if err != nil {
		t.Errorf("Unable to make request")
	}
	if response.StatusCode != 200 {
		t.Errorf("Request returned HTTP error %d", response.StatusCode)
	}
	if !reflect.DeepEqual(received_event, event_info_updated) {
		t.Errorf("Wrong event info. Expected %v, got %v", event_info_updated, received_event)
	}

	// 3. Fetch event
	endpoint_address := fmt.Sprintf("/event/%s/", event_id)

	received_event = event_models.EventDB{}
	response, err = staff_client.New().Get(endpoint_address).ReceiveSuccess(&received_event)

	if err != nil {
		t.Errorf("Unable to make request")
	}
	if response.StatusCode != 200 {
		t.Errorf("Request returned HTTP error %d", response.StatusCode)
	}
	if !reflect.DeepEqual(received_event, event_info_updated) {
		t.Errorf("Wrong event info. Expected %v, got %v", event_info_updated, received_event)
	}
}
