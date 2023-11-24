package usecase

import (
	"simpleshop/domain"
	"simpleshop/domain/model"
)

type productUsecase struct {
	repo domain.ProductRepo
}

func NewProductUsecase(repo domain.ProductRepo) domain.ProductUsecase {
	return &productUsecase{
		repo: repo,
	}
}

func (p *productUsecase) Create(params model.ProductModelNew) error {
	return p.repo.Create(params)
}
