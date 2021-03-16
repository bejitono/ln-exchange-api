package http

import (
	exchange "github.com/bejitono/ln-exchange-api/src/domain"
	"github.com/gin-gonic/gin"
)

type ExchangeHTTPHandler interface {
	FillIn(*gin.Context)
}

type exchangeHTTPHandler struct {
	service exchange.Service
}

func NewHandler(s exchange.Service) ExchangeHTTPHandler {
	return &exchangeHTTPHandler{
		service: s,
	}
}

func (h *exchangeHTTPHandler) FillIn(c *gin.Context) {

}