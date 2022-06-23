package database

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/HackIllinois/api/common/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
	MongoDatabase struct which implements the Database interface for a mongo database
*/
type MongoDatabase struct {
	client *mongo.Client
	name   string
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
	client_options := options.Client().ApplyURI("mongodb://" + host + ":27017")

	{
		err := client_options.Validate()

		if err != nil {
			// problems parsing host
			return ErrConnection
		}
	}

	if config.IS_PRODUCTION {
		tls_config := &tls.Config{}
		client_options.SetTLSConfig(tls_config)
		client_options.SetSocketTimeout(60 * time.Second)
	}

	client_options.SetMaxPoolSize(25) // default is 100, but this was set to 25 by us on the old driver

	client, err := mongo.Connect(context.TODO(), client_options)

	if err != nil {
		// failed to connect to database
		return ErrConnection
	}

	db.client = client

	return nil
}

/*
	Close the global session to the given mongo database
*/
func (db *MongoDatabase) Close() {
	db.client.Disconnect(context.TODO())
}

/*
	Returns a copy of the global session for use by a connection
*/
func (db *MongoDatabase) GetSession() (*mongo.Session, error) {
	session, err := db.client.StartSession(nil)

	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (db *MongoDatabase) StartSession() (*mongo.Session, error) {
	return db.GetSession()
}

/*
	Find one element matching the given query parameters
*/
func (db *MongoDatabase) FindOne(collection_name string, query interface{}, result interface{}, session *mongo.SessionContext) error {
	var s *mongo.Session
	if session == nil {
		var err error
		s, err = db.GetSession()

		if err != nil {
			return convertMgoError(err)
		}

		defer (*s).EndSession(context.TODO())
		sess_ctx := mongo.NewSessionContext(context.TODO(), *s)
		session = &sess_ctx
	}

	query = nilToEmptyBson(query)

	res := db.client.Database(db.name).Collection(collection_name).FindOne(*session, query)

	err := res.Decode(result)

	return convertMgoError(err)
}

/*
	Find all elements matching the given query parameters
*/
func (db *MongoDatabase) FindAll(collection_name string, query interface{}, result interface{}, session *mongo.SessionContext) error {
	var s *mongo.Session
	if session == nil {
		var err error
		s, err = db.GetSession()

		if err != nil {
			return convertMgoError(err)
		}

		defer (*s).EndSession(context.TODO())
		sess_ctx := mongo.NewSessionContext(context.TODO(), *s)
		session = &sess_ctx
	}

	query = nilToEmptyBson(query)

	cursor, err := db.client.Database(db.name).Collection(collection_name).Find(*session, query)

	if err != nil {
		return convertMgoError(err)
	}

	if err = cursor.All(context.TODO(), result); err != nil {
		return convertMgoError(err)
	}

	return nil
}

/*
	Find all elements matching the given query parameters, and sorts them based on given sort fields
        The first sort field is highest priority, each subsequent field breaks ties
*/
func (db *MongoDatabase) FindAllSorted(collection_name string, query interface{}, sort_fields bson.D, result interface{}, session *mongo.SessionContext) error {
	var s *mongo.Session
	if session == nil {
		var err error
		s, err = db.GetSession()

		if err != nil {
			return convertMgoError(err)
		}

		defer (*s).EndSession(context.TODO())
		sess_ctx := mongo.NewSessionContext(context.TODO(), *s)
		session = &sess_ctx
	}

	query = nilToEmptyBson(query)

	options := options.Find().SetSort(sort_fields)

	cursor, err := db.client.Database(db.name).Collection(collection_name).Find(*session, query, options)

	if err != nil {
		return convertMgoError(err)
	}

	if err = cursor.All(context.TODO(), result); err != nil {
		return convertMgoError(err)
	}

	return nil
}

/*
	Remove one element matching the given query parameters
*/
func (db *MongoDatabase) RemoveOne(collection_name string, query interface{}, session *mongo.SessionContext) error {
	var s *mongo.Session
	if session == nil {
		var err error
		s, err = db.GetSession()

		if err != nil {
			return convertMgoError(err)
		}

		defer (*s).EndSession(context.TODO())
		sess_ctx := mongo.NewSessionContext(context.TODO(), *s)
		session = &sess_ctx
	}

	query = nilToEmptyBson(query)

	_, err := db.client.Database(db.name).Collection(collection_name).DeleteOne(*session, query)

	return convertMgoError(err)
}

/*
	Remove all elements matching the given query parameters
*/
func (db *MongoDatabase) RemoveAll(collection_name string, query interface{}, session *mongo.SessionContext) (*ChangeResults, error) {
	var s *mongo.Session
	if session == nil {
		var err error
		s, err = db.GetSession()

		if err != nil {
			return nil, convertMgoError(err)
		}

		defer (*s).EndSession(context.TODO())
		sess_ctx := mongo.NewSessionContext(context.TODO(), *s)
		session = &sess_ctx
	}

	query = nilToEmptyBson(query)

	res, err := db.client.Database(db.name).Collection(collection_name).DeleteMany(*session, query)

	if err != nil {
		return nil, convertMgoError(err)
	}

	change_results := ChangeResults{
		Updated: 0,
		Deleted: int(res.DeletedCount),
	}

	return &change_results, nil
}

/*
	Insert the given item into the collection
*/
func (db *MongoDatabase) Insert(collection_name string, item interface{}, session *mongo.SessionContext) error {
	var s *mongo.Session
	if session == nil {
		var err error
		s, err = db.GetSession()

		if err != nil {
			return convertMgoError(err)
		}

		defer (*s).EndSession(context.TODO())
		sess_ctx := mongo.NewSessionContext(context.TODO(), *s)
		session = &sess_ctx
	}

	item = nilToEmptyBson(item)

	_, err := db.client.Database(db.name).Collection(collection_name).InsertOne(*session, item)

	return convertMgoError(err)
}

/*
	Upsert the given item into the collection i.e.,
	if the item exists, it is updated with the given values, else a new item with those values is created.
*/
func (db *MongoDatabase) Upsert(collection_name string, selector interface{}, update interface{}, session *mongo.SessionContext) (*ChangeResults, error) {
	var s *mongo.Session
	if session == nil {
		var err error
		s, err = db.GetSession()

		if err != nil {
			return nil, convertMgoError(err)
		}

		defer (*s).EndSession(context.TODO())
		sess_ctx := mongo.NewSessionContext(context.TODO(), *s)
		session = &sess_ctx
	}

	selector = nilToEmptyBson(selector)
	update = addReplaceWrapper(update)

	options := options.Update().SetUpsert(true)

	res, err := db.client.Database(db.name).Collection(collection_name).UpdateOne(*session, selector, update, options)

	if err != nil {
		return nil, convertMgoError(err)
	}

	change_results := ChangeResults{
		Updated: int(res.UpsertedCount),
		Deleted: 0,
	}

	return &change_results, nil
}

/*
	Finds an item based on the given selector and updates it with the data in update
*/
func (db *MongoDatabase) Update(collection_name string, selector interface{}, update interface{}, session *mongo.SessionContext) error {
	var s *mongo.Session
	if session == nil {
		var err error
		s, err = db.GetSession()

		if err != nil {
			return convertMgoError(err)
		}

		defer (*s).EndSession(context.TODO())
		sess_ctx := mongo.NewSessionContext(context.TODO(), *s)
		session = &sess_ctx
	}

	selector = nilToEmptyBson(selector)
	update = addReplaceWrapper(update)

	_, err := db.client.Database(db.name).Collection(collection_name).UpdateOne(*session, selector, update)

	return convertMgoError(err)
}

/*
	Finds all items based on the given selector and updates them with the data in update
*/
func (db *MongoDatabase) UpdateAll(collection_name string, selector interface{}, update interface{}, session *mongo.SessionContext) (*ChangeResults, error) {
	var s *mongo.Session
	if session == nil {
		var err error
		s, err = db.GetSession()

		if err != nil {
			return nil, convertMgoError(err)
		}

		defer (*s).EndSession(context.TODO())
		sess_ctx := mongo.NewSessionContext(context.TODO(), *s)
		session = &sess_ctx
	}

	selector = nilToEmptyBson(selector)
	update = addReplaceWrapper(update)

	res, err := db.client.Database(db.name).Collection(collection_name).UpdateMany(*session, selector, update)

	if err != nil {
		return nil, convertMgoError(err)
	}

	change_results := ChangeResults{
		Updated: int(res.ModifiedCount),
		Deleted: 0,
	}

	return &change_results, nil
}

/*
	Drops the entire database
*/
func (db *MongoDatabase) DropDatabase(session *mongo.SessionContext) error {
	var s *mongo.Session
	if session == nil {
		var err error
		s, err = db.GetSession()

		if err != nil {
			return convertMgoError(err)
		}

		defer (*s).EndSession(context.TODO())
		sess_ctx := mongo.NewSessionContext(context.TODO(), *s)
		session = &sess_ctx
	}

	err := db.client.Database(db.name).Drop(*session)

	return convertMgoError(err)
}

/*
	Returns a map of statistics for a given collection
*/
func (db *MongoDatabase) GetStats(collection_name string, fields []string, session *mongo.SessionContext) (map[string]interface{}, error) {
	var s *mongo.Session
	if session == nil {
		var err error
		s, err = db.GetSession()

		if err != nil {
			return nil, convertMgoError(err)
		}

		defer (*s).EndSession(context.TODO())
		sess_ctx := mongo.NewSessionContext(context.TODO(), *s)
		session = &sess_ctx
	}

	cursor, err := db.client.Database(db.name).Collection(collection_name).Find(*session, bson.D{})

	if err != nil {
		return nil, convertMgoError(err)
	}

	stats := GetDefaultStats()
	count := 0

	for cursor.Next(*session) {
		var result map[string]interface{}

		if err := cursor.Decode(&result); err != nil {
			count += 1
			err := AddEntryToStats(stats, result, fields)

			if err != nil {
				return nil, convertMgoError(err)
			}
		}
	}

	err = cursor.Close(*session)

	if err != nil {
		return nil, convertMgoError(err)
	}

	stats["count"] = count

	return stats, nil
}

/*
	Adds the $replaceWith update operator to make passing in direct structs safe to use
	($replaceWith is effectlively the behavior we used whenever updating an item)
*/
func addReplaceWrapper(update interface{}) interface{} {
	// TODO: Right now, we are using bson.M as the main type to use $ operators, but
	// it's probably better to just make this a separate type
	switch update.(type) {
	case nil:
		return bson.D{}
	case *primitive.M:
		return update
	default:
		return bson.A{bson.D{{"$replaceWith", update}}}
	}
}

/*
	Prevents passing nil to any CRUD function by swapping for empty BSON
*/
func nilToEmptyBson(input interface{}) interface{} {
	if input == nil {
		return bson.D{}
	}
	return input
}
