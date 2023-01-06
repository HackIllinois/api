package tests

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/services/notifications/config"
	"github.com/HackIllinois/api/services/notifications/models"
	"github.com/HackIllinois/api/services/notifications/service"
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
		ID:      "User",
		UserIDs: []string{"test_user"},
	}

	err := db.Insert("topics", &topic, nil)

	if err != nil {
		t.Fatal(err)
	}

	notification := models.Notification{
		ID:    "test_id",
		Title: "test title",
		Body:  "test body",
		Topic: "User",
		Time:  2000,
	}

	err = db.Insert("notifications", &notification, nil)

	if err != nil {
		t.Fatal(err)
	}

	user := models.User{
		ID:      "test_user",
		Devices: []string{"test_arn"},
	}

	err = db.Insert("users", &user, nil)

	if err != nil {
		t.Fatal(err)
	}
}

/*
	Drop test db
*/
func CleanupTestDB(t *testing.T) {
	err := db.DropDatabase(nil)

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

	expected_topics := []string{"User"}

	if !reflect.DeepEqual(topics, expected_topics) {
		t.Errorf("Wrong topics.\nExpected %v\ngot %v\n", expected_topics, topics)
	}

	CleanupTestDB(t)
}

/*
	Tests retrieving a topic
*/
func TestGetTopic(t *testing.T) {
	SetupTestDB(t)

	topic, err := service.GetTopic("User")

	if err != nil {
		t.Fatal(err)
	}

	expected_topic := models.Topic{
		ID:      "User",
		UserIDs: []string{"test_user"},
	}

	if !reflect.DeepEqual(topic, &expected_topic) {
		t.Errorf("Wrong topics.\nExpected %v\ngot %v\n", &expected_topic, topic)
	}

	CleanupTestDB(t)
}

/*
	Tests creating a topic
*/
func TestCreateTopic(t *testing.T) {
	SetupTestDB(t)

	err := service.CreateTopic("User2")

	if err != nil {
		t.Fatal(err)
	}

	topic, err := service.GetTopic("User2")

	if err != nil {
		t.Fatal(err)
	}

	expected_topic := models.Topic{
		ID:      "User2",
		UserIDs: []string{},
	}

	if !reflect.DeepEqual(topic, &expected_topic) {
		t.Errorf("Wrong topics.\nExpected %v\ngot %v\n", &expected_topic, topic)
	}

	CleanupTestDB(t)
}

/*
	Tests deleting a topic
*/
func TestDeleteTopic(t *testing.T) {
	SetupTestDB(t)

	err := service.DeleteTopic("User")

	if err != nil {
		t.Fatal(err)
	}

	_, err = service.GetTopic("User")

	if err != database.ErrNotFound {
		t.Fatal(err)
	}

	CleanupTestDB(t)
}

/*
	Tests getting all notifications for a topic
*/
func TestGetAllNotificationsForTopic(t *testing.T) {
	SetupTestDB(t)

	notifications, err := service.GetAllNotificationsForTopic("User")

	if err != nil {
		t.Fatal(err)
	}

	expected_notifications := []models.Notification{
		{
			ID:    "test_id",
			Title: "test title",
			Body:  "test body",
			Topic: "User",
			Time:  2000,
		},
	}

	if !reflect.DeepEqual(notifications, expected_notifications) {
		t.Errorf("Wrong topics.\nExpected %v\ngot %v\n", expected_notifications, notifications)
	}

	CleanupTestDB(t)
}

/*
	Tests getting all notifications for a set of topics
*/
func TestGetAllNotifications(t *testing.T) {
	SetupTestDB(t)

	notifications, err := service.GetAllNotifications([]string{"User"})

	if err != nil {
		t.Fatal(err)
	}

	expected_notifications := []models.Notification{
		{
			ID:    "test_id",
			Title: "test title",
			Body:  "test body",
			Topic: "User",
			Time:  2000,
		},
	}

	if !reflect.DeepEqual(notifications, expected_notifications) {
		t.Errorf("Wrong topics.\nExpected %v\ngot %v\n", expected_notifications, notifications)
	}

	CleanupTestDB(t)
}

/*
	Tests getting all public notifications
*/
func TestGetAllPublicNotifications(t *testing.T) {
	SetupTestDB(t)

	notifications, err := service.GetAllPublicNotifications()

	if err != nil {
		t.Fatal(err)
	}

	expected_notifications := []models.Notification{
		{
			ID:    "test_id",
			Title: "test title",
			Body:  "test body",
			Topic: "User",
			Time:  2000,
		},
	}

	if !reflect.DeepEqual(notifications, expected_notifications) {
		t.Errorf("Wrong topics.\nExpected %v\ngot %v\n", expected_notifications, notifications)
	}

	CleanupTestDB(t)
}

/*
	Tests subscriptioning user to a topic
*/
func TestSubscribeToTopic(t *testing.T) {
	SetupTestDB(t)

	err := service.SubscribeToTopic("test_user2", "User")

	if err != nil {
		t.Fatal(err)
	}

	CleanupTestDB(t)
}

/*
	Tests unsubscriptioning user to a topic
*/
func TestUnsubscribeToTopic(t *testing.T) {
	SetupTestDB(t)

	err := service.UnsubscribeToTopic("test_user", "User")

	if err != nil {
		t.Fatal(err)
	}

	CleanupTestDB(t)
}

