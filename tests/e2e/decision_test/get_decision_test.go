package tests

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/HackIllinois/api/services/decision/models"
)

func TestGetDecision(t *testing.T) {
	defer DropDatabase()

	decision_info := AddGeneratedDecision("123")

	client.Database(decision_db_name).Collection("profile").InsertOne(context.Background(), &decision_info)

	endpoint_address := fmt.Sprintf("/decision/%s/", decision_info.ID)

	received_decision := models.DecisionHistory{}
	response, err := admin_client.New().Get(endpoint_address).ReceiveSuccess(&received_decision)

	if err != nil {
		t.Fatalf("Unable to make request")
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %s", response.Status)
	}
	if !reflect.DeepEqual(decision_info, received_decision) {
		t.Fatalf("Wrong decision info. Expected %v, got %v", decision_info, received_decision)
	}
}

func TestGetDecisionNonExistent(t *testing.T) {
	defer DropDatabase()
	endpoint_address := fmt.Sprintf("/decision/%s/", "some_invalid_id")

	received_decision := models.DecisionHistory{}
	response, err := admin_client.New().Get(endpoint_address).ReceiveSuccess(&received_decision)

	if err != nil {
		t.Fatalf("Unable to make request")
	}
	if response.StatusCode != http.StatusInternalServerError {
		t.Fatalf("Request returned HTTP %s", response.Status)
	}
}

func TestGetDecisionUnauthenticated(t *testing.T) {
	defer DropDatabase()
	decision_info := AddGeneratedDecision("123")

	endpoint_address := fmt.Sprintf("/decision/%s/", decision_info.ID)

	received_decision := models.DecisionHistory{}
	response, err := unauthenticated_client.New().Get(endpoint_address).ReceiveSuccess(&received_decision)

	if err != nil {
		t.Fatalf("Unable to make request")
	}
	if response.StatusCode != http.StatusForbidden {
		t.Fatalf("Request returned HTTP %s", response.Status)
	}
}

func TestGetDecisionAnotherUser(t *testing.T) {
	defer DropDatabase()
	decision_info := AddGeneratedDecision("123")

	endpoint_address := fmt.Sprintf("/decision/%s/", decision_info.ID)

	received_decision := models.DecisionHistory{}
	response, err := user_client.New().Get(endpoint_address).ReceiveSuccess(&received_decision)

	if err != nil {
		t.Fatalf("Unable to make request")
	}
	if response.StatusCode != http.StatusForbidden {
		t.Fatalf("Request returned HTTP %s", response.Status)
	}
}
