package database

import (
	"crypto/tls"
	"net"
	"time"

	"github.com/HackIllinois/api/common/config"
	"gopkg.in/mgo.v2"
)

/*
	Database interface exposing the methods necessary to querying, inserting, updating, upserting, and removing records
*/
type Database interface {
	Connect(host string) error
	FindOne(collection_name string, query interface{}, result interface{}) error
	FindAll(collection_name string, query interface{}, result interface{}) error
	RemoveOne(collection_name string, query interface{}) error
	RemoveAll(collection_name string, query interface{}) (*mgo.ChangeInfo, error)
	Insert(collection_name string, item interface{}) error
	Upsert(collection_name string, selector interface{}, update interface{}) (*mgo.ChangeInfo, error)
	Update(collection_name string, selector interface{}, update interface{}) error
	UpdateAll(collection_name string, selector interface{}, update interface{}) (*mgo.ChangeInfo, error)
}

/*
	MongoDatabase struct which implements the Database interface for a mongo database
*/
type MongoDatabase struct {
	global_session *mgo.Session
	name           string
}

/*
	Initialize connection to mongo database
*/
func InitMongoDatabase(host string, db_name string) (MongoDatabase, error) {
	dial_info, err := mgo.ParseURL(host)

	if err != nil {
		return MongoDatabase{}, ErrConnection
	}

	if config.IS_PRODUCTION {
		dial_info.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
			tls_config := &tls.Config{}
			connection, err := tls.Dial("tcp", addr.String(), tls_config)
			return connection, err
		}
		dial_info.Timeout = 60 * time.Second
	}

	session, err := mgo.DialWithInfo(dial_info)

	if err != nil {
		return MongoDatabase{}, ErrConnection
	}

	db := MongoDatabase{
		global_session: session,
		name:           db_name,
	}

	return db, nil
}

/*
	Returns a copy of the global session for use by a connection
*/
func (db MongoDatabase) GetSession() *mgo.Session {
	return db.global_session.Copy()
}

/*
	Find one element matching the given query parameters
*/
func (db MongoDatabase) FindOne(collection_name string, query interface{}, result interface{}) error {
	current_session := db.GetSession()
	defer current_session.Close()

	collection := current_session.DB(db.name).C(collection_name)

	err := collection.Find(query).One(result)

	return convertMgoError(err)
}

/*
	Find all elements matching the given query parameters
*/
func (db MongoDatabase) FindAll(collection_name string, query interface{}, result interface{}) error {
	current_session := db.GetSession()
	defer current_session.Close()

	collection := current_session.DB(db.name).C(collection_name)

	err := collection.Find(query).All(result)

	return convertMgoError(err)
}

/*
	Remove one element matching the given query parameters
*/
func (db MongoDatabase) RemoveOne(collection_name string, query interface{}) error {
	current_session := db.GetSession()
	defer current_session.Close()

	collection := current_session.DB(db.name).C(collection_name)

	err := collection.Remove(query)

	return convertMgoError(err)
}

/*
	Remove all elements matching the given query parameters
*/
func (db MongoDatabase) RemoveAll(collection_name string, query interface{}) (*mgo.ChangeInfo, error) {
	current_session := db.GetSession()
	defer current_session.Close()

	collection := current_session.DB(db.name).C(collection_name)

	change_info, err := collection.RemoveAll(query)

	return change_info, convertMgoError(err)
}

/*
	Insert the given item into the collection
*/
func (db MongoDatabase) Insert(collection_name string, item interface{}) error {
	current_session := db.GetSession()
	defer current_session.Close()

	collection := current_session.DB(db.name).C(collection_name)

	err := collection.Insert(item)

	return convertMgoError(err)
}

/*
	Upsert the given item into the collection i.e.,
	if the item exists, it is updated with the given values, else a new item with those values is created.
*/
func (db MongoDatabase) Upsert(collection_name string, selector interface{}, update interface{}) (*mgo.ChangeInfo, error) {
	current_session := db.GetSession()
	defer current_session.Close()

	collection := current_session.DB(db.name).C(collection_name)

	change_info, err := collection.Upsert(selector, update)

	return change_info, convertMgoError(err)
}

/*
	Finds an item based on the given selector and updates it with the data in update
*/
func (db MongoDatabase) Update(collection_name string, selector interface{}, update interface{}) error {
	current_session := db.GetSession()
	defer current_session.Close()

	collection := current_session.DB(db.name).C(collection_name)

	err := collection.Update(selector, update)

	return convertMgoError(err)
}

/*
	Finds all items based on the given selector and updates them with the data in update
*/
func (db MongoDatabase) UpdateAll(collection_name string, selector interface{}, update interface{}) (*mgo.ChangeInfo, error) {
	current_session := db.GetSession()
	defer current_session.Close()

	collection := current_session.DB(db.name).C(collection_name)

	change_info, err := collection.UpdateAll(selector, update)

	return change_info, convertMgoError(err)
}
