package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func newValidator() *validator.Validate {
	validate := validator.New()

	_ = validate.RegisterValidation("uuid", func (fl validator.FieldLevel) bool  {
		fieldValue := fl.Field().String()
		if _, err := uuid.Parse(fieldValue); err != nil {
			return true
		}
		return false
	})

	return validate
}

type validatorErr map[string]string

func validatorErrors(err error) validatorErr {
	vE := validatorErr{}
	for _, err := range err.(validator.ValidationErrors) {
		vE[err.Field()] = err.Error()
	}

	return vE
}

func generatePasswordHash(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err.Error()
	}

	return string(hash)
}


func comparePasswords(hashedPassword, inputPassword string) bool {
	hp := []byte(hashedPassword)
	ip := []byte(inputPassword)

	if err := bcrypt.CompareHashAndPassword(hp, ip); err != nil {
		return false
	}

	return true
}

