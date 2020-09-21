package models

type UserCheckin struct {
	ID              string                 `json:"id"`
	Override        bool                   `json:"override"`
	HasCheckedIn    bool                   `json:"hasCheckedIn"`
	RsvpData        map[string]interface{} `json:"rsvpData"`
}
