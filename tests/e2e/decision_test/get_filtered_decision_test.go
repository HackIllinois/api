package tests

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/HackIllinois/api/services/decision/models"
)

func TestGetFilteredDecision(t *testing.T) {
	defer DropDatabase()
	decision_info := AddGeneratedDecision("123")
	_ = AddGeneratedDecision("456")
	_ = AddGeneratedDecision("789")

	endpoint_address := fmt.Sprintf("/decision/filter/?id=%s", decision_info.ID)

	expected_filtered_decision := models.FilteredDecisions{
		Decisions: []models.DecisionHistory{decision_info},
	}

	received_filtered_decision := models.FilteredDecisions{}
	response, err := admin_client.New().Get(endpoint_address).ReceiveSuccess(&received_filtered_decision)

	if err != nil {
		t.Fatalf("Unable to make request")
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %s", response.Status)
	}
	if !reflect.DeepEqual(expected_filtered_decision, received_filtered_decision) {
		t.Fatalf("Wrong decision info. Expected %v, got %v", decision_info, received_filtered_decision)
	}
}

func TestGetFilteredDecisionEmpty(t *testing.T) {
	defer DropDatabase()
	endpoint_address := fmt.Sprintf("/decision/filter/?id=%s", "123")

	expected_filtered_decision := models.FilteredDecisions{}

	received_filtered_decision := models.FilteredDecisions{}
	response, err := admin_client.New().Get(endpoint_address).ReceiveSuccess(&received_filtered_decision)

	if err != nil {
		t.Fatalf("Unable to make request")
	}

	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %s", response.Status)
	}

	if !reflect.DeepEqual(expected_filtered_decision, received_filtered_decision) {
		t.Fatalf("Wrong decision info. Expected %v, got %v", expected_filtered_decision, received_filtered_decision)
	}
}

func TestGetFilteredDecisionBadArg(t *testing.T) {
	defer DropDatabase()
	endpoint_address := "/decision/filter/?someInvalidRequesterwefix!"

	received_filtered_decision := models.FilteredDecisions{}
	response, err := admin_client.New().Get(endpoint_address).ReceiveSuccess(&received_filtered_decision)

	if err != nil {
		t.Fatalf("Unable to make request")
	}

	if response.StatusCode != http.StatusInternalServerError {
		t.Fatalf("Request returned HTTP error %s", response.Status)
	}
}

func TestGetFilteredDecisionUnauthenticated(t *testing.T) {
	defer DropDatabase()
	decision_info := AddGeneratedDecision("123")

	endpoint_address := fmt.Sprintf("/decision/filter/?id=%s", decision_info.ID)

	received_filtered_decision := models.FilteredDecisions{}
	response, err := unauthenticated_client.New().Get(endpoint_address).ReceiveSuccess(&received_filtered_decision)

	if err != nil {
		t.Fatalf("Unable to make request")
	}
	if response.StatusCode != http.StatusForbidden {
		t.Fatalf("Request returned HTTP error %s", response.Status)
	}
}
