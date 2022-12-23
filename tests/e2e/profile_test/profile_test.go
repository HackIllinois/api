package tests

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/HackIllinois/api/common/configloader"
	"github.com/HackIllinois/api/tests/common"
	"github.com/dghubble/sling"
	"go.mongodb.org/mongo-driver/mongo"
)

var admin_client *sling.Sling
var client *mongo.Client
var profile_db_name string
var unauthenticated_client *sling.Sling

func TestMain(m *testing.M) {

	cfg, err := configloader.Load(os.Getenv("HI_CONFIG"))

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	admin_client = common.GetSlingClient("Admin")
	unauthenticated_client = common.GetSlingClient("Unauthenticated")

	client = common.GetLocalMongoSession()

	profile_db_name, err = cfg.Get("PROFILE_DB_NAME")
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}
	client.Database(profile_db_name).Drop(context.Background())

	return_code := m.Run()
	os.Exit(return_code)
}
