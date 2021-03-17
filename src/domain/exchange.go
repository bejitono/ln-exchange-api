package exchange

type Exchange struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Wallet Wallet `json:"wallet"`
}

type Wallet struct {
	Id       int64   `json:"id"`
	Name     string  `json:"name"`
	Balance  float64 `json:"balance"`
	Currency string  `json:"Currency"`
}

type WithdrawalRequest struct {
	ExchangeId int64   `json:"exchangeId"`
	Currency   string  `json:"currency"`
	Address    string  `json:"address"`
	Amount     float64 `json:"amount"`
}
