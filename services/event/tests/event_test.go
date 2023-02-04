package tests

import (
	"fmt"
	"math"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/services/event/config"
	"github.com/HackIllinois/api/services/event/models"
	"github.com/HackIllinois/api/services/event/service"
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

	db, err = database.InitDatabase(config.EVENT_DB_HOST, config.EVENT_DB_NAME)

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	return_code := m.Run()

	os.Exit(return_code)
}

var TestTime = time.Now().Unix()

/*
	Initialize db with a test event
*/
func SetupTestDB(t *testing.T) {
	CleanupTestDB(t) // This prevents tests failing to cleanup from affecting other tests
	event := models.EventDB{
		EventPublic: models.EventPublic{
			ID:          "testid",
			Name:        "testname",
			Description: "testdescription",
			StartTime:   TestTime,
			EndTime:     TestTime + 60000,
			Sponsor:     "testsponsor",
			EventType:   "WORKSHOP",
			Locations: []models.EventLocation{
				{
					Description: "testlocationdescription",
					Tags:        []string{"ECEB1"},
					Latitude:    123.456,
					Longitude:   123.456,
				},
			},
			Points: 10,
		},
		IsPrivate:             false,
		DisplayOnStaffCheckin: false,
	}

	err := db.Insert("events", &event, nil)
	if err != nil {
		t.Fatal(err)
	}

	event_tracker := models.EventTracker{
		EventID: "testid",
		Users:   []string{},
	}

	err = db.Insert("eventtrackers", &event_tracker, nil)

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
	Service level test for getting all events from db
*/
func TestGetAllEventsService(t *testing.T) {
	SetupTestDB(t)
	defer CleanupTestDB(t)

	event := models.EventDB{
		EventPublic: models.EventPublic{
			ID:          "testid2",
			Name:        "testname2",
			Description: "testdescription2",
			StartTime:   TestTime,
			EndTime:     TestTime + 60000,
			Sponsor:     "testsponsor",
			EventType:   "WORKSHOP",
			Locations: []models.EventLocation{
				{
					Description: "testlocationdescription",
					Tags:        []string{"SIEBEL0", "ECEB1"},
					Latitude:    123.456,
					Longitude:   123.456,
				},
			},
		},
		IsPrivate:             false,
		DisplayOnStaffCheckin: false,
	}

	event2 := models.EventDB{
		EventPublic: models.EventPublic{
			ID:          "testid3",
			Name:        "testname3",
			Description: "testdescription3",
			StartTime:   TestTime,
			EndTime:     TestTime + 60000,
			Sponsor:     "",
			EventType:   "OTHER",
			Locations:   []models.EventLocation{},
			Points:      100,
		},
		IsPrivate:             true,
		DisplayOnStaffCheckin: true,
	}

	err := db.Insert("events", &event, nil)
	if err != nil {
		t.Fatal(err)
	}

	err = db.Insert("events", &event2, nil)

	if err != nil {
		t.Fatal(err)
	}

	actual_event_list, err := service.GetAllEvents[models.EventDB]()
	if err != nil {
		t.Fatal(err)
	}

	expected_event_list := models.EventList[models.EventDB]{
		Events: []models.EventDB{
			{
				EventPublic: models.EventPublic{
					ID:          "testid",
					Name:        "testname",
					Description: "testdescription",
					StartTime:   TestTime,
					EndTime:     TestTime + 60000,
					Sponsor:     "testsponsor",
					EventType:   "WORKSHOP",
					Locations: []models.EventLocation{
						{
							Description: "testlocationdescription",
							Tags:        []string{"ECEB1"},

							Latitude:  123.456,
							Longitude: 123.456,
						},
					},
					Points: 10,
				},
				IsPrivate:             false,
				DisplayOnStaffCheckin: false,
			},
			{
				EventPublic: models.EventPublic{
					ID:          "testid2",
					Name:        "testname2",
					Description: "testdescription2",
					StartTime:   TestTime,
					EndTime:     TestTime + 60000,
					Sponsor:     "testsponsor",
					EventType:   "WORKSHOP",
					Locations: []models.EventLocation{
						{
							Description: "testlocationdescription",
							Tags:        []string{"SIEBEL0", "ECEB1"},

							Latitude:  123.456,
							Longitude: 123.456,
						},
					},
					Points: 0,
				},
				IsPrivate:             false,
				DisplayOnStaffCheckin: false,
			},
			{
				EventPublic: models.EventPublic{
					ID:          "testid3",
					Name:        "testname3",
					Description: "testdescription3",
					StartTime:   TestTime,
					EndTime:     TestTime + 60000,
					Sponsor:     "",
					EventType:   "OTHER",
					Locations:   []models.EventLocation{},
					Points:      100,
				},
				IsPrivate:             true,
				DisplayOnStaffCheckin: true,
			},
		},
	}

	if !reflect.DeepEqual(actual_event_list, &expected_event_list) {
		t.Errorf("Wrong event list. Expected %v, got %v", expected_event_list, actual_event_list)
	}

	actual_public_list, err := service.GetAllEvents[models.EventPublic]()
	if err != nil {
		t.Fatal(err)
	}

	expected_public_list := models.EventList[models.EventPublic]{
		Events: []models.EventPublic{
			{
				ID:          "testid",
				Name:        "testname",
				Description: "testdescription",
				StartTime:   TestTime,
				EndTime:     TestTime + 60000,
				Sponsor:     "testsponsor",
				EventType:   "WORKSHOP",
				Locations: []models.EventLocation{
					{
						Description: "testlocationdescription",
						Tags:        []string{"ECEB1"},

						Latitude:  123.456,
						Longitude: 123.456,
					},
				},
				Points: 10,
			},
			{
				ID:          "testid2",
				Name:        "testname2",
				Description: "testdescription2",
				StartTime:   TestTime,
				EndTime:     TestTime + 60000,
				Sponsor:     "testsponsor",
				EventType:   "WORKSHOP",
				Locations: []models.EventLocation{
					{
						Description: "testlocationdescription",
						Tags:        []string{"SIEBEL0", "ECEB1"},

						Latitude:  123.456,
						Longitude: 123.456,
					},
				},
				Points: 0,
			},
		},
	}

	if !reflect.DeepEqual(actual_public_list, &expected_public_list) {
		t.Errorf("Wrong event list. Expected %v, got %v", expected_public_list, actual_public_list)
	}

	db.RemoveAll("events", nil, nil)

	actual_event_list, err = service.GetAllEvents[models.EventDB]()

	if err != nil {
		t.Fatal(err)
	}

	expected_event_list = models.EventList[models.EventDB]{
		Events: []models.EventDB{},
	}

	if !reflect.DeepEqual(actual_event_list, &expected_event_list) {
		t.Errorf("Wrong event list. Expected %v, got %v", expected_event_list, actual_event_list)
	}
}

/*
	Service level test for getting a filtered list of events from the db
*/
func TestGetFilteredEventsService(t *testing.T) {
	SetupTestDB(t)
	defer CleanupTestDB(t)

	event := models.EventDB{
		EventPublic: models.EventPublic{
			ID:          "testid2",
			Name:        "testname2",
			Description: "testdescription2",
			StartTime:   TestTime,
			EndTime:     TestTime + 60000,
			Sponsor:     "testsponsor",
			EventType:   "WORKSHOP",
			Locations: []models.EventLocation{
				{
					Description: "testlocationdescription",
					Tags:        []string{"SIEBEL0", "ECEB1"},

					Latitude:  123.456,
					Longitude: 123.456,
				},
			},
			Points: 0,
		},
		IsPrivate:             false,
		DisplayOnStaffCheckin: false,
	}

	err := db.Insert("events", &event, nil)
	if err != nil {
		t.Fatal(err)
	}

	event2 := models.EventDB{
		EventPublic: models.EventPublic{
			ID:          "testid3",
			Name:        "testname3",
			Description: "testdescription3",
			StartTime:   TestTime,
			EndTime:     TestTime + 60000,
			Sponsor:     "",
			EventType:   "OTHER",
			Locations:   []models.EventLocation{},
			Points:      100,
		},
		IsPrivate:             true,
		DisplayOnStaffCheckin: true,
	}

	err = db.Insert("events", &event2, nil)

	if err != nil {
		t.Fatal(err)
	}

	// Filter to one event
	parameters := map[string][]string{
		"name": {"testname2"},
	}
	actual_event_list, err := service.GetFilteredEvents[models.EventDB](parameters)
	if err != nil {
		t.Fatal(err)
	}

	expected_event_list := models.EventList[models.EventDB]{
		Events: []models.EventDB{
			{
				EventPublic: models.EventPublic{
					ID:          "testid2",
					Name:        "testname2",
					Description: "testdescription2",
					StartTime:   TestTime,
					EndTime:     TestTime + 60000,
					Sponsor:     "testsponsor",
					EventType:   "WORKSHOP",
					Locations: []models.EventLocation{
						{
							Description: "testlocationdescription",
							Tags:        []string{"SIEBEL0", "ECEB1"},

							Latitude:  123.456,
							Longitude: 123.456,
						},
					},
					Points: 0,
				},
				IsPrivate:             false,
				DisplayOnStaffCheckin: false,
			},
		},
	}

	if !reflect.DeepEqual(actual_event_list, &expected_event_list) {
		t.Errorf("Wrong event list. Expected %v, got %v", expected_event_list, actual_event_list)
	}

	// Filter to multiple (all) events
	parameters = map[string][]string{
		"sponsor": {"testsponsor"},
	}
	actual_public_list, err := service.GetFilteredEvents[models.EventPublic](parameters)
	if err != nil {
		t.Fatal(err)
	}

	expected_public_list := models.EventList[models.EventPublic]{
		Events: []models.EventPublic{
			{
				ID:          "testid",
				Name:        "testname",
				Description: "testdescription",
				StartTime:   TestTime,
				EndTime:     TestTime + 60000,
				Sponsor:     "testsponsor",
				EventType:   "WORKSHOP",
				Locations: []models.EventLocation{
					{
						Description: "testlocationdescription",
						Tags:        []string{"ECEB1"},

						Latitude:  123.456,
						Longitude: 123.456,
					},
				},
				Points: 10,
			},
			{
				ID:          "testid2",
				Name:        "testname2",
				Description: "testdescription2",
				StartTime:   TestTime,
				EndTime:     TestTime + 60000,
				Sponsor:     "testsponsor",
				EventType:   "WORKSHOP",
				Locations: []models.EventLocation{
					{
						Description: "testlocationdescription",
						Tags:        []string{"SIEBEL0", "ECEB1"},

						Latitude:  123.456,
						Longitude: 123.456,
					},
				},
				Points: 0,
			},
		},
	}

	if !reflect.DeepEqual(actual_public_list, &expected_public_list) {
		t.Errorf("Wrong event list. Expected %v, got %v", expected_public_list, actual_public_list)
	}

	// Try to filter an event that is private
	parameters = map[string][]string{
		"name": {"testname3"},
	}
	actual_public_list, err = service.GetFilteredEvents[models.EventPublic](parameters)

	if err != nil {
		t.Fatal(err)
	}

	expected_public_list = models.EventList[models.EventPublic]{
		Events: []models.EventPublic{},
	}

	if !reflect.DeepEqual(actual_public_list, &expected_public_list) {
		t.Errorf("Wrong event list. Expected %v, got %v", expected_public_list, actual_public_list)
	}

	db.RemoveAll("events", nil, nil)

	// Filter again, with no events remaining
	actual_event_list, err = service.GetFilteredEvents[models.EventDB](parameters)

	if err != nil {
		t.Fatal(err)
	}

	expected_event_list = models.EventList[models.EventDB]{
		Events: []models.EventDB{},
	}

	if !reflect.DeepEqual(actual_event_list, &expected_event_list) {
		t.Errorf("Wrong event list. Expected %v, got %v", expected_event_list, actual_event_list)
	}
}

/*
	Service level test for getting event from db
*/
func TestGetEventService(t *testing.T) {
	SetupTestDB(t)
	defer CleanupTestDB(t)

	hidden_event := models.EventDB{
		EventPublic: models.EventPublic{
			ID:          "secreteventid",
			Name:        "secretevent",
			Description: "You should not be able to see me",
			StartTime:   TestTime,
			EndTime:     TestTime + 60000,
			Sponsor:     "testsponsor",
			EventType:   "OTHER",
			Locations:   []models.EventLocation{},
			Points:      1337,
		},
		IsPrivate:             true,
		DisplayOnStaffCheckin: true,
	}

	err := db.Insert("events", hidden_event, nil)
	if err != nil {
		t.Fatal(err)
	}

	event, err := service.GetEvent[models.EventDB]("testid")
	if err != nil {
		t.Fatal(err)
	}

	expected_event := models.EventDB{
		EventPublic: models.EventPublic{
			ID:          "testid",
			Name:        "testname",
			Description: "testdescription",
			StartTime:   TestTime,
			EndTime:     TestTime + 60000,
			Sponsor:     "testsponsor",
			EventType:   "WORKSHOP",
			Locations: []models.EventLocation{
				{
					Description: "testlocationdescription",
					Tags:        []string{"ECEB1"},
					Latitude:    123.456,
					Longitude:   123.456,
				},
			},
			Points: 10,
		},
		IsPrivate:             false,
		DisplayOnStaffCheckin: false,
	}

	if !reflect.DeepEqual(event, &expected_event) {
		t.Errorf("Wrong event info. Expected %v, got %v", &expected_event, event)
	}

	_, err = service.GetEvent[models.EventPublic]("secreteventid")

	if err != database.ErrNotFound {
		t.Fatalf("Error found was not ErrNotFound (Got: %v)", err)
	}
}

/*
	Service level test for creating an event in the db
*/
func TestCreateEventService(t *testing.T) {
	SetupTestDB(t)
	defer CleanupTestDB(t)

	new_event := models.EventDB{
		EventPublic: models.EventPublic{
			ID:          "testid2",
			Name:        "testname2",
			Description: "testdescription2",
			StartTime:   TestTime,
			EndTime:     TestTime + 60000,
			Sponsor:     "testsponsor",
			EventType:   "WORKSHOP",
			Locations: []models.EventLocation{
				{
					Description: "testlocationdescription",
					Tags:        []string{"SIEBEL0", "ECEB1"},
					Latitude:    123.456,
					Longitude:   123.456,
				},
			},
		},
		IsPrivate:             true,
		DisplayOnStaffCheckin: false,
	}

	err := service.CreateEvent("testid2", "testcode2", new_event)
	if err != nil {
		t.Fatal(err)
	}

	event, err := service.GetEvent[models.EventDB]("testid2")
	if err != nil {
		t.Fatal(err)
	}

	expected_event := models.EventDB{
		EventPublic: models.EventPublic{
			ID:          "testid2",
			Name:        "testname2",
			Description: "testdescription2",
			StartTime:   TestTime,
			EndTime:     TestTime + 60000,
			Sponsor:     "testsponsor",
			EventType:   "WORKSHOP",
			Locations: []models.EventLocation{
				{
					Description: "testlocationdescription",
					Tags:        []string{"SIEBEL0", "ECEB1"},
					Latitude:    123.456,
					Longitude:   123.456,
				},
			},
			Points: 0,
		},
		IsPrivate:             true,
		DisplayOnStaffCheckin: false,
	}

	if !reflect.DeepEqual(event, &expected_event) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_event, event)
	}

	CleanupTestDB(t)

	SetupTestDB(t)

	new_event_async := models.EventDB{
		EventPublic: models.EventPublic{
			ID:          "testid2",
			Name:        "testname2",
			Description: "testdescription2",
			Sponsor:     "testsponsor",
			EventType:   "WORKSHOP",
			Locations: []models.EventLocation{
				{
					Description: "testlocationdescription",
					Tags:        []string{"SIEBEL0", "ECEB1"},
					Latitude:    123.456,
					Longitude:   123.456,
				},
			},
			IsAsync: true,
		},
		IsPrivate:             false,
		DisplayOnStaffCheckin: false,
	}

	err = service.CreateEvent("testid2", "testcode2", new_event_async)

	if err != nil {
		t.Fatal(err)
	}

	event_async, err := service.GetEvent[models.EventDB]("testid2")
	if err != nil {
		t.Fatal(err)
	}

	expected_event_async := models.EventDB{
		EventPublic: models.EventPublic{
			ID:          "testid2",
			Name:        "testname2",
			Description: "testdescription2",
			Sponsor:     "testsponsor",
			EventType:   "WORKSHOP",
			Locations: []models.EventLocation{
				{
					Description: "testlocationdescription",
					Tags:        []string{"SIEBEL0", "ECEB1"},
					Latitude:    123.456,
					Longitude:   123.456,
				},
			},
			IsAsync: true,
			Points:  0,
		},
		IsPrivate:             false,
		DisplayOnStaffCheckin: false,
	}

	if !reflect.DeepEqual(event_async, &expected_event_async) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_event, event)
	}
}

