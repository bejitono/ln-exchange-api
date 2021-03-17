package exchange

import (
	db "github.com/bejitono/ln-exchange-api/src/repository/db"
	rest "github.com/bejitono/ln-exchange-api/src/repository/rest"
	"github.com/stdemicheli/bookstore_oauth-api/src/utils/errors"
)

type Service interface {
	GetExchangeById(int64) (*Exchange, *errors.RestErr)
	GetExchanges() ([]Exchange, *errors.RestErr)
	Withdraw(WithdrawalRequest) *errors.RestErr
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

func (s *service) GetExchangeById(id int64) (*Exchange, *errors.RestErr) {
	return nil, nil
}

func (s *service) GetExchanges() ([]Exchange, *errors.RestErr) {
	return nil, nil
}

func (s *service) Withdraw(req WithdrawalRequest) *errors.RestErr {
	// TODO: get exchange from exchange id

	switch req.ExchangeId {
	case 1:

	}
	return nil
}
