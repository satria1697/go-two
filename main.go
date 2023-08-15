package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-two/v1/handler"
	hellousecase "go-two/v1/hello/usecase"
	storaeusecase "go-two/v1/storage/usecase"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	helloUseCase := hellousecase.NewHelloUseCase()
	storageUseCase := storaeusecase.NewStorageUseCase()
	handler.NewV1Handler(e, helloUseCase, storageUseCase)
	e.Logger.Fatal(e.Start(":1323"))
}