/*
	Service level test for deleting an event in the db
*/
func TestDeleteEventService(t *testing.T) {
	SetupTestDB(t)
	defer CleanupTestDB(t)

	event_id := "testid"

	// Mark 3 users as attending the event

	err := service.MarkUserAsAttendingEvent(event_id, "user0")
	if err != nil {
		t.Fatal(err)
	}

	err = service.MarkUserAsAttendingEvent(event_id, "user1")

	if err != nil {
		t.Fatal(err)
	}

	err = service.MarkUserAsAttendingEvent(event_id, "user2")

	if err != nil {
		t.Fatal(err)
	}

	// Try to delete the event

	_, err = service.DeleteEvent(event_id)

	if err != nil {
		t.Fatal(err)
	}

	// Try to find the event in the events db
	event, err := service.GetEvent[models.EventDB](event_id)

	if err == nil {
		t.Errorf("Found event %v in events database.", event)
	}

	// Try to find the event in the eventtrackers db
	event_tracker, err := service.GetEventTracker(event_id)

	if err == nil {
		t.Errorf("Found event in the eventtracker %v.", event_tracker)
	}

	// Try to find the event in the usertrackers db
	var user_trackers []models.UserTracker
	db.FindAll("usertrackers", nil, &user_trackers, nil)

	for _, user_tracker := range user_trackers {
		for _, event := range user_tracker.Events {
			if event == event_id {
				t.Errorf("Found event in the usertracker %v.", user_tracker)
			}
		}
	}
}

