package service

import (
	"errors"
	"reflect"
	"time"

	common_config "github.com/HackIllinois/api/common/config"
	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/common/utils"
	"github.com/HackIllinois/api/services/event/config"
	"github.com/HackIllinois/api/services/event/models"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
)

var validate *validator.Validate

var db database.Database

func Initialize() error {
	if db != nil {
		db.Close()
		db = nil
	}

	var err error
	db, err = database.InitDatabase(config.EVENT_DB_HOST, config.EVENT_DB_NAME)

	if err != nil {
		return err
	}

	validate = validator.New()

	return nil
}

/*
	Returns the event with the given id
*/
func GetEvent[T models.Event](id string) (*T, error) {
	query := database.QuerySelector{
		"id": id,
	}

	switch reflect.TypeOf(*new(T)).Name() {
	case reflect.TypeOf(*new(models.EventPublic)).Name():
		query["isprivate"] = false
	}

	var event T
	err := db.FindOne("events", query, &event, nil)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

/*
	Deletes the event with the given id.
	Removes the event from event trackers and every user's tracker.
	Returns the event that was deleted.
*/
func DeleteEvent(id string) (*models.EventDB, error) {
	// Gets event to be able to return it later

	event, err := GetEvent[models.EventDB](id)
	if err != nil {
		return nil, err
	}

	query := database.QuerySelector{
		"id": id,
	}

	// Remove event from events database

	err = db.RemoveOne("events", query, nil)

	if err != nil {
		return nil, err
	}

	// Remove from event trackers database

	event_selector := database.QuerySelector{
		"eventid": id,
	}

	err = db.RemoveOne("eventtrackers", event_selector, nil)

	if err != nil {
		return nil, err
	}

	// Find all elements, and remove `id` from the Events slice
	// All the updates are individually atomic

	update_expression := bson.M{
		"$pull": bson.M{
			"events": id,
		},
	}

	_, err = db.UpdateAll("usertrackers", nil, &update_expression, nil)

	return event, err
}

/*
	Returns all the events
*/
func GetAllEvents[T models.Event]() (*models.EventList[T], error) {
	var query database.QuerySelector = nil

	switch reflect.TypeOf(*new(T)).Name() {
	case reflect.TypeOf(*new(models.EventPublic)).Name():
		query = database.QuerySelector{
			"isprivate": false,
		}
	}

	events := []T{}
	// nil implies there are no filters on the query, therefore everything in the "events" collection is returned.
	err := db.FindAll("events", query, &events, nil)
	if err != nil {
		return nil, err
	}

	event_list := models.EventList[T]{
		Events: events,
	}

	return &event_list, nil
}

/*
	Returns all the events
*/
func GetFilteredEvents[T models.Event](parameters map[string][]string) (*models.EventList[T], error) {
	query, err := database.CreateFilterQuery(parameters, *new(T))
	if err != nil {
		return nil, err
	}

	switch reflect.TypeOf(*new(T)).Name() {
	case reflect.TypeOf(*new(models.EventPublic)).Name():
		query["isprivate"] = false
	}

	events := []T{}
	filtered_events := models.EventList[T]{Events: events}
	err = db.FindAll("events", query, &filtered_events.Events, nil)

	if err != nil {
		return nil, err
	}

	return &filtered_events, nil
}

/*
	Creates an event with the given id
*/
func CreateEvent(id string, code string, event models.EventDB) error {
	err := validate.Struct(event)
	if err != nil {
		return err
	}

	_, err = GetEvent[models.EventDB](id)

	if err != database.ErrNotFound {
		if err != nil {
			return err
		}
		return errors.New("Event already exists")
	}

	err = db.Insert("events", &event, nil)

	if err != nil {
		return err
	}

	event_tracker := models.EventTracker{
		EventID: id,
		Users:   []string{},
	}

	err = db.Insert("eventtrackers", &event_tracker, nil)

	if err != nil {
		return err
	}

	event_code := models.EventCode{
		ID:         id,
		Code:       code,
		Expiration: event.EndTime,
	}

	err = db.Insert("eventcodes", &event_code, nil)

	return err
}

/*
	Updates the event with the given id
*/
func UpdateEvent(id string, event models.EventDB) error {
	err := validate.Struct(event)
	if err != nil {
		return err
	}

	selector := database.QuerySelector{
		"id": id,
	}

	err = db.Replace("events", selector, &event, false, nil)

	return err
}

/*
	Returns the event tracker for the specified event
*/
func GetEventTracker(event_id string) (*models.EventTracker, error) {
	query := database.QuerySelector{
		"eventid": event_id,
	}

	var tracker models.EventTracker
	err := db.FindOne("eventtrackers", query, &tracker, nil)
	if err != nil {
		return nil, err
	}

	return &tracker, nil
}

/*
	Returns the user tracker for the specified user
*/
func GetUserTracker(user_id string) (*models.UserTracker, error) {
	query := database.QuerySelector{
		"userid": user_id,
	}

	var tracker models.UserTracker
	err := db.FindOne("usertrackers", query, &tracker, nil)
	if err != nil {
		if err == database.ErrNotFound {
			return &models.UserTracker{
				UserID: user_id,
				Events: []string{},
			}, nil
		}
		return nil, err
	}

	return &tracker, nil
}

/*
	Returns true is the user has already been marked as attending
	the specified event, false otherwise
*/
func IsUserAttendingEvent(event_id string, user_id string) (bool, error) {
	tracker, err := GetEventTracker(event_id)
	if err != nil {
		return false, err
	}

	for _, id := range tracker.Users {
		if user_id == id {
			return true, nil
		}
	}

	return false, nil
}

/*
	Marks the specified user as attending the specified event
	The user must not already marked as attending for this to return successfully
*/
func MarkUserAsAttendingEvent(event_id string, user_id string) error {
	is_attending, err := IsUserAttendingEvent(event_id, user_id)
	if err != nil {
		return err
	}

	if is_attending {
		return errors.New("User has already been marked as attending")
	}

	if config.EVENT_CHECKIN_TIME_RESTRICTED {
		is_event_active, err := IsEventActive(event_id)
		if err != nil {
			return err
		}

		if !is_event_active {
			return errors.New("People cannot be checked-in for the event at this time.")
		}
	}

	event_selector := database.QuerySelector{
		"eventid": event_id,
	}

	event_modifier := bson.M{
		"$addToSet": bson.M{
			"users": user_id,
		},
	}

	_, err = db.Upsert("eventtrackers", event_selector, &event_modifier, nil)

	if err != nil {
		return err
	}

	user_selector := database.QuerySelector{
		"userid": user_id,
	}

	user_modifier := bson.M{
		"$addToSet": bson.M{
			"events": event_id,
		},
	}

	_, err = db.Upsert("usertrackers", user_selector, &user_modifier, nil)

	if err == database.ErrNotFound {
		user_tracker := models.UserTracker{
			UserID: user_id,
			Events: []string{event_id},
		}
		err = db.Insert("usertrackers", &user_tracker, nil)
	}

	return err
}

const (
	PreEventCheckinIntervalInMinutes = 15
	PreEventCheckinIntervalInSeconds = PreEventCheckinIntervalInMinutes * 60
)

/*
	Check if an event is active, i.e., that check-ins are allowed for the event at the current time.
	Returns true if the current time is between `PreEventCheckinIntervalInMinutes` number of minutes before the event, and the end of event.
*/
func IsEventActive(event_id string) (bool, error) {
	event, err := GetEvent[models.EventDB](event_id)
	if err != nil {
		return false, err
	}

	start_time := event.StartTime
	end_time := event.EndTime
	current_time := time.Now().Unix()

	if current_time < start_time {
		return start_time-current_time <= PreEventCheckinIntervalInSeconds, nil
	} else {
		return current_time < end_time, nil
	}
}

/*
	Returns the event favorites for the user with the given id
*/
func GetEventFavorites(id string) (*models.EventFavorites, error) {
	query := database.QuerySelector{
		"id": id,
	}

	var event_favorites models.EventFavorites
	err := db.FindOne("favorites", query, &event_favorites, nil)
	if err != nil {
		if err == database.ErrNotFound {
			err = db.Insert("favorites", &models.EventFavorites{
				ID:     id,
				Events: []string{},
			}, nil)

			if err != nil {
				return nil, err
			}

			err = db.FindOne("favorites", query, &event_favorites, nil)

			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return &event_favorites, nil
}

/*
	Adds the given event to the favorites for the user with the given id
*/
func AddEventFavorite(id string, event string) error {
	selector := database.QuerySelector{
		"id": id,
	}

	_, err := GetEvent[models.EventPublic](event)
	if err != nil {
		return errors.New("Could not find event with the given id.")
	}

	event_favorites, err := GetEventFavorites(id)
	if err != nil {
		return err
	}

	if !utils.ContainsString(event_favorites.Events, event) {
		event_favorites.Events = append(event_favorites.Events, event)
	}

	err = db.Replace("favorites", selector, event_favorites, false, nil)

	return err
}

/*
	Removes the given event from the favorites for the user with the given id
*/
func RemoveEventFavorite(id string, event string) error {
	selector := database.QuerySelector{
		"id": id,
	}

	event_favorites, err := GetEventFavorites(id)
	if err != nil {
		return err
	}

	event_favorites.Events, err = utils.RemoveString(event_favorites.Events, event)

	if err != nil {
		return errors.New("User's event favorites does not have specified event")
	}

	err = db.Replace("favorites", selector, event_favorites, false, nil)

	return err
}

/*
	Returns all event stats
*/
func GetStats() (map[string]interface{}, error) {
	query := database.QuerySelector{}

	var trackers []models.EventTracker
	err := db.FindAll("eventtrackers", query, &trackers, nil)
	if err != nil {
		return nil, err
	}

	stats := make(map[string]interface{})

	for _, tracker := range trackers {
		stats[tracker.EventID] = len(tracker.Users)
	}

	return stats, nil
}

/*
	Check if an event can be redeemed for points, i.e., that the point timeout has not been reached
	Returns true if the current time is between `PreEventCheckinIntervalInMinutes` number of minutes before the event, and the end of event.
*/
func CanRedeemPoints(event_code string) (bool, string, error) {
	query := database.QuerySelector{
		"code": event_code,
	}

	var eventCode models.EventCode
	err := db.FindOne("eventcodes", query, &eventCode, nil)
	if err != nil {
		return false, "invalid", err
	}

	expiration_time := eventCode.Expiration
	current_time := time.Now().Unix()

	return current_time < expiration_time, eventCode.ID, nil
}

/*
	Returns the eventcode struct for the event with the given id
*/
func GetEventCode(id string) (*models.EventCode, error) {
	query := database.QuerySelector{
		"id": id,
	}

	var eventCode models.EventCode
	err := db.FindOne("eventcodes", query, &eventCode, nil)
	if err != nil {
		return nil, err
	}

	return &eventCode, nil
}

/*
	Updates the event code and end time with the given id
*/
func UpdateEventCode(id string, eventCode models.EventCode) error {
	selector := database.QuerySelector{
		"id": id,
	}

	err := db.Replace("eventcodes", selector, &eventCode, false, nil)

	return err
}

/*
	Returns a CheckinResponse with NewPoints and TotalPoints defaulted to -1, and a status of status
*/
func NewCheckinResponseFailed(status string) *models.CheckinResponse {
	return &models.CheckinResponse{
		NewPoints:   -1,
		TotalPoints: -1,
		Status:      status,
	}
}

/*
	Attempts to checkin a user id to an event id by seeing if they can checkin,
	and then attempting to award them points for doing so
*/
func PerformCheckin(user_id string, event_id string) (*models.CheckinResponse, error) {
	redemption_status, err := RedeemEvent(user_id, event_id)

	if err != nil || redemption_status == nil {
		return nil, errors.New("Failed to verify if user already had redeemed event points")
	}

	if redemption_status.Status != "Success" {
		return NewCheckinResponseFailed("AlreadyCheckedIn"), nil
	}

	// Determine the current event and its point value
	event, err := GetEvent[models.EventDB](event_id)
	if err != nil {
		return nil, errors.New("Could not fetch the event specified")
	}

	// Add this point value to given profile
	profile, err := AwardPoints(user_id, event.Points)
	if err != nil {
		return nil, errors.New("Failed to award user with points")
	}

	return &models.CheckinResponse{
		Status:      "Success",
		NewPoints:   event.Points,
		TotalPoints: profile.Points,
	}, nil
}

/*
	Attempts to checkin a user to an event determined by a code
*/
func CheckinUserByCode(user_id string, code string) (*models.CheckinResponse, error) {
	// Check if we can redeem points for this given code still
	valid, event_id, err := CanRedeemPoints(code)

	// For this specific error, we know the issue was the code doesn't exist / is not valid
	if err == database.ErrNotFound {
		return NewCheckinResponseFailed("InvalidCode"), nil
	} else if err != nil {
		return nil, errors.New("Failed to receive event code information from database")
	}

	if !valid {
		return NewCheckinResponseFailed("ExpiredOrProspective"), nil
	}

	// We've gotten the user id and event id, now we need to Checkin
	return PerformCheckin(user_id, event_id)
}

/*
	Attempts to checkin a user determined by a JWT token to an event
*/
func CheckinUserTokenToEvent(user_token string, event_id string) (*models.CheckinResponse, error) {
	// Validate user_token, extract userId
	value_arr, err := utils.ExtractFieldFromJWT(common_config.TOKEN_SECRET, user_token, "userId")

	if len(value_arr) != 1 || err != nil {
		return NewCheckinResponseFailed("BadUserToken"), nil
	}

	user_id := value_arr[0]

	// Validate event exists
	_, err = GetEvent[models.EventDB](event_id)

	if err != nil {
		return NewCheckinResponseFailed("InvalidEventId"), nil
	}

	return PerformCheckin(user_id, event_id)
}
