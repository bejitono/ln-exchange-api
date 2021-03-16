package exchange

type Repository interface {
}

type Service interface {
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
