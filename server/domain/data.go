package domain

import "fmt"

var ErrorNotFound = fmt.Errorf("initial link not found")
var ErrorInvalidURL = fmt.Errorf("your URL is not valid")
var ErrorInvalidShortURL = fmt.Errorf("your short URL is not valid")
var ErrorInternal = fmt.Errorf("internal server error")
