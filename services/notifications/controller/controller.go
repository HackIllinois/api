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

	router.Handle("/topics/", alice.New().ThenFunc(GetAllTopics)).Methods("GET")
	router.Handle("/topics/create/", alice.New().ThenFunc(CreateTopic)).Methods("POST")
	router.Handle("/topics/delete/", alice.New().ThenFunc(DeleteTopic)).Methods("POST")
	router.Handle("/publish/", alice.New().ThenFunc(PublishNotification)).Methods("POST")
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
	Endpoint to create a new SNS topic
*/
func CreateTopic(w http.ResponseWriter, r *http.Request) {
	var topic_name models.TopicName
	json.NewDecoder(r.Body).Decode(&topic_name)

	topic_arn, err := service.CreateTopic(topic_name.Name)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(topic_arn)
}

/*
	Endpoint to delete a SNS topic
*/
func DeleteTopic(w http.ResponseWriter, r *http.Request) {
	var topic_arn models.TopicArn
	json.NewDecoder(r.Body).Decode(&topic_arn)

	err := service.DeleteTopic(topic_arn.Arn)

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
	var notification models.Notification
	json.NewDecoder(r.Body).Decode(&notification)

	message_id, err := service.PublishNotification(notification)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(message_id)
}
