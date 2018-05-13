package database

import (
	"github.com/HackIllinois/api-registration/config"
	"gopkg.in/mgo.v2"
)

var global_session *mgo.Session

/*
	Initialize connection to mongo database
*/
func init() {
	session, err := mgo.Dial(config.REGISTRATION_DB_HOST)

	global_session = session

	if err != nil {
		panic(err)
	}
}

/*
	Returns a copy of the global session for use by a connection
*/
func GetSession() *mgo.Session {
	return global_session.Copy()
}

/*
	Find one element matching the given query parameters
*/
func FindOne(collection_name string, query interface{}, result interface{}) error {
	current_session := GetSession()
	defer current_session.Close()

	collection := current_session.DB(config.REGISTRATION_DB_NAME).C(collection_name)

	err := collection.Find(query).One(result)

	return err
}

/*
	Find all elements matching the given query parameters
*/
func FindAll(collection_name string, query interface{}, result interface{}) error {
	current_session := GetSession()
	defer current_session.Close()

	collection := current_session.DB(config.REGISTRATION_DB_NAME).C(collection_name)

	err := collection.Find(query).All(result)

	return err
}

/*
	Insert the given item into the collection
*/
func Insert(collection_name string, item interface{}) error {
	current_session := GetSession()
	defer current_session.Close()

	collection := current_session.DB(config.REGISTRATION_DB_NAME).C(collection_name)

	err := collection.Insert(item)

	return err
}

/*
	Finds an item based on the given selector and updates it with the data in update
*/
func Update(collection_name string, selector interface{}, update interface{}) error {
	current_session := GetSession()
	defer current_session.Close()

	collection := current_session.DB(config.REGISTRATION_DB_NAME).C(collection_name)

	err := collection.Update(selector, update)

	return err
}
