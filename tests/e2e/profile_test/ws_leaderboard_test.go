package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"testing"
	"time"

	api_errors "github.com/HackIllinois/api/common/errors"
	event_models "github.com/HackIllinois/api/services/event/models"
	profile_models "github.com/HackIllinois/api/services/profile/models"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	current_time              = time.Now()
	expected_full_leaderboard = []profile_models.LeaderboardEntry{
		{
			ID:      "testuser5",
			Points:  13337,
			Discord: "testuser#0005",
		},
		{
			ID:      "testuser19",
			Points:  2050,
			Discord: "testuser#0019",
		},
		{
			ID:      "testuser0",
			Points:  1337,
			Discord: "testuser#0000",
		},
		{
			ID:      "testuser13",
			Points:  1001,
			Discord: "testuser#0013",
		},
		{
			ID:      "testuser4",
			Points:  1000,
			Discord: "testuser#0004",
		},
		{
			ID:      "testuser8",
			Points:  990,
			Discord: "testuser#0008",
		},
		{
			ID:      "testuser6",
			Points:  500,
			Discord: "testuser#0006",
		},
		{
			ID:      "testuser7",
			Points:  469,
			Discord: "testuser#0007",
		},
		{
			ID:      "testuser1",
			Points:  420,
			Discord: "testuser#0001",
		},
		{
			ID:      "testuser12",
			Points:  401,
			Discord: "testuser#0012",
		},
	}
)

func CreateProfiles() {
	point_values := []int{1337, 420, 100, 0, 1000, 13337, 500, 469, 990, 320, 0, 400, 401, 1001, 42, 10, 0, 99, 120, 2050}
	for i := 0; i < len(point_values); i++ {
		id := fmt.Sprintf("testuser%d", i)
		profile := profile_models.Profile{
			ID:        id,
			FirstName: fmt.Sprintf("Test%d", i),
			LastName:  fmt.Sprintf("User%d", i),
			Points:    point_values[i],
			Discord:   fmt.Sprintf("testuser#%04d", i),
			AvatarUrl: "someimage.uri",
		}

		profile_attendance := profile_models.AttendanceTracker{
			ID:     id,
			Events: []string{},
		}

		profile_id_map := profile_models.IdMap{
			UserID:    id,
			ProfileID: id,
		}

		client.Database(profile_db_name).Collection("profiles").InsertOne(context.Background(), profile)
		client.Database(profile_db_name).Collection("profileattendance").InsertOne(context.Background(), profile_attendance)
		client.Database(profile_db_name).Collection("profileids").InsertOne(context.Background(), profile_id_map)
	}
}

func CreateEvents() {
	test_event := event_models.EventDB{
		EventPublic: event_models.EventPublic{
			ID:          "testevent12345",
			Name:        "An epic workshop",
			Description: "This is an epic workshop that you can learn from!",
			StartTime:   current_time.Unix(),
			EndTime:     current_time.Add(time.Hour).Unix(),
			Locations: []event_models.EventLocation{
				{
					Description: "Siebel Center for Computer Science Room 2404",
					Tags:        []string{"SEIBEL2"},
					Latitude:    40.1138038,
					Longitude:   -88.2254524,
				},
			},
			Sponsor: "",

			EventType: "WORKSHOP",
			Points:    200,
			IsAsync:   false,
		},
		IsPrivate:             false,
		DisplayOnStaffCheckin: false,
	}

	test_eventcode := event_models.EventCode{
		ID:         "testevent12345",
		Code:       "abc123",
		Expiration: current_time.Add(time.Hour).Unix(),
	}
	client.Database(event_db_name).Collection("events").InsertOne(context.Background(), test_event)
	client.Database(event_db_name).Collection("eventcodes").InsertOne(context.Background(), test_eventcode)
}

func ClearEvents() {
	client.Database(event_db_name).Collection("events").DeleteMany(context.Background(), bson.D{})
	client.Database(event_db_name).Collection("eventcodes").InsertOne(context.Background(), bson.D{})
}

func ClearProfiles() {
	client.Database(profile_db_name).Collection("profiles").DeleteMany(context.Background(), bson.D{})
	client.Database(profile_db_name).Collection("profileattendance").DeleteOne(context.Background(), bson.D{})
	client.Database(profile_db_name).Collection("profileids").DeleteOne(context.Background(), bson.D{})
}

