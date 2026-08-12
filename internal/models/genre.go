package models

type Genre struct {
	ID   int	`validate:"require,gt=0"`
	Name string	`validate:"require"`
}
