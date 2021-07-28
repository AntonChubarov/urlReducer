package console

import (
	"bufio"
	"client/domain"
	"fmt"
	"log"
	"os"
	"strings"
)

type UIConsole struct {}

func (c *UIConsole) GetCommand() string {
	fmt.Println("To get short link use command \"reduce\" (example: reduce https://www.somesite.com/somethingelse)")
	fmt.Println("To open destination using short link use command \"open\" (example: open" + domain.WebHost + "/cFiuRTe)")

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
	}

	text = strings.Replace(text, "\n", "", -1)

	return text
}

func (c *UIConsole) ShowMessage(message string) {
	fmt.Println(message)
}

func NewConsoleUI() *UIConsole {
	return &UIConsole{}
}
