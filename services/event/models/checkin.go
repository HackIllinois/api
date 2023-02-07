package models

type StaffCheckinRequest struct {
	UserToken string `json:"userToken"`
	EventID   string `json:"eventId"`
}
type CheckinRequest struct {
	Code string `json:"code"`
}

type CheckinResponse struct {
	NewPoints   int                    `default:"-1" json:"newPoints"`
	TotalPoints int                    `default:"-1" json:"totalPoints"`
	Status      string                 `             json:"status"`
	RsvpData    map[string]interface{} `             json:"rsvpData,omitempty"`
}

type RedeemEventRequest struct {
	ID      string `json:"id"`
	EventID string `json:"eventID"`
}

type RedeemEventResponse struct {
	Status string `json:"status"`
}

type AwardPointsRequest struct {
	ID     string `json:"id"`
	Points int    `json:"points"`
}
