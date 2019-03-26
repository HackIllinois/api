package controller

import (
	"encoding/json"
	"github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/common/utils"
	"github.com/HackIllinois/api/services/notifications/models"
	"github.com/HackIllinois/api/services/notifications/service"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"net/http"
	"time"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.Handle("/topic/", alice.New().ThenFunc(GetAllTopics)).Methods("GET")
	router.Handle("/topic/", alice.New().ThenFunc(CreateTopic)).Methods("POST")
	router.Handle("/topic/all/", alice.New().ThenFunc(GetAllNotifications)).Methods("GET")
	router.Handle("/topic/public/", alice.New().ThenFunc(GetAllPublicNotifications)).Methods("GET")
	router.Handle("/topic/{id}/", alice.New().ThenFunc(GetNotificationsForTopic)).Methods("GET")
	router.Handle("/topic/{id}/", alice.New().ThenFunc(PublishNotificationToTopic)).Methods("POST")
	router.Handle("/topic/{id}/", alice.New().ThenFunc(DeleteTopic)).Methods("DELETE")
	router.Handle("/topic/{id}/subscribe/", alice.New().ThenFunc(SubscribeToTopic)).Methods("POST")
	router.Handle("/topic/{id}/unsubscribe/", alice.New().ThenFunc(UnsubscribeToTopic)).Methods("POST")
	router.Handle("/device/", alice.New().ThenFunc(RegisterDeviceToUser)).Methods("POST")
	router.Handle("/order/{id}/", alice.New().ThenFunc(GetNotificationOrder)).Methods("GET")
}

/*
	Returns all topics that notifications can be published to
*/
func GetAllTopics(w http.ResponseWriter, r *http.Request) {
	topics, err := service.GetAllTopicIDs()

	if err != nil {
		panic(errors.DatabaseError(err.Error(), "Could not retrieve topics."))
	}

	role_topics, err := service.GetValidRoles()

	if err != nil {
		panic(errors.DatabaseError(err.Error(), "Could not retrieve role based topics."))
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
		panic(errors.DatabaseError(err.Error(), "Could not create a new topic."))
	}

	created_topic, err := service.GetTopic(topic.ID)

	if err != nil {
		panic(errors.DatabaseError(err.Error(), "Could not retrieve topic."))
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
		panic(errors.DatabaseError(err.Error(), "Could not retrieve user subscriptions."))
	}

	notifications, err := service.GetAllNotifications(topics)

	if err != nil {
		panic(errors.DatabaseError(err.Error(), "Could not retrieve notifications."))
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
		panic(errors.DatabaseError(err.Error(), "Could not retrieve notifications."))
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
		panic(errors.DatabaseError(err.Error(), "Could not retrieve notifications."))
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
		panic(errors.InternalError(err.Error(), "Could not publish notification."))
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
		panic(errors.InternalError(err.Error(), "Could not publish notification."))
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
		panic(errors.DatabaseError(err.Error(), "Failed to subscribe user to topic."))
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
		panic(errors.DatabaseError(err.Error(), "Failed to unsubscribe user to topic."))
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
		panic(errors.InternalError(err.Error(), "Failed to register device to user."))
	}

	devices, err := service.GetUserDevices(id)

	if err != nil {
		panic(errors.DatabaseError(err.Error(), "Failed to retrieve user's devices."))
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
		panic(errors.DatabaseError(err.Error(), "Could not retrieve notification order."))
	}

	json.NewEncoder(w).Encode(order)
}
