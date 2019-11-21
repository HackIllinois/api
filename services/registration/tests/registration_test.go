package tests

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/common/datastore"
	"github.com/HackIllinois/api/services/registration/config"
	"github.com/HackIllinois/api/services/registration/models"
	"github.com/HackIllinois/api/services/registration/service"
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

	db, err = database.InitDatabase(config.REGISTRATION_DB_HOST, config.REGISTRATION_DB_NAME)

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	return_code := m.Run()

	os.Exit(return_code)
}

/*
	Initialize db with test user and mentor info
*/
func SetupTestDB(t *testing.T) {
	user_registration := getBaseUserRegistration()
	err := db.Insert("attendees", &user_registration)

	if err != nil {
		t.Fatal(err)
	}

	mentor_registration := getBaseMentorRegistration()
	err = db.Insert("mentors", &mentor_registration)

	if err != nil {
		t.Fatal(err)
	}
}

/*
	Drop test db
*/
func CleanupTestDB(t *testing.T) {
	err := db.DropDatabase()

	if err != nil {
		t.Fatal(err)
	}
}

/*
	Service level test for getting user registration from db
*/
func TestGetUserRegistrationService(t *testing.T) {
	SetupTestDB(t)

	user_registration, err := service.GetUserRegistration("testid")

	if err != nil {
		t.Fatal(err)
	}

	expected_registration := getBaseUserRegistration()

	if !reflect.DeepEqual(user_registration.Data["firstName"], expected_registration.Data["firstName"]) {
		t.Errorf("Wrong user info.\nExpected %v\ngot %v\n", expected_registration.Data["firstName"], user_registration.Data["firstName"])
	}

	CleanupTestDB(t)
}

/*
	Service level test for creating user registration in the db
*/
func TestCreateUserRegistrationService(t *testing.T) {
	SetupTestDB(t)

	new_registration := getBaseUserRegistration()
	new_registration.Data["id"] = "testid2"
	new_registration.Data["firstName"] = "first2"
	new_registration.Data["lastName"] = "last2"
	err := service.CreateUserRegistration("testid2", new_registration)

	if err != nil {
		t.Fatal(err)
	}

	user_registration, err := service.GetUserRegistration("testid2")

	if err != nil {
		t.Fatal(err)
	}

	expected_registration := getBaseUserRegistration()
	expected_registration.Data["id"] = "testid2"
	expected_registration.Data["firstName"] = "first2"
	expected_registration.Data["lastName"] = "last2"

	if !reflect.DeepEqual(user_registration.Data["firstName"], expected_registration.Data["firstName"]) {
		t.Errorf("Wrong user info.\nExpected %v\ngot %v\n", expected_registration.Data["firstName"], user_registration.Data["firstName"])
	}

	CleanupTestDB(t)
}

/*
	Service level test for updating user registration in the db
*/
func TestUpdateUserRegistrationService(t *testing.T) {
	SetupTestDB(t)

	updated_registration := getBaseUserRegistration()
	updated_registration.Data["id"] = "testid"
	updated_registration.Data["firstName"] = "first2"
	updated_registration.Data["lastName"] = "last2"
	err := service.UpdateUserRegistration("testid", updated_registration)

	if err != nil {
		t.Fatal(err)
	}

	user_registration, err := service.GetUserRegistration("testid")

	if err != nil {
		t.Fatal(err)
	}

	expected_registration := getBaseUserRegistration()
	expected_registration.Data["id"] = "testid"
	expected_registration.Data["firstName"] = "first2"
	expected_registration.Data["lastName"] = "last2"

	// if !reflect.DeepEqual(user_registration.Data["firstName"], expected_registration.Data["firstName"]) {
	// 	t.Errorf("Wrong user info.\nExpected %v\ngot %v\n", expected_registration.Data["firstName"], user_registration.Data["firstName"])
	// }

	if !reflect.DeepEqual(user_registration.Data, expected_registration.Data) {
		t.Errorf("Wrong user info.\nExpected %v\ngot %v\n", expected_registration.Data, user_registration.Data)
	}

	CleanupTestDB(t)
}

