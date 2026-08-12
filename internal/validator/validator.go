package validator

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate = validator.New()


func ValidateStruct(entity any) error {
	return validate.Struct(entity)
}
