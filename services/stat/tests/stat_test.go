package tests

import (
	"github.com/pattyjogal/api/common/database"
	"github.com/pattyjogal/api/services/stat/config"
	"github.com/pattyjogal/api/services/stat/models"
	"github.com/pattyjogal/api/services/stat/service"
	"reflect"
	"testing"
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
	Initialize db with a test service
*/
func SetupTestDB(t *testing.T) {
	api_service := models.Service{
		Name: "testname",
		URL:  "http://localhost:8050",
	}

	err := db.Insert("services", &api_service)

	if err != nil {
		t.Fatal(err)
	}

	additional_api_service := models.Service{
		Name: "additionaltestname",
		URL:  "http://localhost:8059",
	}

	err = db.Insert("services", &additional_api_service)

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

	err := session.DB(config.STAT_DB_NAME).DropDatabase()

	if err != nil {
		t.Fatal(err)
	}
}

/*
	Service level test for getting service from db
*/
func TestGetServiceService(t *testing.T) {
	SetupTestDB(t)

	api_service, err := service.GetService("testname")

	if err != nil {
		t.Fatal(err)
	}

	expected_api_service := models.Service{
		Name: "testname",
		URL:  "http://localhost:8050",
	}

	if !reflect.DeepEqual(api_service, &expected_api_service) {
		t.Errorf("Wrong stat info. Expected %v, got %v", expected_api_service, api_service)
	}

	CleanupTestDB(t)
}

/*
	Service level test for registering a service in the db
*/
func TestRegisterServiceService(t *testing.T) {
	SetupTestDB(t)

	new_api_service := models.Service{
		Name: "testname2",
		URL:  "http://localhost:8051",
	}

	err := service.RegisterService("testname2", new_api_service)

	if err != nil {
		t.Fatal(err)
	}

	api_service, err := service.GetService("testname2")

	if err != nil {
		t.Fatal(err)
	}

	expected_api_service := models.Service{
		Name: "testname2",
		URL:  "http://localhost:8051",
	}

	if !reflect.DeepEqual(api_service, &expected_api_service) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_api_service, api_service)
	}

	CleanupTestDB(t)
}

/*
	Service level test for registering an existing service in the db
*/
func TestReregisterServiceService(t *testing.T) {
	SetupTestDB(t)

	api_service := models.Service{
		Name: "testname",
		URL:  "http://localhost:8052",
	}

	err := service.RegisterService("testname", api_service)

	if err != nil {
		t.Fatal(err)
	}

	updated_api_service, err := service.GetService("testname")

	if err != nil {
		t.Fatal(err)
	}

	expected_api_service := models.Service{
		Name: "testname",
		URL:  "http://localhost:8052",
	}

	if !reflect.DeepEqual(updated_api_service, &expected_api_service) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_api_service, updated_api_service)
	}

	CleanupTestDB(t)
}

/*
	Service level test for retreiving all registered services
*/
func TestGetAllServiceService(t *testing.T) {
	SetupTestDB(t)

	all_services, err := service.GetAllServices()

	if err != nil {
		t.Fatal(err)
	}

	expected_services := make([]models.Service, 2)

	expected_services[0] = models.Service{
		Name: "testname",
		URL:  "http://localhost:8050",
	}

	expected_services[1] = models.Service{
		Name: "additionaltestname",
		URL:  "http://localhost:8059",
	}

	ordered_equals := reflect.DeepEqual(all_services, expected_services)

	expected_services[0], expected_services[1] = expected_services[1], expected_services[0]
	reversed_equals := reflect.DeepEqual(all_services, expected_services)

	if !(ordered_equals || reversed_equals) {
		t.Errorf("Wrong stat info. Expected %v, got %v", expected_services, all_services)
	}

	CleanupTestDB(t)
}
