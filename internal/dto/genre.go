package dto

type CreateGenreRequest struct {
	Name string `json:"name" validate:"required"`
}

type UpdateGenreRequest struct {
	Name *string `json:"name,omitempty" validate:"omitempty,min=1"`
}

type GenreResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
