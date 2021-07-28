package app

import (
	"log"
	"regexp"
)

const urlRegex = "https?:\\/\\/(www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{1,256}\\.[a-zA-Z0-9()]{1,6}\\b([-a-zA-Z0-9()@:%_\\+.~#?&//=]*)"

type LinkValidator struct {

}

func NewLinkValidator() *LinkValidator {
	return &LinkValidator{}
}

func (l *LinkValidator) Validate(url string) bool {
	isValid, err := regexp.MatchString(urlRegex, url)
	if err != nil {
		log.Println(err)
	}
	return isValid
}