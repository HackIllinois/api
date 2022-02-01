package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/common/metrics"
	"github.com/HackIllinois/api/common/utils"
	"github.com/HackIllinois/api/services/notifications/models"
	"github.com/HackIllinois/api/services/notifications/service"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.Handle("/internal/metrics/", promhttp.Handler()).Methods("GET")

	metrics.RegisterHandler("/topic/", GetAllTopics, "GET", router)
	metrics.RegisterHandler("/topic/", CreateTopic, "POST", router)
	metrics.RegisterHandler("/topic/all/", GetAllNotifications, "GET", router)
	metrics.RegisterHandler("/topic/public/", GetAllPublicNotifications, "GET", router)
	metrics.RegisterHandler("/topic/{id}/", GetNotificationsForTopic, "GET", router)
	metrics.RegisterHandler("/topic/{id}/", PublishNotificationToTopic, "POST", router)
	metrics.RegisterHandler("/topic/{id}/", DeleteTopic, "DELETE", router)
	metrics.RegisterHandler("/topic/{id}/subscribe/", SubscribeToTopic, "POST", router)
	metrics.RegisterHandler("/topic/{id}/unsubscribe/", UnsubscribeToTopic, "POST", router)
	metrics.RegisterHandler("/device/", RegisterDeviceToUser, "POST", router)
	metrics.RegisterHandler("/order/{id}/", GetNotificationOrder, "GET", router)
}

/*
	Returns all topics that notifications can be published to
*/
func GetAllTopics(w http.ResponseWriter, r *http.Request) {
	topics, err := service.GetAllTopicIDs()

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not retrieve topics."))
		return
	}

	role_topics, err := service.GetValidRoles()

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not retrieve role based topics."))
		return
	}

	topics = append(topics, role_topics.Roles...)

	topic_list := models.TopicList{
		Topics: topics,
	}

	json.NewEncoder(w).Encode(topic_list)
}

/*
	Creates a topic with the given id and returns it
*/
func CreateTopic(w http.ResponseWriter, r *http.Request) {
	var topic models.Topic
	json.NewDecoder(r.Body).Decode(&topic)

	err := service.CreateTopic(topic.ID)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not create a new topic."))
		return
	}

	created_topic, err := service.GetTopic(topic.ID)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not retrieve topic."))
		return
	}

	json.NewEncoder(w).Encode(created_topic)
}

/*
	Returns all notifications to topics the user is subscribed to
*/
func GetAllNotifications(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	topics, err := service.GetSubscriptions(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not retrieve user subscriptions."))
		return
	}

	notifications, err := service.GetAllNotifications(topics)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not retrieve notifications."))
		return
	}

	notification_list := models.NotificationList{
		Notifications: notifications,
	}

	json.NewEncoder(w).Encode(notification_list)
}

/*
	Returns all public notifications
*/
func GetAllPublicNotifications(w http.ResponseWriter, r *http.Request) {
	notifications, err := service.GetAllPublicNotifications()

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not retrieve notifications."))
		return
	}

	notification_list := models.NotificationList{
		Notifications: notifications,
	}

	json.NewEncoder(w).Encode(notification_list)
}

/*
	Returns all notifications for the specified topic
*/
func GetNotificationsForTopic(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	notifications, err := service.GetAllNotificationsForTopic(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not retrieve notifications."))
		return
	}

	notification_list := models.NotificationList{
		Notifications: notifications,
	}

	json.NewEncoder(w).Encode(notification_list)
}

/*
	Publishes a notification to the specied topic and returns the notification
*/
func PublishNotificationToTopic(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var notification models.Notification
	json.NewDecoder(r.Body).Decode(&notification)

	notification.Topic = id
	notification.ID = utils.GenerateUniqueID()
	notification.Time = time.Now().Unix()

	order, err := service.PublishNotificationToTopic(notification)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not publish notification."))
		return
	}

	json.NewEncoder(w).Encode(order)
}

/*
	Deletes the specified topic and returns it
*/
func DeleteTopic(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := service.DeleteTopic(id)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not publish notification."))
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{})
}

/*
	Subscribes a user to the specied topic and returns their updated subscriptions
*/
func SubscribeToTopic(w http.ResponseWriter, r *http.Request) {
	topicId := mux.Vars(r)["id"]
	userId := r.Header.Get("HackIllinois-Identity")

	err := service.SubscribeToTopic(userId, topicId)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Failed to subscribe user to topic."))
		return
	}

	subscriptions, err := service.GetSubscriptions(userId)

	topic_list := models.TopicList{
		Topics: subscriptions,
	}

	json.NewEncoder(w).Encode(topic_list)
}

/*
	Unsubscribes a user to the specied topic and returns their updated subscriptions
*/
func UnsubscribeToTopic(w http.ResponseWriter, r *http.Request) {
	topicId := mux.Vars(r)["id"]
	userId := r.Header.Get("HackIllinois-Identity")

	err := service.UnsubscribeToTopic(userId, topicId)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Failed to unsubscribe user to topic."))
		return
	}

	subscriptions, err := service.GetSubscriptions(userId)

	topic_list := models.TopicList{
		Topics: subscriptions,
	}

	json.NewEncoder(w).Encode(topic_list)
}

/*
	Registered the specified device token to the user
*/
func RegisterDeviceToUser(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	var device_registration models.DeviceRegistration
	json.NewDecoder(r.Body).Decode(&device_registration)

	err := service.RegisterDeviceToUser(device_registration.Token, device_registration.Platform, id)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Failed to register device to user."))
		return
	}

	devices, err := service.GetUserDevices(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Failed to retrieve user's devices."))
		return
	}

	device_list := models.DeviceList{
		Devices: devices,
	}

	json.NewEncoder(w).Encode(device_list)
}

/*
	Returns the notification order with the specified id
*/
func GetNotificationOrder(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	order, err := service.GetNotificationOrder(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not retrieve notification order."))
		return
	}

	json.NewEncoder(w).Encode(order)
}
