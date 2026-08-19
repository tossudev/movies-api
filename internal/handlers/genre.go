package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"movies-api/internal/dto"
	"movies-api/internal/response"
	"movies-api/internal/service"
)

type GenreHandler struct {
	service *service.GenreService
}

func NewGenreHandler(service *service.GenreService) *GenreHandler {
	return &GenreHandler{service: service}
}

func (h *GenreHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	genres, err := h.service.GetAll()
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "failed to retrieve genres")
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
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "Malformed ID")
		return
	}
	genre, err := h.service.GetByID(id)
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "Genre does not exist")
		return
	}
	res := dto.GenreResponse{
		ID:   genre.ID,
		Name: genre.Name,
	}

	response.WriteJSON(w, http.StatusOK, res)
}

func (h *GenreHandler) Create(w http.ResponseWriter, r *http.Request) {
	req, err := h.decodeCreateRequest(r)
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "Malformed query")
		return
	}

	if err := h.service.Create(req); err != nil {
		response.WriteError(w, http.StatusInternalServerError, "Failed creating genre")
		return
	}

	response.WriteJSON(w, http.StatusOK, "Successfully added genre")
}

// Get ID from PathValue and rest from the query
// /api/{entity}/{id}
func (h *GenreHandler) Update(w http.ResponseWriter, r *http.Request) {
	req, err := h.decodeUpdateRequest(r)
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "Malformed query")
		return
	}
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "Malformed ID")
		return
	}

	if err := h.service.Update(id, req); err != nil {
		response.WriteError(w, http.StatusInternalServerError, "Failed creating genre")
		return
	}

	response.WriteJSON(w, http.StatusOK, "Successfully updated genre")
}

func (h *GenreHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "Malformed ID")
		return
	}

	if err := h.service.Delete(id); err != nil {
		response.WriteError(w, http.StatusInternalServerError, "Failed deleting genre")
		return
	}

	response.WriteJSON(w, http.StatusOK, "Successfull deleted genre")
}

func (h *GenreHandler) decodeUpdateRequest(r *http.Request) (dto.UpdateGenreRequest, error) {
	decoder := json.NewDecoder(r.Body)
	var req dto.UpdateGenreRequest
	err := decoder.Decode(&req)

	return req, err
}

func (h *GenreHandler) decodeCreateRequest(r *http.Request) (dto.CreateGenreRequest, error) {
	decoder := json.NewDecoder(r.Body)
	var req dto.CreateGenreRequest
	err := decoder.Decode(&req)

	return req, err
}
