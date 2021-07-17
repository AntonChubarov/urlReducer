package app

import (
	"log"
	"regexp"
)

type LinkValidator struct {}

const (
	urlRegEx = "https?:\\/\\/(www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{1,256}\\.[a-zA-Z0-9()]{1,6}\\b([-a-zA-Z0-9()@:%_\\+.~#?&//=]*)"
	idRegEx = "^[a-zA-Z0-9]*$"
)

func NewLinkValidator() *LinkValidator {
	return &LinkValidator{}
}

func (l *LinkValidator) ValidateInitialURL(url string) (isValid bool) {
	isValid, err := regexp.MatchString(urlRegEx, url)
	if err != nil {
		log.Println(err)
	}
	return
}

func (l *LinkValidator) ValidateID(id string) (isValid bool) {
	isValid, err := regexp.MatchString(idRegEx, id)
	if err != nil {
		log.Println(err)
	}
	return
}