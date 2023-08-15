package http

import (
	"github.com/labstack/echo/v4"
	"go-two/common"
	storagedomain "go-two/v1/storage/domain"
	"net/http"
)

type storageHandler struct {
	echo           *echo.Group
	storageUseCase storagedomain.StorageUseCaseInterface
}

func NewStorageHandler(echo *echo.Group, storageUseCase storagedomain.StorageUseCaseInterface) {
	handler := &storageHandler{
		echo:           echo,
		storageUseCase: storageUseCase,
	}
	g := handler.echo.Group("/storage")
	g.GET("/:uuid", handler.getFile)
}

func (h storageHandler) getFile(c echo.Context) error {
	param := c.Param("uuid")
	file, err := h.storageUseCase.GetFile(param)
	if err != nil {
		return c.JSON(http.StatusForbidden, common.ErrorResponse([]error{err}))
	}
	return c.JSON(http.StatusOK, common.SuccessResponse(file))
}
