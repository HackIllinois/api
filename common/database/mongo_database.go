package database

import (
	"crypto/tls"
	"net"
	"time"

	"github.com/HackIllinois/api/common/config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

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
func InitMongoDatabase(host string, db_name string) (*MongoDatabase, error) {
	db := MongoDatabase{}
	err := db.Connect(host)

	if err != nil {
		return &db, err
	}

	db.name = db_name

	return &db, nil
}

/*
	Open a session to the given mongo database
*/
func (db *MongoDatabase) Connect(host string) error {
	dial_info, err := mgo.ParseURL(host)

	if err != nil {
		return ErrConnection
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
		return ErrConnection
	}

	db.global_session = session

	return nil
}

/*
	Close the global session to the given mongo database
*/
func (db *MongoDatabase) Close() {
	db.global_session.Close()
}

/*
	Returns a copy of the global session for use by a connection
*/
func (db *MongoDatabase) GetSession() *mgo.Session {
	return db.global_session.Copy()
}

/*
	Find one element matching the given query parameters
*/
func (db *MongoDatabase) FindOne(collection_name string, query interface{}, result interface{}) error {
	current_session := db.GetSession()
	defer current_session.Close()

	collection := current_session.DB(db.name).C(collection_name)

	err := collection.Find(query).One(result)

	return convertMgoError(err)
}

/*
	Find all elements matching the given query parameters
*/
func (db *MongoDatabase) FindAll(collection_name string, query interface{}, result interface{}) error {
	current_session := db.GetSession()
	defer current_session.Close()

	collection := current_session.DB(db.name).C(collection_name)

	err := collection.Find(query).All(result)

	return convertMgoError(err)
}

/*
	Remove one element matching the given query parameters
*/
func (db *MongoDatabase) RemoveOne(collection_name string, query interface{}) error {
	current_session := db.GetSession()
	defer current_session.Close()

	collection := current_session.DB(db.name).C(collection_name)

	err := collection.Remove(query)

	return convertMgoError(err)
}

/*
	Remove all elements matching the given query parameters
*/
func (db *MongoDatabase) RemoveAll(collection_name string, query interface{}) (*ChangeResults, error) {
	current_session := db.GetSession()
	defer current_session.Close()

	collection := current_session.DB(db.name).C(collection_name)

	change_info, err := collection.RemoveAll(query)

	change_results := ChangeResults{
		Updated: change_info.Updated,
		Deleted: change_info.Removed,
	}

	return &change_results, convertMgoError(err)
}

/*
	Insert the given item into the collection
*/
func (db *MongoDatabase) Insert(collection_name string, item interface{}) error {
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
func (db *MongoDatabase) Upsert(collection_name string, selector interface{}, update interface{}) (*ChangeResults, error) {
	current_session := db.GetSession()
	defer current_session.Close()

	collection := current_session.DB(db.name).C(collection_name)

	change_info, err := collection.Upsert(selector, update)

	change_results := ChangeResults{
		Updated: change_info.Updated,
		Deleted: change_info.Removed,
	}

	return &change_results, convertMgoError(err)
}

// /*
// 	Finds an item based on the given selector and updates it with the data in update
// */
// func (db *MongoDatabase) Update(collection_name string, selector interface{}, update interface{}) error {
// 	current_session := db.GetSession()
// 	defer current_session.Close()

// 	collection := current_session.DB(db.name).C(collection_name)

// 	// Check the variable type of update .
// 	// If it's a map, it's passed by the patch method, so we only update certain fields.
// 	// If it's not a map, it's passed by other methods, and we update all fields.
// 	update_as_map, ok := update.(*map[string]interface{})
// 	if ok {
// 		// Use the mongodb set operation to update only given fields.
// 		err := collection.Update(selector, bson.M{"$set": update_as_map})
// 		return convertMgoError(err)
// 	} else {
// 		err := collection.Update(selector, update)
// 		return convertMgoError(err)
// 	}
// }

/*
	Finds an item based on the given selector and updates it with the data in update
*/
func (db *MongoDatabase) Update(collection_name string, selector interface{}, update interface{}) error {
	current_session := db.GetSession()
	defer current_session.Close()

	collection := current_session.DB(db.name).C(collection_name)

	err := collection.Update(selector, update)

	return convertMgoError(err)
}

/*
	Finds an item based on the given selector and patches it with the data in update
*/
func (db *MongoDatabase) Patch(collection_name string, selector interface{}, update *map[string]interface{}) error {
	current_session := db.GetSession()
	defer current_session.Close()

	collection := current_session.DB(db.name).C(collection_name)

	// Use the mongodb set operation to update only the given fields.
	err := collection.Update(selector, bson.M{"$set": update})
	return convertMgoError(err)
}

/*
	Finds all items based on the given selector and updates them with the data in update
*/
func (db *MongoDatabase) UpdateAll(collection_name string, selector interface{}, update interface{}) (*ChangeResults, error) {
	current_session := db.GetSession()
	defer current_session.Close()

	collection := current_session.DB(db.name).C(collection_name)

	change_info, err := collection.UpdateAll(selector, update)

	change_results := ChangeResults{
		Updated: change_info.Updated,
		Deleted: change_info.Removed,
	}

	return &change_results, convertMgoError(err)
}

/*
	Drops the entire database
*/
func (db *MongoDatabase) DropDatabase() error {
	current_session := db.GetSession()
	defer current_session.Close()

	err := current_session.DB(db.name).DropDatabase()

	return convertMgoError(err)
}

/*
	Returns a map of statistics for a given collection
*/
func (db *MongoDatabase) GetStats(collection_name string, fields []string) (map[string]interface{}, error) {
	current_session := db.GetSession()
	defer current_session.Close()

	collection := current_session.DB(db.name).C(collection_name)

	iter := collection.Find(nil).Iter()

	stats := GetDefaultStats()

	count := 0

	var result map[string]interface{}
	for iter.Next(&result) {
		count += 1

		err := AddEntryToStats(stats, result, fields)

		if err != nil {
			return nil, convertMgoError(err)
		}
	}

	err := iter.Err()

	if err != nil {
		return nil, convertMgoError(err)
	}

	err = iter.Close()

	if err != nil {
		return nil, convertMgoError(err)
	}

	stats["count"] = count

	return stats, nil
}
