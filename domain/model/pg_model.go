package model

type TransactionDetails struct {
	GrossAmount float64 `json:"gross_amount"`
	OrderID     string  `json:"order_id"`
}

type BankTransfer struct {
	Bank string `json:"bank"`
}

type PaymentRequest struct {
	PaymentType        string             `json:"payment_type"`
	TransactionDetails TransactionDetails `json:"transaction_details"`
	BankTransfer       BankTransfer       `json:"bank_transfer"`
}

// Response model
type BankTransferTransaction struct {
	StatusCode        *string    `json:"status_code"`
	StatusMessage     *string    `json:"status_message"`
	TransactionID     *string    `json:"transaction_id"`
	OrderID           *string    `json:"order_id"`
	MerchantID        *string    `json:"merchant_id"`
	GrossAmount       *string    `json:"gross_amount"`
	Currency          *string    `json:"currency"`
	PaymentType       *string    `json:"payment_type"`
	TransactionTime   *string    `json:"transaction_time"`
	TransactionStatus *string    `json:"transaction_status"`
	FraudStatus       *string    `json:"fraud_status"`
	VANumbers         []VANumber `json:"va_numbers"`
	ExpiryTime        *string    `json:"expiry_time"`
}

type VANumber struct {
	Bank     *string `json:"bank"`
	VANumber *string `json:"va_number"`
}

// Callback
type Callback struct {
	StatusCode        string     `json:"status_code"`
	Message           string     `json:"status_message"`
	TransactionStatus string     `json:"transaction_status"`
	VANumbers         []VANumber `json:"va_numbers"`
}
