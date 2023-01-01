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

	expected_code := models.EventCode{
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
