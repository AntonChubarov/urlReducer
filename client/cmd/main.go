package main

import (
	"client/app"
	"client/infrastructure/console"
	"client/infrastructure/httpclient"
)

func main() {
	userInterface := console.NewConsoleUI()
	commandParser := app.NewCommandParser()
	client := httpclient.NewHTTPClient()

	service := app.NewService(userInterface, commandParser, client)

	service.Run()
}
