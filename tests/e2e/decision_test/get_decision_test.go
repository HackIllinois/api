package tests

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/HackIllinois/api/services/decision/models"
)

func TestGetDecision(t *testing.T) {
	decision_info := models.DecisionHistory{
		Finalized: false,
		ID:        "12345",
		Status:    "PENDING",
		Wave:      1,
		Reviewer:  "bob",
		Timestamp: 123,
		ExpiresAt: 1234,
		History: []models.Decision{
			{
				Finalized: false,
				ID:        "12345",
				Status:    "PENDING",
				Wave:      1,
				Reviewer:  "bob",
				Timestamp: 123,
				ExpiresAt: 1234,
			},
		},
	}

	client.Database(decision_db_name).Collection("profile").InsertOne(context.Background(), &decision_info)

	endpoint_address := fmt.Sprintf("/decision/%s/", "12345")

	received_decision := models.DecisionHistory{}
	response, err := admin_client.New().Get(endpoint_address).ReceiveSuccess(&received_decision)

	if err != nil {
		t.Errorf("Unable to make request")
	}
	if response.StatusCode != 200 {
		t.Errorf("Request returned HTTP error %s", response.Status)
	}
	if !reflect.DeepEqual(decision_info, received_decision) {
		t.Errorf("Wrong decision info. Expected %v, got %v", decision_info, received_decision)
	}
}
