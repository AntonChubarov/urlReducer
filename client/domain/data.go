package domain

import "fmt"

const (
	WebHost = "http://localhost:8080"
	)

var ErrorInputDivide = fmt.Errorf("Command and link should be divided by a single space")
var ErrorInvalidURL = fmt.Errorf("Entered link is invalid")
var ErrorInvalidCommand = fmt.Errorf("Entered command is invalid")

var Commands = map[string]struct{} {
	"red": {},
	"get": {},
	"open": {},
}