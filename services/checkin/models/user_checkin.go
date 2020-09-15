package models

type UserCheckin struct {
	ID              string                 `json:"id"`
	Override        bool                   `json:"override"`
	HasCheckedIn    bool                   `json:"hasCheckedIn"`
	HasPickedUpSwag bool                   `json:"hasPickedUpSwag"`
	RsvpData        map[string]interface{} `json:"rsvpData"`
}
