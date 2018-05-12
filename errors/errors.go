package errors

type APIError struct {
    Status     int      `json:"status,omitempty"`
    Title      string   `json:"title,omitempty"`
    Message    string   `json:"message,omitempty"`
}
