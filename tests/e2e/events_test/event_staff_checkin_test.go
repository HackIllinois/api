package tests

import (
	"context"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/services/event/models"
	profile_models "github.com/HackIllinois/api/services/profile/models"
	rsvp_models "github.com/HackIllinois/api/services/rsvp/models"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var rsvp_data = rsvp_models.UserRsvp{
	Data: map[string]interface{}{
		"id":          "localadmin",
		"isAttending": true,
	},
}

func GenerateRsvpData() {
	client.Database(rsvp_db_name).Collection("rsvps").InsertOne(context.Background(), &rsvp_data)
}

func ClearRsvps() {
	client.Database(rsvp_db_name).Collection("rsvps").DeleteMany(context.Background(), bson.D{})
}

func GenerateValidUserToken(t *testing.T) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":    time.Now().Add(time.Hour).Unix(),
		"userId": TEST_USER_ID,
	})

	signed, err := token.SignedString(TOKEN_SECRET)
	if err != nil {
		t.Fatalf("Failed to generate signed token: %v", err)
	}

	return signed
}

func GenerateExpiredUserToken(t *testing.T) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":    time.Now().Add(-time.Hour).Unix(),
		"userId": TEST_USER_ID,
	})

	signed, err := token.SignedString(TOKEN_SECRET)
	if err != nil {
		t.Fatalf("Failed to generate signed token: %v", err)
	}

	return signed
}

func GenerateInvalidUserToken(t *testing.T) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":    time.Now().Add(time.Hour).Unix(),
		"userId": TEST_USER_ID,
	})

	// Note: please do not ever use nonsense as a secret or this will break
	signed, err := token.SignedString([]byte("nonsense"))
	if err != nil {
		t.Fatalf("Failed to generate signed token: %v", err)
	}

	return signed
}

func TestStaffCheckinNormal(t *testing.T) {
	CreateEvents()
	defer ClearEvents()
	CreateProfile()
	defer ClearProfiles()
	GenerateRsvpData()
	defer ClearRsvps()

	req := models.StaffCheckinRequest{
		EventID:   TEST_EVENT_1_ID,
		UserToken: GenerateValidUserToken(t),
	}
	received_res := models.CheckinResponse{}
	api_err := errors.ApiError{}
	response, err := staff_client.New().Post("/event/staff/checkin/").BodyJSON(req).Receive(&received_res, &api_err)
	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d: %v", response.StatusCode, api_err)
		return
	}

	expected_res := models.CheckinResponse{
		NewPoints:   50,
		TotalPoints: 50,
		Status:      "Success",
		RsvpData:    rsvp_data.Data,
	}

	if !reflect.DeepEqual(received_res, expected_res) {
		t.Fatalf("Wrong result received. Expected %v, got %v", expected_res, received_res)
	}

	res := client.Database(profile_db_name).Collection("profiles").FindOne(context.Background(), bson.M{"id": TEST_PROFILE_ID})

	profile := profile_models.Profile{}
	err = res.Decode(&profile)

	if err != nil {
		t.Fatalf("Had trouble finding profile in database: %v", err)
		return
	}

	// it's not crucial we check every field in profile. The profile E2E tests should be doing that
	if expected_res.TotalPoints != profile.Points {
		t.Fatalf("Wrong amount of points in profile database. Expected %v, got %v", expected_res.TotalPoints, profile.Points)
	}
}

