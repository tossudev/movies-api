package handlers

import (
	"net/http"

	"movies-api/internal/dto"
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
		response.WriteJSON(w, err.Code, err)
		return
	}

	genres, err := h.service.GetAll(page, size)
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "Failed to retrieve genres")
		return
	}

	res := make([]dto.GenreResponse, 0, len(genres))

	for _, genre := range genres {
		res = append(res, dto.GenreResponse{
			ID:   genre.ID,
			Name: genre.Name,
		})
	}

	response.WriteJSON(w, http.StatusOK, res)
}

func (h *GenreHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	genre, err := h.service.GetByID(id)
	if err != nil {
		response.WriteError(w, http.StatusNotFound, "Genre does not exist")
		return
	}

	res := dto.GenreResponse{
		ID:   genre.ID,
		Name: genre.Name,
	}

	response.WriteJSON(w, http.StatusOK, res)
}

func (h *GenreHandler) Create(w http.ResponseWriter, r *http.Request) {
	req, err := decodeRequest[dto.CreateGenreRequest](r)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, "Malformed JSON")
		return
	}

	if err := h.validator.Struct(req); err != nil {
		response.WriteError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	genre, err := h.service.Create(req)
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "Failed creating genre")
		return
	}

	response.WriteJSON(w, http.StatusCreated, dto.GenreResponse{
		ID:   genre.ID,
		Name: genre.Name,
	})
}

func (h *GenreHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	req, err := decodeRequest[dto.UpdateGenreRequest](r)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, "Malformed JSON")
		return
	}

	if err := h.validator.Struct(req); err != nil {
		response.WriteError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	if err := h.service.Update(id, req); err != nil {
		response.WriteError(w, http.StatusInternalServerError, "Failed updating genre")
		return
	}

	response.WriteJSON(w, http.StatusOK, "Successfully updated genre")
}

func (h *GenreHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := h.service.Delete(id); err != nil {
		response.WriteError(w, http.StatusNotFound, "Failed deleting genre")
		return
	}

	response.WriteJSON(w, http.StatusOK, "Successfully deleted genre")
}
