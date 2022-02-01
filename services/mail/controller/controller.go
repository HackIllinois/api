package controller

import (
	"encoding/json"
	"net/http"

	"github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/common/metrics"
	"github.com/HackIllinois/api/services/mail/models"
	"github.com/HackIllinois/api/services/mail/service"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.Handle("/internal/metrics/", promhttp.Handler()).Methods("GET")

	metrics.RegisterHandler("/send/", SendMail, "POST", router)
	metrics.RegisterHandler("/send/list/", SendMailList, "POST", router)
	metrics.RegisterHandler("/list/", GetAllMailLists, "GET", router)
	metrics.RegisterHandler("/list/create/", CreateMailList, "POST", router)
	metrics.RegisterHandler("/list/add/", AddToMailList, "POST", router)
	metrics.RegisterHandler("/list/remove/", RemoveFromMailList, "POST", router)
	metrics.RegisterHandler("/list/{id}/", GetMailList, "GET", router)
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
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not send email by ID."))
		return
	}

	json.NewEncoder(w).Encode(mail_status)
}

/*
	Endpoint to send mail to a specified mail list,
	based on a specified template
*/
func SendMailList(w http.ResponseWriter, r *http.Request) {
	var mail_order_list models.MailOrderList
	json.NewDecoder(r.Body).Decode(&mail_order_list)

	mail_status, err := service.SendMailByList(mail_order_list)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not send email by list."))
		return
	}

	json.NewEncoder(w).Encode(mail_status)
}

/*
	Endpoint to create a mailing list
*/
func CreateMailList(w http.ResponseWriter, r *http.Request) {
	var mail_list models.MailList
	json.NewDecoder(r.Body).Decode(&mail_list)

	err := service.CreateMailList(mail_list)

	if err != nil {
		errors.WriteError(w, r, errors.InternalError(err.Error(), "Could not create the specified mail list."))
		return
	}

	created_list, err := service.GetMailList(mail_list.ID)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get mail list."))
		return
	}

	json.NewEncoder(w).Encode(created_list)
}

/*
	Endpoint to add to a mailing list
*/
func AddToMailList(w http.ResponseWriter, r *http.Request) {
	var mail_list models.MailList
	json.NewDecoder(r.Body).Decode(&mail_list)

	err := service.AddToMailList(mail_list)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not add user to mail list."))
		return
	}

	modified_list, err := service.GetMailList(mail_list.ID)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get modified mail list."))
		return
	}

	json.NewEncoder(w).Encode(modified_list)
}

/*
	Endpoint to remove from a mailing list
*/
func RemoveFromMailList(w http.ResponseWriter, r *http.Request) {
	var mail_list models.MailList
	json.NewDecoder(r.Body).Decode(&mail_list)

	err := service.RemoveFromMailList(mail_list)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not remove user from mailing list."))
		return
	}

	modified_list, err := service.GetMailList(mail_list.ID)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get modified mail list."))
		return
	}

	json.NewEncoder(w).Encode(modified_list)
}

/*
	Endpoint to get a mailing list by id
*/
func GetMailList(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	mail_list, err := service.GetMailList(id)

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get mail list."))
		return
	}

	json.NewEncoder(w).Encode(mail_list)
}

/*
	Endpoint to get all mailing lists
*/
func GetAllMailLists(w http.ResponseWriter, r *http.Request) {
	mail_lists, err := service.GetAllMailLists()

	if err != nil {
		errors.WriteError(w, r, errors.DatabaseError(err.Error(), "Could not get all mail lists."))
		return
	}

	json.NewEncoder(w).Encode(mail_lists)
}
