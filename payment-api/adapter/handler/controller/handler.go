package controller

import (
	"github.com/gin-gonic/gin"
)

// Handler provides order funcionalities
type Handler struct{}

// NewHandler is the payment handler builder
func NewHandler() *Handler {
	return &Handler{}
}

// RegisterRoutes register routes
func (h *Handler) RegisterRoutes(routes *gin.RouterGroup) {
	paymentRoute := routes.Group("/payments")
	paymentRoute.POST("/:id", h.MakePayment)
}
