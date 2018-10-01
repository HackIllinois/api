package models

type UserRoles struct {
	ID    string `json:"id"`
	Roles []Role `json:"roles"`
}
