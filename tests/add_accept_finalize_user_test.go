package tests

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"testing"

	decision_models "github.com/HackIllinois/api/services/decision/models"
	user_models "github.com/HackIllinois/api/services/user/models"
	"github.com/dghubble/sling"
)

var admin_client *sling.Sling

func GetAdminClient() *sling.Sling {
	// First, get an admin authorization token by running `make setup`.
	path, err := os.Getwd()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	cmd := exec.Command("make", "setup")
	cmd.Dir = filepath.Dir(path)
	out, err := cmd.Output()

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	out_lines := strings.Split(string(out[:]), "\n")
	admin_token := out_lines[len(out_lines)-3]

	fmt.Printf(admin_token)

	return sling.New().Base("http://localhost:8000").Client(nil).Add("Authorization", admin_token)
}

func TestMain(m *testing.M) {
	fmt.Printf("in testmain")
	admin_client = GetAdminClient()

	return_code := m.Run()
	os.Exit(return_code)
}

func TestAddApproveFinalizeUsers(t *testing.T) {
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
