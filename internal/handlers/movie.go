package handlers

import (
	"encoding/json"
	"net/http"

	"movies-api/internal/dto"
	"movies-api/internal/service"
)

type MovieHandler struct {
	service *service.MovieService
}

func NewMovieHandler(service *service.MovieService) *MovieHandler {
	return &MovieHandler{service: service}
}

func (h *MovieHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	Movies, err := h.service.GetAll()
	if err != nil {
		http.Error(w, "failed to retrieve Movies", http.StatusInternalServerError)
		return
	}

	res := make([]dto.MovieResponse, 0, len(Movies))

	for _, movie := range Movies {
		res = append(res, dto.MovieResponse{
			ID:          movie.ID,
			Title:       movie.Title,
			ReleaseYear: movie.Releaseyear,
			Duration:    movie.Duration,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
