package app

import (
	"client/domain"
	"strings"
)

type CommandParser struct {}

func (c *CommandParser) Parse(input string) (command, content string, err error) {
	input = strings.Trim(input, " ")

	spaceCount := strings.Count(input, " ")

	if spaceCount == 0 || spaceCount > 1 {
		return "", "", domain.ErrorInputDivide
	}

	splitted := strings.Split(input, " ")

	return splitted[0], splitted[1], nil
}

func NewCommandParser() *CommandParser {
	return &CommandParser{}
}
