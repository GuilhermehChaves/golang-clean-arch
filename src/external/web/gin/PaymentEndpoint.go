package gin

import (
	"cleanarch/src/controller"
	"cleanarch/src/controller/model"
	paymentModel "cleanarch/src/controller/model/payment"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PaymentEndpoint struct {
	paymentController controller.PaymentController
}

func NewGinPaymentEndpoint(paymentController controller.PaymentController) *PaymentEndpoint {
	return &PaymentEndpoint{
		paymentController: paymentController,
	}
}

// PayRoute POST /pay
func (p *PaymentEndpoint) PayRoute(ctx *gin.Context) {
	init := time.Now()

	requestBody := new(paymentModel.PaymentRequest)
	//var requestBody *paymentModel.PaymentRequest

	if err := ctx.BindJSON(requestBody); err != nil {
		errorResponse := model.NewErrorResponse("error parsing body", err.Error())
		restResponse := model.NewRestResponseBuilder().
			WithTimestamps().
			WithData(errorResponse).
			Build()

		ctx.JSON(http.StatusBadRequest, restResponse)
		return
	}

	response := p.paymentController.Pay(requestBody)
	ctx.JSON(response.Status, response)
	log.Println("request took", time.Since(init).Seconds()*1000, "ms")
	return
}
