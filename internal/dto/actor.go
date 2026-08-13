package dto

type CreateActorRequest struct {
	Name      string `json:"name" validate:"required"`
	BirthDate string `json:"birthDate" validate:"required,datetime=2006-01-02"`
}

type UpdateActorRequest struct {
	Name      *string `json:"name,omitempty"`
	BirthDate *string `json:"birthDate,omitempty"`
}

type ActorResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	BirthDate string `json:"birthDate"`
}
