package usecase

import (
	"errors"
	"go-two/common"
	"go-two/v1/hello/domain"
)

type HelloUseCase struct {
}

func NewHelloUseCase() domain.HandlerUseCaseInterface {
	return HelloUseCase{}
}

func (c HelloUseCase) GetHello(trigger string) (bool, string, error) {
	if trigger == "foo" {
		return true, "Hello World", nil
	}
	return false, "Hello World", errors.New(common.ValidationError)
}
