package dto

import "net/http"

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

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
