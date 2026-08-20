package dto

import (
	"github.com/google/uuid"
)

type CreateIncomeInput struct {
	Description string
	ProductType string
	UnitPrice   string
	Quantity    int
	CustomerID  uuid.UUID
	Status      string
}

type IncomeRequest struct {
	Description string    `json:"description" binding:"omitempty"`
	ProductType string    `json:"product_type" binding:"required,oneof=unit tray"`
	UnitPrice   string    `json:"unit_price" binding:"required"`
	Quantity    int       `json:"quantity" binding:"required,gt=0"`
	CustomerID  uuid.UUID `json:"customer_id" binding:"required"`
	Status      string    `json:"status" binding:"required,oneof=pending paid"`
}
