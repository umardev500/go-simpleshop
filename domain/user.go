package domain

import (
	"simpleshop/domain/model"

	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	Create(c *fiber.Ctx) error
}

type UserUsecase interface {
	Create(payload model.NewUserModel) error
}

type UserRepo interface {
	Create(payload model.NewUserModel) error
}
