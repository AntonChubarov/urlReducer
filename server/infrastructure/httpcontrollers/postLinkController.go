package httpcontrollers

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"server/domain"
)

type PostController struct {
	Service domain.IService
}

func NewPostController(service domain.IService) *PostController {
	return &PostController{Service: service}
}

func (p *PostController) Post(c echo.Context) error {
	var request domain.Request

	err := c.Bind(&request)
	if err != nil {
		log.Println(err)
		return echo.ErrBadRequest
	}

	shortURL, err := p.Service.SaveLink(request.InitialURL)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, domain.Response{URL: shortURL})
}