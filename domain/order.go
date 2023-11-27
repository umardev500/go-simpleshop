package domain

import (
	"simpleshop/domain/model"

	"github.com/gofiber/fiber/v2"
)

type OrderHandler interface {
	Create(c *fiber.Ctx) error
}

type OrderUsecase interface {
	Create(payload model.NewOrderModel) error
}

type OrderRepo interface {
	Create(payload model.Order) error
}
