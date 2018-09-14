package tests

import (
	"github.com/pattyjogal/api/common/database"
	"github.com/pattyjogal/api/services/checkin/config"
	"github.com/pattyjogal/api/services/checkin/models"
	"github.com/pattyjogal/api/services/checkin/service"
	"net/url"
	"reflect"
	"testing"
)

var db database.MongoDatabase

func init() {
	db_connection, err := database.InitMongoDatabase(config.CHECKIN_DB_HOST, config.CHECKIN_DB_NAME)

	if err != nil {
		panic(err)
	}

	db = db_connection
}

/*
	Initialize db with test checkin info
*/
func SetupTestDB(t *testing.T) {
	checkin := models.UserCheckin{
		ID:              "testid",
		HasCheckedIn:    true,
		HasPickedUpSwag: true,
	}

	err := db.Insert("checkins", &checkin)

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

	err := session.DB(config.CHECKIN_DB_NAME).DropDatabase()

	if err != nil {
		t.Fatal(err)
	}
}

/*
	Service level test for getting user checkin from db
*/
func TestGetUserCheckinService(t *testing.T) {
	SetupTestDB(t)

	checkin, err := service.GetUserCheckin("testid")

	if err != nil {
		t.Fatal(err)
	}

	expected_checkin := models.UserCheckin{
		ID:              "testid",
		HasCheckedIn:    true,
		HasPickedUpSwag: true,
	}

	if !reflect.DeepEqual(checkin, &expected_checkin) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_checkin, checkin)
	}

	CleanupTestDB(t)
}

/*
	Service level test for creating user checkin in the db
*/
func TestCreateUserCheckinService(t *testing.T) {
	SetupTestDB(t)

	new_checkin := models.UserCheckin{
		ID:              "testid2",
		HasCheckedIn:    true,
		HasPickedUpSwag: false,
	}

	err := service.CreateUserCheckin("testid2", new_checkin)

	if err != nil {
		t.Fatal(err)
	}

	checkin, err := service.GetUserCheckin("testid2")

	if err != nil {
		t.Fatal(err)
	}

	expected_checkin := models.UserCheckin{
		ID:              "testid2",
		HasCheckedIn:    true,
		HasPickedUpSwag: false,
	}

	if !reflect.DeepEqual(checkin, &expected_checkin) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_checkin, checkin)
	}

	CleanupTestDB(t)
}

/*
	Service level test for updating user checkin in the db
*/
func TestUpdateUserCheckinService(t *testing.T) {
	SetupTestDB(t)

	checkin := models.UserCheckin{
		ID:              "testid",
		HasCheckedIn:    true,
		HasPickedUpSwag: false,
	}

	err := service.UpdateUserCheckin("testid", checkin)

	if err != nil {
		t.Fatal(err)
	}

	updated_checkin, err := service.GetUserCheckin("testid")

	if err != nil {
		t.Fatal(err)
	}

	expected_checkin := models.UserCheckin{
		ID:              "testid",
		HasCheckedIn:    true,
		HasPickedUpSwag: false,
	}

	if !reflect.DeepEqual(updated_checkin, &expected_checkin) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_checkin, updated_checkin)
	}

	CleanupTestDB(t)
}

/*
	Service level test for generating QR code URI
*/
func TestGetQrInfo(t *testing.T) {
	SetupTestDB(t)

	actual_uri, err := service.GetQrInfo("testid")

	if err != nil {
		t.Fatal(err)
	}

	parsed_uri, err := url.Parse(actual_uri)

	if err != nil {
		t.Fatal(err)
	}

	actual_query_params, err := url.ParseQuery(parsed_uri.RawQuery)

	if err != nil {
		t.Fatal(err)
	}

	expected_query_params := url.Values{
		"userId":          []string{"testid"},
		"hasCheckedIn":    []string{"true"},
		"hasPickedUpSwag": []string{"true"},
	}

	if !reflect.DeepEqual(expected_query_params, actual_query_params) {
		t.Errorf("Wrong QR code URI. Expected %v, got %v", expected_query_params, actual_query_params)
	}

	CleanupTestDB(t)
}
