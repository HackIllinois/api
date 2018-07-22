package service

import (
	"github.com/HackIllinois/api-commons/database"
	"github.com/HackIllinois/api-stat/config"
	"github.com/HackIllinois/api-stat/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var db database.MongoDatabase

func init() {
	db_connection, err := database.InitMongoDatabase(config.STAT_DB_HOST, config.STAT_DB_NAME)

	if err != nil {
		panic(err)
	}

	db = db_connection
}

/*
	Returns the service with the given name
*/
func GetService(name string) (*models.Service, error) {
	query := bson.M{
		"name": name,
	}

	var service models.Service
	err := db.FindOne("services", query, &service)

	if err != nil {
		return nil, err
	}

	return &service, nil
}

/*
	Registers the service with the given name
	The service will be created if it doesn't exist
*/
func RegisterService(name string, service models.Service) error {
	_, err := GetService(name)

	if err == mgo.ErrNotFound {
		err = db.Insert("services", &service)

		return err
	}

	if err != nil {
		return err
	}

	selector := bson.M{
		"name": name,
	}

	err = db.Update("services", selector, &service)

	return err
}
