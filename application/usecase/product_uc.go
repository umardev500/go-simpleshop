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

func (p *productUsecase) Find() ([]model.ProductModel, error) {
	data, err := p.repo.Find()
	return data, err
}

// Find product by id
func (p *productUsecase) FindById(id string) (model.ProductModel, error) {
	return p.repo.FindById(id)
}
