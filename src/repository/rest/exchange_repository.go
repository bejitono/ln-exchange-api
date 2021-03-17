package rest

type RestExchangeRepository interface {
}

type bitfinexExchangeRepository struct {
}

func NewRestRepository() RestExchangeRepository {
	return &bitfinexExchangeRepository{}
}