func TestStaffCheckinAddToExistingPoints(t *testing.T) {
	CreateEvents()
	defer ClearEvents()
	CreateProfile()
	defer ClearProfiles()
	GenerateRsvpData()
	defer ClearRsvps()

	client.Database(profile_db_name).Collection("profiles").UpdateOne(
		context.Background(),
		bson.M{"id": TEST_PROFILE_ID},
		bson.M{"$set": bson.M{
			"points": 19,
		}},
	)

	req := models.StaffCheckinRequest{
		EventID:   TEST_EVENT_1_ID,
		UserToken: GenerateValidUserToken(t),
	}
	received_res := models.CheckinResponse{}
	response, err := staff_client.New().Post("/event/staff/checkin/").BodyJSON(req).ReceiveSuccess(&received_res)
	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_res := models.CheckinResponse{
		NewPoints:   50,
		TotalPoints: 69,
		Status:      "Success",
		RsvpData:    rsvp_data.Data,
	}

	if !reflect.DeepEqual(received_res, expected_res) {
		t.Fatalf("Wrong result received. Expected %v, got %v", expected_res, received_res)
	}

	res := client.Database(profile_db_name).Collection("profiles").FindOne(context.Background(), bson.M{"id": TEST_PROFILE_ID})

	profile := profile_models.Profile{}
	err = res.Decode(&profile)

	if err != nil {
		t.Fatalf("Had trouble finding profile in database: %v", err)
		return
	}

	// it's not crucial we check every field in profile. The profile E2E tests should be doing that
	if expected_res.TotalPoints != profile.Points {
		t.Fatalf("Wrong amount of points in profile database. Expected %v, got %v", expected_res.TotalPoints, profile.Points)
	}
}

func TestStaffCheckinInvalidEvent(t *testing.T) {
	CreateEvents()
	defer ClearEvents()
	CreateProfile()
	defer ClearProfiles()
	GenerateRsvpData()
	defer ClearRsvps()

	req := models.StaffCheckinRequest{
		EventID:   "bogus",
		UserToken: GenerateValidUserToken(t),
	}
	received_res := models.CheckinResponse{}
	response, err := staff_client.New().Post("/event/staff/checkin/").BodyJSON(req).ReceiveSuccess(&received_res)
	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_res := models.CheckinResponse{
		NewPoints:   -1,
		TotalPoints: -1,
		Status:      "InvalidEventId",
	}

	if !reflect.DeepEqual(received_res, expected_res) {
		t.Fatalf("Wrong result received. Expected %v, got %v", expected_res, received_res)
	}

	res := client.Database(profile_db_name).Collection("profiles").FindOne(context.Background(), bson.M{"id": TEST_PROFILE_ID})

	profile := profile_models.Profile{}
	err = res.Decode(&profile)

	if err != nil {
		t.Fatalf("Had trouble finding profile in database: %v", err)
		return
	}

	// it's not crucial we check every field in profile. The profile E2E tests should be doing that
	expected_points := 0
	if expected_points != profile.Points {
		t.Fatalf("Wrong amount of points in profile database. Expected %v, got %v", expected_points, profile.Points)
	}

	// Need to make sure profile attendance was not stored
	res = client.Database(profile_db_name).
		Collection("profileattendance").
		FindOne(context.Background(), bson.M{"id": TEST_PROFILE_ID})

	profile_attendance := profile_models.AttendanceTracker{}
	err = res.Decode(&profile_attendance)

	if err == nil {
		t.Fatalf("Found stored profile attendance when there should be none: %v", profile_attendance)
		return
	}
}

func TestStaffCheckinBadUserTokenInvalidToken(t *testing.T) {
	CreateEvents()
	defer ClearEvents()
	CreateProfile()
	defer ClearProfiles()
	GenerateRsvpData()
	defer ClearRsvps()

	req := models.StaffCheckinRequest{
		EventID:   TEST_EVENT_1_ID,
		UserToken: GenerateInvalidUserToken(t),
	}
	received_res := models.CheckinResponse{}
	response, err := staff_client.New().Post("/event/staff/checkin/").BodyJSON(req).ReceiveSuccess(&received_res)
	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_res := models.CheckinResponse{
		NewPoints:   -1,
		TotalPoints: -1,
		Status:      "BadUserToken",
	}

	if !reflect.DeepEqual(received_res, expected_res) {
		t.Fatalf("Wrong result received. Expected %v, got %v", expected_res, received_res)
	}

	res := client.Database(profile_db_name).Collection("profiles").FindOne(context.Background(), bson.M{"id": TEST_PROFILE_ID})

	profile := profile_models.Profile{}
	err = res.Decode(&profile)

	if err != nil {
		t.Fatalf("Had trouble finding profile in database: %v", err)
		return
	}

	// it's not crucial we check every field. The profile E2E tests should be doing that
	expected_points := 0
	if expected_points != profile.Points {
		t.Fatalf("Wrong amount of points in profile database. Expected %v, got %v", expected_points, profile.Points)
	}
}

