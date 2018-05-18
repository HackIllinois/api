package tests

import (
	"github.com/HackIllinois/api-decision/config"
	"github.com/HackIllinois/api-decision/database"
	"github.com/HackIllinois/api-decision/models"
	"github.com/HackIllinois/api-decision/service"
	"reflect"
	"testing"
)

/*
	Initialize databse with test decision info
*/
func SetupTestDB(t *testing.T) {
	err := database.Insert("decision", &models.DecisionHistory{
		ID:        "testid",
		Status:    "PENDING",
		Wave:      0,
		Reviewer:  "reviewerid",
		Timestamp: 1,
		History: []models.Decision{
			models.Decision{
				ID:        "testid",
				Status:    "PENDING",
				Wave:      0,
				Reviewer:  "reviewerid",
				Timestamp: 1,
			},
		},
	})

	if err != nil {
		t.Fatal(err)
	}
}

/*
	Drop test database
*/
func CleanupTestDB(t *testing.T) {
	session := database.GetSession()
	defer session.Close()

	err := session.DB(config.DECISION_DB_NAME).DropDatabase()

	if err != nil {
		t.Fatal(err)
	}
}

/*
	Service level test for getting decision info from database
*/
func TestGetDecisionService(t *testing.T) {
	SetupTestDB(t)

	decision, err := service.GetDecision("testid")

	if err != nil {
		t.Fatal(err)
	}

	expected_decision := &models.DecisionHistory{
		ID:        "testid",
		Status:    "PENDING",
		Wave:      0,
		Reviewer:  "reviewerid",
		Timestamp: 1,
		History: []models.Decision{
			models.Decision{
				ID:        "testid",
				Status:    "PENDING",
				Wave:      0,
				Reviewer:  "reviewerid",
				Timestamp: 1,
			},
		},
	}

	if !reflect.DeepEqual(decision, expected_decision) {
		t.Errorf("Wrong decision info. Expected %v, got %v", expected_decision, decision)
	}

	CleanupTestDB(t)
}

/*
	Service level test for updating decision in the database
*/
func TestUpdateDecisionService(t *testing.T) {
	SetupTestDB(t)

	err := service.UpdateDecision("testid", models.Decision{
		ID:        "testid",
		Status:    "ACCEPTED",
		Wave:      1,
		Reviewer:  "reviewerid",
		Timestamp: 2,
	})

	if err != nil {
		t.Fatal(err)
	}

	decision, err := service.GetDecision("testid")

	if err != nil {
		t.Fatal(err)
	}

	expected_decision := &models.DecisionHistory{
		ID:        "testid",
		Status:    "ACCEPTED",
		Wave:      1,
		Reviewer:  "reviewerid",
		Timestamp: 2,
		History: []models.Decision{
			models.Decision{
				ID:        "testid",
				Status:    "PENDING",
				Wave:      0,
				Reviewer:  "reviewerid",
				Timestamp: 1,
			},
			models.Decision{
				ID:        "testid",
				Status:    "ACCEPTED",
				Wave:      1,
				Reviewer:  "reviewerid",
				Timestamp: 2,
			},
		},
	}

	if !reflect.DeepEqual(decision, expected_decision) {
		t.Errorf("Wrong decision info. Expected %v, got %v", expected_decision, decision)
	}

	CleanupTestDB(t)
}
