package utils

import "net/http"

type AppError struct {
	Status  int
	Message string
}

func (a *AppError) Error() string {
	return a.Message
}

func BadRequest(msg string) *AppError {
	appErr := AppError{
		Status:  http.StatusBadRequest,
		Message: msg,
	}
	return &appErr
}

func Unauthorized(msg string) *AppError {
	return &AppError{Status: http.StatusUnauthorized, Message: msg}
}

func NotFound(msg string) *AppError {
	return &AppError{Status: http.StatusNotFound, Message: msg}
}

func InternalError(msg string) *AppError {
	return &AppError{Status: http.StatusInternalServerError, Message: msg}
}
