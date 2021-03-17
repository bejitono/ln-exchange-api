package service

import (
	exchange "github.com/bejitono/ln-exchange-api/src/domain"
	db "github.com/bejitono/ln-exchange-api/src/repository/db"
	rest "github.com/bejitono/ln-exchange-api/src/repository/rest"
	"github.com/stdemicheli/bookstore_oauth-api/src/utils/errors"
)

type Service interface {
	GetExchangeById(int64) (*exchange.Exchange, *errors.RestErr)
	GetExchanges() ([]exchange.Exchange, *errors.RestErr)
	Withdraw(exchange.WithdrawalRequest) *errors.RestErr
}

type service struct {
	dbRepository   db.DbRepository
	restRepository rest.RestExchangeRepository
}

func NewService(restRepo rest.RestExchangeRepository, dbRepo db.DbRepository) Service {
	return &service{
		dbRepository:   dbRepo,
		restRepository: restRepo,
	}
}

func (s *service) GetExchangeById(id int64) (*exchange.Exchange, *errors.RestErr) {
	return nil, nil
}

func (s *service) GetExchanges() ([]exchange.Exchange, *errors.RestErr) {
	return nil, nil
}

func (s *service) Withdraw(req exchange.WithdrawalRequest) *errors.RestErr {
	// TODO: get exchange from exchange id
	if err := s.restRepository.Withdraw(req); err != nil {
		return errors.NewNotFoundError(err.Error)
	}
	return nil
}