/*
	Service level test for updating user registration in the db
*/
func TestPatchUserRegistrationService(t *testing.T) {
	SetupTestDB(t)
	fmt.Print("Running the patch test")
	updated_registration := getEmptyUserRegistration()
	updated_registration.Data["email"] = "edited@gmail.com"
	updated_registration.Data["isBeginner"] = true
	updated_registration.Data["priorAttendance"] = false
	updated_registration.Data["age"] = 22

	err := service.PatchUserRegistration("testid", updated_registration)

	if err != nil {
		t.Fatal(err)
	}

	user_registration, err := service.GetUserRegistration("testid")

	if err != nil {
		t.Fatal(err)
	}

	expected_registration := getBaseUserRegistration()
	expected_registration.Data["id"] = "testid"
	expected_registration.Data["firstName"] = "first"
	expected_registration.Data["lastName"] = "last"

	expected_registration.Data["shirtSize"] = "M"
	expected_registration.Data["github"] = "githubusername"
	expected_registration.Data["linkedin"] = "linkedinusername"
	expected_registration.Data["age"] = 22
	expected_registration.Data["createdAt"] = int64(10)

	expected_registration.Data["email"] = "edited@gmail.com"
	expected_registration.Data["isBeginner"] = true
	expected_registration.Data["priorAttendance"] = false
	expected_registration.Data["age"] = 22

	// user_registration := getBaseUserRegistration()
	// user_registration.Data["id"] = "testid"
	// user_registration.Data["firstName"] = "first"
	// user_registration.Data["lastName"] = "last"
	// user_registration.Data["email"] = "edited@gmail.com"
	// user_registration.Data["shirtSize"] = "M"
	// user_registration.Data["github"] = "githubusername"
	// user_registration.Data["linkedin"] = "linkedinusername"
	// user_registration.Data["age"] = 22
	// user_registration.Data["createdAt"] = 10
	// user_registration.Data["priorAttendance"] = false
	// user_registration.Data["isBeginner"] = true
	fmt.Printf("\nType of user_registration.Data[createdAt]: %T Type of expected_registration.Data[createdAt]: %T\n", user_registration.Data["createdAt"], expected_registration.Data["createdAt"])
	fmt.Print("\nType of user_registration.Data: ", reflect.TypeOf(user_registration.Data), " Type of expected_registration.Data: ", reflect.TypeOf(expected_registration.Data), "\n")
	fmt.Print("\nLength of user_registration.Data: ", len(user_registration.Data), " Length of expected_registration.Data: ", len(expected_registration.Data), "\n")
	if !reflect.DeepEqual(user_registration.Data, expected_registration.Data) {
		t.Errorf("Wrong user info.\nExpected %v\nGot %v\n", expected_registration.Data, user_registration.Data)
	}

	// if !reflect.DeepEqual(user_registration.Data["firstName"], expected_registration.Data["firstName"]) {
	// 	t.Errorf("Wrong user info.\nExpected %v\ngot %v\n", expected_registration.Data["firstName"], user_registration.Data["firstName"])
	// }

	CleanupTestDB(t)
}

// map[age:22 createdAt:10 email:edited@gmail.com firstName:first github:githubusername id:testid isBeginner:true lastName:last linkedin:linkedinusername priorAttendance:false shirtSize:M updatedAt:15]
// map[age:22 createdAt:10 email:edited@gmail.com firstName:first github:githubusername id:testid isBeginner:true lastName:last linkedin:linkedinusername priorAttendance:false shirtSize:M updatedAt:15]
/*
	Service level test for getting mentor registration from db
*/
func TestGetMentorRegistrationService(t *testing.T) {
	SetupTestDB(t)

	mentor_registration, err := service.GetMentorRegistration("testid")

	if err != nil {
		t.Fatal(err)
	}

	expected_registration := getBaseMentorRegistration()

	if !reflect.DeepEqual(mentor_registration.Data["firstName"], expected_registration.Data["firstName"]) {
		t.Errorf("Wrong mentor info.\nExpected %v\ngot %v\n", expected_registration.Data["firstName"], mentor_registration.Data["firstName"])
	}

	CleanupTestDB(t)
}

/*
	Service level test for creating mentor registration in the db
*/
func TestCreateMentorRegistrationService(t *testing.T) {
	SetupTestDB(t)

	new_registration := getBaseMentorRegistration()
	new_registration.Data["id"] = "testid2"
	new_registration.Data["firstName"] = "first2"
	new_registration.Data["lastName"] = "last2"
	err := service.CreateMentorRegistration("testid2", new_registration)

	if err != nil {
		t.Fatal(err)
	}

	mentor_registration, err := service.GetMentorRegistration("testid2")

	if err != nil {
		t.Fatal(err)
	}

	expected_registration := getBaseMentorRegistration()
	expected_registration.Data["id"] = "testid2"
	expected_registration.Data["firstName"] = "first2"
	expected_registration.Data["lastName"] = "last2"

	if !reflect.DeepEqual(mentor_registration.Data["firstName"], expected_registration.Data["firstName"]) {
		t.Errorf("Wrong mentor info.\nExpected %v\ngot %v\n", expected_registration.Data["firstName"], mentor_registration.Data["firstName"])
	}

	CleanupTestDB(t)
}

