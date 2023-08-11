package http

import (
	"github.com/labstack/echo/v4"
	"go-two/common"
	"go-two/v1/hello/usecase"
	"net/http"
)

type HelloHandler struct {
	echo         *echo.Group
	helloUseCase usecase.HelloUseCase
}

func NewHelloHandler(echo *echo.Group, helloUseCase usecase.HelloUseCase) {
	handler := &HelloHandler{
		echo:         echo,
		helloUseCase: helloUseCase,
	}
	g := handler.echo.Group("/hello")
	g.GET("", handler.GetHello)
}

func (h HelloHandler) GetHello(c echo.Context) error {
	trigger := c.QueryParam("trigger")
	status, message, err := h.helloUseCase.GetHello(trigger)
	if status {
		return c.JSON(http.StatusOK, common.SuccessResponse(message))
	}
	return c.JSON(http.StatusOK, common.ErrorResponse(err))
}
