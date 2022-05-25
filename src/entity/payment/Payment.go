package paymentEntity

type Payment struct {
	Url    string
	Amount float32
}

func NewPayment(url string, amount float32) *Payment {
	return &Payment{
		Url:    url,
		Amount: amount,
	}
}
