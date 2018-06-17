package controller

import (
	"encoding/json"
	"github.com/HackIllinois/api-commons/errors"
	"github.com/HackIllinois/api-mail/models"
	"github.com/HackIllinois/api-mail/service"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"net/http"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.Handle("/mail/send/", alice.New().ThenFunc(SendMail)).Methods("POST")
}

/*
	Endpoint to send mail to a specified set of users,
	based on a specified template
*/
func SendMail(w http.ResponseWriter, r *http.Request) {
	var mail_order models.MailOrder
	json.NewDecoder(r.Body).Decode(&mail_order)

	mail_status, err := service.SendMailByID(mail_order)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(mail_status)
}
