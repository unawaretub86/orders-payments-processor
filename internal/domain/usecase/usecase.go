package usecase

type (
	UseCase interface {
		CreateOrderRequest(body string, requestId string) error
	}
)
