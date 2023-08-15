package domain

type HandlerUseCaseInterface interface {
	GetHello(trigger string) (bool, string, error)
}
