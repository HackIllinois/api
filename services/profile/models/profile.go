package models

type Profile struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Github    string   `json:"github"`
	Linkedin  string   `json:"linkedin"`
	Interests []string `json:"interests"`
}
