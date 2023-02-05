package models

type Profile struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName"  validate:"required"`
	Points    int    `json:"points"`
	Discord   string `json:"discord"   validate:"required"`
	AvatarUrl string `json:"avatarUrl" validate:"required"`
	FoodWave  int    `json:"foodWave"`
}
