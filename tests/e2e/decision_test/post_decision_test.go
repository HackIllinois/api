package tests

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/HackIllinois/api/services/decision/models"
	"go.mongodb.org/mongo-driver/bson"
)

func TestPostDecision(t *testing.T) {
	defer DropDatabase()

	post_decision := GetGeneratedDecision("123")
	// Arbitrarily set values
	post_decision.Reviewer = "localadmin"
	post_decision.Wave = 0

	received_decision_history := models.DecisionHistory{}
	response, err := admin_client.New().Post("/decision/").BodyJSON(post_decision).ReceiveSuccess(&received_decision_history)
	expected_decision_history := DecisionToDecisionHistory(post_decision)

	// Since timestamp is determined by the server and not us, we need to updated our expected to what we get
	expected_decision_history.Timestamp = received_decision_history.Timestamp
	expected_decision_history.ExpiresAt = received_decision_history.ExpiresAt
	expected_decision_history.History[0].Timestamp = received_decision_history.Timestamp
	expected_decision_history.History[0].ExpiresAt = received_decision_history.ExpiresAt

	if err != nil {
		t.Fatalf("Unable to make request")
	}

	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %s", response.Status)
	}

	if !reflect.DeepEqual(expected_decision_history, received_decision_history) {
		t.Fatalf("Wrong decision info. Expected %v, got %v", expected_decision_history, received_decision_history)
	}

	read_from_database_decision_history := models.DecisionHistory{}

	res := client.Database(decision_db_name).Collection("decision").FindOne(context.Background(), bson.M{"id": post_decision.ID})

	err = res.Decode(&read_from_database_decision_history)

	if err != nil {
		t.Fatalf("Failed to get element: %v", err)
		return
	}

	if !reflect.DeepEqual(read_from_database_decision_history, expected_decision_history) {
		t.Fatalf("Wrong result received from database. Expected %v, got %v", read_from_database_decision_history, expected_decision_history)
	}
}

func TestPostDecisionBadArg(t *testing.T) {
	defer DropDatabase()

	received_decision_history := models.DecisionHistory{}

	bad_args := map[string]interface{}{
		"someRandomNonsense": "Buy cheese and bread for breakfast.",
	}
	response, err := admin_client.New().Post("/decision/").BodyJSON(bad_args).ReceiveSuccess(&received_decision_history)

	if err != nil {
		t.Fatalf("Unable to make request")
	}
	if response.StatusCode != http.StatusUnprocessableEntity {
		t.Fatalf("Request returned HTTP error %s", response.Status)
	}

	read_from_database_decision_history := models.DecisionHistory{}

	res := client.Database(decision_db_name).Collection("decision").FindOne(context.Background(), bson.D{})

	err = res.Decode(&read_from_database_decision_history)

	if err == nil {
		t.Fatalf("Request added element!")
	}
}

func TestPostDecisionUnauthenticated(t *testing.T) {
	defer DropDatabase()

	post_decision := GetGeneratedDecision("123")
	// Arbitrarily set values
	post_decision.Reviewer = "localadmin"
	post_decision.Wave = 0

	received_decision_history := models.DecisionHistory{}
	response, err := unauthenticated_client.New().Post("/decision/").BodyJSON(post_decision).ReceiveSuccess(&received_decision_history)

	if err != nil {
		t.Fatalf("Unable to make request")
	}
	if response.StatusCode != http.StatusForbidden {
		t.Fatalf("Expected forbidden status code, request returned HTTP %s", response.Status)
	}

	read_from_database_decision_history := models.DecisionHistory{}

	res := client.Database(decision_db_name).Collection("decision").FindOne(context.Background(), bson.D{})

	err = res.Decode(&read_from_database_decision_history)

	if err == nil {
		t.Fatalf("Request added element!")
	}
}
