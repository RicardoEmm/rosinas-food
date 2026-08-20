package dto

import (
	"github.com/RicardoEmm/rosinas-food/internal/domain/incomes"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
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

func (i *CreateIncomeInput) ResolvePrice(customerPrice decimal.Decimal) (decimal.Decimal, error) {
	if i.ProductType == string(incomes.ProductTypeUnit) {
		unitPrice, err := decimal.NewFromString(i.UnitPrice)

		if err != nil {
			return decimal.Zero, err
		}

		return unitPrice, nil
	}
	return customerPrice, nil
}
