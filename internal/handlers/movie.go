package handlers

import (
	"net/http"

	"movies-api/internal/dto"
	"movies-api/internal/response"
	"movies-api/internal/service"
)

type MovieHandler struct {
	service *service.MovieService
}

func NewMovieHandler(service *service.MovieService) *MovieHandler {
	return &MovieHandler{service: service}
}

func (h *MovieHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	movies, err := h.service.GetAll()
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "failed to retrieve movies")
		return
	}

	res := make([]dto.MovieResponse, 0, len(movies))

	for _, movie := range movies {
		res = append(res, dto.MovieResponse{
			ID:          movie.ID,
			Title:       movie.Title,
			ReleaseYear: movie.Releaseyear,
			Duration:    movie.Duration,
		})
	}

	response.WriteJSON(w, http.StatusOK, res)
}
