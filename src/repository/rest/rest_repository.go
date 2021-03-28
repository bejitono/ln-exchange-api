package rest

import (
	exchange "github.com/bejitono/ln-exchange-api/src/domain"
	"github.com/bejitono/ln-exchange-api/src/utils/errors"
)

type RestExchangeRepository interface {
	Withdraw(exchange.WithdrawalRequest) *errors.RestErr
	GetInvoice(exchange.InvoiceRequest) (*exchange.Invoice, *errors.RestErr)
	GetAddress(exchange.AddressRequest) (*exchange.Address, *errors.RestErr)
}

type restExchangeRepository struct {
	bitfinexRepository RestExchangeRepository
}

func NewRestRepository() RestExchangeRepository {
	return &restExchangeRepository{
		bitfinexRepository: &bitfinexRepository{},
	}
}

func (r *restExchangeRepository) Withdraw(req exchange.WithdrawalRequest) *errors.RestErr {
	var err *errors.RestErr
	switch req.ExchangeId {
	default:
		err = r.bitfinexRepository.Withdraw(req)
	}
	return err
}

func (r *restExchangeRepository) GetInvoice(req exchange.InvoiceRequest) (*exchange.Invoice, *errors.RestErr) {
	var err *errors.RestErr
	var invoice *exchange.Invoice
	switch req.ExchangeId {
	default:
		invoice, err = r.bitfinexRepository.GetInvoice(req)
	}
	return invoice, err
}

func (r *restExchangeRepository) GetAddress(req exchange.AddressRequest) (*exchange.Address, *errors.RestErr) {
	var err *errors.RestErr
	var address *exchange.Address
	switch req.ExchangeId {
	default:
		address, err = r.bitfinexRepository.GetAddress(req)
	}
	return address, err
}
