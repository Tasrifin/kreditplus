package repositories

import (
	"github.com/Tasrifin/kreditplus/models"
	"gorm.io/gorm"
)

type TransactionRepo interface {
	StoreTransaction(data *models.Transaction) (*models.Transaction, error)
}

type transactionRepo struct {
	db *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) TransactionRepo {
	return &transactionRepo{db}
}

func (t *transactionRepo) StoreTransaction(data *models.Transaction) (*models.Transaction, error) {
	return data, t.db.Create(&data).Error
}
