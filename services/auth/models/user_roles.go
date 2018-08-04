package models

type UserRoles struct {
	ID    string   `bson:"id"    json:"id"`
	Roles []string `bson:"roles" json:"roles"`
}
