package console

import (
	"bufio"
	"client/domain"
	"fmt"
	"log"
	"os"
	"strings"
)

type consoleUI struct {}

func (c *consoleUI) GetCommand() string {
	fmt.Println("To get short link use command \"red\" (example: red https://www.somesite.com/somethingelse)")
	fmt.Println("To get full link using short one use command \"get\" (example: get "+ domain.WebHost +"/cFiuRTe)")
	fmt.Println("To open destination using short link use command \"open\" (example: open" + domain.WebHost + "/cFiuRTe)")

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
	}

	text = strings.Replace(text, "\n", "", -1)

	return text
}

func (c *consoleUI) ShowMessage(message string) {
	fmt.Println(message)
}

func NewConsoleUI() *consoleUI {
	return &consoleUI{}
}
