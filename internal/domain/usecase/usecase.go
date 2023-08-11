package usecase

type (
	UseCase interface {
		CreateOrder()
	}

	useCase struct{}
)

func NewUseCase() UseCase {
	return &useCase{}
}
