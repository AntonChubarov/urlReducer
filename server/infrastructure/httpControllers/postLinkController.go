package httpControllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"server/app"
	"server/domain"
)

func PostController(c echo.Context) error {
	var request domain.Request

	err := c.Bind(&request)
	if err != nil {
		return echo.ErrBadRequest
	}

	response, err := app.SaveLink(request)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}