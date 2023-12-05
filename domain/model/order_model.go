package model

import "time"

type Order struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	OrderNumber string    `json:"order_number"`
	Total       float64   `json:"total"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type NewOrderModel struct {
	UserID   int64   `json:"user_id"`
	Products []int64 `json:"products"`
}

type OrderStatusSet struct {
	ID int64 `json:"id"`
}
