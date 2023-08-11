package entities

type (
	CreateOrderRequest struct {
		UserID     string `json:"user_id"`
		Item       string `json:"item"`
		Quantity   int    `json:"quantity"`
		TotalPrice int64  `json:"total_price"`
	}

	CreateOrderEvent struct {
		OrderID    string `json:"order_id"`
		TotalPrice int64  `json:"total_price"`
	}
)
