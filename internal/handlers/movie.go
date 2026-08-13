package handlers

import (
	"fmt"
	"net/http"

	"movies-api/internal/service"
)

type MovieHandler struct {
	service *service.MovieService
}

func NewMovieHandler(service *service.MovieService) *MovieHandler {
	return &MovieHandler{service: service}
}

func (h *MovieHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "out of the ashes")
}