/*
	Service level test for updating mentor registration in the db
*/
func TestUpdateMentorRegistrationService(t *testing.T) {
	SetupTestDB(t)

	updated_registration := getBaseMentorRegistration()
	updated_registration.Data["id"] = "testid"
	updated_registration.Data["firstName"] = "first2"
	updated_registration.Data["lastName"] = "last2"
	err := service.UpdateMentorRegistration("testid", updated_registration)

	if err != nil {
		t.Fatal(err)
	}

	mentor_registration, err := service.GetMentorRegistration("testid")

	if err != nil {
		t.Fatal(err)
	}

	expected_registration := getBaseMentorRegistration()
	expected_registration.Data["id"] = "testid"
	expected_registration.Data["firstName"] = "first2"
	expected_registration.Data["lastName"] = "last2"

	if !reflect.DeepEqual(mentor_registration.Data["firstName"], expected_registration.Data["firstName"]) {
		t.Errorf("Wrong mentor info.\nExpected %v\ngot %v\n", expected_registration.Data["firstName"], mentor_registration.Data["firstName"])
	}

	CleanupTestDB(t)
}

/*
	Service level test for filtering user registrations in the db
*/
func TestGetFilteredUserRegistrationsService(t *testing.T) {
	SetupTestDB(t)

	registration_1 := getBaseUserRegistration()

	registration_2 := getBaseUserRegistration()
	registration_2.Data["id"] = "testid2"

	err := service.CreateUserRegistration(registration_2.Data["id"].(string), registration_2)
	if err != nil {
		t.Fatal(err)
	}

	// Test single value and one keys
	parameters := map[string][]string{
		"id": {"testid"},
	}
	user_registrations, err := service.GetFilteredUserRegistrations(parameters)
	if err != nil {
		t.Fatal(err)
	}

	expected_registrations := models.FilteredUserRegistrations{
		[]models.UserRegistration{
			registration_1,
		},
	}

	if len(user_registrations.Registrations) != len(expected_registrations.Registrations) {
		t.Errorf("Wrong number of registrations.\nExpected %v\ngot %v\n", len(expected_registrations.Registrations), len(user_registrations.Registrations))
	}

	if !reflect.DeepEqual(user_registrations.Registrations[0].Data["firstName"], expected_registrations.Registrations[0].Data["firstName"]) {
		t.Errorf("Wrong user info.\nExpected %v\ngot %v\n", expected_registrations.Registrations[0].Data["firstName"], user_registrations.Registrations[0].Data["firstName"])
	}

	// Test multiple values
	parameters = map[string][]string{
		"id": {"testid,testid2"},
	}
	user_registrations, err = service.GetFilteredUserRegistrations(parameters)
	if err != nil {
		t.Fatal(err)
	}

	expected_registrations = models.FilteredUserRegistrations{
		[]models.UserRegistration{
			registration_1,
			registration_2,
		},
	}

	if len(user_registrations.Registrations) != len(expected_registrations.Registrations) {
		t.Errorf("Wrong number of registrations.\nExpected %v\ngot %v\n", len(expected_registrations.Registrations), len(user_registrations.Registrations))
	}

	if !reflect.DeepEqual(user_registrations.Registrations[1].Data["firstName"], expected_registrations.Registrations[1].Data["firstName"]) {
		t.Errorf("Wrong user info.\nExpected %v\ngot %v\n", expected_registrations.Registrations[1].Data["firstName"], user_registrations.Registrations[1].Data["firstName"])
	}

	// Test type casting
	parameters = map[string][]string{
		"firstName": {"first"},
		"age":       {"20"},
	}
	user_registrations, err = service.GetFilteredUserRegistrations(parameters)
	if err != nil {
		t.Fatal(err)
	}

	expected_registrations = models.FilteredUserRegistrations{
		[]models.UserRegistration{
			registration_1,
			registration_2,
		},
	}

	if len(user_registrations.Registrations) != len(expected_registrations.Registrations) {
		t.Errorf("Wrong number of registrations.\nExpected %v\ngot %v\n", len(expected_registrations.Registrations), len(user_registrations.Registrations))
	}

	if !reflect.DeepEqual(user_registrations.Registrations[1].Data["firstName"], expected_registrations.Registrations[1].Data["firstName"]) {
		t.Errorf("Wrong user info.\nExpected %v\ngot %v\n", expected_registrations.Registrations[1].Data["firstName"], user_registrations.Registrations[1].Data["firstName"])
	}

	CleanupTestDB(t)
}

