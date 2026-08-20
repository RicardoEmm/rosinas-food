package incomes

type ProductType string

const (
	ProductTypeUnit ProductType = "unit"
	ProductTypeTray ProductType = "tray"
)

func (p ProductType) IsValid() bool {
	switch p {
	case ProductTypeUnit, ProductTypeTray:
		return true
	default:
		return false
	}
}
