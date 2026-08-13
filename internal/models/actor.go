package models

type Actor struct {
	ID        int		`validate:"required,gt=0"`
	Name      string	`validate:"required"`
	BirthDate string	`validate:"required,datetime=2006-01-02"`
}
