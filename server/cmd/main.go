package main

import (
	"github.com/labstack/echo/v4"
	"server/infrastructure/httpControllers"
)

func main() {
	e := echo.New()
	e.GET("/:id", httpControllers.GetController)
	e.POST("/", httpControllers.PostController)
	e.Logger.Fatal(e.Start(":8080"))
}