/*
	Service level test for filtering mentor registrations in the db
*/
func TestGetFilteredMentorRegistrationsService(t *testing.T) {
	SetupTestDB(t)

	registration_1 := getBaseMentorRegistration()

	registration_2 := getBaseMentorRegistration()
	registration_2.Data["id"] = "testid2"

	err := service.CreateMentorRegistration(registration_2.Data["id"].(string), registration_2)
	if err != nil {
		t.Fatal(err)
	}

	// Test single value and one keys
	parameters := map[string][]string{
		"id": {"testid"},
	}
	mentor_registrations, err := service.GetFilteredMentorRegistrations(parameters)
	if err != nil {
		t.Fatal(err)
	}

	expected_registrations := models.FilteredMentorRegistrations{
		[]models.MentorRegistration{
			registration_1,
		},
	}

	if len(mentor_registrations.Registrations) != len(expected_registrations.Registrations) {
		t.Errorf("Wrong number of registrations.\nExpected %v\ngot %v\n", len(expected_registrations.Registrations), len(mentor_registrations.Registrations))
	}

	if !reflect.DeepEqual(mentor_registrations.Registrations[0].Data["firstName"], expected_registrations.Registrations[0].Data["firstName"]) {
		t.Errorf("Wrong user info.\nExpected %v\ngot %v\n", expected_registrations.Registrations[0].Data["firstName"], mentor_registrations.Registrations[0].Data["firstName"])
	}

	// Test multiple values
	parameters = map[string][]string{
		"id": {"testid,testid2"},
	}
	mentor_registrations, err = service.GetFilteredMentorRegistrations(parameters)
	if err != nil {
		t.Fatal(err)
	}

	expected_registrations = models.FilteredMentorRegistrations{
		[]models.MentorRegistration{
			registration_1,
			registration_2,
		},
	}

	if len(mentor_registrations.Registrations) != len(expected_registrations.Registrations) {
		t.Errorf("Wrong number of registrations.\nExpected %v\ngot %v\n", len(expected_registrations.Registrations), len(mentor_registrations.Registrations))
	}

	if !reflect.DeepEqual(mentor_registrations.Registrations[1].Data["firstName"], expected_registrations.Registrations[1].Data["firstName"]) {
		t.Errorf("Wrong user info.\nExpected %v\ngot %v\n", expected_registrations.Registrations[1].Data["firstName"], mentor_registrations.Registrations[1].Data["firstName"])
	}

	CleanupTestDB(t)
}

/*
	Returns a basic user registration
*/
func getBaseUserRegistration() datastore.DataStore {
	base_user_registration := datastore.NewDataStore(config.REGISTRATION_DEFINITION)
	json.Unmarshal([]byte(user_registration_data), &base_user_registration)
	return base_user_registration
}

/*
	Returns a basic mentor registration
*/
func getBaseMentorRegistration() datastore.DataStore {
	base_mentor_registration := datastore.NewDataStore(config.MENTOR_REGISTRATION_DEFINITION)
	json.Unmarshal([]byte(user_registration_data), &base_mentor_registration)
	return base_mentor_registration
}

/*
	Returns an empty user registration
*/
func getEmptyUserRegistration() datastore.DataStore {
	empty_user_registration := datastore.NewDataStore(config.REGISTRATION_DEFINITION)
	json.Unmarshal([]byte(empty_registration_data), &empty_user_registration)
	return empty_user_registration
}

var empty_registration_data string = `{}`
var user_registration_data string = `
{
	"id": "testid",
	"firstName": "first",
	"lastName": "last",
	"email": "test@gmail.com",
	"shirtSize": "M",
	"github": "githubusername",
	"linkedin": "linkedinusername",
	"age": 20,
	"createdAt": 10,
	"updatedAt": 15
}
`

var mentor_registration_data string = `
{
	"id": "testid",
	"firstName": "first",
	"lastName": "last",
	"email": "test@gmail.com",
	"shirtSize": "M",
	"github": "githubusername",
	"linkedin": "linkedinusername",
	"createdAt": 10,
	"updatedAt": 15
}
`
