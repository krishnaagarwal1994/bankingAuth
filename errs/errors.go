package errs

import (
	"net/http"
)

type AppError struct {
	Message string `json:"message"`
	Code    int    `json:"code,omitempty"`
}

func (err AppError) AsMessage() *AppError {
	return &AppError{Message: err.Message}
}

func NewAppError(message string) *AppError {
	return &AppError{Message: message}
}

func NewDBError(message string) *AppError {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}

func NewBadRequest(message string) *AppError {
	return &AppError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

func NewInternalServerError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}
