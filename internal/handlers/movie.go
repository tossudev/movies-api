package handlers

import (
	"net/http"

	"movies-api/internal/dto"
	"movies-api/internal/response"
	"movies-api/internal/service"

	"github.com/go-playground/validator/v10"
)

type MovieHandler struct {
	service   *service.MovieService
	validator *validator.Validate
}

func NewMovieHandler(service *service.MovieService, validator *validator.Validate) *MovieHandler {
	return &MovieHandler{service: service, validator: validator}
}

func (h *MovieHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	movies, err := h.service.GetAll()
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "Failed to retrieve movies")
		return
	}

	res := make([]dto.MovieResponse, 0, len(movies))

	for _, movie := range movies {
		res = append(res, dto.MovieResponse{
			ID:          movie.ID,
			Title:       movie.Title,
			ReleaseYear: movie.Releaseyear,
			Duration:    movie.Duration,
			Actors:      movie.Actors,
			Genres:      movie.Genres,
		})
	}

	response.WriteJSON(w, http.StatusOK, res)
}
