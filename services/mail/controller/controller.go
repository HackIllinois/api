package controller

import (
	"encoding/json"
	"github.com/HackIllinois/api/common/errors"
	"github.com/HackIllinois/api/services/mail/models"
	"github.com/HackIllinois/api/services/mail/service"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"net/http"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.Handle("/send/", alice.New().ThenFunc(SendMail)).Methods("POST")
	router.Handle("/send/list/", alice.New().ThenFunc(SendMailList)).Methods("POST")
	router.Handle("/list/", alice.New().ThenFunc(GetAllMailLists)).Methods("GET")
	router.Handle("/list/create/", alice.New().ThenFunc(CreateMailList)).Methods("POST")
	router.Handle("/list/add/", alice.New().ThenFunc(AddToMailList)).Methods("POST")
	router.Handle("/list/remove/", alice.New().ThenFunc(RemoveFromMailList)).Methods("POST")
	router.Handle("/list/{id}/", alice.New().ThenFunc(GetMailList)).Methods("GET")
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
		panic(errors.INTERNAL_ERROR("Could not send email by ID."))
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
		panic(errors.INTERNAL_ERROR("Could not send email by list."))
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
		panic(errors.INTERNAL_ERROR(err.Error()))
	}

	created_list, err := service.GetMailList(mail_list.ID)

	if err != nil {
		panic(errors.DATABASE_ERROR("Could not get mail list."))
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
		panic(errors.DATABASE_ERROR("Could not add user to mail list."))
	}

	modified_list, err := service.GetMailList(mail_list.ID)

	if err != nil {
		panic(errors.DATABASE_ERROR("Could not get modified mail list."))
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
		panic(errors.DATABASE_ERROR("Could not remove user from mailing list."))
	}

	modified_list, err := service.GetMailList(mail_list.ID)

	if err != nil {
		panic(errors.DATABASE_ERROR("Could not get modified mail list."))
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
		panic(errors.DATABASE_ERROR("Could not get mail list."))
	}

	json.NewEncoder(w).Encode(mail_list)
}

/*
	Endpoint to get all mailing lists
*/
func GetAllMailLists(w http.ResponseWriter, r *http.Request) {
	mail_lists, err := service.GetAllMailLists()

	if err != nil {
		panic(errors.DATABASE_ERROR("Could not get all mail lists."))
	}

	json.NewEncoder(w).Encode(mail_lists)
}
