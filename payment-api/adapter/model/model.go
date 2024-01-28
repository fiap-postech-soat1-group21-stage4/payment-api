package model

import "github.com/google/uuid"

type PaymentDTO struct {
	OrderID uuid.UUID     `json:"order_id"`
	Status  PaymentStatus `json:"status"`
}

type PaymentStatus string

const (
	Paid     = "paid"
	Rejected = "rejected"
)

func (s PaymentStatus) String() string {
	switch s {
	case Paid:
		return Paid
	case Rejected:
		return Rejected
	}
	return "unknown"
}
