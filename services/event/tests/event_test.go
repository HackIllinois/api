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
	event := models.Event{
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
	}

	err := db.Insert("events", &event)

	if err != nil {
		t.Fatal(err)
	}

	event_tracker := models.EventTracker{
		EventID: "testid",
		Users:   []string{},
	}

	err = db.Insert("eventtrackers", &event_tracker)

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
	Service level test for getting all events from db
*/
func TestGetAllEventsService(t *testing.T) {
	SetupTestDB(t)

	event := models.Event{
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
	}

	err := db.Insert("events", &event)

	if err != nil {
		t.Fatal(err)
	}

	actual_event_list, err := service.GetAllEvents()

	if err != nil {
		t.Fatal(err)
	}

	expected_event_list := models.EventList{
		Events: []models.Event{
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
			},
		},
	}

	if !reflect.DeepEqual(actual_event_list, &expected_event_list) {
		t.Errorf("Wrong event list. Expected %v, got %v", expected_event_list, actual_event_list)
	}

	db.RemoveAll("events", nil)

	actual_event_list, err = service.GetAllEvents()

	if err != nil {
		t.Fatal(err)
	}

	expected_event_list = models.EventList{
		Events: []models.Event{},
	}

	if !reflect.DeepEqual(actual_event_list, &expected_event_list) {
		t.Errorf("Wrong event list. Expected %v, got %v", expected_event_list, actual_event_list)
	}

	CleanupTestDB(t)

}

/*
	Service level test for getting a filtered list of events from the db
*/
func TestGetFilteredEventsService(t *testing.T) {
	SetupTestDB(t)

	event := models.Event{
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
	}

	err := db.Insert("events", &event)

	if err != nil {
		t.Fatal(err)
	}

	// Filter to one event
	parameters := map[string][]string{
		"name": {"testname2"},
	}
	actual_event_list, err := service.GetFilteredEvents(parameters)

	if err != nil {
		t.Fatal(err)
	}

	expected_event_list := models.EventList{
		Events: []models.Event{
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
	actual_event_list, err = service.GetFilteredEvents(parameters)

	if err != nil {
		t.Fatal(err)
	}

	expected_event_list = models.EventList{
		Events: []models.Event{
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
			},
		},
	}

	if !reflect.DeepEqual(actual_event_list, &expected_event_list) {
		t.Errorf("Wrong event list. Expected %v, got %v", expected_event_list, actual_event_list)
	}

	db.RemoveAll("events", nil)

	// Filter again, with no events remaining
	actual_event_list, err = service.GetFilteredEvents(parameters)

	if err != nil {
		t.Fatal(err)
	}

	expected_event_list = models.EventList{
		Events: []models.Event{},
	}

	if !reflect.DeepEqual(actual_event_list, &expected_event_list) {
		t.Errorf("Wrong event list. Expected %v, got %v", expected_event_list, actual_event_list)
	}

	CleanupTestDB(t)

}

/*
	Service level test for getting event from db
*/
func TestGetEventService(t *testing.T) {
	SetupTestDB(t)

	event, err := service.GetEvent("testid")

	if err != nil {
		t.Fatal(err)
	}

	expected_event := models.Event{
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
	}

	if !reflect.DeepEqual(event, &expected_event) {
		t.Errorf("Wrong event info. Expected %v, got %v", &expected_event, event)
	}

	CleanupTestDB(t)
}

/*
	Service level test for creating an event in the db
*/
func TestCreateEventService(t *testing.T) {
	SetupTestDB(t)

	new_event := models.Event{
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
	}

	err := service.CreateEvent("testid2", new_event)

	if err != nil {
		t.Fatal(err)
	}

	event, err := service.GetEvent("testid2")

	if err != nil {
		t.Fatal(err)
	}

	expected_event := models.Event{
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
	}

	if !reflect.DeepEqual(event, &expected_event) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_event, event)
	}

	CleanupTestDB(t)
}

/*
	Service level test for deleting an event in the db
*/
func TestDeleteEventService(t *testing.T) {
	SetupTestDB(t)

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
	event, err := service.GetEvent(event_id)

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
	db.FindAll("usertrackers", nil, &user_trackers)

	for _, user_tracker := range user_trackers {
		for _, event := range user_tracker.Events {
			if event == event_id {
				t.Errorf("Found event in the usertracker %v.", user_tracker)
			}
		}
	}

	CleanupTestDB(t)
}

/*
	Service level test for updating an event in the db
*/
func TestUpdateEventService(t *testing.T) {
	SetupTestDB(t)

	event := models.Event{
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
	}

	err := service.UpdateEvent("testid", event)

	if err != nil {
		t.Fatal(err)
	}

	updated_event, err := service.GetEvent("testid")

	if err != nil {
		t.Fatal(err)
	}

	expected_event := models.Event{
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
	}

	if !reflect.DeepEqual(updated_event, &expected_event) {
		t.Errorf("Wrong user info. Expected %v, got %v", expected_event, updated_event)
	}

	CleanupTestDB(t)
}

/*
	Service level test for marking a user as attending an event
*/
func TestMarkUserAsAttendingEventService(t *testing.T) {
	SetupTestDB(t)

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

	CleanupTestDB(t)
}

/*
	Service level test for marking a user as attending an event
	when they have already been marked as attending
*/
func TestMarkUserAsAttendingEventErrorService(t *testing.T) {
	SetupTestDB(t)

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

	CleanupTestDB(t)
}

/*
	Adds an event with the current time, and checks if it is active.
	Confirms if an event that is known to be inactive (time is in the past), is inactive.
*/
func TestIsEventActive(t *testing.T) {

	// Creating a 30 minute long event that SHOULD NOT be active
	const ONE_MINUTE_IN_SECONDS = 60
	new_event := models.Event{
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
	}

	service.CreateEvent(new_event.ID, new_event)

	is_active, err := service.IsEventActive("testid3")

	if err != nil {
		t.Fatal(err)
	}

	if is_active {
		current_time := TestTime
		t.Errorf("Event was incorrectly deemed active. Current time: %v, event start time: %v, time difference: %v", current_time, new_event.StartTime, math.Abs((float64)(current_time-new_event.StartTime)))
	}

	// Creating a 20 minute long event that SHOULD be active
	new_event.ID = "testid4"
	new_event.Name = "test2iseventactive"
	new_event.StartTime = TestTime
	new_event.EndTime = TestTime + ONE_MINUTE_IN_SECONDS*20

	service.CreateEvent(new_event.ID, new_event)

	is_active, err = service.IsEventActive("testid4")

	if err != nil {
		t.Fatal(err)
	}

	if !is_active {
		current_time := TestTime
		t.Errorf("Event was incorrectly deemed inactive. Current time: %v, event start time: %v, time difference: %v", current_time, new_event.StartTime, math.Abs((float64)((current_time - new_event.StartTime))))
	}

	CleanupTestDB(t)
}

/*
	Tests that getting event favorites works correctly at the service level
*/
func TestGetEventFavorites(t *testing.T) {
	SetupTestDB(t)

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

	CleanupTestDB(t)
}

/*
	Tests that adding event favorites works correctly at the service level
*/
func TestAddEventFavorite(t *testing.T) {
	SetupTestDB(t)

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

	CleanupTestDB(t)
}

/*
	Tests that removing event favorites works correctly at the service level
*/
func TestRemoveEventFavorite(t *testing.T) {
	SetupTestDB(t)

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

	CleanupTestDB(t)
}
