package app

import (
	"log"
	"regexp"
	"server/domain"
)

type LinkValidator struct {
	URLRegex string
}

func NewLinkValidator(config domain.ServerConfig) *LinkValidator {
	return &LinkValidator{
		URLRegex: config.URLRegex,
	}
}

func (l *LinkValidator) Validate(url string) bool {
	isValid, err := regexp.MatchString(l.URLRegex, url)
	if err != nil {
		log.Println(err)
	}
	return isValid
}