package handlers

import (
	"net/http"

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
