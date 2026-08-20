package incomes

import (
	"time"

	"github.com/RicardoEmm/rosinas-food/internal/domain"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Income struct {
	ID          uuid.UUID       `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Description string          `gorm:"size:250" json:"description"`
	ProductType ProductType     `gorm:"size:10;not null" json:"product_type"`
	Quantity    int             `gorm:"not null" json:"quantity"`
	UnitPrice   decimal.Decimal `gorm:"type:numeric(10,2);not null" json:"unit_price"`
	Total       decimal.Decimal `gorm:"type:numeric(10,2);not null" json:"total"`
	CustomerID  uuid.UUID       `gorm:"not null" json:"customer_id"`
	Customer    domain.Customer `gorm:"foreignKey:CustomerID" json:"-"`
	Status      PaymentStatus   `gorm:"size:15;not null" json:"status"`
	CreatedAt   time.Time       `json:"created_at"`
}

func (i *Income) BeforeCreate(tx *gorm.DB) error {
	if i.ID == uuid.Nil {
		i.ID = uuid.New()
	}
	return nil
}
