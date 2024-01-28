package controller

import (
	"net/http"

	"github.com/fiap-postech-soat1-group21-stage4/payment-api/payment-api/adapter/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) MakePayment(ctx *gin.Context) {
	orderID := ctx.Param("id")
	id := uuid.MustParse(orderID)

	output := &model.PaymentDTO{
		OrderID: id,
		Status:  model.Paid,
	}

	ctx.JSON(http.StatusOK, output)
}
