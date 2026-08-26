package handlers

import (
	"database/sql"
	"errors"
	"log/slog"
	"net/http"
	"strings"

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
	page, size, err := getPagination(r)
	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, dto.BadRequest(err.Error()))
		return
	}

	movies, err := h.service.GetAll(page, size)
	if err != nil {
		slog.ErrorContext(r.Context(), "failed to retrieve movies", "err", err)
		response.WriteJSON(w, http.StatusInternalServerError, dto.InternalServerError("failed to retrieve movies"))
		return
	}

	res := make([]dto.MovieResponse, 0, len(movies))

	for _, movie := range movies {
		res = append(res, toMovieResponse(movie))
	}

	response.WriteJSON(w, http.StatusOK, res)
}

func (h *MovieHandler) Search(w http.ResponseWriter, r *http.Request) {
	page, size, err := getPagination(r)
	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, dto.BadRequest(err.Error()))
		return
	}
	query := strings.TrimSpace(r.URL.Query().Get("title"))
	if query == "" {
		response.WriteJSON(w, http.StatusBadRequest, dto.BadRequest("search title must be provided"))
		return
	}

	movies, err := h.service.Search(query, page, size)
	if err != nil {
		slog.ErrorContext(r.Context(), "failed to search movies", "err", err)
		response.WriteJSON(w, http.StatusInternalServerError, dto.InternalServerError("failed to search movies"))
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
		response.WriteJSON(w, http.StatusBadRequest, dto.BadRequest(err.Error()))
		return
	}

	movie, err := h.service.GetByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.WriteJSON(w, http.StatusNotFound, dto.NotFound("movie not found"))
		} else {
			slog.ErrorContext(r.Context(), "failed to retrieve movie", "err", err)
			response.WriteJSON(w, http.StatusInternalServerError, dto.InternalServerError("failed to retrieve movie"))
		}
		return
	}

	response.WriteJSON(w, http.StatusOK, toMovieResponse(movie))
}

func (h *MovieHandler) Create(w http.ResponseWriter, r *http.Request) {
	req, err := decodeRequest[dto.CreateMovieRequest](r)
	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, dto.BadRequest("malformed json"))
		return
	}

	if err := h.validator.Struct(req); err != nil {
		response.WriteJSON(w, http.StatusBadRequest, dto.BadRequest("invalid request"))
		return
	}

	movie, err := h.service.Create(req)
	if err != nil {
		slog.ErrorContext(r.Context(), "failed to create movie", "err", err)
		response.WriteJSON(w, http.StatusInternalServerError, dto.InternalServerError("failed creating movie"))
		return
	}

	response.WriteJSON(w, http.StatusCreated, toMovieResponse(movie))
}

func toMovieResponse(model models.Movie) dto.MovieResponse {
	return dto.MovieResponse{
		ID:          model.ID,
		Title:       model.Title,
		ReleaseYear: model.ReleaseYear,
		Duration:    model.Duration,
		Actors:      model.Actors,
		Genres:      model.Genres,
	}
}

func (h *MovieHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r)
	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, dto.BadRequest("invalid id"))
		return
	}

	req, err := decodeRequest[dto.UpdateMovieRequest](r)
	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, dto.BadRequest("malformed json"))
		return
	}

	if req.Title == nil && req.ReleaseYear == nil && req.Duration == nil && req.GenreIDs == nil && req.ActorIDs == nil {
		response.WriteJSON(w, http.StatusBadRequest, dto.BadRequest("at least one field must be provided"))
		return
	}

	if err := h.validator.Struct(req); err != nil {
		response.WriteJSON(w, http.StatusBadRequest, dto.BadRequest("invalid request"))
		return
	}

	if err := h.service.Update(id, req); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.WriteJSON(w, http.StatusNotFound, dto.NotFound("movie not found"))
		} else {
			slog.ErrorContext(r.Context(), "failed to update movie", "err", err)
			response.WriteJSON(w, http.StatusInternalServerError, dto.InternalServerError("failed to update movie"))
		}
		return
	}

	response.WriteJSON(w, http.StatusOK, "successfully updated movie")
}

func (h *MovieHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r)
	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, dto.BadRequest("invalid id"))
		return
	}

	force := r.URL.Query().Get("force") == "true"
	if err := h.service.Delete(id, force); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.WriteJSON(w, http.StatusNotFound, dto.NotFound("movie not found"))
		} else {
			slog.ErrorContext(r.Context(), "failed to delete movie", "err", err)
			response.WriteJSON(w, http.StatusInternalServerError, dto.InternalServerError("failed to delete movie"))
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
