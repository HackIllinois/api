package tests

import (
	"reflect"
	"testing"

	"github.com/ethan-lord/api/common/database"
	"github.com/ethan-lord/api/services/decision/config"
	"github.com/ethan-lord/api/services/decision/models"
	"github.com/ethan-lord/api/services/decision/service"
)

var db database.MongoDatabase

func init() {
	db_connection, err := database.InitMongoDatabase(config.DECISION_DB_HOST, config.DECISION_DB_NAME)

	if err != nil {
		panic(err)
	}

	db = db_connection
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
		History: []models.Decision{
			models.Decision{
				Finalized: false,
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
	Drop test db
*/
func CleanupTestDB(t *testing.T) {
	session := db.GetSession()
	defer session.Close()

	err := session.DB(config.DECISION_DB_NAME).DropDatabase()

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
		History: []models.Decision{
			models.Decision{
				Finalized: false,
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
		History: []models.Decision{
			models.Decision{
				Finalized: false,
				ID:        "testid",
				Status:    "PENDING",
				Wave:      0,
				Reviewer:  "reviewerid",
				Timestamp: 1,
			},
			models.Decision{
				Finalized: false,
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
		History: []models.Decision{
			models.Decision{
				ID:        "testid2",
				Status:    "PENDING",
				Wave:      1,
				Reviewer:  "reviewerid",
				Timestamp: 2,
			},
		},
	}
	err := db.Insert("decision", &decision2)

	if err != nil {
		t.Fatal(err)
	}

	parameters := map[string][]string{
		"id":   []string{"testid2"},
		"wave": []string{"1"},
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
