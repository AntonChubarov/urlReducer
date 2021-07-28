package app

import (
	"log"
	"regexp"
)

const idRegex = "^[a-zA-Z0-9]*$"

type IDValidator struct {

}

func NewIDValidator() *IDValidator {
	return &IDValidator{}
}

func (i *IDValidator) Validate(id string) bool {
	isValid, err := regexp.MatchString(idRegex, id)
	if err != nil {
		log.Println(err)
	}
	return isValid
}