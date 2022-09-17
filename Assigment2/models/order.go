package models

type Order struct {
	ID           uint   `gorm:"primaryKey" json:"order_id"`
	CustomerName string `json:"customer_name"`
	Quantity     string `json:"quantity"`
}
