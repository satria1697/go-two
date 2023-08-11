package handler

import (
	"github.com/labstack/echo/v4"
	"go-two/v1/hello/handler/http"
	"go-two/v1/hello/usecase"
)

func NewV1Handler(e *echo.Echo, useCase usecase.HelloUseCase) {
	g := e.Group("/v1")
	http.NewHelloHandler(g, useCase)
}
