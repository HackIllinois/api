package service

import (
	"errors"

	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/services/decision/config"
	"github.com/HackIllinois/api/services/decision/models"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
)

var validate *validator.Validate

var db database.Database

func Initialize() error {
	if db != nil {
		db.Close()
		db = nil
	}

	var err error
	db, err = database.InitDatabase(config.DECISION_DB_HOST, config.DECISION_DB_NAME)

	if err != nil {
		return err
	}

	validate = validator.New()

	return nil
}

/*
	Returns the decision associated with the given user id
*/
func GetDecision(id string) (*models.DecisionHistory, error) {
	query := database.QuerySelector{"id": id}

	var decision models.DecisionHistory
	err := db.FindOne("decision", query, &decision, nil)

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

	selector := database.QuerySelector{"id": id}

	// TODO: Change to some dedi struct for mongo operations
	decision_history := bson.M{
		"$set": bson.M{
			"finalized": decision.Finalized,
			"status":    decision.Status,
			"wave":      decision.Wave,
			"reviewer":  decision.Reviewer,
			"timestamp": decision.Timestamp,
			"expiresAt": decision.ExpiresAt,
		},
		"$addToSet": bson.M{
			"history": decision,
		},
	}

	_, err = db.Upsert("decision", selector, &decision_history, nil)

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

/*
	Returns decisions based on a filter
*/
func GetFilteredDecisions(parameters map[string][]string) (*models.FilteredDecisions, error) {
	query, err := database.CreateFilterQuery(parameters, models.DecisionHistory{})

	if err != nil {
		return nil, err
	}

	var filtered_decisions models.FilteredDecisions
	err = db.FindAll("decision", query, &filtered_decisions.Decisions, nil)
	if err != nil {
		return nil, err
	}

	return &filtered_decisions, nil
}

/*
	Returns all decision stats
*/
func GetStats() (map[string]interface{}, error) {
	return db.GetStats("decision", []string{"status", "finalized", "wave"}, nil)
}
