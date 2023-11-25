package router

import (
	"simpleshop/application/handler"
	"simpleshop/application/repo"
	"simpleshop/application/usecase"
	"simpleshop/config"

	"github.com/gofiber/fiber/v2"
)

func NewRouter() {
	app := fiber.New()
	db := config.NewConn()

	// User routers
	userRepo := repo.NewUserRepo(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)
	app.Post("/users", userHandler.Create)

	// Product routers
	productRepo := repo.NewProductRepo(db)
	productUsecase := usecase.NewProductUsecase(productRepo)
	productHandler := handler.NewProductHandler(productUsecase)
	app.Get("/products", productHandler.Find)
	app.Get("/products/:id", productHandler.FindById)
	app.Post("/products", productHandler.Create)
	app.Delete("/products/:id", productHandler.Delete)

	app.Listen(":9000")
}
