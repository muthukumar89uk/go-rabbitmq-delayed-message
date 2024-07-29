package models

type PostRequest struct {
	Platform string `json:"platform"`
	Message  string `json:"message"`
	Delay    string `json:"delay"` // delay in milliseconds
}
