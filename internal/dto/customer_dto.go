package dto

import "github.com/shopspring/decimal"

type CustomerDTO struct {
	FullName string          `json:"full_name" binding:"required,min=3,max=75"`
	Phone    string          `json:"phone" binding:"required,ie164"`
	Price    decimal.Decimal `json:"price" binding:"required,gt=0"`
}

type CreateCustomerInput struct {
	FullName string
	Phone    string
	Price    decimal.Decimal
}
