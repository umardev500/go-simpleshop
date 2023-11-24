package repo

import (
	"database/sql"
	"simpleshop/domain"
	"simpleshop/domain/model"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) domain.UserRepo {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) Create(payload model.NewUserModel) error {
	queryStr := `
	INSERT INTO users (username, email, password) VALUES ($1, $2, $3);
	`
	_, err := u.db.Exec(queryStr, payload.Username, payload.Email, payload.Password)
	return err
}
