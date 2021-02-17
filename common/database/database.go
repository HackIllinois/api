package database

/*
	Database interface exposing the methods necessary to querying, inserting, updating, upserting, and removing records
*/
type Database interface {
	Connect(host string) error
	Close()
	FindOne(collection_name string, query interface{}, result interface{}) error
	FindAll(collection_name string, query interface{}, result interface{}) error
	FindAllSorted(collection_name string, query interface{}, sort_fields []SortField, result interface{}) error
	RemoveOne(collection_name string, query interface{}) error
	RemoveAll(collection_name string, query interface{}) (*ChangeResults, error)
	Insert(collection_name string, item interface{}) error
	Upsert(collection_name string, selector interface{}, update interface{}) (*ChangeResults, error)
	Update(collection_name string, selector interface{}, update interface{}) error
	UpdateAll(collection_name string, selector interface{}, update interface{}) (*ChangeResults, error)
	DropDatabase() error
	GetStats(collection_name string, fields []string) (map[string]interface{}, error)
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
