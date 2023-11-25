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
	SELECT id, name, price, stock, created_at FROM products;
	`
	rows, err := p.db.Query(queryStr)
	if err != nil {
		return nil, err
	}

	var result []model.ProductModel

	for rows.Next() {
		var each model.ProductModel
		err := rows.Scan(&each.Id, &each.Name, &each.Price, &each.Stock, &each.CreatedAt)
		if err != nil {
			return nil, err
		}

		result = append(result, each)
	}

	return result, nil
}

// Find product by id
func (p *productRepo) FindById(id string) (model.ProductModel, error) {
	queryStr := `--sql
	SELECT id, name, price, stock, created_at FROM products
	WHERE id = $1;
	`

	var product model.ProductModel
	row := p.db.QueryRow(queryStr, id)
	err := row.Scan(&product.Id, &product.Name, &product.Price, &product.Stock, &product.CreatedAt)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (p *productRepo) Delete(id string) (int64, error) {
	queryStr := `--sql
	DELETE FROM products WHERE id = $1;
	`

	result, err := p.db.Exec(queryStr, id)
	if err != nil {
		return 0, err
	}

	affected, _ := result.RowsAffected()
	return affected, nil
}
