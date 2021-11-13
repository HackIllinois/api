package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/common/utils"
	"github.com/HackIllinois/api/services/notifications/models"
	"github.com/HackIllinois/api/services/notifications/service"
	"github.com/gorilla/mux"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.HandleFunc("/topic/", GetAllTopics).Methods("GET")
	router.HandleFunc("/topic/", CreateTopic).Methods("POST")
	router.HandleFunc("/topic/all/", GetAllNotifications).Methods("GET")
	router.HandleFunc("/topic/public/", GetAllPublicNotifications).Methods("GET")
	router.HandleFunc("/topic/{id}/", GetNotificationsForTopic).Methods("GET")
	router.HandleFunc("/topic/{id}/", PublishNotificationToTopic).Methods("POST")
	router.HandleFunc("/topic/{id}/", DeleteTopic).Methods("DELETE")
	router.HandleFunc("/topic/{id}/subscribe/", SubscribeToTopic).Methods("POST")
	router.HandleFunc("/topic/{id}/unsubscribe/", UnsubscribeToTopic).Methods("POST")
	router.HandleFunc("/device/", RegisterDeviceToUser).Methods("POST")
	router.HandleFunc("/order/{id}/", GetNotificationOrder).Methods("GET")
}

/*
	Returns all topics a user is subscribed to
*/
func GetUserSubscriptions(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")
	topics, err := service.GetSubscriptions(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not retrieve user's subscriptions."))
		return
	}

	topic_list := models.TopicList{
		Topics: topics,
	}

	json.NewEncoder(w).Encode(topic_list)
}

/*
	Returns devices registered of a specific user
*/
func GetUserRegisteredDevices(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")
	devices, err := service.GetUserDevices(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not retrieve user's registered devices."))
		return
	}

	device_list := models.DeviceList{
		Devices: devices,
	}

	json.NewEncoder(w).Encode(device_list)
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
