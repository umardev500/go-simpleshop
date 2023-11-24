package handler

import (
	"simpleshop/domain"
	"simpleshop/domain/model"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	uc domain.UserUsecase
}

func NewUserHandler(uc domain.UserUsecase) domain.UserHandler {
	return &userHandler{
		uc: uc,
	}
}

func (u *userHandler) Create(c *fiber.Ctx) error {
	var user model.NewUserModel
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(500).JSON("Failed to parsing data")
	}

	var payload model.NewUserModel = model.NewUserModel{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}

	err = u.uc.Create(payload)
	if err != nil {
		return c.Status(500).JSON("Error failed to create user")
	}

	return c.JSON("Succesfuly create user")
}
