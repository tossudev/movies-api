package handlers

import (
	"net/http"

	"movies-api/internal/dto"
	"movies-api/internal/models"
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
		res = append(res, toMovieResponse(movie))
	}

	response.WriteJSON(w, http.StatusOK, res)
}

func (h *MovieHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	movie, err := h.service.GetByID(id)
	if err != nil {
		response.WriteError(w, http.StatusNotFound, "Movie does not exist")
		return
	}

	response.WriteJSON(w, http.StatusOK, toMovieResponse(movie))
}

func (h *MovieHandler) Create(w http.ResponseWriter, r *http.Request) {
	req, err := decodeRequest[dto.CreateMovieRequest](r)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, "Malformed JSON")
		return
	}

	if err := h.validator.Struct(req); err != nil {
		response.WriteError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	movie, err := h.service.Create(req)
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "Failed creating movie")
		return
	}

	response.WriteJSON(w, http.StatusCreated, toMovieResponse(movie))
}

func toMovieResponse(model models.Movie) dto.MovieResponse {
	return dto.MovieResponse{
		ID:          model.ID,
		Title:       model.Title,
		ReleaseYear: model.Releaseyear,
		Duration:    model.Duration,
		Actors:      model.Actors,
		Genres:      model.Genres,
	}
}
