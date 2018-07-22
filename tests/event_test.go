package tests

import (
	"github.com/HackIllinois/api-commons/database"
	"github.com/HackIllinois/api-event/config"
	"github.com/HackIllinois/api-event/models"
	"github.com/HackIllinois/api-event/service"
	"reflect"
	"testing"
)

var db database.MongoDatabase

func init() {
	db_connection, err := database.InitMongoDatabase(config.EVENT_DB_HOST, config.EVENT_DB_NAME)

	if err != nil {
		panic(err)
	}

	db = db_connection
}

/*
	Initialize db with a test event
*/
func SetupTestDB(t *testing.T) {
	event := models.Event{
		Name: "testname",
		Description: "testdescription",
		StartTime: 1000,
		EndTime: 2000,
		LocationDescription: "testlocationdescription",
		Latitude: 123.456,
		Longitude: 123.456,
		Sponsor: "testsponsor",
		EventType: "WORKSHOP",
	}

	err := db.Insert("events", &event)

	if err != nil {
		t.Fatal(err)
	}
}

/*
	Drop test db
*/
func CleanupTestDB(t *testing.T) {
	session := db.GetSession()
	defer session.Close()

	err := session.DB(config.EVENT_DB_NAME).DropDatabase()

	if err != nil {
		t.Fatal(err)
	}
}

/*
	Service level test for getting event from db
*/
func TestGetEventService(t *testing.T) {
	SetupTestDB(t)

	event, err := service.GetEvent("testname")

	if err != nil {
		t.Fatal(err)
	}

	expected_event := models.Event{
		Name: "testname",
		Description: "testdescription",
		StartTime: 1000,
		EndTime: 2000,
		LocationDescription: "testlocationdescription",
		Latitude: 123.456,
		Longitude: 123.456,
		Sponsor: "testsponsor",
		EventType: "WORKSHOP",
	}

	if !reflect.DeepEqual(event, &expected_event) {
		t.Errorf("Wrong event info. Expected %v, got %v", expected_event, event)
	}

	CleanupTestDB(t)
}

/*
	Service level test for creating an event in the db
*/
func TestCreateEventService(t *testing.T) {
	SetupTestDB(t)

	new_event := models.Event{
		Name: "testname2",
		Description: "testdescription2",
		StartTime: 1000,
		EndTime: 2000,
		LocationDescription: "testlocationdescription",
		Latitude: 123.456,
		Longitude: 123.456,
		Sponsor: "testsponsor",
		EventType: "WORKSHOP",
	}

	err := service.CreateEvent("testname2", new_event)

	if err != nil {
		t.Fatal(err)
	}

	event, err := service.GetEvent("testname2")

	if err != nil {
		t.Fatal(err)
	}

	expected_event := models.Event{
		Name: "testname2",
		Description: "testdescription2",
		StartTime: 1000,
		EndTime: 2000,
		LocationDescription: "testlocationdescription",
		Latitude: 123.456,
		Longitude: 123.456,
		Sponsor: "testsponsor",
		EventType: "WORKSHOP",
	}

	if !reflect.DeepEqual(event, &expected_event) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_event, event)
	}

	CleanupTestDB(t)
}

/*
	Service level test for updating an event in the db
*/
func TestUpdateEventService(t *testing.T) {
	SetupTestDB(t)

	event := models.Event{
		Name: "testname",
		Description: "testdescription2",
		StartTime: 1000,
		EndTime: 2000,
		LocationDescription: "testlocationdescription",
		Latitude: 123.456,
		Longitude: 123.456,
		Sponsor: "testsponsor",
		EventType: "WORKSHOP",
	}

	err := service.UpdateEvent("testname", event)

	if err != nil {
		t.Fatal(err)
	}

	updated_event, err := service.GetEvent("testname")

	if err != nil {
		t.Fatal(err)
	}

	expected_event := models.Event{
		Name: "testname",
		Description: "testdescription2",
		StartTime: 1000,
		EndTime: 2000,
		LocationDescription: "testlocationdescription",
		Latitude: 123.456,
		Longitude: 123.456,
		Sponsor: "testsponsor",
		EventType: "WORKSHOP",
	}

	if !reflect.DeepEqual(updated_event, &expected_event) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_event, updated_event)
	}

	CleanupTestDB(t)
}
