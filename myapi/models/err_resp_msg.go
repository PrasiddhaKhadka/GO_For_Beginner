package models

type ErrorResponseMessage struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
