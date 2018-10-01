package models

type UserRoles struct {
	ID    string `bson:"id"    json:"id"`
	Roles []Role `bson:"roles" json:"roles"`
}
