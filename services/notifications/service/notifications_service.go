package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/services/notifications/config"
	"github.com/HackIllinois/api/services/notifications/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"strings"
	"time"
)

const APPLICATION_PROTOCOL = "application"
const MESSAGE_STRUCTURE = "json"

var ERR_DEVICE_NOT_SUBSCRIBED = errors.New("Device not subscribed to topic")

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

	if config.IS_PRODUCTION {
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

	current_time := time.Now().Unix()

	past_notification := models.PastNotification{TopicName: topic_name, Title: notification.Title, Body: notification.Body, Time: current_time}

	time_selector := database.QuerySelector{
		"time": past_notification.Time,
	}

	err = db.FindOne("notifications", time_selector, nil)

	if err != database.ErrNotFound {
		return nil, errors.New("Notification will same timestamp found")
	}

	notification_json_str, err := GenerateNotificationJson(notification, past_notification)

	if err != nil {
		return nil, err
	}

	arn := topic.Arn
	message_structure := MESSAGE_STRUCTURE

	if config.IS_PRODUCTION {
		_, err = client.Publish(&sns.PublishInput{
			TopicArn:         &arn,
			Message:          notification_json_str,
			MessageStructure: &message_structure,
		})

		if err != nil {
			return nil, err
		}
	}

	err = db.Insert("notifications", &past_notification)

	return &past_notification, err
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
   Subscribes a user to topics corresponding to their roles, and unsubscribes a user from all other topics
*/
func UpdateUserSubscriptions(user_id string) (*models.TopicList, error) {
	user_roles, err := GetRoles(user_id)

	if err != nil {
		return nil, err
	}

	topics_list, err := GetAllTopics()

	if err != nil {
		return nil, err
	}

	user_list := models.UserIDList{UserIDs: []string{user_id}}

	for _, topic := range topics_list.Topics {
		err = RemoveUsersFromTopic(topic, user_list)
		if err != nil {
			if err == ERR_DEVICE_NOT_SUBSCRIBED {
				continue
			}
			return nil, err
		}
	}

	var topics []string

	for _, role := range user_roles.Roles {
		if topic, ok := config.GROUP_TOPIC_MAP[role]; ok {
			topics = append(topics, topic)
			err = AddUsersToTopic(topic, user_list)
			if err != nil {
				return nil, err
			}
		}
	}

	subscribed_topics := models.TopicList{
		Topics: topics,
	}

	return &subscribed_topics, nil
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
	query := database.QuerySelector{
		"userid": database.QuerySelector{
			"$in": userid_list.UserIDs,
		},
	}

	var devices []models.Device
	err = db.FindAll("devices", query, &devices)

	if err != nil {
		return err
	}

	err_chan := make(chan error)
	for _, device := range devices {
		go SubscribeDeviceAsync(topic, device, err_chan)
	}

	for range devices {
		err := <-err_chan

		if err != nil {
			return err
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
	query := database.QuerySelector{
		"userid": database.QuerySelector{
			"$in": userid_list.UserIDs,
		},
	}

	var devices []models.Device
	err = db.FindAll("devices", query, &devices)

	if err != nil {
		return err
	}

	err_chan := make(chan error)
	for _, device := range devices {
		go UnsubscribeDeviceAsync(topic, device, err_chan)
	}

	for range devices {
		err := <-err_chan

		if err != nil {
			return err
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
	switch strings.ToLower(device_reg.Platform) {
	case "android":
		platform_arn = config.ANDROID_PLATFORM_ARN
	case "ios":
		platform_arn = config.IOS_PLATFORM_ARN
	default:
		return errors.New("Invalid platform")
	}

	if config.IS_PRODUCTION {
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

	if config.IS_PRODUCTION {
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

	if !ok {
		return ERR_DEVICE_NOT_SUBSCRIBED
	}

	if config.IS_PRODUCTION {
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

	return nil
}

/*
	Attempts to subscribe a device to the given topic, returning an error or nil in the
	channel
*/
func SubscribeDeviceAsync(topic models.Topic, device models.Device, err_chan chan error) {
	err := SubscribeDeviceToTopic(topic, device)
	err_chan <- err
}

/*
	Attempts to unsubscribe a device to the given topic, returning an error or nil
	in the channel
*/
func UnsubscribeDeviceAsync(topic models.Topic, device models.Device, err_chan chan error) {
	err := UnsubscribeDeviceFromTopic(topic, device)
	err_chan <- err
}

/*
	Returns a list of registered devices
*/
func GetAllDevices() (*[]models.Device, error) {
	var devices []models.Device
	err := db.FindAll("devices", nil, &devices)

	if err != nil {
		return nil, err
	}

	return &devices, nil
}

func GenerateNotificationJson(notification models.Notification, past_notification models.PastNotification) (*string, error) {
	apns_payload := models.APNSPayload{
		Container: models.APNSContainer{
			Alert: models.APNSAlert{
				Title: notification.Title,
				Body:  notification.Body,
			},
			Sound: "default",
		},
		Data: past_notification,
	}

	gcm_payload := models.GCMPayload{
		Notification: models.GCMNotification{
			Title: notification.Title,
			Body:  notification.Body,
		},
	}

	apns_payload_json, err := json.Marshal(apns_payload)

	if err != nil {
		return nil, err
	}

	gcm_payload_json, err := json.Marshal(gcm_payload)

	if err != nil {
		return nil, err
	}

	notification_payload := models.NotificationPayload{
		APNS:        string(apns_payload_json),
		APNSSandbox: string(apns_payload_json),
		GCM:         string(gcm_payload_json),
		Default:     notification.Body,
	}

	notification_json, err := json.Marshal(notification_payload)

	if err != nil {
		return nil, err
	}

	notification_json_str := string(notification_json)

	return &notification_json_str, nil
}
