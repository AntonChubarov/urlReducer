package domain

import "fmt"

const (
	WebHost = "http://localhost:8080"
)

var ErrorNotFound = fmt.Errorf("Initial link not found")
var ErrorInvalidURL = fmt.Errorf("Your URL is not valid")
var ErrorInvalidShortURL = fmt.Errorf("Your short URL is not valid")
var ErrorInternal = fmt.Errorf("Internal server error")
