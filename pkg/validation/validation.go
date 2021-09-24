package validation

import "github.com/go-playground/validator/v10"

// ValidateStruct will validate request
func ValidateStruct(i interface{}) error {
	validate := validator.New()
	RegisterCustomValidator(validate)
	return validate.Struct(i)
}

func RegisterCustomValidator(validator *validator.Validate) {

}
