package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

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
	page, size, err := getPagination(r)
	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, dto.BadRequest(err.Error()))
		return
	}

	filters := make(map[string]string)

	for key, values := range r.URL.Query() {
		value := strings.TrimSpace(values[0])
		switch key {
		case "name", "birth_date":
			filters[key] = value
		case "page", "size":
			continue
		default:
			response.WriteJSON(w, http.StatusBadRequest, fmt.Sprintf("Invalid filter: %s", key))
			return
		}
	}

	actors, err := h.service.GetAll(page, size, filters)

	if err != nil {
		slog.ErrorContext(r.Context(), "failed to retrieve actors", "err", err)
		response.WriteJSON(w, http.StatusInternalServerError, dto.InternalServerError("failed to retrieve actors"))
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
		response.WriteJSON(w, http.StatusBadRequest, dto.BadRequest(err.Error()))
		return
	}

	actor, err := h.service.GetByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.WriteJSON(w, http.StatusNotFound, dto.NotFound("actor not found"))
		} else {
			slog.ErrorContext(r.Context(), "failed to retrieve actor", "err", err)
			response.WriteJSON(w, http.StatusInternalServerError, dto.InternalServerError("failed to retrieve actor"))
		}
		return
	}

	response.WriteJSON(w, http.StatusOK, toActorResponse(actor))
}

func (h *ActorHandler) Create(w http.ResponseWriter, r *http.Request) {
	req, err := decodeRequest[dto.CreateActorRequest](r)
	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, dto.BadRequest("malformed json"))
		return
	}

	if err := h.validator.Struct(req); err != nil {
		response.WriteJSON(w, http.StatusBadRequest, dto.BadRequest("invalid request"))
		return
	}

	actor, err := h.service.Create(req)
	if err != nil {
		slog.ErrorContext(r.Context(), "failed to create actor", "err", err)
		response.WriteJSON(w, http.StatusInternalServerError, dto.InternalServerError("failed creating actor"))
		return
	}

	response.WriteJSON(w, http.StatusCreated, toActorResponse(actor))
}

func (h *ActorHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r)
	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, dto.BadRequest("invalid id"))
		return
	}

	req, err := decodeRequest[dto.UpdateActorRequest](r)
	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, dto.BadRequest("malformed json"))
		return
	}

	if req.Name == nil && req.BirthDate == nil {
		response.WriteJSON(w, http.StatusBadRequest, dto.BadRequest("at least one field must be provided"))
		return
	}

	if err := h.validator.Struct(req); err != nil {
		response.WriteJSON(w, http.StatusBadRequest, dto.BadRequest("invalid request"))
		return
	}

	if err := h.service.Update(id, req); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.WriteJSON(w, http.StatusNotFound, dto.NotFound("actor not found"))
		} else {
			slog.ErrorContext(r.Context(), "failed to update actor", "err", err)
			response.WriteJSON(w, http.StatusInternalServerError, dto.InternalServerError("failed to update actor"))
		}
		return
	}

	response.WriteJSON(w, http.StatusOK, "successfully updated actor")
}

func (h *ActorHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r)
	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, dto.BadRequest("invalid id"))
		return
	}

	force := r.URL.Query().Get("force") == "true"
	if err := h.service.Delete(id, force); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows): // actor with id doesn't exist
			response.WriteJSON(w, http.StatusNotFound, dto.NotFound("actor not found"))
		case errors.Is(err, service.ErrAssociatedMoviesActor): // actor has associated movies
			response.WriteJSON(w, http.StatusConflict, dto.Conflict("cannot delete actor because it has associated movies"))
		default: // unexpected error
			slog.ErrorContext(r.Context(), "failed to delete actor", "err", err)
			response.WriteJSON(w, http.StatusInternalServerError, dto.InternalServerError("failed to delete actor"))
		}

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
