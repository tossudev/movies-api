package dto

import (
	"errors"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

var ActorNotExist = errors.New("actor does not exist")
var GenreNotExist = errors.New("genre does not exist")

func BadRequest(message string) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		Code:    http.StatusBadRequest,
	}
}

func InternalServerError(message string) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

func NotFound(message string) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func Conflict(message string) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		Code:    http.StatusConflict,
	}
}
