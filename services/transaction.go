package services

import (
	"fmt"
	"net/http"

	"github.com/Tasrifin/kreditplus/helpers"
	"github.com/Tasrifin/kreditplus/models"
	"github.com/Tasrifin/kreditplus/params"
	"github.com/Tasrifin/kreditplus/repositories"
	"github.com/gin-gonic/gin"
)

type TransactionService struct {
	transactionRepo repositories.TransactionRepo
	limitRepo       repositories.LimitRepo
}

func NewTransactionService(repo repositories.TransactionRepo, limitRepo repositories.LimitRepo) *TransactionService {
	return &TransactionService{
		transactionRepo: repo,
		limitRepo:       limitRepo,
	}
}

func (t *TransactionService) StoreTransaction(request params.TransactionPayload) *params.Response {
	// check if have tenor
	checkTenor, err := t.limitRepo.CheckLimit(request.CustomerID, request.Tenor)
	if err != nil {
		return &params.Response{
			Status: http.StatusInternalServerError,
			Payload: gin.H{
				"error": err.Error(),
			},
		}
	}

	if checkTenor.ID == 0 {
		return &params.Response{
			Status: http.StatusInternalServerError,
			Payload: gin.H{
				"message": fmt.Sprintf("user doesn't have tenor %d", request.Tenor),
			},
		}
	}

	// check otr price - tenor - amount
	totalPrice := request.OTRPrice + request.AdminFee + request.TotalInterest
	tenorAmount := float64(checkTenor.LimitMonth) * checkTenor.LimitValue
	if totalPrice > tenorAmount {
		return &params.Response{
			Status: http.StatusInternalServerError,
			Payload: gin.H{
				"message": "tenor amount exceeds limit",
			},
		}
	}

	trx := models.Transaction{
		ContractNumber:  helpers.GenerateUUID(),
		CustomerID:      request.CustomerID,
		OTRPrice:        request.OTRPrice,
		AdminFee:        request.AdminFee,
		TotalInstalment: request.Tenor,
		TotalInterest:   request.TotalInterest,
		ProductName:     request.ProductName,
		LimitID:         checkTenor.ID,
	}

	storedData, err := t.transactionRepo.StoreTransaction(&trx)
	if err != nil {
		return &params.Response{
			Status: http.StatusInternalServerError,
			Payload: gin.H{
				"error": err.Error(),
			},
		}
	}

	return &params.Response{
		Status: http.StatusCreated,
		Payload: gin.H{
			"contract number": storedData.ContractNumber,
		},
	}
}
