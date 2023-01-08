package tests

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/HackIllinois/api/services/event/models"
	profile_models "github.com/HackIllinois/api/services/profile/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestCheckinNormal(t *testing.T) {
	CreateEvents()
	defer ClearEvents()
	CreateProfile()
	defer ClearProfiles()

	req := models.CheckinRequest{
		Code: TEST_EVENT_1_CODE,
	}
	received_res := models.CheckinResponse{}
	response, err := staff_client.New().Post("/event/checkin/").BodyJSON(req).ReceiveSuccess(&received_res)

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
		TotalPoints: 50,
		Status:      "Success",
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

func TestCheckinAddToExistingPoints(t *testing.T) {
	CreateEvents()
	defer ClearEvents()
	CreateProfile()
	defer ClearProfiles()

	client.Database(profile_db_name).Collection("profiles").UpdateOne(
		context.Background(),
		bson.M{"id": TEST_PROFILE_ID},
		bson.M{"$set": bson.M{
			"points": 19,
		}},
	)

	req := models.CheckinRequest{
		Code: TEST_EVENT_1_CODE,
	}
	received_res := models.CheckinResponse{}
	response, err := staff_client.New().Post("/event/checkin/").BodyJSON(req).ReceiveSuccess(&received_res)

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

func TestCheckinInvalidCode(t *testing.T) {
	CreateEvents()
	defer ClearEvents()
	CreateProfile()
	defer ClearProfiles()

	req := models.CheckinRequest{
		Code: "wrongcode",
	}
	received_res := models.CheckinResponse{}
	response, err := staff_client.New().Post("/event/checkin/").BodyJSON(req).ReceiveSuccess(&received_res)

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
		Status:      "InvalidCode",
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

func TestCheckinExpiredOrProspective(t *testing.T) {
	CreateEvents()
	defer ClearEvents()
	CreateProfile()
	defer ClearProfiles()

	req := models.CheckinRequest{
		Code: TEST_EVENT_2_CODE,
	}
	received_res := models.CheckinResponse{}
	response, err := staff_client.New().Post("/event/checkin/").BodyJSON(req).ReceiveSuccess(&received_res)

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
		Status:      "ExpiredOrProspective",
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

	// Need to make sure profile attendance was not stored
	res = client.Database(profile_db_name).Collection("profileattendance").FindOne(context.Background(), bson.M{"id": TEST_PROFILE_ID})

	profile_attendance := profile_models.AttendanceTracker{}
	err = res.Decode(&profile_attendance)

	if err == nil {
		t.Fatalf("Found stored profile attendance when there should be none: %v", profile_attendance)
		return
	}
}

func TestCheckinAlreadyCheckedIn(t *testing.T) {
	CreateEvents()
	defer ClearEvents()
	CreateProfile()
	defer ClearProfiles()

	client.Database(profile_db_name).Collection("profileattendance").UpdateOne(
		context.Background(),
		bson.M{"id": TEST_PROFILE_ID},
		bson.M{"$addToSet": bson.M{
			"events": TEST_EVENT_1_ID,
		}},
		options.Update().SetUpsert(true),
	)

	req := models.CheckinRequest{
		Code: TEST_EVENT_1_CODE,
	}
	received_res := models.CheckinResponse{}
	response, err := staff_client.New().Post("/event/checkin/").BodyJSON(req).ReceiveSuccess(&received_res)

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
