package dto

import (
	"movies-api/internal/models"
)

type CreateMovieRequest struct {
	Title       string `json:"title" validate:"required"`
	ReleaseYear int    `json:"releaseYear" validate:"required,gt=0"`
	Duration    int    `json:"duration" validate:"required,gt=0"`
	GenreIDs    []int  `json:"genreIds"`
	ActorIDs    []int  `json:"actorIds"`
}

type UpdateMovieRequest struct {
	Title       *string `json:"title,omitempty"`
	ReleaseYear *int    `json:"releaseYear,omitempty"`
	Duration    *int    `json:"duration,omitempty"`
	GenreIDs    *[]int  `json:"genreIds,omitempty"`
	ActorIDs    *[]int  `json:"actorIds,omitempty"`
}

type MovieResponse struct {
	ID          int            `json:"id"`
	Title       string         `json:"title"`
	ReleaseYear int            `json:"releaseYear"`
	Duration    int            `json:"duration"`
	Genres      []models.Genre `json:"genres,omitempty"`
	Actors      []models.Actor `json:"actors,omitempty"`
}
