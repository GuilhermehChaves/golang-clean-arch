package model

type PaymentRequest struct {
	Amount     float32 `json:"amount"`
	SuccessUrl string  `json:"success_url"`
	CancelUrl  string  `json:"cancel_url"`
}
