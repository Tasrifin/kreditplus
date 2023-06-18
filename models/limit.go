package models

type Limit struct {
	ID         int     `gorm:"primaryKey" json:"id"`
	LimitMonth int     `gorm:"not null" json:"limit_month"`
	LimitValue float64 `gorm:"not null" json:"limit_value"`
	CustomerID int
	Customer   *Customer
}
