package controller_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fiap-postech-soat1-group21-stage4/payment-api/payment-api/adapter/handler/controller"
	"github.com/fiap-postech-soat1-group21-stage4/payment-api/payment-api/adapter/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestMakePayment(t *testing.T) {
	// Setup
	r := gin.Default()
	handler := controller.NewHandler()
	r.POST("/payments/:id", handler.MakePayment)

	t.Run("MakePayment", func(t *testing.T) {
		// Create a request
		orderID := "8c2b51bf-7b4c-4a4b-a024-f283576cf190" // Replace with a valid UUID
		req, err := http.NewRequest("POST", "/payments/"+orderID, nil)
		assert.NoError(t, err)

		// Create a response recorder to record the response
		w := httptest.NewRecorder()

		// Serve the request to the recorder
		r.ServeHTTP(w, req)

		// Check the response code
		assert.Equal(t, http.StatusOK, w.Code)

		// Decode the response body
		var paymentDTO model.PaymentDTO
		err = json.Unmarshal(w.Body.Bytes(), &paymentDTO)
		assert.NoError(t, err)

		// Validate the response
		expectedOrderID, _ := uuid.Parse(orderID)
		expectedOutput := &model.PaymentDTO{
			OrderID: expectedOrderID,
			Status:  model.Paid,
		}

		assert.Equal(t, expectedOutput, &paymentDTO)
	})
}
