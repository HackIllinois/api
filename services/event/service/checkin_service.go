package service

import (
	"errors"
	"net/http"

	"github.com/HackIllinois/api/common/apirequest"
	"github.com/HackIllinois/api/common/database"
	"github.com/HackIllinois/api/common/utils"
	gateway_config "github.com/HackIllinois/api/gateway/config"
	"github.com/HackIllinois/api/services/event/config"
	"github.com/HackIllinois/api/services/event/models"
)

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
	Checks if the user has been checked in with the checkin service
*/
func IsUserCheckedIn(id string) (bool, error) {
	status, err := apirequest.Get(config.CHECKIN_SERVICE+"/checkin/"+id+"/", nil)

	if err != nil {
		return false, err
	}

	return status == http.StatusOK, nil
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
	event, err := GetEvent(event_id)

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
	user_id, err := utils.ExtractFieldFromJWT(gateway_config.TOKEN_SECRET, user_token, "userId")

	if err != nil {
		return NewCheckinResponseFailed("ExpiredOrProspective"), nil
	}

	// Note: the event id will be validated in PerformCheckin
	return PerformCheckin(user_id[0], event_id)
}
