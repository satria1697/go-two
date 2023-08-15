package domain

type StorageUseCaseInterface interface {
	GetFile(uuid string) (string, error)
}
