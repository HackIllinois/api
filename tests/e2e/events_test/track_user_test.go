package tests

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	checkin_models "github.com/HackIllinois/api/services/checkin/models"
	event_models "github.com/HackIllinois/api/services/event/models"
	"go.mongodb.org/mongo-driver/bson"
)

func TestTrackUserNormal(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	checkin_info := checkin_models.UserCheckin{
		ID:              "localadmin",
		Override:        false,
		HasCheckedIn:    true,
		HasPickedUpSwag: false,
		RsvpData:        map[string]interface{}{},
	}

	client.Database(checkin_db_name).Collection("checkins").InsertOne(context.Background(), checkin_info)
	defer client.Database(checkin_db_name).Collection("checkins").DeleteMany(context.Background(), bson.D{})

	tracking_info := event_models.TrackingInfo{
		EventID: "testeventid12345",
		UserID:  "localadmin",
	}

	res_tracking := event_models.TrackingStatus{}
	response, err := staff_client.New().Post("/event/track/").BodyJSON(tracking_info).ReceiveSuccess(&res_tracking)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_trackers := event_models.TrackingStatus{
		EventTracker: event_models.EventTracker{
			EventID: "testeventid12345",
			Users:   []string{"localadmin"},
		},
		UserTracker: event_models.UserTracker{
			UserID: "localadmin",
			Events: []string{"testeventid67890", "testeventid12345"},
		},
	}

	if !reflect.DeepEqual(res_tracking, expected_trackers) {
		t.Fatalf("Wrong result received. Expected %v, got %v", expected_trackers, res_tracking)
	}

	expected_event_tracker := event_models.EventTracker{
		EventID: "testeventid12345",
		Users:   []string{"localadmin"},
	}

	var actual_event_tracker event_models.EventTracker

	res := client.Database(events_db_name).Collection("eventtrackers").FindOne(context.Background(), bson.M{"eventid": "testeventid12345"})

	err = res.Decode(&actual_event_tracker)

	if err != nil {
		t.Fatalf("Had trouble finding event tracker in database: %v", err)
		return
	}

	if !reflect.DeepEqual(actual_event_tracker, expected_event_tracker) {
		t.Fatalf("Wrong event tracker received from database. Expected %v, got %v", expected_event_tracker, actual_event_tracker)
	}

	expected_user_tracker := event_models.UserTracker{
		UserID: "localadmin",
		Events: []string{"testeventid67890", "testeventid12345"},
	}

	var actual_user_tracker event_models.UserTracker

	res = client.Database(events_db_name).Collection("usertrackers").FindOne(context.Background(), bson.M{"userid": "localadmin"})

	err = res.Decode(&actual_user_tracker)

	if err != nil {
		t.Fatalf("Had trouble finding user tracker in database: %v", err)
		return
	}

	if !reflect.DeepEqual(actual_user_tracker, expected_user_tracker) {
		t.Fatalf("Wrong user tracker received from database. Expected %v, got %v", expected_user_tracker, actual_user_tracker)
	}
}

func TestTrackUserNotCheckedin(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	tracking_info := event_models.TrackingInfo{
		EventID: "testeventid12345",
		UserID:  "localadmin",
	}

	res_tracking := event_models.TrackingStatus{}
	response, err := staff_client.New().Post("/event/track/").BodyJSON(tracking_info).ReceiveSuccess(&res_tracking)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusUnprocessableEntity {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}
}
