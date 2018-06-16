package models

type MailInfo struct {
	Recipients []Recipient `json:"recipients"`
	Content    Content     `json:"content"`
}

type Recipient struct {
	Address       string                 `json:"address"`
	Substitutions map[string]interface{} `json:"substitution_data,omitempty"`
}

type Content struct {
	TemplateID string `json:"template_id"`
}
