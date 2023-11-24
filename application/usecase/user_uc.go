package usecase

import (
	"simpleshop/domain"
	"simpleshop/domain/model"
)

type userUsecase struct {
	repo domain.UserRepo
}

func NewUserUsecase(repo domain.UserRepo) domain.UserUsecase {
	return &userUsecase{
		repo: repo,
	}
}

func (u *userUsecase) Create(user model.NewUserModel) error {
	return u.repo.Create(user)
}
