package utils

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"strconv"
	"strings"
)

func ValidateInput(dataSet any) (bool, map[string]string) {
	var validate = validator.New()

	err := validate.Struct(dataSet)

	if err != nil {
		if err, ok := err.(*validator.InvalidValidationError); ok {
			panic(err)
		}

		errors := make(map[string]string)
		reflectedDataset := reflect.ValueOf(dataSet)
		var countError = 0
		for _, err := range err.(validator.ValidationErrors) {
			field, _ := reflectedDataset.Type().FieldByName(err.StructField())
			var name string
			if name = field.Tag.Get("json"); name == "" {
				name = strings.ToLower(err.StructField())
			}
			var msg string
			switch err.Tag() {
			case "required":
				msg = "The " + name + " field is required"
			case "email":
				msg = "The " + name + " field must be a valid email address"
			case "gte":
				msg = "The " + name + " field must be greater than or equal to " + err.Param()
			case "lte":
				msg = "The " + name + " field must be less than or equal to " + err.Param()
			case "eqfield":
				msg = "The " + name + " field must be equal to " + err.Param()
			default:
				msg = "The " + name + " field is invalid"
			}
			var idxStr = strconv.Itoa(countError)
			errors[idxStr] = msg
			countError++
		}
		return false, errors
	}
	return true, nil
}
