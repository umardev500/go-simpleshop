package domain

import (
	"simpleshop/domain/model"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler interface {
	Create(c *fiber.Ctx) error
	Find(c *fiber.Ctx) error
}

type ProductUsecase interface {
	Create(params model.ProductModelNew) error
	Find() ([]model.ProductModel, error)
}

type ProductRepo interface {
	Create(params model.ProductModelNew) error
	Find() ([]model.ProductModel, error)
}
