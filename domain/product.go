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
	Update(c *fiber.Ctx) error
}

type ProductUsecase interface {
	Create(params model.ProductModelNew) error
	Find() ([]model.ProductModel, error)
	FindById(id int64) (model.ProductModel, error)
	Delete(id int64) error
	Update(id int64, payload model.ProductModelNew) error
}

type ProductRepo interface {
	Create(params model.ProductModelNew) error
	Find() ([]model.ProductModel, error)
	// newer
	FindById(id int64) (model.ProductModel, error)
	Delete(id int64) (int64, error)
	Update(id int64, payload model.ProductModelNew) (int64, error)
}
