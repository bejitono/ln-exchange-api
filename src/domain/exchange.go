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
