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
	var payload model.NewUserModel = model.NewUserModel{
		Username: "root",
		Email:    "email@example.com",
		Password: "root",
	}

	u.uc.Create(payload)

	return nil
}
