package models

type UserRoleModification struct {
	ID    string   `bson:"id"    json:"id"`
	Role  string   `bson:"role"  json:"role"`
}
