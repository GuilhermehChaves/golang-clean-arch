package controller

import (
	"cleanarch/src/controller/model"
	paymentModel "cleanarch/src/controller/model/payment"
)

type PaymentController interface {
	Pay(body *paymentModel.PaymentRequest) *model.RestResponse
}
