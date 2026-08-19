package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"movies-api/internal/dto"
	"movies-api/internal/models"
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

func (h *ActorHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "Malformed ID")
		return
	}
	actor, err := h.service.GetByID(id)
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "Actor does not exist")
		return
	}
	res := dto.ActorResponse{
		ID:        actor.ID,
		Name:      actor.Name,
		BirthDate: actor.BirthDate,
	}

	response.WriteJSON(w, http.StatusOK, res)
}

func (h *ActorHandler) Create(w http.ResponseWriter, r *http.Request) {
	actor, err := h.decodeQuery(r)
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "Malformed query")
		return
	}

	if err := h.service.Create(actor); err != nil {
		response.WriteError(w, http.StatusInternalServerError, "Failed creating actor")
		return
	}

	response.WriteJSON(w, http.StatusOK, "Successfully added actor")
}

// Get ID from PathValue and rest from the query
// /api/{entity}/{id}
func (h *ActorHandler) Update(w http.ResponseWriter, r *http.Request) {
	actor, err := h.decodeQuery(r)
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "Malformed query")
		return
	}
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "Malformed ID")
		return
	}
	actor.ID = id

	if err := h.service.Update(actor); err != nil {
		response.WriteError(w, http.StatusInternalServerError, "Failed creating actor")
		return
	}

	response.WriteJSON(w, http.StatusOK, "Successfully updated actor")
}

func (h *ActorHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "Malformed ID")
		return
	}

	if err := h.service.Delete(id); err != nil {
		response.WriteError(w, http.StatusInternalServerError, "Failed deleting actor")
		return
	}

	response.WriteJSON(w, http.StatusOK, "Successfull deleted actor")
}

func (h *ActorHandler) decodeQuery(r *http.Request) (models.Actor, error) {
	decoder := json.NewDecoder(r.Body)
	var actor models.Actor
	err := decoder.Decode(&actor)

	return actor, err
}