/*
	Service level test for updating an event in the db
*/
func TestUpdateEventService(t *testing.T) {
	SetupTestDB(t)
	defer CleanupTestDB(t)

	event := models.EventDB{
		EventPublic: models.EventPublic{
			ID:          "testid",
			Name:        "testname",
			Description: "testdescription2",
			StartTime:   TestTime,
			EndTime:     TestTime + 60000,
			Sponsor:     "testsponsor",
			EventType:   "WORKSHOP",
			Locations: []models.EventLocation{
				{
					Description: "testlocationdescription",
					Tags:        []string{"SIEBEL3", "ECEB2"},
					Latitude:    123.456,
					Longitude:   123.456,
				},
			},
			Points: 100,
		},
		IsPrivate:             false,
		DisplayOnStaffCheckin: false,
	}

	err := service.UpdateEvent("testid", event)
	if err != nil {
		t.Fatal(err)
	}

	updated_event, err := service.GetEvent[models.EventDB]("testid")
	if err != nil {
		t.Fatal(err)
	}

	expected_event := models.EventDB{
		EventPublic: models.EventPublic{
			ID:          "testid",
			Name:        "testname",
			Description: "testdescription2",
			StartTime:   TestTime,
			EndTime:     TestTime + 60000,
			Sponsor:     "testsponsor",
			EventType:   "WORKSHOP",
			Locations: []models.EventLocation{
				{
					Description: "testlocationdescription",
					Tags:        []string{"SIEBEL3", "ECEB2"},
					Latitude:    123.456,
					Longitude:   123.456,
				},
			},
			Points: 100,
		},
		IsPrivate:             false,
		DisplayOnStaffCheckin: false,
	}

	if !reflect.DeepEqual(updated_event, &expected_event) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_event, updated_event)
	}
}

