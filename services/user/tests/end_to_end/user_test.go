package end_to_end

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/HackIllinois/api/services/user/config"
	"github.com/HackIllinois/api/services/user/models"
	"github.com/HackIllinois/api/services/user/service"
	"github.com/dghubble/sling"
)

var api_base *sling.Sling

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

	api_base = sling.New().Base("http://localhost:8000").Client(nil).Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImxvY2FsYWRtaW5AbG9jYWwubG9jYWwiLCJleHAiOjI1MjQ2MDgwMDAsImlkIjoibG9jYWxhZG1pbiIsInJvbGVzIjpbIkFkbWluIiwiVXNlciJdfQ.ZZItxIHsr-8XtZvGcQzxwlaIPWokMwwU2bio-QcAA9o")

	return_code := m.Run()

	os.Exit(return_code)
}

func TestGetLocalAdminUser(t *testing.T) {
	got := models.UserInfo{}
	response, err := api_base.Path("user/").ReceiveSuccess(&got)
	if err != nil {
		t.Errorf("Unable to make request")
	}
	if response.StatusCode != 200 {
		t.Errorf(response.Status)
	}

	want := models.UserInfo{
		ID:        "localadmin",
		Username:  "localadmin",
		FirstName: "local",
		LastName:  "admin",
		Email:     "localadmin@local.local",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Wrong user info. Expected %v, got %v", want, got)
	}
}
