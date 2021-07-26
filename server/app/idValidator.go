package app

import (
	"log"
	"regexp"
)

type IDValidator struct {}

const idRegEx = "^[a-zA-Z0-9]*$"

func NewIDValidator() *IDValidator {
	return &IDValidator{}
}

func (l *IDValidator) Validate(id string) bool {
	isValid, err := regexp.MatchString(idRegEx, id)
	if err != nil {
		log.Println(err)
	}
	return isValid
}