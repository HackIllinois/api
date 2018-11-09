package tests

import (
	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/services/notifications/config"
	"github.com/HackIllinois/api/services/notifications/models"
	"github.com/HackIllinois/api/services/notifications/service"
	"reflect"
	"testing"
)

var db database.Database

func init() {
	db_connection, err := database.InitDatabase(config.NOTIFICATIONS_DB_HOST, config.NOTIFICATIONS_DB_NAME)

	if err != nil {
		panic(err)
	}

	db = db_connection
}

/*
	Initialize db with a test topic and notification
*/
func SetupTestDB(t *testing.T) {
	topic := models.Topic{
		Name: "test_topic",
		Arn:  "arn:test",
	}

	err := db.Insert("topics", &topic)

	if err != nil {
		t.Fatal(err)
	}

	notification := models.PastNotification{
		Message:   "test message",
		TopicName: "test_topic",
		Time:      2000,
	}

	err = db.Insert("notifications", &notification)

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
		Name: "test_topic_2",
		Arn:  "arn:test2",
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
		Topics: []models.Topic{
			models.Topic{
				Name: "test_topic",
				Arn:  "arn:test",
			},
			models.Topic{
				Name: "test_topic_2",
				Arn:  "arn:test2",
			},
		},
	}

	if !reflect.DeepEqual(actual_topic_list, &expected_topic_list) {
		t.Errorf("Wrong topic list. Expected %v, got %v", expected_topic_list, actual_topic_list)
	}

	CleanupTestDB(t)
}

/*
	Service level test for creating a notification topic
*/
func TestCreateTopicService(t *testing.T) {
	SetupTestDB(t)

	topic_details, err := service.CreateTopic("test_topic_2")

	if err != nil {
		t.Fatal(err)
	}

	actual_topic_list, err := service.GetAllTopics()

	if err != nil {
		t.Fatal(err)
	}

	expected_topic_list := models.TopicList{
		Topics: []models.Topic{
			models.Topic{
				Name: "test_topic",
				Arn:  "arn:test",
			},
			models.Topic{
				Name: "test_topic_2",
				Arn:  topic_details.Arn,
			},
		},
	}

	if !reflect.DeepEqual(actual_topic_list, &expected_topic_list) {
		t.Errorf("Wrong topic list. Expected %v, got %v", expected_topic_list, actual_topic_list)
	}

	CleanupTestDB(t)
}

/*
	Service level test for getting all notifications from db
*/
func TestGetAllNotificationsService(t *testing.T) {
	SetupTestDB(t)

	notification := models.PastNotification{
		Message:   "test message 2",
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
			models.PastNotification{
				Message:   "test message",
				TopicName: "test_topic",
				Time:      2000,
			},
			models.PastNotification{
				Message:   "test message 2",
				TopicName: "test_topic_2",
				Time:      3000,
			},
		},
	}

	if !reflect.DeepEqual(actual_notification_list, &expected_notification_list) {
		t.Errorf("Wrong notification list. Expected %v, got %v", expected_notification_list, actual_notification_list)
	}

	CleanupTestDB(t)
}

/*
	Service level test for getting notifications for a specific topic from db
*/
func TestGetNotificationsForTopicService(t *testing.T) {
	SetupTestDB(t)

	notification := models.PastNotification{
		Message:   "test message again",
		TopicName: "test_topic",
		Time:      5000,
	}

	err := db.Insert("notifications", &notification)

	if err != nil {
		t.Fatal(err)
	}

	notification = models.PastNotification{
		Message:   "test message 2",
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
			models.PastNotification{
				Message:   "test message",
				TopicName: "test_topic",
				Time:      2000,
			},
			models.PastNotification{
				Message:   "test message again",
				TopicName: "test_topic",
				Time:      5000,
			},
		},
	}

	if !reflect.DeepEqual(actual_notification_list, &expected_notification_list) {
		t.Errorf("Wrong notification list. Expected %v, got %v", expected_notification_list, actual_notification_list)
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
		Topics: []models.Topic{
			models.Topic{
				Name: "test_topic",
				Arn:  "arn:test",
			},
		},
	}

	if !reflect.DeepEqual(actual_topic_list, &expected_topic_list) {
		t.Errorf("Wrong topic list. Expected %v, got %v", expected_topic_list, actual_topic_list)
	}

	CleanupTestDB(t)
}

/*
	Service level test for creating a notification
*/
func TestCreateNotificationService(t *testing.T) {
	SetupTestDB(t)

	notification := models.Notification{
		Message: "test message 2",
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
			models.PastNotification{
				Message:   "test message",
				TopicName: "test_topic",
				Time:      2000,
			},
			models.PastNotification{
				Message:   "test message 2",
				TopicName: "test_topic",
				Time:      past_notification.Time,
			},
		},
	}

	if !reflect.DeepEqual(actual_notification_list, &expected_notification_list) {
		t.Errorf("Wrong notification list. Expected %v, got %v", expected_notification_list, actual_notification_list)
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
		Name: "test_topic",
		Arn:  "arn:test",
	}

	if !reflect.DeepEqual(actual_topic_info, &expected_topic_info) {
		t.Errorf("Wrong topic info. Expected %v, got %v", expected_topic_info, actual_topic_info)
	}

	CleanupTestDB(t)
}
