package service

import (
	"context"

	apperors "github.com/RicardoEmm/rosinas-food/internal/apperrors"
	"github.com/RicardoEmm/rosinas-food/internal/domain/incomes"
	"github.com/RicardoEmm/rosinas-food/internal/dto"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type IncomeService struct {
	incomeRepo   IncomeRepo
	customerRepo CustomerRepo
}

func NewIncomeService(incomeRepo IncomeRepo, customerRepo CustomerRepo) *IncomeService {
	return &IncomeService{incomeRepo: incomeRepo, customerRepo: customerRepo}
}

func (s *IncomeService) FindById(ctx context.Context, id uuid.UUID) (*incomes.Income, error) {
	income, err := s.incomeRepo.FindById(ctx, id)

	if err != nil {
		return nil, apperors.ErrIncomeNotFound
	}

	return income, nil
}

func (s *IncomeService) FindAll(ctx context.Context) ([]*incomes.Income, error) {
	incomes, err := s.incomeRepo.FindAll(ctx)

	if err != nil {
		return nil, apperors.ErrInternalServer
	}

	return incomes, nil
}

func (s *IncomeService) FindAllByCustomerId(ctx context.Context, customerId uuid.UUID) ([]*incomes.Income, error) {
	incomes, err := s.incomeRepo.FindAllByCustomerId(ctx, customerId)

	if err != nil {
		return nil, apperors.ErrInternalServer
	}

	return incomes, nil
}

func (s *IncomeService) Create(ctx context.Context, input dto.CreateIncomeInput) error {
	customer, err := s.customerRepo.FindById(ctx, input.CustomerID)

	if err != nil {
		return apperors.ErrCustomerNotFound
	}

	unitPrice, err := input.ResolvePrice(customer.Price)
	if err != nil {
		return apperors.ErrInvalidUnitPrice
	}

	if err := s.incomeRepo.Create(ctx, &incomes.Income{
		Description: input.Description,
		ProductType: incomes.ProductType(input.ProductType),
		Quantity:    input.Quantity,
		UnitPrice:   unitPrice,
		Total:       unitPrice.Mul(decimal.NewFromInt(int64(input.Quantity))),
		CustomerID:  customer.ID,
		Status:      incomes.PaymentStatus(input.Status),
	}); err != nil {
		return apperors.ErrIncomeNotCreated
	}

	return nil
}

func (s *IncomeService) ChangeToPaid(ctx context.Context, id uuid.UUID) error {
	if err := s.incomeRepo.ChangeToPaid(ctx, id); err != nil {
		return apperors.ErrInternalServer
	}
	return nil
}

func (s *IncomeService) DeleteById(ctx context.Context, id uuid.UUID) error {
	if err := s.incomeRepo.DeleteById(ctx, id); err != nil {
		return apperors.ErrInternalServer
	}
	return nil
}

func (s *IncomeService) DeleteAllByCustomerId(ctx context.Context, customerId uuid.UUID) error {
	if err := s.incomeRepo.DeleteAllByCustomerId(ctx, customerId); err != nil {
		return apperors.ErrInternalServer
	}
	return nil
}
