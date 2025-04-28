package utils

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

var validate = validator.New()

func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}

func init() {
	// validate username
	validate.RegisterValidation("username", ValidateUsername)
	//validate strong password
	validate.RegisterValidation("strongpassword", ValidateStrongPassword)

}

func ValidateUsername(fl validator.FieldLevel) bool {
	username := fl.Field().String()
	return regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(username)
}

func ValidateStrongPassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	if len(password) < 8 {
		return false
	}
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`).MatchString(password)
	return hasUpper && hasLower && hasNumber && hasSpecial
}

func GetValidationErrors(err error) map[string]string {
	if _, ok := err.(*validator.InvalidValidationError); ok {
		return map[string]string{
			"message": err.Error(),
		}
	}

	errors := make(map[string]string)

	for _, vErr := range err.(validator.ValidationErrors) {
		errors[vErr.Field()] = GetErrMessage(vErr)
	}
	return errors
}

func GetErrMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return "Value is too short"
	case "max":
		return "Value is too long"
	case "username":
		return "Username must contain only letters and number"
	case "strongpassword":
		return "Password must be at least 8 characters, with uppercase, lowercase, number, and special character"
	default:
		return err.Error()
	}
}
