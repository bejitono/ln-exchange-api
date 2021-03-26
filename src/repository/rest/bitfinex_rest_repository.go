package rest

import (
	"os"
	"strings"

	exchange "github.com/bejitono/ln-exchange-api/src/domain"
	"github.com/bejitono/ln-exchange-api/src/utils/errors"
	"github.com/bitfinexcom/bitfinex-api-go/pkg/models/notification"
	bitfinex "github.com/bitfinexcom/bitfinex-api-go/v2/rest"
)

const (
	bitfinexWallet                = "exchange"
	bitfinex_api_key_env          = "BITFINEX_API_KEY"
	bitfinex_api_sec_env          = "BITFINEX_API_SEC"
	bitfinex_success_status       = "SUCCESS"
	bitfinex_invalid_invoice_text = "Invalid Invoice"
)

var (
	bitfinexClient *bitfinex.Client
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

type bitfinexRepository struct {
	RestExchangeRepository
}

func init() {
	bitfinex_api_key := os.Getenv(bitfinex_api_key_env)
	bitfinex_api_sec := os.Getenv(bitfinex_api_sec_env)
	bitfinexClient = bitfinex.NewClient().Credentials(bitfinex_api_key, bitfinex_api_sec)
}

func (r *bitfinexRepository) Withdraw(req exchange.WithdrawalRequest) *errors.RestErr {
	notification, err := bitfinexClient.Wallet.Withdraw(bitfinexWallet, req.Currency, req.Amount, req.Address)

	if err != nil {
		return errors.NewNotFoundError(err.Error())
	}

	switch notification.Status {
	case bitfinex_success_status:
		return handleSuccess(notification)
	default:
		return errors.NewInternalServerError(notification.Text)
	}
}

func (r *bitfinexRepository) GetInvoice(req exchange.InvoiceRequest) (*exchange.Invoice, *errors.RestErr) {
	depositReq := newDepositRequest(req)
	if err := validate(depositReq.Currency); err != nil {
		return nil, err
	}
	invoice, err := bitfinexClient.Invoice.GenerateInvoice(depositReq)
	if err != nil {
		return nil, errors.NewNotFoundError(err.Error())
	}
	// TODO: Validate invoice (only lnx)
	return &exchange.Invoice{
		InvoiceHash: invoice.InvoiceHash,
		Invoice:     invoice.Invoice,
		Amount:      invoice.Amount,
	}, nil
}

func handleSuccess(n *notification.Notification) *errors.RestErr {
	switch n.Text {
	case bitfinex_invalid_invoice_text:
		return errors.NewInternalServerError(n.Text)
	default:
		return nil
	}
}

func newDepositRequest(req exchange.InvoiceRequest) bitfinex.DepositInvoiceRequest {
	return bitfinex.DepositInvoiceRequest{
		Currency: strings.ToUpper(req.Currency),
		Wallet:   req.Wallet,
		Amount:   req.Amount,
	}
}

func validate(currency string) *errors.RestErr {
	if _, ok := validCurrencies[currency]; !ok {
		return errors.NewBadRequestError("Currency not supported")
	}
	return nil
}
