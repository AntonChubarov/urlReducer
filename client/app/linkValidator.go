package app

import (
	"log"
	"regexp"
)

type linkValidator struct {}

const (
	urlRegEx = "https?:\\/\\/(www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{1,256}\\.[a-zA-Z0-9()]{1,6}\\b([-a-zA-Z0-9()@:%_\\+.~#?&//=]*)"
)

func (l *linkValidator) Validate(url string) (isValid bool) {
	isValid, err := regexp.MatchString(urlRegEx, url)
	if err != nil {
		log.Println(err)
	}
	return
}

func NewLinkValidator() *linkValidator {
	return &linkValidator{}
}
