package service

import (
	"errors"
	"net/http"

	"github.com/HackIllinois/api/common/apirequest"
	"github.com/HackIllinois/api/services/event/config"
	"github.com/HackIllinois/api/services/profile/models"
)

/*
	Checks if a user can redeem a certain event by seeing if they've already checked in, marking it redeemed if they haven't already
*/
func RedeemEvent(user_id string, event_id string) (*models.RedeemEventResponse, error) {
	var redemption_status models.RedeemEventResponse
	event_info := models.RedeemEventRequest{
		ID:      user_id,
		EventID: event_id,
	}

	status, err := apirequest.Post(config.PROFILE_SERVICE+"/profile/event/checkin/", &event_info, &redemption_status)

	if err != nil {
		return nil, err
	}

	if status != http.StatusOK {
		return nil, errors.New("Unable to check event redemption")
	}

	return &redemption_status, nil
}

/*
	Performs a get and a put operation on the profile to increment the current number of points
*/
func AwardPoints(id string, points int) (*models.Profile, error) {
	var profile models.Profile
	point_info := models.AwardPointsRequest{
		ID:     id,
		Points: points,
	}
	status, err := apirequest.Post(config.PROFILE_SERVICE+"/profile/points/award/", point_info, &profile)

	if err != nil {
		return nil, err
	}

	if status != http.StatusOK {
		return nil, errors.New("Unable to award points")
	}

	return &profile, nil
}
