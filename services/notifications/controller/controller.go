package controller

import (
	"encoding/json"
	"github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/services/notifications/models"
	"github.com/HackIllinois/api/services/notifications/service"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"net/http"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.Handle("/", alice.New().ThenFunc(GetAllTopics)).Methods("GET")
	router.Handle("/", alice.New().ThenFunc(CreateTopic)).Methods("POST")
	router.Handle("/all/", alice.New().ThenFunc(GetAllNotifications)).Methods("GET")
	router.Handle("/device/", alice.New().ThenFunc(RegisterDeviceToUser)).Methods("POST")
	router.Handle("/update/", alice.New().ThenFunc(UpdateUserSubscriptions)).Methods("POST")
	router.Handle("/{name}/", alice.New().ThenFunc(GetNotificationsForTopic)).Methods("GET")
	router.Handle("/{name}/", alice.New().ThenFunc(DeleteTopic)).Methods("DELETE")
	router.Handle("/{name}/", alice.New().ThenFunc(PublishNotification)).Methods("POST")
	router.Handle("/{name}/add/", alice.New().ThenFunc(AddUsersToTopic)).Methods("POST")
	router.Handle("/{name}/remove/", alice.New().ThenFunc(RemoveUsersFromTopic)).Methods("POST")
	router.Handle("/{name}/info/", alice.New().ThenFunc(GetTopicInfo)).Methods("GET")
}

/*
	Endpoint to get all SNS Topics
*/
func GetAllTopics(w http.ResponseWriter, r *http.Request) {
	topics, err := service.GetAllTopics()

	if err != nil {
		panic(errors.DatabaseError(err.Error(), "Could not get all SNS topics."))
	}

	json.NewEncoder(w).Encode(topics)
}

/*
	Endpoint to get all past notifications
*/
func GetAllNotifications(w http.ResponseWriter, r *http.Request) {
	notifications_list, err := service.GetAllNotifications()

	if err != nil {
		panic(errors.DatabaseError(err.Error(), "Could not get all past notifications."))
	}

	json.NewEncoder(w).Encode(notifications_list)
}

/*
	Endpoint to create a new SNS topic
*/
func CreateTopic(w http.ResponseWriter, r *http.Request) {
	var topic_name models.TopicName
	json.NewDecoder(r.Body).Decode(&topic_name)

	err := service.CreateTopic(topic_name.Name)

	if err != nil {
		panic(errors.DatabaseError(err.Error(), "Could not create a new SNS topic."))
	}

	json.NewEncoder(w).Encode(topic_name)
}

/*
	Endpoint to delete a SNS topic
*/
func DeleteTopic(w http.ResponseWriter, r *http.Request) {
	topic_name := mux.Vars(r)["name"]

	err := service.DeleteTopic(topic_name)

	if err != nil {
		panic(errors.DatabaseError(err.Error(), "Could not delete topic."))
	}

	topics, err := service.GetAllTopics()

	if err != nil {
		panic(errors.DatabaseError(err.Error(), "Could not fetch updated topics."))
	}

	json.NewEncoder(w).Encode(topics)
}

/*
	Endpoint to create a new notification
*/
func PublishNotification(w http.ResponseWriter, r *http.Request) {
	topic_name := mux.Vars(r)["name"]
	var notification models.Notification
	json.NewDecoder(r.Body).Decode(&notification)

	past_notification, err := service.PublishNotification(topic_name, notification)

	if err != nil {
		panic(errors.InternalError(err.Error(), "Could not publish new notification."))
	}

	json.NewEncoder(w).Encode(past_notification)
}

/*
	Endpoint to get all past notifications for a given Topic
*/
func GetNotificationsForTopic(w http.ResponseWriter, r *http.Request) {
	topic_name := mux.Vars(r)["name"]

	notifications_list, err := service.GetNotificationsForTopic(topic_name)

	if err != nil {
		panic(errors.DatabaseError(err.Error(), "Could not get all past notifications for the given topic."))
	}

	json.NewEncoder(w).Encode(notifications_list)
}

/*
	Endpoint to get name, ARN for a topic
*/
func GetTopicInfo(w http.ResponseWriter, r *http.Request) {
	topic_name := mux.Vars(r)["name"]

	topic, err := service.GetTopicInfo(topic_name)

	if err != nil {
		panic(errors.DatabaseError(err.Error(), "Could not get name / ARN for topic."))
	}

	var topic_public models.TopicPublic
	if topic != nil {
		topic_public = models.TopicPublic{Name: topic.Name, UserIDs: topic.UserIDs}
	}

	json.NewEncoder(w).Encode(topic_public)
}

/*
	Adds users with given userids to the specified topic
*/
func AddUsersToTopic(w http.ResponseWriter, r *http.Request) {
	topic_name := mux.Vars(r)["name"]

	var userid_list models.UserIDList
	json.NewDecoder(r.Body).Decode(&userid_list)

	err := service.AddUsersToTopic(topic_name, userid_list)

	if err != nil {
		panic(errors.DatabaseError(err.Error(), "Could not add users to specified topic."))
	}

	topic, err := service.GetTopicInfo(topic_name)

	if err != nil {
		panic(errors.DatabaseError(err.Error(), "Could not get name / ARN for topic."))
	}

	var topic_public models.TopicPublic
	if topic != nil {
		topic_public = models.TopicPublic{Name: topic.Name, UserIDs: topic.UserIDs}
	}

	json.NewEncoder(w).Encode(topic_public)
}

/*
	Removes users with given userids from the specified topic
*/
func RemoveUsersFromTopic(w http.ResponseWriter, r *http.Request) {
	topic_name := mux.Vars(r)["name"]

	var userid_list models.UserIDList
	json.NewDecoder(r.Body).Decode(&userid_list)

	err := service.RemoveUsersFromTopic(topic_name, userid_list)

	if err != nil {
		panic(errors.DatabaseError(err.Error(), "Could not remove given users from topic."))
	}

	topic, err := service.GetTopicInfo(topic_name)

	if err != nil {
		panic(errors.DatabaseError(err.Error(), "Could not get name / ARN for topic."))
	}

	var topic_public models.TopicPublic
	if topic != nil {
		topic_public = models.TopicPublic{Name: topic.Name, UserIDs: topic.UserIDs}
	}

	json.NewEncoder(w).Encode(topic_public)
}

/*
	Endpoint to register a device token to a given user
*/
func RegisterDeviceToUser(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	if id == "" {
		panic(errors.MalformedRequestError("Must provide id to register a device token with.", "Must provide id to register a device token with."))
	}

	var device_registration models.DeviceRegistration
	json.NewDecoder(r.Body).Decode(&device_registration)

	err := service.RegisterDeviceToUser(id, device_registration)

	if err != nil {
		panic(errors.InternalError(err.Error(), "Could not register device to user."))
	}

	json.NewEncoder(w).Encode(device_registration)
}

/*
   Subscribes a user to topics corresponding to their roles, and unsubscribes a user from all other topics
*/
func UpdateUserSubscriptions(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	if id == "" {
		panic(errors.MalformedRequestError("Must provide id of user to update subscriptions", "Must provide id of user to update subscriptions"))
	}

	topic_list, err := service.UpdateUserSubscriptions(id)

	if err != nil {
		panic(errors.InternalError(err.Error(), "Could not update user subscriptions."))
	}

	json.NewEncoder(w).Encode(topic_list)
}
