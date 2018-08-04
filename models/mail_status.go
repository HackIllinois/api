package models

type MailStatus struct {
	Results MailStatusResults `json:"results"`
}

type MailStatusResults struct {
	Rejected int `json:"total_rejected_recipients"`
	Accepted int `json:"total_accepted_recipients"`
}
