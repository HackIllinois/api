package tests

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/services/mail/config"
	"github.com/HackIllinois/api/services/mail/models"
	"github.com/HackIllinois/api/services/mail/service"
)

var db database.Database

func TestMain(m *testing.M) {
	err := config.Initialize()

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)

	}

	err = service.Initialize()

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	db, err = database.InitDatabase(config.MAIL_DB_HOST, config.MAIL_DB_NAME)

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	return_code := m.Run()

	os.Exit(return_code)
}

/*
	Initialize databse with test user info
*/
func SetupTestDB(t *testing.T) {
	err := db.Insert("lists", &models.MailList{
		ID:      "testlist",
		UserIDs: []string{"userid1", "userid2"},
	}, nil)

	if err != nil {
		t.Fatal(err)
	}
}

/*
	Drop test db
*/
func CleanupTestDB(t *testing.T) {
	err := db.DropDatabase(nil)

	if err != nil {
		t.Fatal(err)
	}
}

/*
	Tests getting a mail list from the db at the service level
*/
func TestGetMailListService(t *testing.T) {
	SetupTestDB(t)

	mail_list, err := service.GetMailList("testlist")

	if err != nil {
		t.Fatal(err)
	}

	expected := &models.MailList{
		ID:      "testlist",
		UserIDs: []string{"userid1", "userid2"},
	}

	if !reflect.DeepEqual(mail_list, expected) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected, mail_list)
	}

	CleanupTestDB(t)
}

/*
	Tests creating a mailing list
*/
func TestCreateMailListService(t *testing.T) {
	SetupTestDB(t)

	mail_list := models.MailList{
		ID:      "testlist2",
		UserIDs: []string{"userid1", "userid2"},
	}

	err := service.CreateMailList(mail_list)

	if err != nil {
		t.Fatal(err)
	}

	retreived_list, err := service.GetMailList("testlist2")

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(retreived_list, &mail_list) {
		t.Errorf("Wrong user info. Expected %v, got %v", &mail_list, retreived_list)
	}

	CleanupTestDB(t)
}

/*
	Tests adding to a mailing list
*/
func TestAddToMailList(t *testing.T) {
	SetupTestDB(t)

	mail_list := models.MailList{
		ID:      "testlist",
		UserIDs: []string{"userid3", "userid4"},
	}

	err := service.AddToMailList(mail_list)

	if err != nil {
		t.Fatal(err)
	}

	retreived_list, err := service.GetMailList("testlist")

	if err != nil {
		t.Fatal(err)
	}

	expected := &models.MailList{
		ID:      "testlist",
		UserIDs: []string{"userid1", "userid2", "userid3", "userid4"},
	}

	if !reflect.DeepEqual(retreived_list, expected) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected, retreived_list)
	}

	CleanupTestDB(t)
}

/*
	Tests removing from a mailing list
*/
func TestRemoveFromMailList(t *testing.T) {
	SetupTestDB(t)

	mail_list := models.MailList{
		ID:      "testlist",
		UserIDs: []string{"userid2"},
	}

	err := service.RemoveFromMailList(mail_list)

	if err != nil {
		t.Fatal(err)
	}

	retreived_list, err := service.GetMailList("testlist")

	if err != nil {
		t.Fatal(err)
	}

	expected := &models.MailList{
		ID:      "testlist",
		UserIDs: []string{"userid1"},
	}

	if !reflect.DeepEqual(retreived_list, expected) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected, retreived_list)
	}

	CleanupTestDB(t)
}

/*
	Tests getting a list of all mailing lists
*/
func TestGetAllMailLists(t *testing.T) {
	SetupTestDB(t)

	mail_list := models.MailList{
		ID:      "testlist2",
		UserIDs: []string{"userid1", "userid2"},
	}

	err := service.CreateMailList(mail_list)

	if err != nil {
		t.Fatal(err)
	}

	mail_lists, err := service.GetAllMailLists()

	if err != nil {
		t.Fatal(err)
	}

	expected := &models.MailListList{
		MailLists: []models.MailList{
			{
				ID:      "testlist",
				UserIDs: []string{"userid1", "userid2"},
			},
			{
				ID:      "testlist2",
				UserIDs: []string{"userid1", "userid2"},
			},
		},
	}

	if !reflect.DeepEqual(mail_lists, expected) {
		t.Errorf("Wrong set of mailing lists. Expected %v, got %v", expected, mail_lists)
	}

	CleanupTestDB(t)
}
