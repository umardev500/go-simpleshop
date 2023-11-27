package usecase

import (
	"simpleshop/domain"
	"simpleshop/domain/model"
	"simpleshop/utils"
)

type orderUsecase struct {
	orderRepo domain.OrderRepo
}

func NewOrderUsecase(orderRepo domain.OrderRepo) domain.OrderUsecase {
	return &orderUsecase{
		orderRepo: orderRepo,
	}
}

func (u *orderUsecase) Create(payload model.NewOrderModel) error {
	// Your business logic before creating an order goes here, if needed

	order := model.Order{
		UserID:      payload.UserID,
		OrderNumber: utils.GenerateRandomNumber(),
		Total:       44.0,
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
