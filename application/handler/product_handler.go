package handler

import (
	"simpleshop/domain"
	"simpleshop/domain/model"

	"github.com/gofiber/fiber/v2"
)

type productHandler struct {
	uc domain.ProductUsecase
}

func NewProductHandler(uc domain.ProductUsecase) domain.ProductHandler {
	return &productHandler{
		uc: uc,
	}
}

func (p *productHandler) Create(c *fiber.Ctx) error {
	var product model.ProductModelNew
	err := c.BodyParser(&product)
	if err != nil {
		return c.Status(400).JSON("Bad request")
	}

	err = p.uc.Create(product)
	if err != nil {
		return c.Status(500).JSON("failed to create product")
	}

	// succeed block
	return c.JSON("succeed")
}
