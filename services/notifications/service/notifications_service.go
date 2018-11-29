package service

import (
	"errors"
	"fmt"
	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/services/notifications/config"
	"github.com/HackIllinois/api/services/notifications/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"time"
)

const APPLICATION_PROTOCOL = "application"

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
	var topics []models.Topic
	err := db.FindAll("topics", nil, &topics)

	if err != nil {
		return nil, err
	}

	var topic_list models.TopicList
	for _, topic := range topics {
		topic_list.Topics = append(topic_list.Topics, topic.Name)
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
func CreateTopic(name string) error {
	var arn string

	if config.IS_PRODUCTION || true {
		out, err := client.CreateTopic(&sns.CreateTopicInput{Name: &name})

		if err != nil {
			return err
		}

		arn = *out.TopicArn
	}

	_, err := GetTopicInfo(name)

	if err != database.ErrNotFound {
		if err != nil {
			return err
		}
		return errors.New("Topic already exists")
	}

	topic := models.Topic{Arn: arn, Name: name, UserIDs: nil}

	err = db.Insert("topics", &topic)

	if err != nil {
		return err
	}

	return nil
}

/*
	Deletes an SNS Topic
*/
func DeleteTopic(name string) error {

	topic, err := GetTopicInfo(name)

	if err != nil {
		return err
	}

	if config.IS_PRODUCTION || true {
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

	if config.IS_PRODUCTION || true {
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

/*
	Adds the given userids to the specified topic
*/
func AddUsersToTopic(topic_name string, userid_list models.UserIDList) error {
	selector := database.QuerySelector{
		"name": topic_name,
	}

	modifier := database.QuerySelector{
		"$addToSet": database.QuerySelector{
			"userids": database.QuerySelector{
				"$each": userid_list.UserIDs,
			},
		},
	}

	topic_selector := database.QuerySelector{
		"name": topic_name,
	}

	var topic models.Topic
	err := db.FindOne("topics", topic_selector, &topic)

	if err != nil {
		return err
	}

	// Subscribe each of the specified users' devices to this topic
	for _, user_id := range userid_list.UserIDs {
		query := database.QuerySelector{
			"userid": user_id,
		}

		var devices []models.Device
		err := db.FindAll("devices", query, &devices)

		if err != nil {
			return err
		}

		for _, device := range devices {
			err := SubscribeDeviceToTopic(topic, device)

			if err != nil {
				return err
			}
		}
	}

	return db.Update("topics", selector, &modifier)
}

/*
	Removes the given userids from the specified topic
*/
func RemoveUsersFromTopic(topic_name string, userid_list models.UserIDList) error {
	selector := database.QuerySelector{
		"name": topic_name,
	}

	modifier := database.QuerySelector{
		"$pull": database.QuerySelector{
			"userids": database.QuerySelector{
				"$in": userid_list.UserIDs,
			},
		},
	}

	topic_selector := database.QuerySelector{
		"name": topic_name,
	}

	var topic models.Topic
	err := db.FindOne("topics", topic_selector, &topic)

	if err != nil {
		return err
	}

	// Unsubscribe each of the specificed users' devices from this topic
	for _, user_id := range userid_list.UserIDs {
		query := database.QuerySelector{
			"userid": user_id,
		}

		var devices []models.Device
		err = db.FindAll("devices", query, &devices)

		if err != nil {
			return err
		}

		for _, device := range devices {
			err = UnsubscribeDeviceFromTopic(topic, device)

			if err != nil {
				return err
			}
		}
	}

	return db.Update("topics", selector, &modifier)
}

/*
	Links the given device token with a user
*/
func RegisterDeviceToUser(user_id string, device_reg models.DeviceRegistration) error {
	var device_arn string
	var platform_arn string

	// Map platform (android, ios etc) to its ARN
	switch device_reg.Platform {
	case "android":
		platform_arn = config.ANDROID_PLATFORM_ARN
	default:
		return errors.New("Invalid platform")
	}

	if config.IS_PRODUCTION || true {
		out, err := client.CreatePlatformEndpoint(&sns.CreatePlatformEndpointInput{CustomUserData: &user_id, Token: &device_reg.DeviceToken, PlatformApplicationArn: &platform_arn})

		if err != nil {
			return err
		}

		device_arn = *out.EndpointArn
	}

	subs := make(map[string]string)
	device := models.Device{DeviceArn: device_arn, DeviceToken: device_reg.DeviceToken, Platform: device_reg.Platform, UserID: user_id, Subscriptions: subs}

	err := db.Insert("devices", device)

	if err != nil {
		return err
	}

	// Subscribe the device to all of a user's topics

	topic_selector := database.QuerySelector{
		"userids": database.QuerySelector{
			"$all": [1]string{user_id},
		},
	}

	var topics []models.Topic
	err = db.FindAll("topics", topic_selector, &topics)

	if err != nil {
		return err
	}

	for _, topic := range topics {
		err = SubscribeDeviceToTopic(topic, device)

		if err != nil {
			return err
		}
	}

	return nil
}

/*
	Subscribes a given Device to a Topic, in the database and SNS
*/
func SubscribeDeviceToTopic(topic models.Topic, device models.Device) error {
	app_protocol := APPLICATION_PROTOCOL

	var sub_arn string

	if config.IS_PRODUCTION || true {
		out, err := client.Subscribe(&sns.SubscribeInput{Protocol: &app_protocol, TopicArn: &topic.Arn, Endpoint: &device.DeviceArn})

		if err != nil {
			return err
		}

		sub_arn = *out.SubscriptionArn
	}

	device_selector := database.QuerySelector{
		"devicearn": device.DeviceArn,
	}

	set_query := fmt.Sprintf("subscriptions.%s", topic.Name)

	// Keep track of subscription's ARN so we can unsubscribe later
	device_modifier := database.QuerySelector{
		"$set": database.QuerySelector{
			set_query: sub_arn,
		},
	}

	err := db.Update("devices", device_selector, &device_modifier)

	return err
}

/*
	Unsubscribes a given Device from a Topic, both in the database and SNS
*/
func UnsubscribeDeviceFromTopic(topic models.Topic, device models.Device) error {
	sub_arn, ok := device.Subscriptions[topic.Name]
	if ok {
		if config.IS_PRODUCTION || true {
			_, err := client.Unsubscribe(&sns.UnsubscribeInput{SubscriptionArn: &sub_arn})

			if err != nil {
				return err
			}
		}

		device_selector := database.QuerySelector{
			"devicearn": device.DeviceArn,
		}

		set_query := fmt.Sprintf("subscriptions.%s", topic.Name)

		// Unset device's subscription ARN for this topic since it's no longer needed
		device_modifier := database.QuerySelector{
			"$unset": database.QuerySelector{
				set_query: "",
			},
		}

		err := db.Update("devices", device_selector, &device_modifier)

		if err != nil {
			return err
		}
	}
	return nil
}
