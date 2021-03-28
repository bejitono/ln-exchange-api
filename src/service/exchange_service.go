package service

import (
	exchange "github.com/bejitono/ln-exchange-api/src/domain"
	db "github.com/bejitono/ln-exchange-api/src/repository/db"
	rest "github.com/bejitono/ln-exchange-api/src/repository/rest"
	"github.com/bejitono/ln-exchange-api/src/utils/errors"
)

const (
	lnCurrency = "lnx"
)

var validCurrencies = map[string]struct {
	name string
	min  float64
	max  float64
}{
	"LNX": {
		name: "Bitcoin Lightning Network",
		min:  0.000001,
		max:  0.02,
	},
}

type Service interface {
	GetExchangeById(int64) (*exchange.Exchange, *errors.RestErr)
	GetExchanges() ([]exchange.Exchange, *errors.RestErr)
	Withdraw(exchange.WithdrawalRequest) *errors.RestErr
	GetInvoice(exchange.InvoiceRequest) (*exchange.Invoice, *errors.RestErr)
	GetAddress(exchange.AddressRequest) (*exchange.Address, *errors.RestErr)
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
	if err := validate(req.Currency); err != nil {
		return err
	}
	return s.restRepository.Withdraw(req)
}

func (s *service) GetInvoice(req exchange.InvoiceRequest) (*exchange.Invoice, *errors.RestErr) {
	if err := validate(req.Currency); err != nil {
		return nil, err
	}
	return s.restRepository.GetInvoice(req)
}

func (s *service) GetAddress(req exchange.AddressRequest) (*exchange.Address, *errors.RestErr) {
	return s.restRepository.GetAddress(req)
}

func validate(currency string) *errors.RestErr {
	if _, ok := validCurrencies[currency]; !ok {
		return errors.NewBadRequestError("Currency not supported")
	}
	return nil
}
