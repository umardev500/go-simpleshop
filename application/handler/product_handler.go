package handler

import (
	"fmt"
	"simpleshop/constant"
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

func (p *productHandler) Find(c *fiber.Ctx) error {
	data, err := p.uc.Find()
	if err != nil {
		fmt.Println(err)
		return c.Status(500).JSON("failed to get data")
	}

	return c.JSON(data)
}

// New code bellow
func (p *productHandler) FindById(c *fiber.Ctx) error {
	var id = c.Params("id")
	data, err := p.uc.FindById(id)
	if err != nil {
		fmt.Println(err)
		return c.Status(500).JSON("failed to get data")
	}

	return c.JSON(data)
}

func (p *productHandler) Delete(c *fiber.Ctx) error {
	var id = c.Params("id")
	err := p.uc.Delete(id)
	if err != nil {
		if err == constant.ErrNoAffected {
			return c.Status(404).JSON("data to delete not found")
		}

		return c.Status(500).JSON("failed to delete")
	}

	return c.JSON("product deleted")
}

func (p *productHandler) Update(c *fiber.Ctx) error {
	var product model.ProductModelNew
	err := c.BodyParser(&product)
	if err != nil {
		return c.Status(400).JSON("bad request")
	}

	var id = c.Params("id")
	var payload = model.ProductModelNew{
		Name:  product.Name,
		Price: product.Price,
		Stock: product.Stock,
	}
	err = p.uc.Update(id, payload)
	if err != nil {
		if err == constant.ErrNoAffected {
			return c.Status(404).JSON("data to update not found")
		}

		return c.Status(500).JSON("failed to update")
	}

	return c.JSON("product updated")
}
