package main

import (
	"client/app"
	"client/infrastructure/console"
	"client/infrastructure/httpClient"
)

func main() {
	userInterface := console.NewConsoleUI()
	commandParser := app.NewCommandParser()
	linkValidator := app.NewLinkValidator()
	client := httpClient.NewHttpClient()

	service := app.NewService(userInterface, commandParser, linkValidator, client)

	service.Run()
}
