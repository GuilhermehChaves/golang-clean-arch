package paymentUsecase

import paymentEntity "cleanarch/src/entity/payment"

type InsertPaymentUseCase struct {
	repository InsertPaymentRepository
}

func NewInsertPaymentUseCase(repository InsertPaymentRepository) *InsertPaymentUseCase {
	return &InsertPaymentUseCase{
		repository: repository,
	}
}

func (i *InsertPaymentUseCase) Insert(payment *paymentEntity.Payment) error {
	i.repository.InsertPayment(payment)
	return nil
}
