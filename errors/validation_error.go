package errors

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidationErrors(err error) (message string) {
	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for idx, err := range castedObject {
			// field := ToSnakeCase(err.Field())
			field := err.Field()

			// Change Field Name
			switch field {
			}

			// Check Error Type
			switch err.Tag() {
			case "required":
				message = fmt.Sprintf("%s is mandatory",
					field)
			case "email":
				message = fmt.Sprintf("%s must be valid email address",
					field)
			case "latitude":
				message = fmt.Sprintf("%s value must be a valid latitude",
					field)
			case "longitude":
				message = fmt.Sprintf("%s value must be a valid longitude",
					field)
			case "gte":
				message = fmt.Sprintf("%s value must be greater than or equal to %s",
					field, err.Param())
			case "shipment_type":
				message = fmt.Sprintf("%s value must be a valid shipment type",
					field)
			case "required_if":
				message = fmt.Sprintf("%s is mandatory", field)
			case "lte":
				message = fmt.Sprintf("%s value must be lower than %s",
					field, err.Param())
			case "cod_price":
				message = fmt.Sprintf("%s value must be greater or equal than shipment_price + insurance_price", field)
			case "pickup_time":
				message = fmt.Sprintf("%s value must be greater than now", field)
			case "pickup_time_length":
				message = fmt.Sprintf("%s value must be greater than 10 number", field)
			case "startswith":
				message = fmt.Sprintf("%s must starts with %s",
					field, err.Param())
			default:
				message = err.Error()
			}

			if idx == 0 {
				break
			}

		}
	}
	return
}
