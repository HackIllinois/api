package service

import (
	"errors"

	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/services/decision/config"
	"github.com/HackIllinois/api/services/decision/models"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/mgo.v2/bson"

	"strconv"
	"strings"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

var db database.MongoDatabase

func init() {
	db_connection, err := database.InitMongoDatabase(config.DECISION_DB_HOST, config.DECISION_DB_NAME)

	if err != nil {
		panic(err)
	}

	db = db_connection
}

/*
	Returns the decision associated with the given user id
*/
func GetDecision(id string) (*models.DecisionHistory, error) {
	query := bson.M{"id": id}

	var decision models.DecisionHistory
	err := db.FindOne("decision", query, &decision)

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
	} else if decision.Status != "ACCEPTED" && decision.Wave != 0 {
		return errors.New("Cannot set a wave for non-accepted attendee")
	}

	decision_history, err := GetDecision(id)

	if err != nil {
		if err == database.ErrNotFound {
			decision_history = &models.DecisionHistory{
				ID: id,
			}
		} else {
			return err
		}
	}

	decision_history.Finalized = decision.Finalized
	decision_history.Status = decision.Status
	decision_history.Wave = decision.Wave
	decision_history.History = append(decision_history.History, decision)
	decision_history.Reviewer = decision.Reviewer
	decision_history.Timestamp = decision.Timestamp

	selector := bson.M{"id": id}

	err = db.Update("decision", selector, &decision_history)

	if err == database.ErrNotFound {
		err = db.Insert("decision", &decision_history)
	}

	return err
}

/*
	Checks if a decision with the provided id exists.
*/
func HasDecision(id string) (bool, error) {
	_, err := GetDecision(id)

	if err == nil {
		return true, nil
	} else if err == database.ErrNotFound {
		return false, nil
	} else {
		return false, err
	}
}

func Contains(slice []string, str string) bool {
	for _, value := range slice {
		if value == str {
			return true
		}
	}
	return false
}

func AssignValueType(key, value string) (interface{}, error) {
	int_keys := []string{"wave", "timestamp"}
	if Contains(int_keys, key) {
		return strconv.Atoi(value)
	}
	return value, nil
}

/*
	Returns decisions based on a filter
*/
func GetFilteredDecisions(parameters map[string][]string) (*models.FilteredDecisions, error) {
	query := make(map[string]interface{})
	for key, values := range parameters {
		if len(values) > 1 {
			return nil, errors.New("Multiple usage of key " + key)
		}

		key = strings.ToLower(key)
		value_list := strings.Split(values[0], ",")

		correctly_typed_value_list := make([]interface{}, len(value_list))
		for i, value := range value_list {
			correctly_typed_value, err := AssignValueType(key, value)
			if err == nil {
				correctly_typed_value_list[i] = correctly_typed_value
			} else {
				return nil, err
			}
		}
		query[key] = bson.M{"$in": correctly_typed_value_list}
	}

	var filtered_decisions models.FilteredDecisions
	err := db.FindAll("decision", query, &filtered_decisions.Decisions)
	if err != nil {
		return nil, err
	}

	return &filtered_decisions, nil
}
