package paymentUsecase

import "cleanarch/src/entity/payment"

type InsertPaymentRepository interface {
	InsertPayment(payment *paymentEntity.Payment)
}
