package models

type UserCheckin struct {
	ID              string `json:"id"`
	HasCheckedIn    bool   `json:"hasCheckedIn"`
	HasPickedUpSwag bool   `json:"hasPickedUpSwag"`
}
