package route

import (
	"cleanarch/src/factory"

	"github.com/gin-gonic/gin"
)

type GinRoute struct {
}

func NewGinRoute() *GinRoute {
	return new(GinRoute)
}

func (r *GinRoute) GetEngine() *gin.Engine {
	engine := gin.New()
	paymentRoute := factory.NewGinPaymentEndpoint()

	engine.POST("/pay", paymentRoute.PayRoute)
	return engine
}
