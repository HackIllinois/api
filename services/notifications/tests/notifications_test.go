package tests

import (
	"fmt"
	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/services/notifications/config"
	"github.com/HackIllinois/api/services/notifications/models"
	"github.com/HackIllinois/api/services/notifications/service"
	"os"
	"reflect"
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
		Name:    "test_topic",
		Arn:     "arn:test",
		UserIDs: []string{"test_user"},
	}

	err := db.Insert("topics", &topic)

	if err != nil {
		t.Fatal(err)
	}

	notification := models.PastNotification{
		Body:      "test message",
		Title:     "test title",
		TopicName: "test_topic",
		Time:      2000,
	}

	err = db.Insert("notifications", &notification)

	if err != nil {
		t.Fatal(err)
	}

	device := models.Device{
		UserID:        "test_user",
		DeviceToken:   "token1",
		DeviceArn:     "arn:device_test_1",
		Platform:      "android",
		Subscriptions: map[string]string{"test_topic": ""},
	}
	device2 := models.Device{
		UserID:        "test_user_2",
		DeviceToken:   "token2",
		DeviceArn:     "arn:device_test_2",
		Platform:      "android",
		Subscriptions: map[string]string{},
	}

	err = db.Insert("devices", &device)

	if err != nil {
		t.Fatal(err)
	}

	err = db.Insert("devices", &device2)

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
	Service level test for getting all topics from db
*/
func TestGetAllTopicsSerice(t *testing.T) {
	SetupTestDB(t)

	topic := models.Topic{
		Name:    "test_topic_2",
		Arn:     "arn:test2",
		UserIDs: []string{"test_user_2"},
	}

	err := db.Insert("topics", &topic)

	if err != nil {
		t.Fatal(err)
	}

	actual_topic_list, err := service.GetAllTopics()

	if err != nil {
		t.Fatal(err)
	}

	expected_topic_list := models.TopicList{
		Topics: []string{
			"test_topic",
			"test_topic_2",
		},
	}

	if !reflect.DeepEqual(actual_topic_list, &expected_topic_list) {
		t.Errorf("Wrong topic list. Expected %v, got %v", &expected_topic_list, actual_topic_list)
	}

	CleanupTestDB(t)
}

/*
	Service level test for creating a notification topic
*/
func TestCreateTopicService(t *testing.T) {
	SetupTestDB(t)

	err := service.CreateTopic("test_topic_2")

	if err != nil {
		t.Fatal(err)
	}

	actual_topic_list, err := service.GetAllTopics()

	if err != nil {
		t.Fatal(err)
	}

	expected_topic_list := models.TopicList{
		Topics: []string{
			"test_topic",
			"test_topic_2",
		},
	}

	if !reflect.DeepEqual(actual_topic_list, &expected_topic_list) {
		t.Errorf("Wrong topic list. Expected %v, got %v", &expected_topic_list, actual_topic_list)
	}

	CleanupTestDB(t)
}

/*
	Service level test for getting all notifications from db
*/
func TestGetAllNotificationsService(t *testing.T) {
	SetupTestDB(t)

	notification := models.PastNotification{
		Body:      "test message 2",
		Title:     "test title 2",
		TopicName: "test_topic_2",
		Time:      3000,
	}

	err := db.Insert("notifications", &notification)

	if err != nil {
		t.Fatal(err)
	}

	actual_notification_list, err := service.GetAllNotifications()

	if err != nil {
		t.Fatal(err)
	}

	expected_notification_list := models.NotificationList{
		Notifications: []models.PastNotification{
			{
				Body:      "test message",
				Title:     "test title",
				TopicName: "test_topic",
				Time:      2000,
			},
			{
				Body:      "test message 2",
				Title:     "test title 2",
				TopicName: "test_topic_2",
				Time:      3000,
			},
		},
	}

	if !reflect.DeepEqual(actual_notification_list, &expected_notification_list) {
		t.Errorf("Wrong notification list. Expected %v, got %v", &expected_notification_list, actual_notification_list)
	}

	CleanupTestDB(t)
}

