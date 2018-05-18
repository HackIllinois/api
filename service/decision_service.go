package service

import (
	"errors"
	"github.com/HackIllinois/api-decision/database"
	"github.com/HackIllinois/api-decision/models"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

/*
	Returns the decision associated with the given user id
*/
func GetDecision(id string) (*models.Decision, error) {
	query := bson.M{
		"id": id,
	}

	var decision models.Decision
	err := database.FindOne("decision", query, &decision)

	if err != nil {
		return nil, err
	}

	return &decision, nil
}

/*
	Updates the decision associated with the given user id
	If a decision doesn't exist it will be created
*/
func UpdateDecision(id string, decision models.Decision) error {
	err := validate.Struct(decision)

	if err != nil {
		return err
	}

	if decision.Status == "ACCEPTED" && decision.Wave == 0 {
		return errors.New("Must set a wave for accepted attendee")
	}

	selector := bson.M{
		"id": id,
	}

	err = database.Update("decision", selector, &decision)

	if err == mgo.ErrNotFound {
		err = database.Insert("decision", &decision)
	}

	return err
}
