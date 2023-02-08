package models

type UserCheckin struct {
	UserToken       string                 `json:"userToken,omitempty" bson:"-" validate:"required"`
	ID              string                 `json:"id"                           validate:"omitempty"`
	Override        bool                   `json:"override"`
	HasCheckedIn    bool                   `json:"hasCheckedIn"`
	HasPickedUpSwag bool                   `json:"hasPickedUpSwag"`
	RsvpData        map[string]interface{} `json:"rsvpData"`
}
