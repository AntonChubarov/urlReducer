package domain

import "fmt"

const (
	WebHost = "http://localhost:8080"
)

var ErrorNotFound = fmt.Errorf("initial link not found")
var ErrorInvalidURL = fmt.Errorf("your URL is not valid")
var ErrorInvalidShortURL = fmt.Errorf("your short URL is not valid")
var ErrorInternal = fmt.Errorf("internal server error")
