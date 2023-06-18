package models

type Transaction struct {
	ID              int     `gorm:"primaryKey" json:"id"`
	ContractNumber  string  `gorm:"not null" json:"contract_number"`
	OTRPrice        float64 `gorm:"not null" json:"otr_price"`
	AdminFee        float64 `gorm:"not null" json:"admin_fee"`
	TotalInstalment int     `gorm:"not null" json:"total_instalment"`
	TotalInterest   float64 `gorm:"not null" json:"total_interest"`
	ProductName     string  `gorm:"not null" json:"product_name"`
	CustomerID      int
	Customer        *Customer
	LimitID         int
	Limit           *Limit
}
