package service

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/common/utils"
	"github.com/HackIllinois/api/services/notifications/config"
	"github.com/HackIllinois/api/services/notifications/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"go.mongodb.org/mongo-driver/bson"
)

var SNS_MESSAGE_STRUCTURE string = "json"
var WORKER_POOL_SIZE int = 128

var sess *session.Session
var client *sns.SNS
var db database.Database

func Initialize() error {
	sess = session.Must(session.NewSession(&aws.Config{
		Region: aws.String(config.SNS_REGION),
	}))
	client = sns.New(sess)

	if db != nil {
		db.Close()
		db = nil
	}

	var err error
	db, err = database.InitDatabase(config.NOTIFICATIONS_DB_HOST, config.NOTIFICATIONS_DB_NAME)

	if err != nil {
		return err
	}

	return nil
}

/*
	Returns a list of all topic ids
*/
func GetAllTopicIDs() ([]string, error) {
	var topics []models.Topic
	err := db.FindAll("topics", nil, &topics, nil)

	if err != nil {
		return nil, err
	}

	topicIds := make([]string, len(topics))

	for i, topic := range topics {
		topicIds[i] = topic.ID
	}

	return topicIds, nil
}

/*
	Returns the topic with the specified id
*/
func GetTopic(id string) (*models.Topic, error) {
	selector := database.QuerySelector{
		"id": id,
	}

	var topic models.Topic
	err := db.FindOne("topics", selector, &topic, nil)

	if err != nil {
		return nil, err
	}

	return &topic, nil
}

/*
	Creates a topic
*/
func CreateTopic(id string) error {
	_, err := GetTopic(id)

	if err != database.ErrNotFound {
		if err != nil {
			return err
		}

		return errors.New("Topic already exists.")
	}

	topic := models.Topic{
		ID:      id,
		UserIDs: []string{},
	}

	err = db.Insert("topics", &topic, nil)

	if err != nil {
		return err
	}

	return nil
}

/*
	Deletes a topic
*/
func DeleteTopic(id string) error {
	selector := database.QuerySelector{
		"id": id,
	}

	err := db.RemoveOne("topics", selector, nil)

	if err != nil {
		return err
	}

	return nil
}

/*
	Returns all notification for the specified topic
*/
func GetAllNotificationsForTopic(topic string) ([]models.Notification, error) {
	selector := database.QuerySelector{
		"topic": topic,
	}

	var notifications []models.Notification
	err := db.FindAll("notifications", selector, &notifications, nil)

	if err != nil {
		return nil, err
	}

	return notifications, nil
}

/*
	Returns all notifications for the specified topics
*/
func GetAllNotifications(topics []string) ([]models.Notification, error) {
	notifications := make([]models.Notification, 0)

	for _, topic := range topics {
		topic_notifications, err := GetAllNotificationsForTopic(topic)

		if err != nil {
			return nil, err
		}

		notifications = append(notifications, topic_notifications...)
	}

	return notifications, nil
}

/*
	Returns all public notifications
*/
func GetAllPublicNotifications() ([]models.Notification, error) {
	return GetAllNotifications([]string{"User", "Attendee"})
}

/*
	Returns the list of topics the user is subscribed to
*/
func GetSubscriptions(id string) ([]string, error) {
	selector := database.QuerySelector{
		"userids": database.QuerySelector{
			"$elemMatch": database.QuerySelector{
				"$eq": id,
			},
		},
	}

	var topics []models.Topic
	err := db.FindAll("topics", selector, &topics, nil)

	if err != nil {
		return nil, err
	}

	topicIds := make([]string, len(topics))

	for i, topic := range topics {
		topicIds[i] = topic.ID
	}

	roles, err := GetUserRoles(id)

	if err != nil {
		return nil, err
	}

	topicIds = append(topicIds, roles...)

	return topicIds, nil
}

/*
	Subscribes the user to the specified topic
*/
func SubscribeToTopic(userId string, topicId string) error {
	selector := database.QuerySelector{
		"id": topicId,
	}

	modifier := bson.M{
		"$addToSet": bson.M{
			"userids": userId,
		},
	}

	err := db.Update("topics", selector, &modifier, nil)

	if err != nil {
		return err
	}

	return nil
}

/*
	Unsubscribes the user to the specified topic
*/
func UnsubscribeToTopic(userId string, topicId string) error {
	selector := database.QuerySelector{
		"id": topicId,
	}

	modifier := bson.M{
		"$pull": bson.M{
			"userids": userId,
		},
	}

	err := db.Update("topics", selector, &modifier, nil)

	if err != nil {
		return err
	}

	return nil
}

/*
	Gets the list of devices registered to a user
*/
func GetUserDevices(id string) ([]string, error) {
	selector := database.QuerySelector{
		"id": id,
	}

	var user models.User
	err := db.FindOne("users", selector, &user, nil)

	if err != nil {
		if err == database.ErrNotFound {
			err = db.Insert("users", &models.User{
				ID:      id,
				Devices: []string{},
			}, nil)

			if err != nil {
				return nil, err
			}

			err = db.FindOne("users", selector, &user, nil)

			if err != nil {
				return nil, err
			}
		}
		return nil, err
	}

	return user.Devices, nil
}

/*
	Sets the list of devices registered to a user
*/
func SetUserDevices(id string, devices []string) error {
	selector := database.QuerySelector{
		"id": id,
	}

	user := models.User{
		ID:      id,
		Devices: devices,
	}

	err := db.Replace("users", selector, &user, false, nil)

	if err != nil {
		return err
	}

	return nil
}

