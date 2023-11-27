package handler

import (
	"simpleshop/domain"
	"simpleshop/domain/model"

	"github.com/gofiber/fiber/v2"
)

type orderHandler struct {
	uc domain.OrderUsecase
}

func NewOrderHandler(uc domain.OrderUsecase) domain.OrderHandler {
	return &orderHandler{
		uc: uc,
	}
}

func (o *orderHandler) Find(c *fiber.Ctx) error {
	orders, err := o.uc.Find()
	if err != nil {
		return c.Status(500).JSON("failed to get data")
	}

	return c.JSON(orders)
}

func (o *orderHandler) Create(c *fiber.Ctx) error {
	var payload model.NewOrderModel
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON("bad request")
	}

	err := o.uc.Create(payload)
	if err != nil {
		return c.Status(500).JSON("failed to create order")
	}

	// succeed block
	return c.JSON("succeed")
}
