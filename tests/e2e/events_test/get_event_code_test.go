package tests

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/services/event/models"
)

func TestGetEventCodeNormal(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	recieved_code := models.EventCode{}
	id := TEST_EVENT_1_ID
	response, err := staff_client.New().Get(fmt.Sprintf("/event/code/%s/", id)).ReceiveSuccess(&recieved_code)

	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d", response.StatusCode)
		return
	}

	expected_code := models.EventCode{
		ID:         TEST_EVENT_1_ID,
		Code:       TEST_EVENT_1_CODE,
		Expiration: current_unix_time + 60000,
	}

	if !reflect.DeepEqual(recieved_code, expected_code) {
		t.Fatalf("Wrong event code info. Expected %v, got %v", expected_code, recieved_code)
	}
}

func TestGetEventCodeWrongEventId(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	id := "thisisclearlythewrongid"
	api_err := errors.ApiError{}
	response, err := staff_client.New().Get(fmt.Sprintf("/event/code/%s/", id)).Receive(nil, &api_err)

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
		Message:  "Failed to receive event code information from database",
		RawError: "Error: NOT_FOUND",
	}

	if !reflect.DeepEqual(expected_error, api_err) {
		t.Fatalf("Wrong error response received. Expected %v, got %v", expected_error, api_err)
	}
}

func TestGetEventCodeForbidden(t *testing.T) {
	CreateEvents()
	defer ClearEvents()

	id := TEST_EVENT_1_ID
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
