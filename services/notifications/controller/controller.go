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
}

/*
	Returns all topics that notifications can be published to
*/
func GetAllTopics(w http.ResponseWriter, r *http.Request) {

}

/*
	Creates a topic with the given id and returns it
*/
func CreateTopic(w http.ResponseWriter, r *http.Request) {

}

/*
	Returns all notifications to topics the user is subscribed to
*/
func GetAllNotifications(w http.ResponseWriter, r *http.Request) {

}

/*
	Returns all public notifications
*/
func GetAllPublicNotifications(w http.ResponseWriter, r *http.Request) {

}

/*
	Returns all notifications for the specified topic
*/
func GetNotificationsForTopic(w http.ResponseWriter, r *http.Request) {

}

/*
	Publishes a notification to the specied topic and returns the notification
*/
func PublishNotificationToTopic(w http.ResponseWriter, r *http.Request) {

}

/*
	Deletes the specified topic and returns it
*/
func DeleteTopic(w http.ResponseWriter, r *http.Request) {

}

/*
	Subscribes a user to the specied topic and returns their updated subscriptions
*/
func SubscribeToTopic(w http.ResponseWriter, r *http.Request) {

}

/*
	Unsubscribes a user to the specied topic and returns their updated subscriptions
*/
func UnsubscribeToTopic(w http.ResponseWriter, r *http.Request) {

}

/*
	Registered the specified device token to the user
*/
func RegisterDeviceToUser(w http.ResponseWriter, r *http.Request) {

}
