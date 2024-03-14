package model

import "time"

type Order struct {
	OrderID    int       `json:"order_id"`
	UserID     int       `json:"user_id"`
	Quantity   int       `json:"quantity"`
	TotalPrice int       `json:"total_price"`
	ServiceID  int       `json:"service_id"`
	DurationID int       `json:"duration_id"`
	CreatedAt  time.Time `json:"created_at"`
}
