package models

type Movie struct {
	ID          int
	Title       string
	Releaseyear int
	Duration    int
	Actors      []Actor
	Genres      []Genre
}