/*
	Registers the device token with SNS and stores the arn with the associated user
*/
func RegisterDeviceToUser(token string, platform string, id string) error {
	var platform_arn string

	switch strings.ToLower(platform) {
	case "android":
		platform_arn = config.ANDROID_PLATFORM_ARN
	case "ios":
		platform_arn = config.IOS_PLATFORM_ARN
	default:
		return errors.New("Invalid platform")
	}

	var device_arn string

	if config.IS_PRODUCTION {
		response, err := client.CreatePlatformEndpoint(
			&sns.CreatePlatformEndpointInput{
				CustomUserData:         &id,
				Token:                  &token,
				PlatformApplicationArn: &platform_arn,
			},
		)

		if err != nil {
			return err
		}

		device_arn = *response.EndpointArn
	}

	devices, err := GetUserDevices(id)

	if err != nil {
		return err
	}

	if !utils.ContainsString(devices, device_arn) {
		devices = append(devices, device_arn)
	}

	err = SetUserDevices(id, devices)

	if err != nil {
		return err
	}

	return nil
}

/*
	Returns a list of userids to receive a notification to the specified topic
*/
func GetNotificationRecipients(topicId string) ([]string, error) {
	topic, err := GetTopic(topicId)

	if err != nil {
		if err == database.ErrNotFound {
			usersIds, err := GetUsersByRole(topicId)

			if err != nil {
				return nil, err
			}

			return usersIds, nil
		}
		return nil, err
	}

	return topic.UserIDs, nil
}

/*
	Returns a list of arns to receive a notification
*/
func GetNotificationRecipientArns(userIds []string) ([]string, error) {
	device_arns := make([]string, 0)

	for _, userId := range userIds {
		devices, err := GetUserDevices(userId)

		if err != nil {
			return nil, err
		}

		device_arns = append(device_arns, devices...)
	}

	return device_arns, nil
}

/*
	Returns the notification order with the specified id
*/
func GetNotificationOrder(id string) (*models.NotificationOrder, error) {
	selector := database.QuerySelector{
		"id": id,
	}

	var order models.NotificationOrder
	err := db.FindOne("orders", selector, &order, nil)

	if err != nil {
		return nil, err
	}

	return &order, nil
}

/*
	Publishes a notification to the specified topic
*/
func PublishNotificationToTopic(notification models.Notification) (*models.NotificationOrder, error) {
	err := db.Insert("notifications", &notification, nil)

	if err != nil {
		return nil, err
	}

	recipients, err := GetNotificationRecipients(notification.Topic)

	if err != nil {
		return nil, err
	}

	device_arns, err := GetNotificationRecipientArns(recipients)

	if err != nil {
		return nil, err
	}

	notification_payload, err := GenerateNotificationJson(notification)

	if err != nil {
		return nil, err
	}

	if config.IS_PRODUCTION {
		go PublishNotification(notification.ID, notification_payload, device_arns)
	}

	order := models.NotificationOrder{
		ID:         notification.ID,
		Recipients: len(device_arns),
		Success:    0,
		Failure:    0,
		Time:       notification.Time,
	}

	err = db.Insert("orders", &order, nil)

	if err != nil {
		return nil, err
	}

	return &order, nil
}

/*
	Publishes the notification payload to all specified arns
*/
func PublishNotification(id string, payload string, arns []string) error {
	success_count := 0
	failure_count := 0

	queued_devices := make(chan string, len(arns))
	responses := make(chan bool, len(arns))

	for i := 0; i < WORKER_POOL_SIZE; i++ {
		go PublishNotificationWorker(payload, queued_devices, responses)
	}

	for _, device_arn := range arns {
		queued_devices <- device_arn
	}

	close(queued_devices)

	for i := 0; i < len(arns); i++ {
		response := <-responses

		if response {
			success_count++
		} else {
			failure_count++
		}
	}

	close(responses)

	order, err := GetNotificationOrder(id)

	if err != nil {
		return err
	}

	order.Success = success_count
	order.Failure = failure_count

	selector := database.QuerySelector{
		"id": id,
	}

	err = db.Replace("orders", selector, &order, false, nil)

	if err != nil {
		return err
	}

	return nil
}

/*
	Worker go routine to publish notifications
*/
func PublishNotificationWorker(notification string, device_arns <-chan string, responses chan<- bool) {
	for device_arn := range device_arns {
		_, err := client.Publish(&sns.PublishInput{
			TargetArn:        &device_arn,
			Message:          &notification,
			MessageStructure: &SNS_MESSAGE_STRUCTURE,
		})

		responses <- (err == nil)
	}
}

/*
	Generates the notification payload for SNS
*/
func GenerateNotificationJson(notification models.Notification) (string, error) {
	apns_payload := models.APNSPayload{
		Container: models.APNSContainer{
			Alert: models.APNSAlert{
				Title: notification.Title,
				Body:  notification.Body,
			},
			Sound: "default",
		},
		Data: notification,
	}

	gcm_payload := models.GCMPayload{
		Notification: models.GCMNotification{
			Title: notification.Title,
			Body:  notification.Body,
		},
	}

	apns_payload_json, err := json.Marshal(apns_payload)

	if err != nil {
		return "", err
	}

	gcm_payload_json, err := json.Marshal(gcm_payload)

	if err != nil {
		return "", err
	}

	notification_payload := models.NotificationPayload{
		APNS:        string(apns_payload_json),
		APNSSandbox: string(apns_payload_json),
		GCM:         string(gcm_payload_json),
		Default:     notification.Body,
	}

	notification_json, err := json.Marshal(notification_payload)

	if err != nil {
		return "", err
	}

	return string(notification_json), nil
}
