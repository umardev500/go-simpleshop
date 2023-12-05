package main

import (
	"fmt"
	"simpleshop/router"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("failed to load .env")
		return
	}
	router.NewRouter()
}
