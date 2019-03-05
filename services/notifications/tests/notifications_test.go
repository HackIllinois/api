package tests

import (
	"fmt"
	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/services/notifications/config"
	"github.com/HackIllinois/api/services/notifications/models"
	"github.com/HackIllinois/api/services/notifications/service"
	"os"
	"testing"
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

	db, err = database.InitDatabase(config.NOTIFICATIONS_DB_HOST, config.NOTIFICATIONS_DB_NAME)

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	return_code := m.Run()

	os.Exit(return_code)
}

/*
	Initialize db with a test topic and notification
*/
func SetupTestDB(t *testing.T) {
	topic := models.Topic{
		ID:      "test_topic",
		UserIDs: []string{"test_user"},
	}

	err := db.Insert("topics", &topic)

	if err != nil {
		t.Fatal(err)
	}

	notification := models.Notification{
		ID:    "test_id",
		Title: "test title",
		Body:  "test body",
		Topic: "test_topic",
		Time:  2000,
	}

	err = db.Insert("notifications", &notification)

	if err != nil {
		t.Fatal(err)
	}

	user := models.User{
		ID:      "test_userid",
		Devices: []string{"test_arn"},
	}

	err = db.Insert("devices", &user)

	if err != nil {
		t.Fatal(err)
	}
}

/*
	Drop test db
*/
func CleanupTestDB(t *testing.T) {
	err := db.DropDatabase()

	if err != nil {
		t.Fatal(err)
	}
}

/*
	Placeholder test for notifications service
*/
func TestPlaceholder(t *testing.T) {
	SetupTestDB(t)
	CleanupTestDB(t)
}
