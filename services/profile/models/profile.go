package models

type Profile struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Points    int    `json:"points"`
	Timezone  string `json:"timezone"`
	Discord   string `json:"discord"`
	AvatarUrl string `json:"avatarUrl"`
}
