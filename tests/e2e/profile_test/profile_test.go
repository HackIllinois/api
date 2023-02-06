package tests

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/HackIllinois/api/common/configloader"
	profile_models "github.com/HackIllinois/api/services/profile/models"
	"github.com/HackIllinois/api/tests/common"
	"github.com/dghubble/sling"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	admin_client           *sling.Sling
	client                 *mongo.Client
	profile_db_name        string
	unauthenticated_client *sling.Sling
)

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
	DropDatabases()

	return_code := m.Run()
	os.Exit(return_code)
}

func DropDatabases() {
	client.Database(profile_db_name).Drop(context.Background())
}

func CheckDatabaseProfileNotFound(t *testing.T, filter bson.M) {
	actual_profile_db := profile_models.Profile{}
	res := client.Database(profile_db_name).Collection("profiles").FindOne(context.Background(), filter)

	err := res.Decode(&actual_profile_db)

	if err == nil {
		t.Errorf("Profile was found: %v", actual_profile_db)
	} else if err != mongo.ErrNoDocuments {
		t.Fatalf("Failed when decoding profile: %v", err)
	}
}
