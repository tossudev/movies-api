package dto

type CreateActorRequest struct {
	Name      string `json:"name" validate:"required"`
	BirthDate string `json:"birthDate" validate:"required,validatebirthday"`
}

type UpdateActorRequest struct {
	Name      *string `json:"name,omitempty" validate:"omitempty,min=1"`
	BirthDate *string `json:"birthDate,omitempty" validate:"omitempty,validatebirthday"`
}

type ActorResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	BirthDate string `json:"birthDate"`
}
