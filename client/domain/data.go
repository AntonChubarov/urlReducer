package domain

import "fmt"

const WebHost = "http://localhost:8080"

var ErrorInputDivide = fmt.Errorf("command and link should be divided by a single space")
var ErrorInvalidCommand = fmt.Errorf("entered command is invalid")

var Commands = map[string]struct{} {
	"reduce": {},
	"open": {},
}