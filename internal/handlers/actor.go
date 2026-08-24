package handlers

import (
	"net/http"

	"movies-api/internal/dto"
	"movies-api/internal/models"
	"movies-api/internal/response"
	"movies-api/internal/service"

	"github.com/go-playground/validator/v10"
)

type ActorHandler struct {
	service   *service.ActorService
	validator *validator.Validate
}

func NewActorHandler(service *service.ActorService, validator *validator.Validate) *ActorHandler {
	return &ActorHandler{service: service, validator: validator}
}

func (h *ActorHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	actors, err := h.service.GetAll()
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "Failed to retrieve actors")
		return
	}

	res := make([]dto.ActorResponse, 0, len(actors))

	for _, actor := range actors {
		res = append(res, toActorResponse(actor))
	}

	response.WriteJSON(w, http.StatusOK, res)
}

func (h *ActorHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	actor, err := h.service.GetByID(id)
	if err != nil {
		response.WriteError(w, http.StatusNotFound, "Actor does not exist")
		return
	}

	response.WriteJSON(w, http.StatusOK, toActorResponse(actor))
}

func (h *ActorHandler) Create(w http.ResponseWriter, r *http.Request) {
	req, err := decodeRequest[dto.CreateActorRequest](r)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, "Malformed JSON")
		return
	}

	if err := h.validator.Struct(req); err != nil {
		response.WriteError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	actor, err := h.service.Create(req)
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "Failed creating actor")
		return
	}

	response.WriteJSON(w, http.StatusCreated, toActorResponse(actor))
}

func (h *ActorHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	req, err := decodeRequest[dto.UpdateActorRequest](r)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, "Malformed JSON")
		return
	}

	if err := h.validator.Struct(req); err != nil {
		response.WriteError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	if err := h.service.Update(id, req); err != nil {
		response.WriteError(w, http.StatusInternalServerError, "Failed updating actor")
		return
	}

	response.WriteJSON(w, http.StatusOK, "Successfully updated actor")
}

func (h *ActorHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := h.service.Delete(id); err != nil {
		response.WriteError(w, http.StatusNotFound, "Failed deleting actor")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func toActorResponse(model models.Actor) dto.ActorResponse {
	return dto.ActorResponse{
		ID:        model.ID,
		Name:      model.Name,
		BirthDate: model.BirthDate,
	}
}