func TestWSLeaderboardOneConnection(t *testing.T) {
	CreateProfiles()
	defer ClearProfiles()
	limit := 10
	u := url.URL{
		Scheme:   "ws",
		Host:     "localhost:8000",
		Path:     "/profile/live/leaderboard/",
		RawQuery: fmt.Sprintf("limit=%d", limit),
	}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		t.Fatal(err)
	}

	defer c.Close()

	msg_type, message, err := c.ReadMessage()
	if err != nil {
		t.Fatal(err)
	}

	if msg_type != websocket.TextMessage {
		t.Fatalf("Message recieved was not of type text message (got %v)", msg_type)
	}

	var leaderboard profile_models.LeaderboardEntryList

	err = json.Unmarshal(message, &leaderboard)
	if err != nil {
		t.Fatal(err)
	}

	expected_leaderboard := profile_models.LeaderboardEntryList{
		LeaderboardEntries: expected_full_leaderboard[:limit],
	}

	if !reflect.DeepEqual(expected_leaderboard, leaderboard) {
		t.Errorf("Wrong leaderboard. Expected %v, got %v", expected_leaderboard, leaderboard)
	}
}

func TestWSLeaderboardOnRedeem(t *testing.T) {
	CreateProfiles()
	CreateEvents()
	defer ClearProfiles()
	defer ClearEvents()
	limit := 10
	u := url.URL{
		Scheme:   "ws",
		Host:     "localhost:8000",
		Path:     "/profile/live/leaderboard/",
		RawQuery: fmt.Sprintf("limit=%d", limit),
	}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		t.Fatal(err)
	}

	defer c.Close()

	msg_type, message, err := c.ReadMessage()
	if err != nil {
		t.Fatal(err)
	}

	if msg_type != websocket.TextMessage {
		t.Fatalf("Message recieved was not of type text message (got %v)", msg_type)
	}

	var leaderboard profile_models.LeaderboardEntryList

	err = json.Unmarshal(message, &leaderboard)
	if err != nil {
		t.Fatal(err)
	}

	expected_leaderboard := profile_models.LeaderboardEntryList{
		LeaderboardEntries: expected_full_leaderboard[:limit],
	}

	if !reflect.DeepEqual(expected_leaderboard, leaderboard) {
		t.Errorf("Wrong leaderboard. Expected %v, got %v", expected_leaderboard, leaderboard)
	}

	// Attendee redeems event
	req := event_models.CheckinRequest{
		Code: "abc123",
	}
	id := "testuser11"
	received_res := event_models.CheckinResponse{}
	api_err := api_errors.ApiError{}
	response, err := admin_client.New().
		Post("/event/checkin/").
		Set("HackIllinois-Impersonation", id).
		BodyJSON(req).
		Receive(&received_res, &api_err)
	if err != nil {
		t.Fatal("Unable to make request")
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Request returned HTTP error %d, %v", response.StatusCode, api_err)
		return
	}

	expected_res := event_models.CheckinResponse{
		NewPoints:   200,
		TotalPoints: 600,
		Status:      "Success",
	}

	if !reflect.DeepEqual(received_res, expected_res) {
		t.Fatalf("Wrong result received. Expected %v, got %v", expected_res, received_res)
	}

	res := client.Database(profile_db_name).Collection("profiles").FindOne(context.Background(), bson.M{"id": id})

	profile := profile_models.Profile{}
	err = res.Decode(&profile)

	if err != nil {
		t.Fatalf("Had trouble finding profile in database: %v", err)
		return
	}

	if expected_res.TotalPoints != profile.Points {
		t.Fatalf("Wrong amount of points in profile database. Expected %v, got %v", expected_res.TotalPoints, profile.Points)
	}

	// Read from leaderboard socket (because the leaderboard should be updated)
	msg_type, message, err = c.ReadMessage()
	if err != nil {
		t.Fatal(err)
	}

	if msg_type != websocket.TextMessage {
		t.Fatalf("Message recieved was not of type text message (got %v)", msg_type)
	}

	err = json.Unmarshal(message, &leaderboard)
	if err != nil {
		t.Fatal(err)
	}

	modified_leaderboard := append(
		expected_full_leaderboard[:limit],
		profile_models.LeaderboardEntry{
			ID:      "testuser11",
			Points:  600,
			Discord: "testuser#0011",
		},
	)
	sort.Slice(modified_leaderboard[:], func(i, j int) bool {
		return modified_leaderboard[i].Points > modified_leaderboard[j].Points
	})
	expected_leaderboard = profile_models.LeaderboardEntryList{
		LeaderboardEntries: modified_leaderboard[:limit],
	}

	if !reflect.DeepEqual(expected_leaderboard, leaderboard) {
		t.Errorf("Wrong leaderboard. Expected %v, got %v", expected_leaderboard, leaderboard)
	}
}
