package handlers

import (
	"database/sql"
	"errors"
	"log/slog"
	"net/http"

	"movies-api/internal/dto"
	"movies-api/internal/models"
	"movies-api/internal/response"
	"movies-api/internal/service"

	"github.com/go-playground/validator/v10"
)

type GenreHandler struct {
	service   *service.GenreService
	validator *validator.Validate
}

func NewGenreHandler(service *service.GenreService, validator *validator.Validate) *GenreHandler {
	return &GenreHandler{service: service, validator: validator}
}

func (h *GenreHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	page, size, err := getPagination(r)
	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, dto.BadRequest(err.Error()))
		return
	}

	genres, err := h.service.GetAll(page, size)
	if err != nil {
		slog.ErrorContext(r.Context(), "failed to retrieve genres", "err", err)
		response.WriteJSON(w, http.StatusInternalServerError, dto.InternalServerError("failed to retrieve genres"))
		return
	}

	res := make([]dto.GenreResponse, 0, len(genres))

	for _, genre := range genres {
		res = append(res, toGenreResponse(genre))
	}

	response.WriteJSON(w, http.StatusOK, res)
}

func (h *GenreHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r)
	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, dto.BadRequest(err.Error()))
		return
	}

	genre, err := h.service.GetByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.WriteJSON(w, http.StatusNotFound, dto.NotFound("genre not found"))
		} else {
			slog.ErrorContext(r.Context(), "failed to retrieve genre", "err", err)
			response.WriteJSON(w, http.StatusInternalServerError, dto.InternalServerError("failed to retrieve genre"))
		}
		return
	}

	res := toGenreResponse(genre)
	response.WriteJSON(w, http.StatusOK, res)
}

func (h *GenreHandler) GetMovies(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r)
	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, dto.BadRequest(err.Error()))
		return
	}

	movies, err := h.service.GetMovies(id)

	if err != nil {
		slog.ErrorContext(r.Context(), "failed to retrieve movies", "err", err)
		response.WriteJSON(w, http.StatusInternalServerError, dto.InternalServerError("failed to retrieve movies"))
		return
	}

	res := make([]dto.MovieResponse, 0, len(movies))

	for _, movie := range movies {
		res = append(res, toMovieResponse(movie))
	}

	if len(movies) == 0 {
		response.WriteJSON(w, http.StatusNotFound, dto.NotFound("no movies found"))
		return
	}

	response.WriteJSON(w, http.StatusOK, res)
}

func (h *GenreHandler) Create(w http.ResponseWriter, r *http.Request) {
	req, err := decodeRequest[dto.CreateGenreRequest](r)
	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, dto.BadRequest("malformed json"))
		return
	}

	if err := h.validator.Struct(req); err != nil {
		response.WriteJSON(w, http.StatusBadRequest, dto.BadRequest("invalid request"))
		return
	}

	genre, err := h.service.Create(req)
	if err != nil {
		slog.ErrorContext(r.Context(), "failed to create genre", "err", err)
		response.WriteJSON(w, http.StatusInternalServerError, dto.InternalServerError("failed creating genre"))
		return
	}

	response.WriteJSON(w, http.StatusCreated, toGenreResponse(genre))
}

func (h *GenreHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r)
	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, dto.BadRequest("invalid id"))
		return
	}

	req, err := decodeRequest[dto.UpdateGenreRequest](r)
	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, dto.BadRequest("malformed json"))
		return
	}

	if req.Name == nil {
		response.WriteJSON(w, http.StatusBadRequest, dto.BadRequest("at least one field must be provided"))
		return
	}

	if err := h.validator.Struct(req); err != nil {
		response.WriteJSON(w, http.StatusBadRequest, dto.BadRequest("invalid request"))
		return
	}

	genre, err := h.service.Update(id, req)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.WriteJSON(w, http.StatusNotFound, dto.NotFound("genre not found"))
		} else {
			slog.ErrorContext(r.Context(), "failed to update genre", "err", err)
			response.WriteJSON(w, http.StatusInternalServerError, dto.InternalServerError("failed to update genre"))
		}
		return
	}

	response.WriteJSON(w, http.StatusOK, toGenreResponse(genre))
}

func (h *GenreHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r)
	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, dto.BadRequest("invalid id"))
		return
	}

	force := r.URL.Query().Get("force") == "true"
	if err := h.service.Delete(id, force); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows): // genre with id doesn't exist
			response.WriteJSON(w, http.StatusNotFound, dto.NotFound("genre not found"))
		case errors.Is(err, service.ErrAssociatedMoviesGenre): // genre has associated movies
			response.WriteJSON(w, http.StatusConflict, dto.Conflict("cannot delete genre because it has associated movies"))
		default: // unexpected error
			slog.ErrorContext(r.Context(), "failed to delete genre", "err", err)
			response.WriteJSON(w, http.StatusInternalServerError, dto.InternalServerError("failed to delete genre"))
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func toGenreResponse(model models.Genre) dto.GenreResponse {
	return dto.GenreResponse{
		ID:   model.ID,
		Name: model.Name,
	}
}
