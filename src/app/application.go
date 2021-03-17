package app

import (
	exchange "github.com/bejitono/ln-exchange-api/src/domain"
	"github.com/bejitono/ln-exchange-api/src/http"
	db "github.com/bejitono/ln-exchange-api/src/repository/db"
	"github.com/bejitono/ln-exchange-api/src/repository/rest"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	// Start db session

	dbRepository := db.NewDbRepository()
	restRepository := rest.NewRestRepository()
	service := exchange.NewService(restRepository, dbRepository)
	handler := http.NewHandler(service)

	router.GET("/exchanges", handler.GetExchangeById)
	router.GET("/exchanges/:exchange_id", handler.GetExchanges)
	router.POST("/exchanges/transaction", handler.Withdraw)
	router.Run(":8080")
}
