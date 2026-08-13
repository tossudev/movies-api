package handlers

import (
	"encoding/json"
	"net/http"

	"movies-api/internal/dto"
	"movies-api/internal/service"
)

type ActorHandler struct {
	service *service.ActorService
}

func NewActorHandler(service *service.ActorService) *ActorHandler {
	return &ActorHandler{service: service}
}

func (h *ActorHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	actors, err := h.service.GetAll()
	if err != nil {
		http.Error(w, "failed to retrieve actors", http.StatusInternalServerError)
		return
	}

	res := make([]dto.ActorResponse, 0, len(actors))

	for _, actor := range actors {
		res = append(res, dto.ActorResponse{
			ID:        actor.ID,
			Name:      actor.Name,
			BirthDate: actor.BirthDate,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
