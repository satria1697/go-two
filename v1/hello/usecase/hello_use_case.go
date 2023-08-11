package usecase

import (
	"errors"
	"go-two/common"
)

type HelloUseCase struct {
}

func NewHelloUseCase() HelloUseCase {
	return HelloUseCase{}
}

func (c HelloUseCase) GetHello(trigger string) (bool, string, error) {
	if trigger == "foo" {
		return true, "Hello World", nil
	}
	return false, "Hello World", errors.New(common.ValidationError)
}
