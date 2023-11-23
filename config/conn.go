package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewConn() *sql.DB {
	host := "127.0.0.1"
	port := "5432"
	user := "root"
	pass := "root"
	dbName := "simpleshop"
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("connected")

	return db
}
