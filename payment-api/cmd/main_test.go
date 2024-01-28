package main_test

import (
	"testing"

	model "github.com/fiap-postech-soat1-group21-stage4/payment-api/payment-api/adapter/model"
	"github.com/google/uuid"
)

func TestPaymentDTO(t *testing.T) {
	orderID := uuid.New()
	payment := model.PaymentDTO{
		OrderID: orderID,
		Status:  model.Paid,
	}

	// Testa se os valores são atribuídos corretamente
	if payment.OrderID != orderID {
		t.Errorf("Expected OrderID %v, got %v", orderID, payment.OrderID)
	}

	if payment.Status != model.Paid {
		t.Errorf("Expected Status %v, got %v", model.Paid, payment.Status)
	}
}

func TestPaymentStatusString(t *testing.T) {
	// Testa se a conversão para string funciona corretamente
	paidString := model.Paid
	if paidString != "paid" {
		t.Errorf("Expected 'paid', got %v", paidString)
	}

	rejectedString := model.Rejected
	if rejectedString != "rejected" {
		t.Errorf("Expected 'rejected', got %v", rejectedString)
	}

	unknownString := model.PaymentStatus("unknown")
	if unknownString != "unknown" {
		t.Errorf("Expected 'unknown', got %v", unknownString)
	}
}
