package domain

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Customer struct {
	ID       uuid.UUID       `gorm:"primaryKey" json:"id"`
	FullName string          `gorm:"size:75;not null" json:"full_name"`
	Phone    string          `gorm:"type:varchar(20);not null;uniqueIndex" json:"phone"`
	Price    decimal.Decimal `gorm:"type:decimal(10,2);not null" json:"price"`
}
