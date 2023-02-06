package tests

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/HackIllinois/api/common/configloader"
	user_models "github.com/HackIllinois/api/services/user/models"
	"github.com/HackIllinois/api/tests/common"
	"github.com/dghubble/sling"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var user_client *sling.Sling
var client *mongo.Client

var user_db_name string

var TOKEN_SECRET []byte

const TEST_USER_ID = "localadmin"

var current_unix_time = time.Now().Unix()

func TestMain(m *testing.M) {
	cfg, err := configloader.Load(os.Getenv("HI_CONFIG"))

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	user_client = common.GetSlingClient("User")

	client = common.GetLocalMongoSession()

	user_db_name, err = cfg.Get("USER_DB_NAME")
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}
	ResetDatabase()

	token_secret_string, err := cfg.Get("TOKEN_SECRET")
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	TOKEN_SECRET = []byte(token_secret_string)

	return_code := m.Run()
	ResetDatabase()
	os.Exit(return_code)
}

func ResetDatabase() {
	client.Database(user_db_name).Drop(context.Background())
}

func CreateUserInfo() {
	user_info := user_models.UserInfo{
		ID:        TEST_USER_ID,
		Username:  "Bob",
		FirstName: "Bob",
		LastName:  "Ross",
		Email:     "bob@ross.com",
	}

	client.Database(user_db_name).Collection("info").InsertOne(context.Background(), user_info)
}

func ClearUserInfo() {
	client.Database(user_db_name).Collection("info").DeleteMany(context.Background(), bson.D{})
}
