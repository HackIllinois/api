package main

import (
	"./config"
	"gopkg.in/mgo.v2"
)

var global_session *mgo.Session

func init() {
	session, err := mgo.Dial(config.AUTH_DB_HOST)

	global_session = session

	if err != nil {
		panic(err)
	}
}

func GetSession() *mgo.Session {
	return global_session.Copy()
}

func FindOne(collection_name string, query interface{}, result interface{}) error {
	current_session := GetSession()
	defer current_session.Close()

	collection := current_session.DB(config.AUTH_DB_NAME).C(collection_name)

	err := collection.Find(query).One(result)

	return err
}

func FindAll(collection_name string, query interface{}, result interface{}) error {
	current_session := GetSession()
	defer current_session.Close()

	collection := current_session.DB(config.AUTH_DB_NAME).C(collection_name)

	err := collection.Find(query).All(result)

	return err
}

func Insert(collection_name string, item interface{}) error {
	current_session := GetSession()
	defer current_session.Close()

	collection := current_session.DB(config.AUTH_DB_NAME).C(collection_name)

	err := collection.Insert(item)

	return err
}
