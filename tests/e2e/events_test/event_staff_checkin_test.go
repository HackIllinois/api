package tests

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/services/event/models"
	profile_models "github.com/HackIllinois/api/services/profile/models"
	user_models "github.com/HackIllinois/api/services/user/models"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

func TestStaffCheckinAddToExistingPoints(t *testing.T) {
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
	res = client.Database(profile_db_name).Collection("profileattendance").FindOne(context.Background(), bson.M{"id": TEST_PROFILE_ID})

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

// This test + TestGetUserQR verify that an expired token will not checkin and
// that an expiring token will be generated, respectively.
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

// This is a e2e test that encompasses both GET /user/qr/ and /event/staff/checkin/.
// Basically, the prevent-everything-from-breaking-down-because-its-tested-separately test
func TestStaffCheckinFromGetUserQR(t *testing.T) {
	CreateEvents()
	defer ClearEvents()
	CreateProfile()
	defer ClearProfiles()
	CreateUserInfo()
	defer ClearUserInfo()

	received_qr_info_container := user_models.QrInfoContainer{}
	received_qr_info_error := errors.ApiError{}

	response, err := admin_client.New().Get(fmt.Sprintf("/user/qr/%s/", TEST_USER_ID)).Receive(&received_qr_info_container, &received_qr_info_error)

	if err != nil {
		t.Fatalf("Failed to parse qr info container: %v", err)
		return
	}

	if response.StatusCode != http.StatusOK {
		t.Fatalf("QR code request failed: %v, %v", response.Status, received_qr_info_error)
		return
	}

	u, err := url.Parse(received_qr_info_container.QrInfo)

	if err != nil {
		t.Fatalf("Failed to parse url: %v", err)
		return
	}

	query_map, err := url.ParseQuery(u.RawQuery)

	if err != nil {
		t.Fatalf("Failed to parse query string: %v", err)
		return
	}

	user_token := query_map.Get("userToken")

	if user_token == "" {
		t.Fatalf("Failed to get userToken from raw query: %v", query_map.Encode())
		return
	}

	req := models.StaffCheckinRequest{
		EventID:   TEST_EVENT_1_ID,
		UserToken: user_token,
	}

	received_res := models.CheckinResponse{}
	response, err = staff_client.New().Post("/event/staff/checkin/").BodyJSON(req).ReceiveSuccess(&received_res)

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
