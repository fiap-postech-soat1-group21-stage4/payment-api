package manage

import (
	p "github.com/fiap-postech-soat1-group21-stage4/payment-api/payment-api/adapter/handler/controller"
	"github.com/gin-gonic/gin"
)

type apps interface {
	RegisterRoutes(routes *gin.RouterGroup)
}

type Manage struct {
	payment apps
}

type UseCases struct{}

func New(uc *UseCases) *Manage {

	paymentHandler := p.NewHandler()

	return &Manage{
		payment: paymentHandler,
	}
}

func (m *Manage) RegisterRoutes(group *gin.RouterGroup) {
	m.payment.RegisterRoutes(group)
}
