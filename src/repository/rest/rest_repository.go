package rest

import (
	"fmt"
	"os"

	exchange "github.com/bejitono/ln-exchange-api/src/domain"
	"github.com/bejitono/ln-exchange-api/src/utils/errors"
	bitfinex "github.com/bitfinexcom/bitfinex-api-go/v2/rest"
)

const (
	bitfinexWallet       = "exchange"
	bitfinex_api_key_env = "BITFINEX_API_KEY"
	bitfinex_api_sec_env = "BITFINEX_API_SEC"
)

var (
	bitfinexClient *bitfinex.Client
)

type RestExchangeRepository interface {
	Withdraw(exchange.WithdrawalRequest) *errors.RestErr
}

type restExchangeRepository struct {
	bitfinexRepository RestExchangeRepository
}

type bitfinexRepository struct {
	RestExchangeRepository
}

func init() {
	bitfinex_api_key := os.Getenv(bitfinex_api_key_env)
	bitfinex_api_sec := os.Getenv(bitfinex_api_sec_env)
	bitfinexClient = bitfinex.NewClient().Credentials(bitfinex_api_key, bitfinex_api_sec)
}

func NewRestRepository() RestExchangeRepository {
	return &restExchangeRepository{}
}

func (r *restExchangeRepository) Withdraw(req exchange.WithdrawalRequest) *errors.RestErr {
	var err *errors.RestErr
	switch req.ExchangeId {
	default:
		err = r.bitfinexRepository.Withdraw(req)
	}
	return err
}

func (r *bitfinexRepository) Withdraw(req exchange.WithdrawalRequest) *errors.RestErr {
	notification, err := bitfinexClient.Wallet.Withdraw(bitfinexWallet, req.Currency, req.Amount, req.Address)
	fmt.Println(notification)
	if err != nil {
		return errors.NewNotFoundError(err.Error())
	}
	return nil
}
