package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
	Database interface exposing the methods necessary to querying, inserting, updating, upserting, and removing records
*/
type Database interface {
	Connect(host string) error
	Close()
	GetRaw() *mongo.Client
	StartSession() (*mongo.Session, error)
	GetNewContext() (context.Context, context.CancelFunc)
	FindOne(collection_name string, query interface{}, result interface{}, session *mongo.SessionContext) error
	FindOneAndDelete(collection_name string, query interface{}, result interface{}, session *mongo.SessionContext) error
	FindOneAndUpdate(collection_name string, query interface{}, update interface{}, result interface{}, return_new_doc bool, upsert bool, session *mongo.SessionContext) error
	FindOneAndReplace(collection_name string, query interface{}, update interface{}, result interface{}, return_new_doc bool, upsert bool, session *mongo.SessionContext) error
	FindAll(collection_name string, query interface{}, result interface{}, session *mongo.SessionContext) error
	FindAllSorted(collection_name string, query interface{}, sort_fields bson.D, result interface{}, session *mongo.SessionContext) error
	RemoveOne(collection_name string, query interface{}, session *mongo.SessionContext) error
	RemoveAll(collection_name string, query interface{}, session *mongo.SessionContext) (*ChangeResults, error)
	Insert(collection_name string, item interface{}, session *mongo.SessionContext) error
	Upsert(collection_name string, selector interface{}, update interface{}, session *mongo.SessionContext) (*ChangeResults, error)
	Update(collection_name string, selector interface{}, update interface{}, session *mongo.SessionContext) error
	UpdateAll(collection_name string, selector interface{}, update interface{}, session *mongo.SessionContext) (*ChangeResults, error)
	Replace(collection_name string, selector interface{}, update interface{}, upsert bool, session *mongo.SessionContext) error
	DropDatabase(session *mongo.SessionContext) error
	GetStats(collection_name string, fields []string, session *mongo.SessionContext) (map[string]interface{}, error)
}

/*
	An alias of a string -> interface{} map used for database queries and selectors
*/
type QuerySelector map[string]interface{}

/*
   Represents a single field to sort by, including the name and whether the sort should be reversed
*/
type SortField struct {
	Name     string
	Reversed bool
}

/*
	Used to store information about the changes made by a database operation
*/
type ChangeResults struct {
	Updated int
	Deleted int
}

/*
	Initialize a connection to the given database

	This function wraps a database specific initializion function
	This makes it simple to change the database used without rewriting
	code in the microservices
*/
func InitDatabase(host string, db_name string) (Database, error) {
	db, err := InitMongoDatabase(host, db_name)
	return db, err
}
