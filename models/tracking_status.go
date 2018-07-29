package models

type TrackingStatus struct {
	EventTracker EventTracker `json:"eventTracker"`
	UserTracker  UserTracker  `json:"userTracker"`
}
