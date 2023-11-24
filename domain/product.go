package domain

import (
	"simpleshop/domain/model"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler interface {
	Create(c *fiber.Ctx) error
}

type ProductUsecase interface {
	Create(params model.ProductModelNew) error
}

type ProductRepo interface {
	Create(params model.ProductModelNew) error
}
