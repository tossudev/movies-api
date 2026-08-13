package handlers

import (
	"net/http"

	"movies-api/internal/dto"
	"movies-api/internal/response"
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
		response.WriteError(w, http.StatusInternalServerError, "failed to retrieve actors")
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

	response.WriteJSON(w, http.StatusOK, res)
}