/*
	Service level test for marking a user as attending an event
*/
func TestMarkUserAsAttendingEventService(t *testing.T) {
	SetupTestDB(t)
	defer CleanupTestDB(t)

	err := service.MarkUserAsAttendingEvent("testid", "testuser")
	if err != nil {
		t.Fatal(err)
	}

	event_tracker, err := service.GetEventTracker("testid")
	if err != nil {
		t.Fatal(err)
	}

	expected_event_tracker := models.EventTracker{
		EventID: "testid",
		Users:   []string{"testuser"},
	}

	if !reflect.DeepEqual(event_tracker, &expected_event_tracker) {
		t.Errorf("Wrong tracker info. Expected %v, got %v", expected_event_tracker, event_tracker)
	}

	user_tracker, err := service.GetUserTracker("testuser")
	if err != nil {
		t.Fatal(err)
	}

	expected_user_tracker := models.UserTracker{
		UserID: "testuser",
		Events: []string{"testid"},
	}

	if !reflect.DeepEqual(user_tracker, &expected_user_tracker) {
		t.Errorf("Wrong tracker info. Expected %v, got %v", expected_user_tracker, user_tracker)
	}
}

/*
	Service level test for marking a user as attending an event
	when they have already been marked as attending
*/
func TestMarkUserAsAttendingEventErrorService(t *testing.T) {
	SetupTestDB(t)
	defer CleanupTestDB(t)

	err := service.MarkUserAsAttendingEvent("testid", "testuser")
	if err != nil {
		t.Fatal(err)
	}

	err = service.MarkUserAsAttendingEvent("testid", "testuser")

	if err == nil {
		t.Fatal("User was marked as attending event twice")
	}

	event_tracker, err := service.GetEventTracker("testid")
	if err != nil {
		t.Fatal(err)
	}

	expected_event_tracker := models.EventTracker{
		EventID: "testid",
		Users:   []string{"testuser"},
	}

	if !reflect.DeepEqual(event_tracker, &expected_event_tracker) {
		t.Errorf("Wrong tracker info. Expected %v, got %v", expected_event_tracker, event_tracker)
	}

	user_tracker, err := service.GetUserTracker("testuser")
	if err != nil {
		t.Fatal(err)
	}

	expected_user_tracker := models.UserTracker{
		UserID: "testuser",
		Events: []string{"testid"},
	}

	if !reflect.DeepEqual(user_tracker, &expected_user_tracker) {
		t.Errorf("Wrong tracker info. Expected %v, got %v", expected_user_tracker, user_tracker)
	}
}

