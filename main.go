package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-two/v1/handler"
	"go-two/v1/hello/usecase"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	helloUseCase := usecase.NewHelloUseCase()
	handler.NewV1Handler(e, helloUseCase)
	e.Logger.Fatal(e.Start(":1323"))
}