/*
	Tests retrieving a user's devices
*/
func TestGetUserDevices(t *testing.T) {
	SetupTestDB(t)

	devices, err := service.GetUserDevices("test_user")

	if err != nil {
		t.Fatal(err)
	}

	expected_devices := []string{"test_arn"}

	if !reflect.DeepEqual(devices, expected_devices) {
		t.Errorf("Wrong topics.\nExpected %v\ngot %v\n", expected_devices, devices)
	}

	CleanupTestDB(t)
}

/*
	Tests setting a user's devices
*/
func TestSetUserDevices(t *testing.T) {
	SetupTestDB(t)

	err := service.SetUserDevices("test_user", []string{"test_arn", "test_arn2"})

	if err != nil {
		t.Fatal(err)
	}

	devices, err := service.GetUserDevices("test_user")

	if err != nil {
		t.Fatal(err)
	}

	expected_devices := []string{"test_arn", "test_arn2"}

	if !reflect.DeepEqual(devices, expected_devices) {
		t.Errorf("Wrong topics.\nExpected %v\ngot %v\n", expected_devices, devices)
	}

	CleanupTestDB(t)
}

/*
	Tests registering a device to a user
*/
func TestRegisterDeviceToUser(t *testing.T) {
	SetupTestDB(t)

	err := service.RegisterDeviceToUser("test_token", "android", "test_user")

	if err != nil {
		t.Fatal(err)
	}

	devices, err := service.GetUserDevices("test_user")

	if err != nil {
		t.Fatal(err)
	}

	expected_devices := []string{"test_arn", ""}

	if !reflect.DeepEqual(devices, expected_devices) {
		t.Errorf("Wrong topics.\nExpected %v\ngot %v\n", expected_devices, devices)
	}

	// Test deduplication
	err = service.RegisterDeviceToUser("test_token", "android", "test_user")

	if err != nil {
		t.Fatal(err)
	}

	devices, err = service.GetUserDevices("test_user")

	if err != nil {
		t.Fatal(err)
	}

	expected_devices = []string{"test_arn", ""}

	if !reflect.DeepEqual(devices, expected_devices) {
		t.Errorf("Wrong topics.\nExpected %v\ngot %v\n", expected_devices, devices)
	}

	CleanupTestDB(t)
}

/*
	Tests getting the list of userids to receive a notification
*/
func TestGetNotificationRecipients(t *testing.T) {
	SetupTestDB(t)

	userids, err := service.GetNotificationRecipients("User")

	if err != nil {
		t.Fatal(err)
	}

	expected_userids := []string{"test_user"}

	if !reflect.DeepEqual(userids, expected_userids) {
		t.Errorf("Wrong topics.\nExpected %v\ngot %v\n", expected_userids, userids)
	}

	CleanupTestDB(t)
}

/*
	Tests getting the list of arns to send a notification to based on userids
*/
func TestGetNotificationRecipientArns(t *testing.T) {
	SetupTestDB(t)

	arns, err := service.GetNotificationRecipientArns([]string{"test_user"})

	if err != nil {
		t.Fatal(err)
	}

	expected_arns := []string{"test_arn"}

	if !reflect.DeepEqual(arns, expected_arns) {
		t.Errorf("Wrong topics.\nExpected %v\ngot %v\n", expected_arns, arns)
	}

	CleanupTestDB(t)
}

/*
	Tests publishing a notification
*/
func TestPublishNotificationToTopic(t *testing.T) {
	SetupTestDB(t)

	notification := models.Notification{
		ID:    "test_id2",
		Title: "test title 2",
		Body:  "test body 2",
		Topic: "User",
		Time:  3000,
	}

	// Send notification to one user w/ one device
	order, err := service.PublishNotificationToTopic(notification)
	if err != nil {
		t.Fatal(err)
	}

	expected_order := models.NotificationOrder{
		ID:         "test_id2",
		Recipients: 1,
		Success:    0,
		Failure:    0,
		Time:       3000,
	}
	if !reflect.DeepEqual(order, &expected_order) {
		t.Errorf("Wrong order.\nExpected %v\ngot %v\n", &expected_order, order)
	}

	// Add additional user w/ two devices
	user := models.User{
		ID:      "test_user_2",
		Devices: []string{"test_arn2", "test_arn3"},
	}
	err = db.Insert("users", &user, nil)

	selector := database.QuerySelector{
		"id": "User",
	}
	topic := models.Topic{
		ID:      "User",
		UserIDs: []string{"test_user", "test_user_2"},
	}
	err = db.Replace("topics", selector, &topic, false, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Send notification to two users w/ three total devices
	order, err = service.PublishNotificationToTopic(notification)
	if err != nil {
		t.Fatal(err)
	}

	expected_order = models.NotificationOrder{
		ID:         "test_id2",
		Recipients: 3,
		Success:    0,
		Failure:    0,
		Time:       3000,
	}

	if !reflect.DeepEqual(order, &expected_order) {
		t.Errorf("Wrong order.\nExpected %v\ngot %v\n", &expected_order, order)
	}

	CleanupTestDB(t)
}
