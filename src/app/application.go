package app

import (
	exchange "github.com/bejitono/ln-exchange-api/src/domain"
	"github.com/bejitono/ln-exchange-api/src/http"
	db "github.com/bejitono/ln-exchange-api/src/repository"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	// Start db session

	dbRepository := db.NewRepository()
	service := exchange.NewService(dbRepository)
	handler := http.NewHandler(service)

	router.GET("/exchanges", handler.FillIn)
	router.GET("/exchanges/:exchange_id", handler.FillIn)
	router.POST("/exchanges/transaction", handler.FillIn)
	router.Run(":8080")
}
