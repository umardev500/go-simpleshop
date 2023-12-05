package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"simpleshop/constant"
	"simpleshop/domain"
	"simpleshop/domain/model"
	"simpleshop/utils"
	"strconv"
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

	// Call midtrans
	orderID := utils.GenerateRandomNumber() // create order id
	// Create data for midtrans
	payment := model.PaymentRequest{
		PaymentType: constant.BankTransfer,
		TransactionDetails: model.TransactionDetails{
			OrderID:     orderID,
			GrossAmount: total,
		},
		BankTransfer: model.BankTransfer{
			Bank: constant.Bni,
		},
	}
	// Convert to json
	bodyString, _ := json.Marshal(payment)
	baseURL := os.Getenv("PG_URL")
	url := fmt.Sprintf("%s/v2/charge", baseURL)

	// Use http post client
	body, err := utils.Post(url, bodyString)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var transaction model.BankTransferTransaction
	err = json.Unmarshal(body, &transaction) // convert json to struct
	if err != nil {
		return err
	}

	if *transaction.StatusCode != "201" {
		// Check for status code
		return errors.New(*transaction.StatusMessage)
	}

	va := transaction.VANumbers[0] // get va number
	id, _ := strconv.Atoi(orderID)
	order := model.Order{
		ID:          int64(id),
		UserID:      payload.UserID,
		OrderNumber: *va.VANumber,
		Total:       total,
	}

	// Do create order
	err = u.orderRepo.Create(order)
	if err != nil {
		return err
	}
	fmt.Println("order done")

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

func (u *orderUsecase) SetStatus(id int64, status string) error {
	affected, err := u.orderRepo.SetStatus(id, status)
	if err != nil {
		return err
	}

	if affected < 1 {
		return constant.ErrNoAffected
	}

	return nil
}
