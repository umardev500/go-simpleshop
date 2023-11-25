package domain

import (
	"simpleshop/domain/model"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler interface {
	Create(c *fiber.Ctx) error
	Find(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type ProductUsecase interface {
	Create(params model.ProductModelNew) error
	Find() ([]model.ProductModel, error)
	FindById(id string) (model.ProductModel, error)
	Delete(id string) error
}

type ProductRepo interface {
	Create(params model.ProductModelNew) error
	Find() ([]model.ProductModel, error)
	FindById(id string) (model.ProductModel, error)
	Delete(id string) (int64, error)
}
