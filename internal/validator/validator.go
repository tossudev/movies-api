package validator

import (
	"time"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func Init() {
	Validate = validator.New()
	Validate.RegisterValidation("validatebirthday", validateBirthday)
}

func ValidateStruct(entity any) error {
	return Validate.Struct(entity)
}

func validateBirthday(fl validator.FieldLevel) bool {
	t, err := time.Parse("2006-01-02", fl.Field().String())
	return err == nil && t.Before(time.Now())
}
