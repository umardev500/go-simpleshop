package handler

import (
	"fmt"
	"simpleshop/domain"
	"simpleshop/domain/model"
	"strconv"

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

func (o *orderHandler) FindById(c *fiber.Ctx) error {
	var id = c.Params("id")
	order, err := o.uc.FindById(id)
	if err != nil {
		return c.Status(500).JSON("failed to get data")
	}

	return c.JSON(order)
}

func (o *orderHandler) Delete(c *fiber.Ctx) error {
	var id = c.Params("id")
	err := o.uc.Delete(id)
	if err != nil {
		return c.Status(500).JSON("failed to delete order")
	}

	return c.JSON("order deleted")
}

func (o *orderHandler) Create(c *fiber.Ctx) error {
	var payload model.NewOrderModel
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON("bad request")
	}

	err := o.uc.Create(payload)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	// succeed block
	return c.JSON("succeed")
}

func (o *orderHandler) Callback(c *fiber.Ctx) error {
	fmt.Println("calling...")
	var payload model.Callback
	if err := c.BodyParser(&payload); err != nil {
		fmt.Println(err)
		return c.Status(400).JSON("bad request")
	}
	if payload.StatusCode != "201" && payload.StatusCode != "200" {
		fmt.Println(payload.Message)
		// Check for status code
		return c.Status(500).JSON("failed to create payment")
	}

	if payload.TransactionStatus != "pending" {
		// call api
		vaNumber := payload.VANumbers[0].VANumber
		vaNumberInt, _ := strconv.Atoi(*vaNumber)
		err := o.uc.SetStatus(int64(vaNumberInt), payload.TransactionStatus)
		if err != nil {
			fmt.Println(err)
			return c.SendStatus(500)
		}
		fmt.Println("called", payload.TransactionStatus, *vaNumber)
	}

	return c.JSON("ok")
}
