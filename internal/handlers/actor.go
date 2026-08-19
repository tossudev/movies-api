package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

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
	req, err := h.decodeCreateRequest(r)
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "Malformed query")
		return
	}

	if err := h.service.Create(req); err != nil {
		response.WriteError(w, http.StatusInternalServerError, "Failed creating actor")
		return
	}

	response.WriteJSON(w, http.StatusOK, "Successfully added actor")
}

func (h *ActorHandler) Update(w http.ResponseWriter, r *http.Request) {
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
		response.WriteError(w, http.StatusInternalServerError, "Failed updating actor")
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

func (h *ActorHandler) decodeCreateRequest(r *http.Request) (dto.CreateActorRequest, error) {
	decoder := json.NewDecoder(r.Body)
	var req dto.CreateActorRequest
	err := decoder.Decode(&req)

	return req, err
}

func (h *ActorHandler) decodeUpdateRequest(r *http.Request) (dto.UpdateActorRequest, error) {
	decoder := json.NewDecoder(r.Body)
	var req dto.UpdateActorRequest
	err := decoder.Decode(&req)

	return req, err
}
