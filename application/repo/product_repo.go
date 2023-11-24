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
