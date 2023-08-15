package usecase

import (
	"errors"
	"go-two/common"
	"go-two/v1/storage/domain"
)

type StorageUseCase struct {
}

func NewStorageUseCase() domain.StorageUseCaseInterface {
	return StorageUseCase{}
}

func (s StorageUseCase) GetFile(uuid string) (string, error) {
	if uuid == "1234" {
		return "file ini", nil
	}
	return "", errors.New(common.NotFoundError)
}
