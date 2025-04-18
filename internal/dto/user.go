package dto

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type CreateUserPayload struct {
	Name     string `schema:"name" validate:"required,min=2,max=50"`
	Username string `schema:"username" validate:"required,alphanum,min=3,max=20"`
	Password string `schema:"password" validate:"required,min=8"`
	Email    string `schema:"email" validate:"required,email"`
}

func (u *CreateUserPayload) Validate(validate *validator.Validate) map[string]string {
	err := validate.Struct(u)
	if err == nil {
		return nil
	}

	errors := make(map[string]string)
	for _, e := range err.(validator.ValidationErrors) {
		switch e.Field() {
		case "Name":
			switch e.Tag() {
			case "required":
				errors["name"] = "Name is required."
			case "min":
				errors["name"] = "Name must be at least 2 characters."
			case "max":
				errors["name"] = "Name must be at most 50 characters."
			}
		case "Username":
			switch e.Tag() {
			case "required":
				errors["username"] = "Username is required."
			case "alphanum":
				errors["username"] = "Username can only contain letters and numbers."
			case "min":
				errors["username"] = "Username must be at least 3 characters."
			case "max":
				errors["username"] = "Username must be at most 20 characters."
			}
		case "Password":
			switch e.Tag() {
			case "required":
				errors["password"] = "Password is required."
			case "min":
				errors["password"] = "Password must be at least 8 characters."
			}
		case "Email":
			switch e.Tag() {
			case "required":
				errors["email"] = "Email is required."
			case "email":
				errors["email"] = "Email must be a valid email address."
			}
		default:
			errors[e.Field()] = fmt.Sprintf("Invalid value for %s", e.Field())
		}
	}

	return errors
}
