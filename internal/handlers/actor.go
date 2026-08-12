package handlers

import (
	"fmt"
	"net/http"

	"movies-api/internal/service"
)

type ActorHandler struct {
	service *service.ActorService
}

func NewActorHandler(service *service.ActorService) *ActorHandler {
	return &ActorHandler{service: service}
}

func (h *ActorHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "joonas suokko")
}
