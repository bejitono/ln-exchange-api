package rest

import (
	exchange "github.com/bejitono/ln-exchange-api/src/domain"
	"github.com/bejitono/ln-exchange-api/src/utils/errors"
)

type RestExchangeRepository interface {
	Withdraw(exchange.WithdrawalRequest) *errors.RestErr
	GetInvoice(exchange.InvoiceRequest) (*exchange.Invoice, *errors.RestErr)
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

func (r *restExchangeRepository) GetInvoice(exchange.InvoiceRequest) (*exchange.Invoice, *errors.RestErr) {
	return nil, nil
}
