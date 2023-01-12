package tests

import (
	"context"
	"fmt"
	"os"
	"testing"

	"time"

	"github.com/HackIllinois/api/common/configloader"
	"github.com/HackIllinois/api/services/decision/models"
	"github.com/HackIllinois/api/tests/common"
	"github.com/dghubble/sling"
	"go.mongodb.org/mongo-driver/mongo"
)

var current_unix_time int64
var unauthenticated_client *sling.Sling
var user_client *sling.Sling
var admin_client *sling.Sling
var client *mongo.Client
var decision_db_name string

func TestMain(m *testing.M) {
	cfg, err := configloader.Load(os.Getenv("HI_CONFIG"))

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	unauthenticated_client = common.GetSlingClient("Unauthenticated")
	user_client = common.GetSlingClient("User")
	admin_client = common.GetSlingClient("Admin")

	client = common.GetLocalMongoSession()

	decision_db_name, err = cfg.Get("DECISION_DB_NAME")
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	DropDatabase()

	current_unix_time = time.Now().Unix()

	return_code := m.Run()
	os.Exit(return_code)
}

func DropDatabase() {
	client.Database(decision_db_name).Drop(context.Background())
}

func GetGeneratedDecision(id string) models.Decision {
	return models.Decision{
		Finalized: false,
		ID:        id,
		Status:    "PENDING",
		Wave:      1,
		Reviewer:  "bob",
		Timestamp: current_unix_time,
		ExpiresAt: current_unix_time + 60000,
	}
}

func DecisionToDecisionHistory(decision models.Decision) models.DecisionHistory {
	return models.DecisionHistory{
		Finalized: decision.Finalized,
		ID:        decision.ID,
		Status:    decision.Status,
		Wave:      decision.Wave,
		Reviewer:  decision.Reviewer,
		Timestamp: decision.Timestamp,
		ExpiresAt: decision.ExpiresAt,
		History:   []models.Decision{decision},
	}
}

func AddGeneratedDecision(id string) models.DecisionHistory {
	decision_info := DecisionToDecisionHistory(GetGeneratedDecision(id))

	client.Database(decision_db_name).Collection("decision").InsertOne(context.Background(), &decision_info)

	return decision_info
}
