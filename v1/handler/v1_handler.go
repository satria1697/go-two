package handler

import (
	"github.com/labstack/echo/v4"
	hellodomain "go-two/v1/hello/domain"
	hellohandler "go-two/v1/hello/handler/http"
	storagedomain "go-two/v1/storage/domain"
	storagehandler "go-two/v1/storage/handler/http"
)

func NewV1Handler(e *echo.Echo, handlerUseCase hellodomain.HandlerUseCaseInterface, storageUseCase storagedomain.StorageUseCaseInterface) {
	g := e.Group("/v1")
	hellohandler.NewHelloHandler(g, handlerUseCase)
	storagehandler.NewStorageHandler(g, storageUseCase)
}