/*
	Adds an event with the current time, and checks if it is active.
	Confirms if an event that is known to be inactive (time is in the past), is inactive.
*/
func TestIsEventActive(t *testing.T) {
	defer CleanupTestDB(t)
	// Creating a 30 minute long event that SHOULD NOT be active
	const ONE_MINUTE_IN_SECONDS = 60
	new_event := models.EventDB{
		EventPublic: models.EventPublic{
			ID:          "testid3",
			Name:        "testiseventactive",
			Description: "testdescription2",
			StartTime:   TestTime + ONE_MINUTE_IN_SECONDS*40,
			EndTime:     TestTime + ONE_MINUTE_IN_SECONDS*70,
			Sponsor:     "testsponsor",
			EventType:   "WORKSHOP",
			Locations: []models.EventLocation{
				{
					Description: "testlocationdescription",
					Tags:        []string{"SIEBEL3", "ECEB2"},
					Latitude:    123.456,
					Longitude:   123.456,
				},
			},
		},
		IsPrivate:             false,
		DisplayOnStaffCheckin: false,
	}

	service.CreateEvent(new_event.ID, "testcode3", new_event)

	is_active, err := service.IsEventActive("testid3")
	if err != nil {
		t.Fatal(err)
	}

	if is_active {
		current_time := TestTime
		t.Errorf(
			"Event was incorrectly deemed active. Current time: %v, event start time: %v, time difference: %v",
			current_time,
			new_event.StartTime,
			math.Abs((float64)(current_time-new_event.StartTime)),
		)
	}

	// Creating a 20 minute long event that SHOULD be active
	new_event.ID = "testid4"
	new_event.Name = "test2iseventactive"
	new_event.StartTime = TestTime
	new_event.EndTime = TestTime + ONE_MINUTE_IN_SECONDS*20

	service.CreateEvent(new_event.ID, "testcode4", new_event)

	is_active, err = service.IsEventActive("testid4")

	if err != nil {
		t.Fatal(err)
	}

	if !is_active {
		current_time := TestTime
		t.Errorf(
			"Event was incorrectly deemed inactive. Current time: %v, event start time: %v, time difference: %v",
			current_time,
			new_event.StartTime,
			math.Abs((float64)((current_time - new_event.StartTime))),
		)
	}
}

