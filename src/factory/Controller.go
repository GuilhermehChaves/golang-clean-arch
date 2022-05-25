package factory

import (
	"cleanarch/src/controller"
	"cleanarch/src/external/repository"
	"cleanarch/src/external/web/gin"
	usecase "cleanarch/src/usecase/payment"
)

func NewPaymentController() controller.PaymentController {
	paymentRepository := repository.NewPaymentRepositoryImpl()
	insertPaymentUseCase := usecase.NewInsertPaymentUseCase(paymentRepository)
	return controller.NewPaymentControllerImpl(insertPaymentUseCase)
}

func NewGinPaymentEndpoint() *gin.PaymentEndpoint {
	return gin.NewGinPaymentEndpoint(NewPaymentController())
}
