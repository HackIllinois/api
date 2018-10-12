package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/HackIllinois/api/common/apirequest"
	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/services/mail/config"
	"github.com/HackIllinois/api/services/mail/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

var db database.MongoDatabase

func init() {
	db_connection, err := database.InitMongoDatabase(config.MAIL_DB_HOST, config.MAIL_DB_NAME)

	if err != nil {
		panic(err)
	}

	db = db_connection
}

/*
	Send mail to the users in the given mailing list, using the provided template
	Substitution will be generated based on user info
*/
func SendMailByList(mail_order_list models.MailOrderList) (*models.MailStatus, error) {
	mail_list, err := GetMailList(mail_order_list.ListID)

	if err != nil {
		return nil, err
	}

	mail_order := models.MailOrder{
		IDs:      mail_list.UserIDs,
		Template: mail_order_list.Template,
	}

	return SendMailByID(mail_order)
}

/*
	Send mail the the users with the given ids, using the provided template
	Substitution will be generated based on user info
*/
func SendMailByID(mail_order models.MailOrder) (*models.MailStatus, error) {
	var mail_info models.MailInfo

	mail_info.Content = models.Content{
		TemplateID: mail_order.Template,
	}

	mail_info.Recipients = make([]models.Recipient, len(mail_order.IDs))
	for i, id := range mail_order.IDs {
		user_info, err := GetUserInfo(id)

		if err != nil {
			return nil, err
		}

		mail_info.Recipients[i].Address = models.Address{
			Email: user_info.Email,
			Name:  fmt.Sprintf("%s %s", user_info.FirstName, user_info.LastName),
		}
		mail_info.Recipients[i].Substitutions = models.Substitutions{
			"name": user_info.FirstName,
		}
	}

	return SendMail(mail_info)
}

/*
	Send mail based on the given mailing info
	Returns the results of sending the mail
*/
func SendMail(mail_info models.MailInfo) (*models.MailStatus, error) {
	if !config.IS_PRODUCTION {
		return SendMailDev(mail_info)
	}

	body := bytes.Buffer{}
	json.NewEncoder(&body).Encode(&mail_info)

	req, err := http.NewRequest("POST", config.SPARKPOST_API+"/transmissions/", &body)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", config.SPARKPOST_APIKEY)
	req.Header.Set("Content-Type", "application/json")

	var mail_status models.MailStatus
	status, err := apirequest.Do(req, &mail_status)

	if err != nil {
		return nil, err
	}

	if status != http.StatusOK {
		return nil, errors.New("Failed to send mail")
	}

	return &mail_status, nil
}

/*
	Returns the expected success response in the same format as SparkPost
	This is only to be used in development environments
*/
func SendMailDev(mail_info models.MailInfo) (*models.MailStatus, error) {
	mail_status := models.MailStatus{
		Results: models.MailStatusResults{
			Rejected: 0,
			Accepted: len(mail_info.Recipients),
		},
	}
	return &mail_status, nil
}

/*
	Create a mailing list with the given id and initial set of user, if provided.
	Returns an error if a list with given ID already exists.
*/
func CreateMailList(mail_list models.MailList) error {
	if mail_list.UserIDs == nil {
		mail_list.UserIDs = []string{}
	}

	_, err := GetMailList(mail_list.ID)

	if err == mgo.ErrNotFound {
		return db.Insert("lists", &mail_list)
	} else if err != nil {
		return err
	} else {
		return errors.New("Mail list with given ID already exists.")
	}
}

/*
	Adds the given users to the specified mailing list
*/
func AddToMailList(mail_list models.MailList) error {
	selector := bson.M{
		"id": mail_list.ID,
	}

	modifier := bson.M{
		"$addToSet": bson.M{
			"userids": bson.M{
				"$each": mail_list.UserIDs,
			},
		},
	}

	return db.Update("lists", selector, &modifier)
}

/*
	Removes the given users from the specified mailing list
*/
func RemoveFromMailList(mail_list models.MailList) error {
	selector := bson.M{
		"id": mail_list.ID,
	}

	modifier := bson.M{
		"$pull": bson.M{
			"userids": bson.M{
				"$in": mail_list.UserIDs,
			},
		},
	}

	return db.Update("lists", selector, &modifier)
}

/*
	Gets the mail list with the given id
*/
func GetMailList(id string) (*models.MailList, error) {
	query := bson.M{
		"id": id,
	}

	var mail_list models.MailList
	err := db.FindOne("lists", query, &mail_list)

	if err != nil {
		return nil, err
	}

	return &mail_list, nil
}

/*
	Gets all created mailing lists
*/
func GetAllMailLists() (*models.MailListList, error) {
	var mail_lists []models.MailList

	// nil in this case means that we return everything in the lists collection
	err := db.FindAll("lists", nil, &mail_lists)

	if err != nil {
		return nil, err
	}

	mail_list_list := models.MailListList{
		MailLists: mail_lists,
	}

	return &mail_list_list, nil
}
