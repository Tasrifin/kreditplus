package repositories

import (
	"github.com/Tasrifin/kreditplus/models"
	"gorm.io/gorm"
)

type LimitRepo interface {
	CheckLimit(customerId int, tenor int) (*models.Limit, error)
}

type limitRepo struct {
	db *gorm.DB
}

func NewLimitRepo(db *gorm.DB) LimitRepo {
	return &limitRepo{db}
}

func (l *limitRepo) CheckLimit(customerId int, tenor int) (*models.Limit, error) {
	var limit models.Limit
	err := l.db.Preload("Customer").Where("customer_id=?", customerId).Where("limit_month=?", tenor).Find(&limit).Error

	return &limit, err
}
