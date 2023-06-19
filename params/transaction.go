package params

type TransactionPayload struct {
	CustomerID    int     `json:"customer_id" binding:"required"`
	OTRPrice      float64 `json:"otr_price" binding:"required"`
	Tenor         int     `json:"tenor" binding:"required"`
	AdminFee      float64 `json:"admin_fee" `
	TotalInterest float64 `json:"total_interest"`
	ProductName   string  `json:"product_name" binding:"required"`
}
