package models

type Movie struct {
	ID          int		`validate:"required,gt=0"`
	Title       string	`validate:"required"`
	Releaseyear int		`validate:"required,gt=0"`
	Duration    int		`validate:"required,gt=0"`
}
