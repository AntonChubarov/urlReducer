package httpcontrollers

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"server/app"
)

type GetController struct {
	Service *app.Service
}

func NewGetController(service *app.Service) *GetController {
	return &GetController{Service: service}
}

func (g *GetController) Get(c echo.Context) error {
	id := c.Param("id")
	var initialLink string
	var err error
	if initialLink, err = g.Service.GetLink(id); err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, err.Error())
	}
	c.Redirect(http.StatusSeeOther, initialLink)
	return nil
}