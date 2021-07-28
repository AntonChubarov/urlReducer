package main

import (
	"github.com/labstack/echo/v4"
	"server/app"
	"server/config"
	"server/infrastructure/dal"
	"server/infrastructure/httpcontrollers"
)

func main() {
	sConfig := config.NewServerConfig()

	linkValidator := app.NewLinkValidator()
	idValidator := app.NewIDValidator()
	hasher := app.NewLinkHasher(sConfig)

	storage := dal.NewDatabaseConnector(sConfig)
	defer storage.CloseDatabaseConnection()

	service := app.NewService(linkValidator, idValidator, hasher, storage, sConfig)

	getController := httpcontrollers.NewGetController(service)
	postController := httpcontrollers.NewPostController(service)

	e := echo.New()
	e.GET("/:id", getController.Get)
	e.POST("/", postController.Post)
	e.Logger.Fatal(e.Start(sConfig.Host.ServerStartPort))
}