/*
	Tests that getting event favorites works correctly at the service level
*/
func TestGetEventFavorites(t *testing.T) {
	SetupTestDB(t)
	defer CleanupTestDB(t)

	event_favorites, err := service.GetEventFavorites("testid")
	if err != nil {
		t.Fatal(err)
	}

	expected_event_favorites := models.EventFavorites{
		ID:     "testid",
		Events: []string{},
	}

	if !reflect.DeepEqual(event_favorites, &expected_event_favorites) {
		t.Errorf("Wrong tracker info. Expected %v, got %v", &expected_event_favorites, event_favorites)
	}
}

/*
	Tests that adding event favorites works correctly at the service level
*/
func TestAddEventFavorite(t *testing.T) {
	SetupTestDB(t)
	defer CleanupTestDB(t)

	err := service.AddEventFavorite("testid", "testid")
	if err != nil {
		t.Fatal(err)
	}

	event_favorites, err := service.GetEventFavorites("testid")
	if err != nil {
		t.Fatal(err)
	}

	expected_event_favorites := models.EventFavorites{
		ID:     "testid",
		Events: []string{"testid"},
	}

	if !reflect.DeepEqual(event_favorites, &expected_event_favorites) {
		t.Errorf("Wrong tracker info. Expected %v, got %v", &expected_event_favorites, event_favorites)
	}
}

/*
	Tests that removing event favorites works correctly at the service level
*/
func TestRemoveEventFavorite(t *testing.T) {
	SetupTestDB(t)
	defer CleanupTestDB(t)

	err := service.AddEventFavorite("testid", "testid")
	if err != nil {
		t.Fatal(err)
	}

	err = service.RemoveEventFavorite("testid", "testid")

	if err != nil {
		t.Fatal(err)
	}

	event_favorites, err := service.GetEventFavorites("testid")
	if err != nil {
		t.Fatal(err)
	}

	expected_event_favorites := models.EventFavorites{
		ID:     "testid",
		Events: []string{},
	}

	if !reflect.DeepEqual(event_favorites, &expected_event_favorites) {
		t.Errorf("Wrong tracker info. Expected %v, got %v", &expected_event_favorites, event_favorites)
	}
}
