package dto

import (
	"movies-api/internal/models"
)

type CreateMovieRequest struct {
	Title       string `json:"title" validate:"required"`
	ReleaseYear int    `json:"releaseYear" validate:"required,gt=0"`
	Duration    int    `json:"duration" validate:"required,gt=0"`
	GenreIDs    []int  `json:"genreIds,omitempty" validate:"unique,dive,gt=0"`
	ActorIDs    []int  `json:"actorIds,omitempty" validate:"unique,dive,gt=0"`
}

type UpdateMovieRequest struct {
	Title       *string `json:"title,omitempty" validate:"omitempty,min=1"`
	ReleaseYear *int    `json:"releaseYear,omitempty" validate:"omitempty,gt=0"`
	Duration    *int    `json:"duration,omitempty" validate:"omitempty,gt=0"`
	GenreIDs    []int   `json:"genreIds,omitempty" validate:"unique,dive,gt=0"`
	ActorIDs    []int   `json:"actorIds,omitempty" validate:"unique,dive,gt=0"`
}

type MovieResponse struct {
	ID          int            `json:"id"`
	Title       string         `json:"title"`
	ReleaseYear int            `json:"releaseYear"`
	Duration    int            `json:"duration"`
	Genres      []models.Genre `json:"genres,omitempty"`
	Actors      []models.Actor `json:"actors,omitempty"`
}