/*
	Service level test for getting notifications for a specific topic from db
*/
func TestGetNotificationsForTopicService(t *testing.T) {
	SetupTestDB(t)

	notification := models.PastNotification{
		Body:      "test message again",
		Title:     "test title again",
		TopicName: "test_topic",
		Time:      5000,
	}

	err := db.Insert("notifications", &notification)

	if err != nil {
		t.Fatal(err)
	}

	notification = models.PastNotification{
		Body:      "test message 2",
		Title:     "test title 2",
		TopicName: "test_topic_2",
		Time:      3000,
	}

	err = db.Insert("notifications", &notification)

	if err != nil {
		t.Fatal(err)
	}

	actual_notification_list, err := service.GetNotificationsForTopic("test_topic")

	if err != nil {
		t.Fatal(err)
	}

	expected_notification_list := models.NotificationList{
		Notifications: []models.PastNotification{
			{
				Body:      "test message",
				Title:     "test title",
				TopicName: "test_topic",
				Time:      2000,
			},
			{
				Body:      "test message again",
				Title:     "test title again",
				TopicName: "test_topic",
				Time:      5000,
			},
		},
	}

	if !reflect.DeepEqual(actual_notification_list, &expected_notification_list) {
		t.Errorf("Wrong notification list. Expected %v, got %v", &expected_notification_list, actual_notification_list)
	}

	CleanupTestDB(t)
}

/*
	Service level test for deleting a topic from db
*/
func TestDeleteTopicService(t *testing.T) {
	SetupTestDB(t)

	topic := models.Topic{
		Name: "test_topic_2",
		Arn:  "arn:test2",
	}

	err := db.Insert("topics", &topic)

	if err != nil {
		t.Fatal(err)
	}

	err = service.DeleteTopic("test_topic_2")

	if err != nil {
		t.Fatal(err)
	}

	actual_topic_list, err := service.GetAllTopics()

	if err != nil {
		t.Fatal(err)
	}

	expected_topic_list := models.TopicList{
		Topics: []string{
			"test_topic",
		},
	}

	if !reflect.DeepEqual(actual_topic_list, &expected_topic_list) {
		t.Errorf("Wrong topic list. Expected %v, got %v", &expected_topic_list, actual_topic_list)
	}

	CleanupTestDB(t)
}

/*
	Service level test for creating a notification
*/
func TestCreateNotificationService(t *testing.T) {
	SetupTestDB(t)

	notification := models.Notification{
		Body:  "test message 2",
		Title: "test title 2",
	}

	past_notification, err := service.PublishNotification("test_topic", notification)

	if err != nil {
		t.Fatal(err)
	}

	actual_notification_list, err := service.GetNotificationsForTopic("test_topic")

	if err != nil {
		t.Fatal(err)
	}

	expected_notification_list := models.NotificationList{
		Notifications: []models.PastNotification{
			{
				Body:      "test message",
				Title:     "test title",
				TopicName: "test_topic",
				Time:      2000,
			},
			{
				Body:      "test message 2",
				Title:     "test title 2",
				TopicName: "test_topic",
				Time:      past_notification.Time,
			},
		},
	}

	if !reflect.DeepEqual(actual_notification_list, &expected_notification_list) {
		t.Errorf("Wrong notification list. Expected %v, got %v", &expected_notification_list, actual_notification_list)
	}

	CleanupTestDB(t)
}

/*
	Service level test for getting topic info from db
*/
func TestGetTopicInfoService(t *testing.T) {
	SetupTestDB(t)

	actual_topic_info, err := service.GetTopicInfo("test_topic")

	if err != nil {
		t.Fatal(err)
	}

	expected_topic_info := models.Topic{
		Name:    "test_topic",
		Arn:     "arn:test",
		UserIDs: []string{"test_user"},
	}

	if !reflect.DeepEqual(actual_topic_info, &expected_topic_info) {
		t.Errorf("Wrong topic info. Expected %v, got %v", &expected_topic_info, actual_topic_info)
	}

	CleanupTestDB(t)
}

