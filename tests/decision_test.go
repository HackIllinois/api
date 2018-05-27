package tests

import (
	"github.com/HackIllinois/api-commons/database"
	"github.com/HackIllinois/api-decision/config"
	"github.com/HackIllinois/api-decision/models"
	"github.com/HackIllinois/api-decision/service"
	"reflect"
	"testing"
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
	Service level test for updating decision in the db
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
