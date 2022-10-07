package tests

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/services/decision/config"
	"github.com/HackIllinois/api/services/decision/models"
	"github.com/HackIllinois/api/services/decision/service"
)

var db database.Database

func TestMain(m *testing.M) {
	err := config.Initialize()

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)

	}

	err = service.Initialize()

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	db, err = database.InitDatabase(config.DECISION_DB_HOST, config.DECISION_DB_NAME)

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	return_code := m.Run()

	os.Exit(return_code)
}

/*
	Initialize databse with test decision info
*/
func SetupTestDB(t *testing.T) {
	err := db.Insert("decision", &models.DecisionHistory{
		Finalized: false,
		ID:        "testid",
		Status:    "PENDING",
		Wave:      0,
		Reviewer:  "reviewerid",
		Timestamp: 1,
		ExpiresAt: 5,
		History: []models.Decision{
			{
				Finalized: false,
				ID:        "testid",
				Status:    "PENDING",
				Wave:      0,
				Reviewer:  "reviewerid",
				Timestamp: 1,
				ExpiresAt: 5,
			},
		},
	}, nil)

	if err != nil {
		t.Fatal(err)
	}
}

/*
	Drop test db
*/
func CleanupTestDB(t *testing.T) {
	err := db.DropDatabase(nil)

	if err != nil {
		t.Fatal(err)
	}
}

/*
	Service level test for getting decision info from db
*/
func TestGetDecisionService(t *testing.T) {
	SetupTestDB(t)

	decision, err := service.GetDecision("testid")

	if err != nil {
		t.Fatal(err)
	}

	expected_decision := &models.DecisionHistory{
		Finalized: false,
		ID:        "testid",
		Status:    "PENDING",
		Wave:      0,
		Reviewer:  "reviewerid",
		Timestamp: 1,
		ExpiresAt: 5,
		History: []models.Decision{
			{
				Finalized: false,
				ID:        "testid",
				Status:    "PENDING",
				Wave:      0,
				Reviewer:  "reviewerid",
				Timestamp: 1,
				ExpiresAt: 5,
			},
		},
	}

	if !reflect.DeepEqual(decision, expected_decision) {
		t.Errorf("Wrong decision info. Expected %v, got %v", expected_decision, decision)
	}

	CleanupTestDB(t)
}

/*
	Service level test for updating decision in the db
*/
func TestUpdateDecisionService(t *testing.T) {
	SetupTestDB(t)

	err := service.UpdateDecision("testid", models.Decision{
		Finalized: false,
		ID:        "testid",
		Status:    "ACCEPTED",
		Wave:      1,
		Reviewer:  "reviewerid",
		Timestamp: 2,
		ExpiresAt: 7,
	})

	if err != nil {
		t.Fatal(err)
	}

	decision, err := service.GetDecision("testid")

	if err != nil {
		t.Fatal(err)
	}

	expected_decision := &models.DecisionHistory{
		Finalized: false,
		ID:        "testid",
		Status:    "ACCEPTED",
		Wave:      1,
		Reviewer:  "reviewerid",
		Timestamp: 2,
		ExpiresAt: 7,
		History: []models.Decision{
			{
				Finalized: false,
				ID:        "testid",
				Status:    "PENDING",
				Wave:      0,
				Reviewer:  "reviewerid",
				Timestamp: 1,
				ExpiresAt: 5,
			},
			{
				Finalized: false,
				ID:        "testid",
				Status:    "ACCEPTED",
				Wave:      1,
				Reviewer:  "reviewerid",
				Timestamp: 2,
				ExpiresAt: 7,
			},
		},
	}

	if !reflect.DeepEqual(decision, expected_decision) {
		t.Errorf("Wrong decision info. Expected %v, got %v", expected_decision, decision)
	}

	CleanupTestDB(t)
}

/*
	Service level test for getting filtered decision info from db
*/
func TestGetFilteredDecisionsService(t *testing.T) {
	SetupTestDB(t)

	decision2 := models.DecisionHistory{
		ID:        "testid2",
		Status:    "PENDING",
		Wave:      1,
		Reviewer:  "reviewerid",
		Timestamp: 2,
		ExpiresAt: 7,
		History: []models.Decision{
			{
				ID:        "testid2",
				Status:    "PENDING",
				Wave:      1,
				Reviewer:  "reviewerid",
				Timestamp: 2,
				ExpiresAt: 7,
			},
		},
	}
	err := db.Insert("decision", &decision2, nil)

	if err != nil {
		t.Fatal(err)
	}

	parameters := map[string][]string{
		"id":   {"testid2"},
		"wave": {"1"},
	}
	decisions, err := service.GetFilteredDecisions(parameters)
	if err != nil {
		t.Fatal(err)
	}

	expected_decisions := models.FilteredDecisions{
		[]models.DecisionHistory{
			decision2,
		},
	}

	if !reflect.DeepEqual(decisions, &expected_decisions) {
		t.Errorf("Wrong decision info. Expected %v, got %v", expected_decisions, decisions)
	}

	CleanupTestDB(t)
}
