package exchange

import "github.com/stdemicheli/bookstore_oauth-api/src/utils/errors"

type Repository interface {
}

type Service interface {
	GetExchangeById(string) (*Exchange, *errors.RestErr)
	GetExchanges() ([]Exchange, *errors.RestErr)
	Withdraw(exchangeId string, amount float64) *errors.RestErr
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetExchangeById(id string) (*Exchange, *errors.RestErr) {
	return nil, nil
}

func (s *service) GetExchanges() ([]Exchange, *errors.RestErr) {
	return nil, nil
}

func (s *service) Withdraw(exchangeId string, amount float64) *errors.RestErr {
	return nil
}
