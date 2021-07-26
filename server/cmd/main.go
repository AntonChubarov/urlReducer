package main

import (
	"github.com/labstack/echo/v4"
	"server/app"
	"server/domain"
	"server/infrastructure/dal"
	"server/infrastructure/httpcontrollers"
)

func main() {
	config := domain.NewServerConfig()

	linkValidator := app.NewLinkValidator(config)
	idValidator := app.NewIDValidator()
	hasher := app.NewLinkHasher()
	storage := dal.NewDatabaseConnector()

	service := app.NewService(linkValidator, idValidator, hasher, storage)

	getController := httpcontrollers.NewGetController(service)
	postController := httpcontrollers.NewPostController(service)

	e := echo.New()
	e.GET("/:id", getController.Get)
	e.POST("/", postController.Post)
	e.Logger.Fatal(e.Start(":8080"))
}
