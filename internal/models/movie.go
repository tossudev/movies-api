package models

type Movie struct {
	ID          int
	Title       string
	ReleaseYear int
	Duration    int
	Actors      []Actor
	Genres      []Genre
}
