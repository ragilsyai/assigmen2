package params

type CreateItem struct {
	ItemId      uint   `json:"item_id,omitempty"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderId     uint   `json:"order_id,omitempty"`
}

type UpdateItem struct {
	OrderId      uint   `json:"order_id,omitempty"`
	CustomerName string `json:"customer_name,omitempty"`
	Quantity     int    `json:"quantity,omitempty"`
}
