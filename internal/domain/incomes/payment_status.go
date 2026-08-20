package incomes

type PaymentStatus string

const (
	PaymentStatusPending PaymentStatus = "pending"
	PaymentStatusPaid    PaymentStatus = "paid"
)

func (s PaymentStatus) IsValid() bool {
	switch s {
	case PaymentStatusPending, PaymentStatusPaid:
		return true
	default:
		return false
	}
}

func (s PaymentStatus) IsValidToMarkAsPaid() bool {
	if s == PaymentStatusPending {
		return true
	}
	return false
}
