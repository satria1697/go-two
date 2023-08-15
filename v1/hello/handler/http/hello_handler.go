package http

import (
	"github.com/labstack/echo/v4"
	"go-two/common"
	"go-two/v1/hello/domain"
	"net/http"
)

type helloHandler struct {
	echo         *echo.Group
	helloUseCase domain.HandlerUseCaseInterface
}

func NewHelloHandler(echo *echo.Group, helloUseCase domain.HandlerUseCaseInterface) {
	handler := &helloHandler{
		echo:         echo,
		helloUseCase: helloUseCase,
	}
	g := handler.echo.Group("/hello")
	g.GET("", handler.getHello)
}

func (h helloHandler) getHello(c echo.Context) error {
	trigger := c.QueryParam("trigger")
	status, message, err := h.helloUseCase.GetHello(trigger)
	if status {
		return c.JSON(http.StatusOK, common.SuccessResponse(message))
	}
	return c.JSON(http.StatusOK, common.ErrorResponse([]error{err}))
}