func TestStaffCheckinBadUserTokenExpiredToken(t *testing.T) {
	CreateEvents()
	defer ClearEvents()
	CreateProfile()
	defer ClearProfiles()

	req := models.StaffCheckinRequest{
		EventID:   TEST_EVENT_1_ID,
		UserToken: GenerateExpiredUserToken(t),
	}
	received_res := models.CheckinResponse{}
	response, err := staff_client.New().Post("/event/staff/checkin/").BodyJSON(req).ReceiveSuccess(&received_res)
	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_res := models.CheckinResponse{
		NewPoints:   -1,
		TotalPoints: -1,
		Status:      "BadUserToken",
	}

	if !reflect.DeepEqual(received_res, expected_res) {
		t.Fatalf("Wrong result received. Expected %v, got %v", expected_res, received_res)
	}

	res := client.Database(profile_db_name).Collection("profiles").FindOne(context.Background(), bson.M{"id": TEST_PROFILE_ID})

	profile := profile_models.Profile{}
	err = res.Decode(&profile)

	if err != nil {
		t.Fatalf("Had trouble finding profile in database: %v", err)
		return
	}

	// it's not crucial we check every field. The profile E2E tests should be doing that
	expected_points := 0
	if expected_points != profile.Points {
		t.Fatalf("Wrong amount of points in profile database. Expected %v, got %v", expected_points, profile.Points)
	}
}

func TestStaffCheckinAlreadyCheckedIn(t *testing.T) {
	CreateEvents()
	defer ClearEvents()
	CreateProfile()
	defer ClearProfiles()
	GenerateRsvpData()
	defer ClearRsvps()

	client.Database(profile_db_name).Collection("profileattendance").UpdateOne(
		context.Background(),
		bson.M{"id": TEST_PROFILE_ID},
		bson.M{"$addToSet": bson.M{
			"events": TEST_EVENT_1_ID,
		}},
		options.Update().SetUpsert(true),
	)

	req := models.StaffCheckinRequest{
		EventID:   TEST_EVENT_1_ID,
		UserToken: GenerateValidUserToken(t),
	}
	received_res := models.CheckinResponse{}
	response, err := staff_client.New().Post("/event/staff/checkin/").BodyJSON(req).ReceiveSuccess(&received_res)
	if err != nil {
		t.Fatal("Unable to make request: ", err)
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Errorf("Request returned HTTP error %d", response.StatusCode)
		// return
	}

	expected_res := models.CheckinResponse{
		NewPoints:   -1,
		TotalPoints: -1,
		Status:      "AlreadyCheckedIn",
		RsvpData:    rsvp_data.Data,
	}

	if !reflect.DeepEqual(received_res, expected_res) {
		t.Fatalf("Wrong result received. Expected %v, got %v", expected_res, received_res)
	}

	res := client.Database(profile_db_name).Collection("profiles").FindOne(context.Background(), bson.M{"id": TEST_PROFILE_ID})

	profile := profile_models.Profile{}
	err = res.Decode(&profile)

	if err != nil {
		t.Fatalf("Had trouble finding profile in database: %v", err)
		return
	}

	// it's not crucial we check every field. The profile E2E tests should be doing that
	expected_points := 0
	if expected_points != profile.Points {
		t.Fatalf("Wrong amount of points in profile database. Expected %v, got %v", expected_points, profile.Points)
	}
}
