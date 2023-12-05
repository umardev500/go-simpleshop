package usecase

import (
	"fmt"
	"simpleshop/domain"
	"simpleshop/domain/model"
	"simpleshop/utils"
)

type orderUsecase struct {
	orderRepo   domain.OrderRepo
	productRepo domain.ProductRepo
}

func NewOrderUsecase(orderRepo domain.OrderRepo, productRepo domain.ProductRepo) domain.OrderUsecase {
	return &orderUsecase{
		orderRepo:   orderRepo,
		productRepo: productRepo,
	}
}

func (u *orderUsecase) Create(payload model.NewOrderModel) error {
	// Your business logic before creating an order goes here, if needed
	var total float64
	for _, productID := range payload.Products {
		product, err := u.productRepo.FindById(productID)
		if err != nil {
			fmt.Println(err)
			return err
		}

		// calculate
		price := product.Price
		total += price
	}

	order := model.Order{
		UserID:      payload.UserID,
		OrderNumber: utils.GenerateRandomNumber(),
		Total:       total,
	}
	// Delegate the creation to the repository
	err := u.orderRepo.Create(order)
	if err != nil {
		// Handle the error appropriately
		return err
	}

	// Your business logic after creating an order goes here, if needed

	return nil
}

func (u *orderUsecase) Delete(id string) error {
	err := u.orderRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (u *orderUsecase) Find() ([]model.Order, error) {
	orders, err := u.orderRepo.Find()
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (u *orderUsecase) FindById(id string) (model.Order, error) {
	order, err := u.orderRepo.FindById(id)
	if err != nil {
		return model.Order{}, err
	}
	return order, nil
}
