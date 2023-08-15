package usecase

type UseCase interface {
	ConvertOrderRequest(body string, requestId string) error
}
