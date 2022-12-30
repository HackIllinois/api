package tests

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/HackIllinois/api/services/event/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestCheckinNormal(t *testing.T) {
	CreateEvents()
	defer ClearEvents()
	CreateProfile()
	defer ClearProfiles()

	req := models.CheckinRequest{
		Code: "123456",
	}
	received_res := models.CheckinResult{}
	response, err := staff_client.New().Post("/event/checkin/").BodyJSON(req).ReceiveSuccess(&received_res)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_res := models.CheckinResult{
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

	req := models.CheckinRequest{
		Code: "123456",
	}
	received_res := models.CheckinResult{}
	response, err := staff_client.New().Post("/event/checkin/").BodyJSON(req).ReceiveSuccess(&received_res)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_res := models.CheckinResult{
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

	req := models.CheckinRequest{
		Code: "wrongcode",
	}
	received_res := models.CheckinResult{}
	response, err := staff_client.New().Post("/event/checkin/").BodyJSON(req).Receive(nil, &received_res)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusNotFound {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_res := models.CheckinResult{
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

	req := models.CheckinRequest{
		Code: "abcdef",
	}
	received_res := models.CheckinResult{}
	response, err := staff_client.New().Post("/event/checkin/").BodyJSON(req).Receive(nil, &received_res)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusGone {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_res := models.CheckinResult{
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

	req := models.CheckinRequest{
		Code: "123456",
	}
	received_res := models.CheckinResult{}
	response, err := staff_client.New().Post("/event/checkin/").BodyJSON(req).Receive(&received_res, &received_res)

	if err != nil {
		t.Fatal("Unable to make request: ", err)
		return
	}
	if response.StatusCode != http.StatusConflict {
		t.Errorf("Request returned HTTP error %d", response.StatusCode)
		// return
	}

	expected_res := models.CheckinResult{
		NewPoints:   -1,
		TotalPoints: -1,
		Status:      "AlreadyCheckedIn",
	}

	if !reflect.DeepEqual(received_res, expected_res) {
		t.Fatalf("Wrong result received. Expected %v, got %v", expected_res, received_res)
	}
}
