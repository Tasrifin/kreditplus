package main

import (
	"github.com/Tasrifin/kreditplus/config"
	"github.com/Tasrifin/kreditplus/controllers"
	"github.com/Tasrifin/kreditplus/repositories"
	"github.com/Tasrifin/kreditplus/services"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()
	route := gin.Default()

	transactionRepo := repositories.NewTransactionRepo(db)
	limitRepo := repositories.NewLimitRepo(db)
	transactionService := services.NewTransactionService(transactionRepo, limitRepo)
	transactionController := controllers.NewTransactionController(transactionService)

	//route
	route.POST("/transaction/store", transactionController.CreateTransaction)

	route.Run(config.APP_PORT)
}
