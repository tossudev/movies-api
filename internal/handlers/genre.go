package handlers

import (
	"encoding/json"
	"net/http"

	"movies-api/internal/dto"
	"movies-api/internal/service"
)

type GenreHandler struct {
	service *service.GenreService
}

func NewGenreHandler(service *service.GenreService) *GenreHandler {
	return &GenreHandler{service: service}
}

func (h *GenreHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	Genres, err := h.service.GetAll()
	if err != nil {
		http.Error(w, "failed to retrieve Genres", http.StatusInternalServerError)
		return
	}

	res := make([]dto.GenreResponse, 0, len(Genres))

	for _, Genre := range Genres {
		res = append(res, dto.GenreResponse{
			ID:   Genre.ID,
			Name: Genre.Name,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
