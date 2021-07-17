package httpControllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"server/app"
)

func GetController(c echo.Context) error {
	//id := c.QueryParam("id")
	id := c.Param("id")

	if response, err := app.GetLink(id); err == nil {
		return c.JSON(http.StatusOK, response)
	} else {
		return c.String(http.StatusBadRequest, err.Error())
	}
}