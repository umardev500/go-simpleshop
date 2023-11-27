package repo

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"simpleshop/domain"
	"simpleshop/domain/model"
)

type orderRepository struct {
	DB *sql.DB
}

func NewOrderRepository(db *sql.DB) domain.OrderRepo {
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

func (r *orderRepository) Delete(id string) error {
	// Assuming you have a table named "orders" in your database
	query := `
		DELETE FROM orders
		WHERE id = $1
	`

	_, err := r.DB.Exec(query, id)
	return err
}

func (r *orderRepository) Find() ([]model.Order, error) {
	// Assuming you have a table named "orders" in your database
	query := `
		SELECT json_build_object(
			'id', id,
			'user_id', user_id,
			'order_number', order_number,
			'status', status,
			'total', total,
			'created_at', created_at,
			'updated_at', updated_at
		) FROM orders
	`

	// Execute the SQL query
	rows, err := r.DB.Query(query)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to find orders: %v", err)
	}
	defer rows.Close()

	var orders []model.Order
	for rows.Next() {
		var each []byte
		if err := rows.Scan(&each); err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("failed to scan order: %v", err)
		}

		var order model.Order
		if err := json.Unmarshal(each, &order); err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("failed to unmarshal order: %v", err)
		}

		orders = append(orders, order)
	}

	return orders, nil

}

func (r *orderRepository) FindById(id string) (model.Order, error) {
	// Assuming you have a table named "orders" in your database
	query := `
		SELECT json_build_object(
			'id', id,
			'user_id', user_id,
			'order_number', order_number,
			'status', status,
			'total', total,
			'created_at', created_at,
			'updated_at', updated_at
		) FROM orders
		WHERE id = $1
	`

	// Execute the SQL query
	row := r.DB.QueryRow(query, id)
	var each []byte

	if err := row.Scan(&each); err != nil {
		fmt.Println(err)
		return model.Order{}, fmt.Errorf("failed to find order: %v", err)
	}

	var order model.Order
	if err := json.Unmarshal(each, &order); err != nil {
		fmt.Println(err)
		return model.Order{}, fmt.Errorf("failed to unmarshal order: %v", err)
	}

	return order, nil
}
