package repo

import (
	"database/sql"
	"fmt"
	"simpleshop/domain/model"
)

type orderRepository struct {
	DB *sql.DB
}

func NewOrderRepository(db *sql.DB) *orderRepository {
	return &orderRepository{
		DB: db,
	}
}

func (r *orderRepository) Create(payload model.Order) error {
	// Assuming you have a table named "orders" in your database
	query := `
		INSERT INTO orders (
			user_id,
			order_number,
			total
		) VALUES (
			$1, $2, $3
		)
	`

	// Execute the SQL query
	_, err := r.DB.Exec(query, payload.UserID, payload.OrderNumber, payload.Total)

	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("failed to create order: %v", err)
	}

	return nil
}
