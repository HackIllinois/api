package tests

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/HackIllinois/api/common/errors"
	checkin_models "github.com/HackIllinois/api/services/checkin/models"
	event_models "github.com/HackIllinois/api/services/event/models"
	"go.mongodb.org/mongo-driver/bson"
)

func TestTrackUserNormal(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	checkin_info := checkin_models.UserCheckin{
		ID:              TEST_USER_ID,
		Override:        false,
		HasCheckedIn:    true,
		HasPickedUpSwag: false,
		RsvpData:        map[string]interface{}{},
	}

	client.Database(checkin_db_name).Collection("checkins").InsertOne(context.Background(), checkin_info)
	defer client.Database(checkin_db_name).Collection("checkins").DeleteMany(context.Background(), bson.D{})

	tracking_info := event_models.TrackingInfo{
		EventID: TEST_EVENT_1_ID,
		UserID:  TEST_USER_ID,
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
			EventID: TEST_EVENT_1_ID,
			Users:   []string{TEST_USER_ID},
		},
		UserTracker: event_models.UserTracker{
			UserID: TEST_USER_ID,
			Events: []string{TEST_EVENT_2_ID, TEST_EVENT_1_ID},
		},
	}

	if !reflect.DeepEqual(res_tracking, expected_trackers) {
		t.Fatalf("Wrong result received. Expected %v, got %v", expected_trackers, res_tracking)
	}

	expected_event_tracker := event_models.EventTracker{
		EventID: TEST_EVENT_1_ID,
		Users:   []string{TEST_USER_ID},
	}

	var actual_event_tracker event_models.EventTracker

	res := client.Database(events_db_name).Collection("eventtrackers").FindOne(context.Background(), bson.M{"eventid": TEST_EVENT_1_ID})

	err = res.Decode(&actual_event_tracker)

	if err != nil {
		t.Fatalf("Had trouble finding event tracker in database: %v", err)
		return
	}

	if !reflect.DeepEqual(actual_event_tracker, expected_event_tracker) {
		t.Fatalf("Wrong event tracker received from database. Expected %v, got %v", expected_event_tracker, actual_event_tracker)
	}

	expected_user_tracker := event_models.UserTracker{
		UserID: TEST_USER_ID,
		Events: []string{TEST_EVENT_2_ID, TEST_EVENT_1_ID},
	}

	var actual_user_tracker event_models.UserTracker

	res = client.Database(events_db_name).Collection("usertrackers").FindOne(context.Background(), bson.M{"userid": TEST_USER_ID})

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
		EventID: TEST_EVENT_1_ID,
		UserID:  TEST_USER_ID,
	}

	api_err := errors.ApiError{}
	response, err := staff_client.New().Post("/event/track/").BodyJSON(tracking_info).Receive(nil, &api_err)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusUnprocessableEntity {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_error := errors.ApiError{
		Status:   http.StatusUnprocessableEntity,
		Type:     "ATTRIBUTE_MISMATCH_ERROR",
		Message:  "User must be checked-in to attend event.",
		RawError: "User must be checked-in to attend event.",
	}

	if !reflect.DeepEqual(expected_error, api_err) {
		t.Fatalf("Wrong error response received. Expected %v, got %v", expected_error, api_err)
	}
}
