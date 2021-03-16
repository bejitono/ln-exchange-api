package http

import (
	exchange "github.com/bejitono/ln-exchange-api/src/domain"
	"github.com/gin-gonic/gin"
)

type ExchangeHTTPHandler interface {
	GetExchangeById(*gin.Context)
	GetExchanges(*gin.Context)
	Withdraw(*gin.Context)
}

type exchangeHTTPHandler struct {
	service exchange.Service
}

func NewHandler(s exchange.Service) ExchangeHTTPHandler {
	return &exchangeHTTPHandler{
		service: s,
	}
}

func (h *exchangeHTTPHandler) GetExchangeById(c *gin.Context) {

}

func (h *exchangeHTTPHandler) GetExchanges(c *gin.Context) {

}

func (h *exchangeHTTPHandler) Withdraw(c *gin.Context) {

}
