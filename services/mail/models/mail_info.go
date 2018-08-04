package models

type MailInfo struct {
	Recipients []Recipient `json:"recipients"`
	Content    Content     `json:"content"`
}

type Recipient struct {
	Address       Address       `json:"address"`
	Substitutions Substitutions `json:"substitution_data,omitempty"`
}

type Content struct {
	TemplateID string `json:"template_id"`
}

type Address struct {
	Email string `json:"email"`
	Name  string `json:"name,omitempty"`
}

type Substitutions map[string]interface{}
