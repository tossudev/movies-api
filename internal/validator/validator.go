package validator

import (
	"database/sql"

	"movies-api/internal/repository"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate = validator.New()


func ValidateStruct(entity any) error {
	return validate.Struct(entity)
}

func ActorExists(id int) bool {
	_, err := repository.Actor.GetByID(id)
	return err != sql.ErrNoRows
}

func MovieExists(id int) bool {
	_, err := repository.Movie.GetByID(id)
	return err != sql.ErrNoRows
}

func GenreExists(id int) bool {
	_, err := repository.Genre.GetByID(id)
	return err != sql.ErrNoRows
}
