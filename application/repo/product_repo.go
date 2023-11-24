package repo

import (
	"database/sql"
	"simpleshop/domain"
	"simpleshop/domain/model"
)

type productRepo struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) domain.ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (p *productRepo) Create(params model.ProductModelNew) error {
	queryStr := `--sql
		INSERT INTO products (name, price, stock) VALUES ($1, $2, $3);
	`

	_, err := p.db.Exec(queryStr, params.Name, params.Price, params.Stock)
	return err
}

func (p *productRepo) Find() ([]model.ProductModel, error) {
	queryStr := `--sql
	SELECT name, price, stock, created_at FROM products;
	`
	rows, err := p.db.Query(queryStr)
	if err != nil {
		return nil, err
	}

	var result []model.ProductModel

	for rows.Next() {
		var each model.ProductModel
		err := rows.Scan(&each.Name, &each.Price, &each.Stock, &each.CreatedAt)
		if err != nil {
			return nil, err
		}

		result = append(result, each)
	}

	return result, nil
}
