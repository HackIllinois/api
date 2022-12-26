package tests

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"testing"

	"github.com/HackIllinois/api/common/configloader"
	decision_models "github.com/HackIllinois/api/services/decision/models"
	user_models "github.com/HackIllinois/api/services/user/models"
	"github.com/HackIllinois/api/tests/common"
	"github.com/dghubble/sling"
	"go.mongodb.org/mongo-driver/mongo"
)

var admin_client *sling.Sling
var client *mongo.Client
var user_db_name string
var decision_db_name string

func TestMain(m *testing.M) {
	cfg, err := configloader.Load(os.Getenv("HI_CONFIG"))

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	admin_client = common.GetSlingClient("Admin")

	client = common.GetLocalMongoSession()

	user_db_name, err = cfg.Get("USER_DB_NAME")
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	decision_db_name, err = cfg.Get("DECISION_DB_NAME")
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	DropDatabases()

	return_code := m.Run()
	os.Exit(return_code)
}

func DropDatabases() {
	client.Database(user_db_name).Drop(context.Background())
	client.Database(decision_db_name).Drop(context.Background())
}

func TestAddApproveFinalizeUsers(t *testing.T) {
	defer DropDatabases()
	// Make 10 random users
	for i := 0; i < 10; i++ {
		userinfo := user_models.UserInfo{
			ID:        "github000001" + strconv.Itoa(i),
			Username:  "test",
			FirstName: "ExampleFirstName",
			LastName:  "ExampleLastName",
			Email:     "test@gmail.com",
		}

		received_userinfo := user_models.UserInfo{}
		response, err := admin_client.New().Post("user/").BodyJSON(userinfo).ReceiveSuccess(&received_userinfo)
		if err != nil {
			t.Errorf("Unable to make request")
		}
		if response.StatusCode != 200 {
			t.Errorf(response.Status)
		}

		expected_userinfo := user_models.UserInfo{
			ID:        "github000001" + strconv.Itoa(i),
			Username:  "test",
			FirstName: "ExampleFirstName",
			LastName:  "ExampleLastName",
			Email:     "test@gmail.com",
		}

		if !reflect.DeepEqual(received_userinfo, expected_userinfo) {
			t.Errorf("Wrong user info. Expected %v, got %v", expected_userinfo, received_userinfo)
		}
	}

	// Accept them
	for i := 0; i < 10; i++ {

		type DecisionRequest struct {
			ID     string `json:"id"`
			Status string `json:"status"`
			Wave   int    `json:"wave"`
		}

		decision_request := DecisionRequest{
			ID:     "github000001" + strconv.Itoa(i),
			Status: "ACCEPTED",
			Wave:   1,
		}

		received_decision_history := decision_models.DecisionHistory{}
		response, err := admin_client.New().Post("decision/").BodyJSON(decision_request).ReceiveSuccess(&received_decision_history)
		if err != nil {
			t.Errorf("Unable to make request")
		}
		if response.StatusCode != 200 {
			t.Errorf(response.Status)
		}

		// The current status should be "ACCEPTED"
		if received_decision_history.Status != "ACCEPTED" {
			t.Errorf("Failed to accept user. Got %v", received_decision_history)
		}
	}

	// Finalize them
	for i := 0; i < 10; i++ {
		finalize_request := decision_models.DecisionFinalized{
			ID:        "github000001" + strconv.Itoa(i),
			Finalized: true,
		}

		received_decision_history := decision_models.DecisionHistory{}
		response, err := admin_client.New().Post("decision/finalize/").BodyJSON(finalize_request).ReceiveSuccess(&received_decision_history)
		if err != nil {
			t.Errorf("Unable to make request")
		}
		if response.StatusCode != 200 {
			t.Errorf(response.Status)
		}

		// The current status should be "ACCEPTED", and the user should be finalized
		if received_decision_history.Status != "ACCEPTED" || !received_decision_history.Finalized {
			t.Errorf("Failed to finalize user. Got %v", received_decision_history)
		}
	}
}
