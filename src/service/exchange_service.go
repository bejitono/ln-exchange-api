package service

import (
	"strings"

	exchange "github.com/bejitono/ln-exchange-api/src/domain"
	db "github.com/bejitono/ln-exchange-api/src/repository/db"
	rest "github.com/bejitono/ln-exchange-api/src/repository/rest"
	"github.com/bejitono/ln-exchange-api/src/utils/errors"
)

const (
	lnCurrency = "lnx"
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
	return nil, errors.NewNotFoundError("not implemented")
}

func (s *service) GetExchanges() ([]exchange.Exchange, *errors.RestErr) {
	return nil, errors.NewNotFoundError("not implemented")
}

func (s *service) Withdraw(req exchange.WithdrawalRequest) *errors.RestErr {
	// TODO: get exchange from exchange id
	if strings.ToLower(req.Currency) != lnCurrency {
		return errors.NewBadRequestError("Only lightning (lnx) is currently supported")
	}

	return s.restRepository.Withdraw(req)
}
