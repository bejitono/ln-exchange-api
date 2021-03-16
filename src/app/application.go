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

	// Make repo / service / http handler
	dbRepository := db.NewRepository()
	service := exchange.NewService(dbRepository)
	handler := http.NewHandler(service)

	// /wallet
	router.GET("/wallet", handler.FillIn)
	router.POST("/wallet/withdraw", handler.FillIn)
	router.Run(":8080")
}
