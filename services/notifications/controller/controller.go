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
	router.Handle("/{name}/", alice.New().ThenFunc(GetNotificationsForTopic)).Methods("GET")
	router.Handle("/{name}/", alice.New().ThenFunc(DeleteTopic)).Methods("DELETE")
	router.Handle("/{name}/", alice.New().ThenFunc(PublishNotification)).Methods("POST")
	router.Handle("/{name}/info/", alice.New().ThenFunc(GetTopicInfo)).Methods("GET")
}

/*
	Endpoint to get all SNS Topics
*/
func GetAllTopics(w http.ResponseWriter, r *http.Request) {
	topics, err := service.GetAllTopics()

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(topics)
}

/*
	Endpoint to get all past notifications
*/
func GetAllNotifications(w http.ResponseWriter, r *http.Request) {
	notifications_list, err := service.GetAllNotifications()

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
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
		panic(errors.UnprocessableError(err.Error()))
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
		panic(errors.UnprocessableError(err.Error()))
	}

	topics, err := service.GetAllTopics()

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
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
		panic(errors.UnprocessableError(err.Error()))
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
		panic(errors.UnprocessableError(err.Error()))
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
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(topic)
}
