package utils

import (
	"github.com/kurianCoding/EchoApp/models"
	"gopkg.in/go-playground/validator.v8"
)

func ValidateLogin(l *models.User) error {
	return validate(l)
}

func validate(l *models.User) error {
	config := &validate.Config{TagName: "validate"}
	validate := validator.New(config)
	errs := validate.Struct(l)
	return errs[0]
}
