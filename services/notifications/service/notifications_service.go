package service

import (
	"errors"
	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/services/notifications/config"
	"github.com/HackIllinois/api/services/notifications/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"time"
)

var sess *session.Session
var client *sns.SNS
var db database.Database

func init() {
	sess = session.Must(session.NewSession(&aws.Config{
		Region: aws.String(config.SNS_REGION),
	}))
	client = sns.New(sess)

	db_connection, err := database.InitDatabase(config.NOTIFICATIONS_DB_HOST, config.NOTIFICATIONS_DB_NAME)

	if err != nil {
		panic(err)
	}

	db = db_connection
}

/*
	Returns a list of available SNS Topics
*/
func GetAllTopics() (*models.TopicList, error) {
	var topic_list models.TopicList

	err := db.FindAll("topics", nil, &topic_list.Topics)

	if err != nil {
		return nil, err
	}

	return &topic_list, nil
}

/*
	Returns a list of available SNS Topics
*/
func GetAllNotifications() (*models.NotificationList, error) {
	var notifications []models.PastNotification

	err := db.FindAll("notifications", nil, &notifications)

	if err != nil {
		return nil, err
	}

	notifications_list := models.NotificationList{
		Notifications: notifications,
	}

	return &notifications_list, nil
}

/*
	Creates an SNS Topic
*/
func CreateTopic(name string) (*models.Topic, error) {
	var arn string

	if config.IS_PRODUCTION {
		out, err := client.CreateTopic(&sns.CreateTopicInput{Name: &name})

		if err != nil {
			return nil, err
		}

		arn = *out.TopicArn
	}

	_, err := GetTopicInfo(name)

	if err != database.ErrNotFound {
		if err != nil {
			return nil, err
		}
		return nil, errors.New("Topic already exists")
	}

	topic := models.Topic{Arn: arn, Name: name}

	err = db.Insert("topics", &topic)

	if err != nil {
		return nil, err
	}

	return &topic, nil
}

/*
	Deletes an SNS Topic
*/
func DeleteTopic(name string) error {

	topic, err := GetTopicInfo(name)

	if err != nil {
		return err
	}

	if config.IS_PRODUCTION {
		_, err = client.DeleteTopic(&sns.DeleteTopicInput{TopicArn: &topic.Arn})

		if err != nil {
			return err
		}
	}

	topic_selector := database.QuerySelector{
		"name": name,
	}

	err = db.RemoveOne("topics", topic_selector)

	if err != nil {
		return err
	}

	return nil
}

func GetTopicInfo(name string) (*models.Topic, error) {
	topic_selector := database.QuerySelector{
		"name": name,
	}

	var topic models.Topic

	err := db.FindOne("topics", topic_selector, &topic)

	if err != nil {
		return nil, err
	}

	return &topic, nil
}

/*
	Dispatches a notification to a given SNS Topic
*/
func PublishNotification(topic_name string, notification models.Notification) (*models.PastNotification, error) {

	topic, err := GetTopicInfo(topic_name)

	if err != nil {
		return nil, err
	}

	arn := topic.Arn

	if config.IS_PRODUCTION {
		_, err = client.Publish(&sns.PublishInput{
			TopicArn: &arn,
			Message:  &notification.Message,
		})

		if err != nil {
			return nil, err
		}
	}

	current_time := time.Now().Unix()

	past_notification := models.PastNotification{TopicName: topic_name, Message: notification.Message, Time: current_time}
	err = db.Insert("notifications", &past_notification)

	return &past_notification, nil
}

func GetNotificationsForTopic(topic_name string) (*models.NotificationList, error) {
	topic_name_selector := database.QuerySelector{
		"topicname": topic_name,
	}

	var notifications []models.PastNotification

	err := db.FindAll("notifications", topic_name_selector, &notifications)

	if err != nil {
		return nil, err
	}

	notifications_list := models.NotificationList{
		Notifications: notifications,
	}

	return &notifications_list, nil
}
