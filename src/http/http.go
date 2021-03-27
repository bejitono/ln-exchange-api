package http

import (
	"net/http"
	"strconv"

	exchange "github.com/bejitono/ln-exchange-api/src/domain"
	"github.com/bejitono/ln-exchange-api/src/service"
	"github.com/gin-gonic/gin"
	"github.com/stdemicheli/bookstore_oauth-api/src/utils/errors"
)

type ExchangeHTTPHandler interface {
	GetExchangeById(*gin.Context)
	GetExchanges(*gin.Context)
	Withdraw(*gin.Context)
	GetInvoice(*gin.Context)
}

type exchangeHTTPHandler struct {
	service service.Service
}

func NewHandler(s service.Service) ExchangeHTTPHandler {
	return &exchangeHTTPHandler{
		service: s,
	}
}

func (h *exchangeHTTPHandler) GetExchangeById(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, errors.NewInternalServerError("Not implemented"))
}

func (h *exchangeHTTPHandler) GetExchanges(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, errors.NewInternalServerError("Not implemented"))
}

func (h *exchangeHTTPHandler) Withdraw(c *gin.Context) {
	var req exchange.WithdrawalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		restErr := errors.NewBadRequestError("Invalid request")
		c.JSON(restErr.Status, restErr)
		return
	}
	if err := h.service.Withdraw(req); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *exchangeHTTPHandler) GetInvoice(c *gin.Context) {
	var req exchange.InvoiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		restErr := errors.NewBadRequestError("Invalid request")
		c.JSON(restErr.Status, restErr)
		return
	}
	if invoice, err := h.service.GetInvoice(); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, invoice)
}

// Private

func getExchangeId(exchangeIdParam string) (int64, *errors.RestErr) {
	exchangeId, err := strconv.ParseInt(exchangeIdParam, 10, 64)
	if err != nil {
		return -1, errors.NewBadRequestError("Exchange id should be a number")
	}
	return exchangeId, nil
}
