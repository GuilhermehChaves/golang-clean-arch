package repository

import (
	paymentEntity "cleanarch/src/entity/payment"
	"log"
	"time"
)

type PaymentRepositoryImpl struct {
}

func NewPaymentRepositoryImpl() *PaymentRepositoryImpl {
	return new(PaymentRepositoryImpl)
}

func (p *PaymentRepositoryImpl) InsertPayment(payment *paymentEntity.Payment) {
	// insert here
	log.Printf("Payed %s", time.Now().Format("2006-01-02 15:04:05"))
}
