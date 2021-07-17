package app

import (
	"client/domain"
	"strings"
)

type commandParser struct {}

func (c *commandParser) Parse(input string) (string, string, error) {
	input = strings.Trim(input, " ")

	spaceCount := strings.Count(input, " ")

	if spaceCount == 0 || spaceCount > 1 {
		return "", "", domain.ErrorInputDivide
	}

	splitted := strings.Split(input, " ")

	return splitted[0], splitted[1], nil
}

func NewCommandParser() *commandParser {
	return &commandParser{}
}
