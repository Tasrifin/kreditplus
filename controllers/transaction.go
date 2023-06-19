package controllers

import (
	"net/http"

	"github.com/Tasrifin/kreditplus/params"
	"github.com/Tasrifin/kreditplus/services"
	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	transactionService services.TransactionService
}

func NewTransactionController(service *services.TransactionService) *TransactionController {
	return &TransactionController{
		transactionService: *service,
	}
}

func (t *TransactionController) CreateTransaction(c *gin.Context) {
	var req params.TransactionPayload

	err := c.ShouldBind(&req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})

		return
	}

	result := t.transactionService.StoreTransaction(req)

	c.JSON(result.Status, result.Payload)
}
