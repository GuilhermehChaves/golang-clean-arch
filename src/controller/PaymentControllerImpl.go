package controller

import (
	"cleanarch/src/controller/model"
	paymentModel "cleanarch/src/controller/model/payment"
	paymentUsecase "cleanarch/src/usecase/payment"
	"net/http"
)

type PaymentControllerImpl struct {
	usecase *paymentUsecase.InsertPaymentUseCase
}

func NewPaymentControllerImpl(usecase *paymentUsecase.InsertPaymentUseCase) *PaymentControllerImpl {
	return &PaymentControllerImpl{
		usecase: usecase,
	}
}

func (p *PaymentControllerImpl) Pay(body *paymentModel.PaymentRequest) *model.RestResponse {
	if err := p.usecase.Insert(nil); err != nil {
		errorResponse := model.NewErrorResponse("Internal Server Error", err.Error())
		return model.NewRestResponseBuilder().
			WithStatus(http.StatusInternalServerError).
			WithTimestamps().
			WithData(errorResponse).
			Build()
	}

	paymentResponse := paymentModel.NewPaymentResponse("https://www.ciandt.com")
	return model.NewRestResponseBuilder().
		WithStatus(http.StatusOK).
		WithTimestamps().
		WithData(paymentResponse).
		Build()
}
