package tests

import (
	"fmt"
	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/services/notifications/config"
	"github.com/HackIllinois/api/services/notifications/models"
	"github.com/HackIllinois/api/services/notifications/service"
	"os"
	"testing"
	"reflect"
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
	Tests retrieving all topic ids
*/
func TestGetAllTopicIDs(t *testing.T) {
	SetupTestDB(t)

	topics, err := service.GetAllTopicIDs()

	if err != nil {
		t.Fatal(err)
	}

	expected_topics := []string{"test_topic"}

	if !reflect.DeepEqual(topics, expected_topics) {
		t.Errorf("Wrong topcis.\nExpected %v\ngot %v\n", expected_topics, topics)
	}

	CleanupTestDB(t)
}

/*
	Tests retrieving a topic
*/
func TestGetTopic(t *testing.T) {
	SetupTestDB(t)

	topic, err := service.GetTopic("test_topic")

	if err != nil {
		t.Fatal(err)
	}

	expected_topic := models.Topic{
		ID:      "test_topic",
		UserIDs: []string{"test_user"},
	}

	if !reflect.DeepEqual(topic, &expected_topic) {
		t.Errorf("Wrong topcis.\nExpected %v\ngot %v\n", &expected_topic, topic)
	}

	CleanupTestDB(t)
}

/*
	Tests creating a topic
*/
func TestCreateTopic(t *testing.T) {
	SetupTestDB(t)

	err := service.CreateTopic("test_topic2")

	if err != nil {
		t.Fatal(err)
	}

	topic, err := service.GetTopic("test_topic2")

	if err != nil {
		t.Fatal(err)
	}

	expected_topic := models.Topic{
		ID:      "test_topic2",
		UserIDs: []string{},
	}

	if !reflect.DeepEqual(topic, &expected_topic) {
		t.Errorf("Wrong topcis.\nExpected %v\ngot %v\n", &expected_topic, topic)
	}

	CleanupTestDB(t)
}

/*
	Tests deleting a topic
*/
func TestDeleteTopic(t *testing.T) {
	SetupTestDB(t)

	err := service.DeleteTopic("test_topic")

	if err != nil {
		t.Fatal(err)
	}

	_, err = service.GetTopic("test_topic")

	if err != database.ErrNotFound {
		t.Fatal(err)
	}

	CleanupTestDB(t)
}