/*
	Service level test for subscribing users to a topic
*/
func TestSubscribeUserService(t *testing.T) {
	SetupTestDB(t)

	userid_list := models.UserIDList{
		UserIDs: []string{
			"test_user_2",
			"test_user_3",
		},
	}

	err := service.AddUsersToTopic("test_topic", userid_list)

	if err != nil {
		t.Fatal(err)
	}

	actual_topic_info, err := service.GetTopicInfo("test_topic")

	if err != nil {
		t.Fatal(err)
	}

	expected_topic_info := models.Topic{
		Name: "test_topic",
		Arn:  "arn:test",
		UserIDs: []string{
			"test_user",
			"test_user_2",
			"test_user_3",
		},
	}

	if !reflect.DeepEqual(actual_topic_info, &expected_topic_info) {
		t.Errorf("Wrong topic info. Expected %v, got %v", &expected_topic_info, actual_topic_info)
	}

	actual_devices_list, err := service.GetAllDevices()

	if err != nil {
		t.Fatal(err)
	}

	expected_devices_list := []models.Device{
		{
			UserID:        "test_user",
			DeviceToken:   "token1",
			DeviceArn:     "arn:device_test_1",
			Platform:      "android",
			Subscriptions: map[string]string{"test_topic": ""},
		},
		{
			UserID:        "test_user_2",
			DeviceToken:   "token2",
			DeviceArn:     "arn:device_test_2",
			Platform:      "android",
			Subscriptions: map[string]string{"test_topic": ""},
		},
	}

	if !reflect.DeepEqual(actual_devices_list, &expected_devices_list) {
		t.Errorf("Wrong devices list. Expected %v, got %v", &expected_devices_list, actual_devices_list)
	}

	CleanupTestDB(t)
}

/*
	Service level test for unsubscribing users from a topic
*/
func TestUnsubscribeUserService(t *testing.T) {
	SetupTestDB(t)

	userid_list := models.UserIDList{
		UserIDs: []string{
			"test_user",
			"test_user_3",
		},
	}

	err := service.RemoveUsersFromTopic("test_topic", userid_list)

	if err != nil {
		t.Fatal(err)
	}

	actual_topic_info, err := service.GetTopicInfo("test_topic")

	if err != nil {
		t.Fatal(err)
	}

	expected_topic_info := models.Topic{
		Name:    "test_topic",
		Arn:     "arn:test",
		UserIDs: []string{},
	}

	if !reflect.DeepEqual(actual_topic_info, &expected_topic_info) {
		t.Errorf("Wrong topic info. Expected %v, got %v", &expected_topic_info, actual_topic_info)
	}

	actual_devices_list, err := service.GetAllDevices()

	if err != nil {
		t.Fatal(err)
	}

	expected_devices_list := []models.Device{
		{
			UserID:        "test_user",
			DeviceToken:   "token1",
			DeviceArn:     "arn:device_test_1",
			Platform:      "android",
			Subscriptions: map[string]string{},
		},
		{
			UserID:        "test_user_2",
			DeviceToken:   "token2",
			DeviceArn:     "arn:device_test_2",
			Platform:      "android",
			Subscriptions: map[string]string{},
		},
	}

	if !reflect.DeepEqual(actual_devices_list, &expected_devices_list) {
		t.Errorf("Wrong devices list. Expected %v, got %v", &expected_devices_list, actual_devices_list)
	}

	CleanupTestDB(t)
}

/*
	Service level test for registering a device to a user
*/
func TestRegisterDevice(t *testing.T) {
	SetupTestDB(t)

	device_registration := models.DeviceRegistration{
		DeviceToken: "token3",
		Platform:    "android",
	}

	err := service.RegisterDeviceToUser("test_user", device_registration)

	if err != nil {
		t.Fatal(err)
	}

	actual_devices_list, err := service.GetAllDevices()

	if err != nil {
		t.Fatal(err)
	}

	expected_devices_list := []models.Device{
		{
			UserID:        "test_user",
			DeviceToken:   "token1",
			DeviceArn:     "arn:device_test_1",
			Platform:      "android",
			Subscriptions: map[string]string{"test_topic": ""},
		},
		{
			UserID:        "test_user_2",
			DeviceToken:   "token2",
			DeviceArn:     "arn:device_test_2",
			Platform:      "android",
			Subscriptions: map[string]string{},
		},
		{
			UserID:        "test_user",
			DeviceToken:   "token3",
			DeviceArn:     "",
			Platform:      "android",
			Subscriptions: map[string]string{"test_topic": ""},
		},
	}

	if !reflect.DeepEqual(actual_devices_list, &expected_devices_list) {
		t.Errorf("Wrong devices list. Expected %v, got %v", &expected_devices_list, actual_devices_list)
	}

	CleanupTestDB(t)
}
