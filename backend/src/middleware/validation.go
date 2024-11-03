package middleware

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var validate *validator.Validate

func init() {
	validate = validator.New()

	validate.RegisterValidation("unique", validateUnique)
}

func validateUnique(fl validator.FieldLevel) bool {
	field := fl.Field()
	if field.Kind() != reflect.Slice {
		return false
	}

	seen := make(map[interface{}]bool)
	for i := 0; i < field.Len(); i++ {
		val := field.Index(i).Interface()
		if seen[val] {
			return false
		}
		seen[val] = true
	}
	return true
}

type ValidationError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

func Validate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := c.Validate(c.Request().Body); err != nil {
			var errors []ValidationError
			for _, err := range err.(validator.ValidationErrors) {
				errors = append(errors, ValidationError{
					Field: strings.ToLower(err.Field()),
					Error: getErrorMsg(err),
				})
			}
			return c.JSON(http.StatusBadRequest, errors)
		}
		return next(c)
	}
}

func getErrorMsg(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return "Should be at least " + err.Param() + " characters"
	case "max":
		return "Should be at most " + err.Param() + " characters"
	case "alpha":
		return "Should only contain letters"
	case "e164":
		return "Invalid phone number format"
	case "unique":
		return "All elements must be unique"
	case "gtefield":
		return "Must be after " + err.Param()
	default:
		return "Invalid value"
	}
}
